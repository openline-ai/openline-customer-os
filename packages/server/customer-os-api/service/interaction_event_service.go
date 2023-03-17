package service

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"golang.org/x/exp/slices"
	"golang.org/x/net/context"
)

type InteractionEventService interface {
	GetInteractionEventsForInteractionSessions(ctx context.Context, ids []string) (*entity.InteractionEventEntities, error)
	GetSentByParticipantsForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventParticipants, error)
	GetSentToParticipantsForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventParticipants, error)
	GetInteractionEventById(ctx context.Context, id string) (*entity.InteractionEventEntity, error)
	GetInteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*entity.InteractionEventEntity, error)
	GetReplyToInteractionsEventForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventEntities, error)

	mapDbNodeToInteractionEventEntity(node dbtype.Node) *entity.InteractionEventEntity
}

type interactionEventService struct {
	repositories *repository.Repositories
	services     *Services
}

func NewInteractionEventService(repositories *repository.Repositories, services *Services) InteractionEventService {
	return &interactionEventService{
		repositories: repositories,
		services:     services,
	}
}

func (s *interactionEventService) GetInteractionEventsForInteractionSessions(ctx context.Context, ids []string) (*entity.InteractionEventEntities, error) {
	interactionEvents, err := s.repositories.InteractionEventRepository.GetAllForInteractionSessions(ctx, common.GetTenantFromContext(ctx), ids)
	if err != nil {
		return nil, err
	}
	interactionEventEntities := entity.InteractionEventEntities{}
	for _, v := range interactionEvents {
		interactionEventEntity := s.mapDbNodeToInteractionEventEntity(*v.Node)
		interactionEventEntity.DataloaderKey = v.LinkedNodeId
		interactionEventEntities = append(interactionEventEntities, *interactionEventEntity)
	}
	return &interactionEventEntities, nil
}

func (s *interactionEventService) GetSentByParticipantsForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventParticipants, error) {
	records, err := s.repositories.InteractionEventRepository.GetSentByParticipantsForInteractionEvents(ctx, common.GetTenantFromContext(ctx), ids)
	if err != nil {
		return nil, err
	}

	interactionEventParticipants := s.convertDbNodesToInteractionEventParticipants(records)

	return &interactionEventParticipants, nil
}

func (s *interactionEventService) GetSentToParticipantsForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventParticipants, error) {
	records, err := s.repositories.InteractionEventRepository.GetSentToParticipantsForInteractionEvents(ctx, common.GetTenantFromContext(ctx), ids)
	if err != nil {
		return nil, err
	}

	interactionEventParticipants := s.convertDbNodesToInteractionEventParticipants(records)

	return &interactionEventParticipants, nil
}

func (s *interactionEventService) GetReplyToInteractionsEventForInteractionEvents(ctx context.Context, ids []string) (*entity.InteractionEventEntities, error) {
	records, err := s.repositories.InteractionEventRepository.GetReplyToInteractionEventsForInteractionEvents(ctx, common.GetTenantFromContext(ctx), ids)
	if err != nil {
		return nil, err
	}

	interactionEvents := s.convertDbNodesToInteractionEvent(records)

	return &interactionEvents, nil
}

func (s *interactionEventService) GetInteractionEventById(ctx context.Context, id string) (*entity.InteractionEventEntity, error) {
	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (e:InteractionEvent_%s {id:$id}) RETURN e`,
			common.GetTenantFromContext(ctx)),
			map[string]interface{}{
				"id": id,
			})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	return s.mapDbNodeToInteractionEventEntity(queryResult.(dbtype.Node)), nil
}

func (s *interactionEventService) GetInteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*entity.InteractionEventEntity, error) {
	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (e:InteractionEvent_%s {identifier:$identifier}) RETURN e`,
			common.GetTenantFromContext(ctx)),
			map[string]interface{}{
				"identifier": eventIdentifier,
			})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	return s.mapDbNodeToInteractionEventEntity(queryResult.(dbtype.Node)), nil
}

func (s *interactionEventService) mapDbNodeToInteractionEventEntity(node dbtype.Node) *entity.InteractionEventEntity {
	props := utils.GetPropsFromNode(node)
	interactionEventEntity := entity.InteractionEventEntity{
		Id:              utils.GetStringPropOrEmpty(props, "id"),
		CreatedAt:       utils.GetTimePropOrEpochStart(props, "createdAt"),
		EventIdentifier: utils.GetStringPropOrEmpty(props, "identifier"),
		Channel:         utils.GetStringPropOrEmpty(props, "channel"),
		Content:         utils.GetStringPropOrEmpty(props, "content"),
		ContentType:     utils.GetStringPropOrEmpty(props, "contentType"),
		AppSource:       utils.GetStringPropOrEmpty(props, "appSource"),
		Source:          entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth:   entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
	}
	return &interactionEventEntity
}

func (s *interactionEventService) convertDbNodesToInteractionEvent(records []*utils.DbNodeAndId) entity.InteractionEventEntities {
	interactionEvents := entity.InteractionEventEntities{}
	for _, v := range records {
		event := s.mapDbNodeToInteractionEventEntity(*v.Node)
		event.DataloaderKey = v.LinkedNodeId
		interactionEvents = append(interactionEvents, *event)

	}
	return interactionEvents
}

func (s *interactionEventService) convertDbNodesToInteractionEventParticipants(records []*utils.DbNodeWithRelationAndId) entity.InteractionEventParticipants {
	interactionEventParticipants := entity.InteractionEventParticipants{}
	for _, v := range records {
		if slices.Contains(v.Node.Labels, entity.NodeLabel_Email) {
			participant := s.services.EmailService.mapDbNodeToEmailEntity(*v.Node)
			participant.InteractionEventParticipantDetails = s.mapDbRelationshipToParticipantDetails(*v.Relationship)
			participant.DataloaderKey = v.LinkedNodeId
			interactionEventParticipants = append(interactionEventParticipants, participant)
		} else if slices.Contains(v.Node.Labels, entity.NodeLabel_PhoneNumber) {
			participant := s.services.PhoneNumberService.mapDbNodeToPhoneNumberEntity(*v.Node)
			participant.InteractionEventParticipantDetails = s.mapDbRelationshipToParticipantDetails(*v.Relationship)
			participant.DataloaderKey = v.LinkedNodeId
			interactionEventParticipants = append(interactionEventParticipants, participant)
		} else if slices.Contains(v.Node.Labels, entity.NodeLabel_User) {
			participant := s.services.UserService.mapDbNodeToUserEntity(*v.Node)
			participant.InteractionEventParticipantDetails = s.mapDbRelationshipToParticipantDetails(*v.Relationship)
			participant.DataloaderKey = v.LinkedNodeId
			interactionEventParticipants = append(interactionEventParticipants, participant)
		} else if slices.Contains(v.Node.Labels, entity.NodeLabel_Contact) {
			participant := s.services.ContactService.mapDbNodeToContactEntity(*v.Node)
			participant.InteractionEventParticipantDetails = s.mapDbRelationshipToParticipantDetails(*v.Relationship)
			participant.DataloaderKey = v.LinkedNodeId
			interactionEventParticipants = append(interactionEventParticipants, participant)
		}
	}
	return interactionEventParticipants
}

func (s *interactionEventService) mapDbRelationshipToParticipantDetails(relationship dbtype.Relationship) entity.InteractionEventParticipantDetails {
	props := utils.GetPropsFromRelationship(relationship)
	details := entity.InteractionEventParticipantDetails{
		Type: utils.GetStringPropOrEmpty(props, "type"),
	}
	return details
}

func (s *interactionEventService) getNeo4jDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}
