package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
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

// InteractionSessionCreate is the resolver for the interactionSession_Create field.
func (r *mutationResolver) InteractionSessionCreate(ctx context.Context, session model.InteractionSessionInput) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionCreate - interactionSession_Create"))
}

// InteractionEventCreate is the resolver for the interactionEvent_Create field.
func (r *mutationResolver) InteractionEventCreate(ctx context.Context, event model.InteractionEventInput) (*model.InteractionEvent, error) {

	interactionEventCreated, err := r.Services.InteractionEventService.Create(ctx, &service.InteractionEventCreateData{
		InteractionEventEntity: mapper.MapInteractionEventInputToEntity(&event),
		Content:                event.Content,
		ContentType:            event.ContentType,
		SessionIdentifier:      event.InteractionSession,
		SentBy:                 mapper.MapInteractionEventParticipantInputToAddressData(event.SentBy),
		SentTo:                 mapper.MapInteractionEventParticipantInputToAddressData(event.SentBy),

		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create InteractionEvent")
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(interactionEventCreated), nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateInteractionSession(ctx context.Context, sessionIdentifier *string, name *string, status *string, typeArg *string, channel *string, source model.DataSource, sourceOfTruth model.DataSource, appSource string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: CreateInteractionSession - createInteractionSession"))
}
func (r *mutationResolver) CreateInteractionEvent(ctx context.Context, eventIdentifier *string, content *string, contentType *string, channel *string, interactionSession string, sentBy []*model.InteractionEventParticipantInput, sentTo []*model.InteractionEventParticipantInput, source model.DataSource, sourceOfTruth model.DataSource, appSource string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: CreateInteractionEvent - createInteractionEvent"))
}
func (r *queryResolver) InteractionSessionByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionByEventIdentifier - interactionSession_ByEventIdentifier"))
}
