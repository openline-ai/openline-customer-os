package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// InteractionSession is the resolver for the interactionSession field.
func (r *interactionEventResolver) InteractionSession(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionSession, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionSessionEntityNillable, err := dataloader.For(ctx).GetInteractionSessionForInteractionEvent(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get interaction session for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntityNillable), nil
}

// Meeting is the resolver for the meeting field.
func (r *interactionEventResolver) Meeting(ctx context.Context, obj *model.InteractionEvent) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: Meeting - meeting"))
}

// SentBy is the resolver for the sentBy field.
func (r *interactionEventResolver) SentBy(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetSentByParticipantsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEventParticipants(participantEntities), nil
}

// SentTo is the resolver for the sentTo field.
func (r *interactionEventResolver) SentTo(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetSentToParticipantsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEventParticipants(participantEntities), nil
}

// RepliesTo is the resolver for the repliesTo field.
func (r *interactionEventResolver) RepliesTo(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get ReplyTo for interaction event %s", obj.ID)
		return nil, err
	}
	if len(*interactionEventEntities) > 0 {
		return mapper.MapEntityToInteractionEvent(&(*interactionEventEntities)[0]), nil
	}
	return nil, nil
}

// Includes is the resolver for the includes field.
func (r *interactionEventResolver) Includes(ctx context.Context, obj *model.InteractionEvent) ([]*model.Attachment, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	entities, err := r.Services.AttachmentService.GetAttachmentsForNode(ctx, repository.INCLUDED_BY_INTERACTION_EVENT, []string{obj.ID})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get attachment entities for Interaction Event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// Events is the resolver for the events field.
func (r *interactionSessionResolver) Events(ctx context.Context, obj *model.InteractionSession) ([]*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForInteractionSession(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get interaction events for interaction session %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEvents(interactionEventEntities), nil
}

// AttendedBy is the resolver for the attendedBy field.
func (r *interactionSessionResolver) AttendedBy(ctx context.Context, obj *model.InteractionSession) ([]model.InteractionSessionParticipant, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetAttendedByParticipantsForInteractionSession(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionSessionParticipants(participantEntities), nil
}

// Includes is the resolver for the includes field.
func (r *interactionSessionResolver) Includes(ctx context.Context, obj *model.InteractionSession) ([]*model.Attachment, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	entities, err := r.Services.AttachmentService.GetAttachmentsForNode(ctx, repository.INCLUDED_BY_INTERACTION_SESSION, []string{obj.ID})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get attachment entities for Interaction Session %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// InteractionSessionCreate is the resolver for the interactionSession_Create field.
func (r *mutationResolver) InteractionSessionCreate(ctx context.Context, session model.InteractionSessionInput) (*model.InteractionSession, error) {
	interactionSessionEntity, err := r.Services.InteractionSessionService.Create(ctx,
		&service.InteractionSessionCreateData{
			InteractionSessionEntity: mapper.MapInteractionSessionInputToEntity(&session),
			AttendedBy:               service.MapInteractionSessionParticipantInputToAddressData(session.AttendedBy),
		})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create InteractionEvent")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToInteractionSession(interactionSessionEntity)
	return interactionEvent, nil
}

// InteractionSessionLinkAttachment is the resolver for the interactionSession_LinkAttachment field.
func (r *mutationResolver) InteractionSessionLinkAttachment(ctx context.Context, sessionID string, attachmentID string) (*model.InteractionSession, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	session, err := r.Services.InteractionSessionService.InteractionSessionLinkAttachment(ctx, sessionID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(session), nil
}

// InteractionEventCreate is the resolver for the interactionEvent_Create field.
func (r *mutationResolver) InteractionEventCreate(ctx context.Context, event model.InteractionEventInput) (*model.InteractionEvent, error) {
	interactionEventCreated, err := r.Services.InteractionEventService.Create(ctx, &service.InteractionEventCreateData{
		InteractionEventEntity: mapper.MapInteractionEventInputToEntity(&event),
		SessionIdentifier:      event.InteractionSession,
		SentBy:                 service.MapInteractionEventParticipantInputToAddressData(event.SentBy),
		SentTo:                 service.MapInteractionEventParticipantInputToAddressData(event.SentTo),
		RepliesTo:              event.RepliesTo,

		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create InteractionEvent")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToInteractionEvent(interactionEventCreated)
	return interactionEvent, nil
}

// InteractionEventLinkAttachment is the resolver for the interactionEvent_LinkAttachment field.
func (r *mutationResolver) InteractionEventLinkAttachment(ctx context.Context, eventID string, attachmentID string) (*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	event, err := r.Services.InteractionEventService.InteractionEventLinkAttachment(ctx, eventID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(event), nil
}

// InteractionSession is the resolver for the interactionSession field.
func (r *queryResolver) InteractionSession(ctx context.Context, id string) (*model.InteractionSession, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionSessionEntity, err := r.Services.InteractionSessionService.GetInteractionSessionById(ctx, id)
	if err != nil || interactionSessionEntity == nil {
		graphql.AddErrorf(ctx, "InteractionEvent with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntity), nil
}

// InteractionSessionBySessionIdentifier is the resolver for the interactionSession_BySessionIdentifier field.
func (r *queryResolver) InteractionSessionBySessionIdentifier(ctx context.Context, sessionIdentifier string) (*model.InteractionSession, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionSessionEntity, err := r.Services.InteractionSessionService.GetInteractionSessionBySessionIdentifier(ctx, sessionIdentifier)
	if err != nil || interactionSessionEntity == nil {
		graphql.AddErrorf(ctx, "InteractionEvent with identifier %s not found", sessionIdentifier)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntity), nil
}

// InteractionEvent is the resolver for the interactionEvent field.
func (r *queryResolver) InteractionEvent(ctx context.Context, id string) (*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionEventEntity, err := r.Services.InteractionEventService.GetInteractionEventById(ctx, id)
	if err != nil || interactionEventEntity == nil {
		graphql.AddErrorf(ctx, "InteractionEvent with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(interactionEventEntity), nil
}

// InteractionEventByEventIdentifier is the resolver for the interactionEvent_ByEventIdentifier field.
func (r *queryResolver) InteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionEventEntity, err := r.Services.InteractionEventService.GetInteractionEventByEventIdentifier(ctx, eventIdentifier)
	if err != nil || interactionEventEntity == nil {
		graphql.AddErrorf(ctx, "InteractionEvent with EventIdentifier %s not found", eventIdentifier)
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(interactionEventEntity), nil
}

// InteractionEvent returns generated.InteractionEventResolver implementation.
func (r *Resolver) InteractionEvent() generated.InteractionEventResolver {
	return &interactionEventResolver{r}
}

// InteractionSession returns generated.InteractionSessionResolver implementation.
func (r *Resolver) InteractionSession() generated.InteractionSessionResolver {
	return &interactionSessionResolver{r}
}

type interactionEventResolver struct{ *Resolver }
type interactionSessionResolver struct{ *Resolver }
