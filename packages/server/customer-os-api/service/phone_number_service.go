package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

type PhoneNumberService interface {
	FindAllForContact(ctx context.Context, contactId string) (*entity.PhoneNumberEntities, error)
	MergePhoneNumberToContact(ctx context.Context, id string, toEntity *entity.PhoneNumberEntity) (*entity.PhoneNumberEntity, error)
	UpdatePhoneNumberInContact(ctx context.Context, id string, toEntity *entity.PhoneNumberEntity) (*entity.PhoneNumberEntity, error)
	Delete(ctx context.Context, contactId string, e164 string) (bool, error)
	DeleteById(ctx context.Context, contactId string, phoneId string) (bool, error)
}

type phoneNumberService struct {
	repositories *repository.Repositories
}

func NewPhoneNumberService(repositories *repository.Repositories) PhoneNumberService {
	return &phoneNumberService{
		repositories: repositories,
	}
}

func (s *phoneNumberService) getDriver() neo4j.Driver {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *phoneNumberService) FindAllForContact(ctx context.Context, contactId string) (*entity.PhoneNumberEntities, error) {
	session := utils.NewNeo4jReadSession(s.getDriver())
	defer session.Close()

	queryResult, err := s.repositories.PhoneNumberRepository.FindAllForContact(session, common.GetContext(ctx).Tenant, contactId)
	if err != nil {
		return nil, err
	}

	phoneNumberEntities := entity.PhoneNumberEntities{}

	for _, dbRecord := range queryResult.([]*db.Record) {
		phoneNumberEntity := s.mapDbNodeToPhoneNumberEntity(dbRecord.Values[0].(dbtype.Node))
		s.addDbRelationshipToPhoneNumberEntity(dbRecord.Values[1].(dbtype.Relationship), phoneNumberEntity)
		phoneNumberEntities = append(phoneNumberEntities, *phoneNumberEntity)
	}

	return &phoneNumberEntities, nil
}

func (s *phoneNumberService) MergePhoneNumberToContact(ctx context.Context, contactId string, entity *entity.PhoneNumberEntity) (*entity.PhoneNumberEntity, error) {
	session := utils.NewNeo4jWriteSession(s.getDriver())
	defer session.Close()

	var err error
	var phoneNumberNode *dbtype.Node
	var phoneNumberRelationship *dbtype.Relationship

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if entity.Primary == true {
			err := setOtherContactPhoneNumbersNonPrimaryInTx(ctx, contactId, entity.E164, tx)
			if err != nil {
				return nil, err
			}
		}
		phoneNumberNode, phoneNumberRelationship, err = s.repositories.PhoneNumberRepository.MergePhoneNumberToContactInTx(tx, common.GetContext(ctx).Tenant, contactId, *entity)
		return nil, err
	})
	if err != nil {
		return nil, err
	}

	var phoneNumberEntity = s.mapDbNodeToPhoneNumberEntity(*phoneNumberNode)
	s.addDbRelationshipToPhoneNumberEntity(*phoneNumberRelationship, phoneNumberEntity)
	return phoneNumberEntity, nil
}

func (s *phoneNumberService) UpdatePhoneNumberInContact(ctx context.Context, contactId string, entity *entity.PhoneNumberEntity) (*entity.PhoneNumberEntity, error) {
	session := utils.NewNeo4jWriteSession(s.getDriver())
	defer session.Close()

	var err error
	var phoneNumberNode *dbtype.Node
	var phoneNumberRelationship *dbtype.Relationship

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		phoneNumberNode, phoneNumberRelationship, err = s.repositories.PhoneNumberRepository.UpdatePhoneNumberByContactInTx(tx, common.GetContext(ctx).Tenant, contactId, *entity)

		if err != nil {
			return nil, err
		}
		if entity.Primary == true {
			err := setOtherContactPhoneNumbersNonPrimaryInTx(ctx, contactId, entity.E164, tx)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	var phoneNumberEntity = s.mapDbNodeToPhoneNumberEntity(*phoneNumberNode)
	s.addDbRelationshipToPhoneNumberEntity(*phoneNumberRelationship, phoneNumberEntity)
	return phoneNumberEntity, nil
}

func setOtherContactPhoneNumbersNonPrimaryInTx(ctx context.Context, contactId string, e164 string, tx neo4j.Transaction) error {
	_, err := tx.Run(`
			MATCH (c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}),
				 (c)-[r:CALLED_AT]->(p:PhoneNumber)
			WHERE p.e164 <> $e164
            SET r.primary=false`,
		map[string]interface{}{
			"tenant":    common.GetContext(ctx).Tenant,
			"contactId": contactId,
			"e164":      e164,
		})
	return err
}

func (s *phoneNumberService) Delete(ctx context.Context, contactId string, e164 string) (bool, error) {
	session := utils.NewNeo4jWriteSession(s.getDriver())
	defer session.Close()

	queryResult, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`
			MATCH (c:Contact {id:$id})-[:CONTACT_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}),
                  (c:Contact {id:$id})-[:CALLED_AT]->(p:PhoneNumber {e164:$e164})
            DETACH DELETE p
			`,
			map[string]interface{}{
				"id":     contactId,
				"e164":   e164,
				"tenant": common.GetContext(ctx).Tenant,
			})

		return true, err
	})
	if err != nil {
		return false, err
	}

	return queryResult.(bool), nil
}

func (s *phoneNumberService) DeleteById(ctx context.Context, contactId string, phoneId string) (bool, error) {
	session := utils.NewNeo4jWriteSession(s.getDriver())
	defer session.Close()

	queryResult, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`
			MATCH (c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}),
                  (c:Contact {id:$contactId})-[:CALLED_AT]->(p:PhoneNumber {id:$phoneId})
            DETACH DELETE p
			`,
			map[string]interface{}{
				"contactId": contactId,
				"phoneId":   phoneId,
				"tenant":    common.GetContext(ctx).Tenant,
			})

		return true, err
	})
	if err != nil {
		return false, err
	}

	return queryResult.(bool), nil
}

func (s *phoneNumberService) mapDbNodeToPhoneNumberEntity(node dbtype.Node) *entity.PhoneNumberEntity {
	props := utils.GetPropsFromNode(node)
	result := entity.PhoneNumberEntity{
		Id:            utils.GetStringPropOrEmpty(props, "id"),
		E164:          utils.GetStringPropOrEmpty(props, "e164"),
		Label:         utils.GetStringPropOrEmpty(props, "label"),
		Source:        entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
	}
	return &result
}

func (s *phoneNumberService) addDbRelationshipToPhoneNumberEntity(relationship dbtype.Relationship, phoneNumberEntity *entity.PhoneNumberEntity) {
	props := utils.GetPropsFromRelationship(relationship)
	phoneNumberEntity.Primary = utils.GetBoolPropOrFalse(props, "primary")
}
