package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	events_processing_email "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/proto/email"
	"github.com/sirupsen/logrus"
)

type EmailService interface {
	GetAllFor(ctx context.Context, entityType entity.EntityType, entityId string) (*entity.EmailEntities, error)
	GetAllForEntityTypeByIds(ctx context.Context, entityType entity.EntityType, entityIds []string) (*entity.EmailEntities, error)
	MergeEmailTo(ctx context.Context, entityType entity.EntityType, entityId string, entity *entity.EmailEntity) (*entity.EmailEntity, error)
	UpdateEmailFor(ctx context.Context, entityType entity.EntityType, entityId string, entity *entity.EmailEntity) (*entity.EmailEntity, error)
	DetachFromEntity(ctx context.Context, entityType entity.EntityType, entityId, email string) (bool, error)
	DetachFromEntityById(ctx context.Context, entityType entity.EntityType, entityId, emailId string) (bool, error)
	DeleteById(ctx context.Context, emailId string) (bool, error)

	mapDbNodeToEmailEntity(node dbtype.Node) *entity.EmailEntity

	UpsertInEventStore(ctx context.Context, size int) (int, int, error)
}

type emailService struct {
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
}

func NewEmailService(repository *repository.Repositories, grpcClients *grpc_client.Clients) EmailService {
	return &emailService{
		repositories: repository,
		grpcClients:  grpcClients,
	}
}

func (s *emailService) getDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *emailService) GetAllFor(ctx context.Context, entityType entity.EntityType, entityId string) (*entity.EmailEntities, error) {
	session := utils.NewNeo4jReadSession(ctx, s.getDriver())
	defer session.Close(ctx)

	queryResult, err := s.repositories.EmailRepository.GetAllFor(ctx, session, common.GetContext(ctx).Tenant, entityType, entityId)
	if err != nil {
		return nil, err
	}

	emailEntities := entity.EmailEntities{}

	for _, dbRecord := range queryResult.([]*db.Record) {
		emailEntity := s.mapDbNodeToEmailEntity(dbRecord.Values[0].(dbtype.Node))
		s.addDbRelationshipToEmailEntity(dbRecord.Values[1].(dbtype.Relationship), emailEntity)
		emailEntities = append(emailEntities, *emailEntity)
	}

	return &emailEntities, nil
}

func (s *emailService) GetAllForEntityTypeByIds(ctx context.Context, entityType entity.EntityType, entityIds []string) (*entity.EmailEntities, error) {
	emails, err := s.repositories.EmailRepository.GetAllForIds(ctx, common.GetContext(ctx).Tenant, entityType, entityIds)
	if err != nil {
		return nil, err
	}

	emailEntities := entity.EmailEntities{}
	for _, v := range emails {
		emailEntity := s.mapDbNodeToEmailEntity(*v.Node)
		s.addDbRelationshipToEmailEntity(*v.Relationship, emailEntity)
		emailEntity.DataloaderKey = v.LinkedNodeId
		emailEntities = append(emailEntities, *emailEntity)
	}
	return &emailEntities, nil
}

func (s *emailService) MergeEmailTo(ctx context.Context, entityType entity.EntityType, entityId string, inputEntity *entity.EmailEntity) (*entity.EmailEntity, error) {
	session := utils.NewNeo4jWriteSession(ctx, s.getDriver())
	defer session.Close(ctx)

	var err error
	var emailNode *dbtype.Node
	var emailRelationship *dbtype.Relationship

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		emailNode, emailRelationship, err = s.repositories.EmailRepository.MergeEmailToInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
		if err != nil {
			return nil, err
		}
		emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
		if inputEntity.Primary == true {
			err := s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}

	var emailEntity = s.mapDbNodeToEmailEntity(*emailNode)
	s.addDbRelationshipToEmailEntity(*emailRelationship, emailEntity)
	return emailEntity, nil
}

func (s *emailService) UpdateEmailFor(ctx context.Context, entityType entity.EntityType, entityId string, inputEntity *entity.EmailEntity) (*entity.EmailEntity, error) {
	session := utils.NewNeo4jWriteSession(ctx, s.getDriver())
	defer session.Close(ctx)

	var err error
	var emailNode *dbtype.Node
	var emailRelationship *dbtype.Relationship
	var detachCurrentEmail = false

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		currentEmailNode, err := s.repositories.EmailRepository.GetByIdAndRelatedEntity(ctx, entityType, common.GetTenantFromContext(ctx), inputEntity.Id, entityId)
		if err != nil {
			return nil, err
		}

		currentEmail := utils.GetStringPropOrEmpty(utils.GetPropsFromNode(*currentEmailNode), "email")
		currentRawEmail := utils.GetStringPropOrEmpty(utils.GetPropsFromNode(*currentEmailNode), "rawEmail")

		if len(inputEntity.RawEmail) == 0 || inputEntity.RawEmail == currentEmail || inputEntity.RawEmail == currentRawEmail {
			// email address replace not requested, proceed with update
			emailNode, emailRelationship, err = s.repositories.EmailRepository.UpdateEmailForInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
			if err != nil {
				return nil, err
			}
			emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
			if inputEntity.Primary == true {
				err := s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
				if err != nil {
					return nil, err
				}
			}
		} else {
			// proceed with email address replace
			// merge new email address
			emailNode, emailRelationship, err = s.repositories.EmailRepository.MergeEmailToInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
			if err != nil {
				return nil, err
			}
			emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
			if inputEntity.Primary == true {
				err := s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
				if err != nil {
					return nil, err
				}
			}
			detachCurrentEmail = true
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	if detachCurrentEmail {
		_, err = s.DetachFromEntityById(ctx, entityType, entityId, inputEntity.Id)
	}

	var emailEntity = s.mapDbNodeToEmailEntity(*emailNode)
	s.addDbRelationshipToEmailEntity(*emailRelationship, emailEntity)
	return emailEntity, nil
}

func (s *emailService) DetachFromEntity(ctx context.Context, entityType entity.EntityType, entityId, email string) (bool, error) {
	err := s.repositories.EmailRepository.RemoveRelationship(ctx, entityType, common.GetTenantFromContext(ctx), entityId, email)
	return err == nil, err
}

func (s *emailService) DetachFromEntityById(ctx context.Context, entityType entity.EntityType, entityId, emailId string) (bool, error) {
	err := s.repositories.EmailRepository.RemoveRelationshipById(ctx, entityType, common.GetTenantFromContext(ctx), entityId, emailId)
	return err == nil, err
}

func (s *emailService) DeleteById(ctx context.Context, emailId string) (bool, error) {
	err := s.repositories.EmailRepository.DeleteById(ctx, common.GetTenantFromContext(ctx), emailId)
	return err == nil, err
}

func (s *emailService) UpsertInEventStore(ctx context.Context, size int) (int, int, error) {
	processedRecords := 0
	failedRecords := 0
	for size > 0 {
		batchSize := constants.Neo4jBatchSize
		if size < constants.Neo4jBatchSize {
			batchSize = size
		}
		records, err := s.repositories.EmailRepository.GetAllCrossTenants(ctx, batchSize)
		if err != nil {
			return 0, 0, err
		}
		for _, v := range records {
			_, err := s.grpcClients.EmailClient.UpsertEmail(context.Background(), &events_processing_email.UpsertEmailGrpcRequest{
				Id:            utils.GetStringPropOrEmpty(v.Node.Props, "id"),
				Tenant:        v.LinkedNodeId,
				RawEmail:      utils.GetStringPropOrEmpty(v.Node.Props, "rawEmail"),
				Source:        utils.GetStringPropOrEmpty(v.Node.Props, "source"),
				SourceOfTruth: utils.GetStringPropOrEmpty(v.Node.Props, "sourceOfTruth"),
				AppSource:     utils.GetStringPropOrEmpty(v.Node.Props, "appSource"),
				CreatedAt:     utils.ConvertTimeToTimestampPtr(utils.GetTimePropOrNil(v.Node.Props, "createdAt")),
				UpdatedAt:     utils.ConvertTimeToTimestampPtr(utils.GetTimePropOrNil(v.Node.Props, "updatedAt")),
			})
			if err != nil {
				failedRecords++
				logrus.Errorf("Failed to call method: %v", err)
			} else {
				processedRecords++
			}
		}

		size -= batchSize
	}

	return processedRecords, failedRecords, nil
}

func (s *emailService) mapDbNodeToEmailEntity(node dbtype.Node) *entity.EmailEntity {
	props := utils.GetPropsFromNode(node)
	result := entity.EmailEntity{
		Id:            utils.GetStringPropOrEmpty(props, "id"),
		Email:         utils.GetStringPropOrEmpty(props, "email"),
		RawEmail:      utils.GetStringPropOrEmpty(props, "rawEmail"),
		Validated:     utils.GetBoolPropOrFalse(props, "validated"),
		Source:        entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
		AppSource:     utils.GetStringPropOrEmpty(props, "appSource"),
		CreatedAt:     utils.GetTimePropOrEpochStart(props, "createdAt"),
		UpdatedAt:     utils.GetTimePropOrEpochStart(props, "updatedAt"),
	}
	return &result
}

func (s *emailService) addDbRelationshipToEmailEntity(relationship dbtype.Relationship, emailEntity *entity.EmailEntity) {
	props := utils.GetPropsFromRelationship(relationship)
	emailEntity.Primary = utils.GetBoolPropOrFalse(props, "primary")
	emailEntity.Label = utils.GetStringPropOrEmpty(props, "label")
}
