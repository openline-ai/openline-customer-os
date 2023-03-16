package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"
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

// CreateInteractionSession is the resolver for the createInteractionSession field.
func (r *mutationResolver) CreateInteractionSession(ctx context.Context, sessionIdentifier *string, name *string, status *string, typeArg *string, channel *string, source model.DataSource, sourceOfTruth model.DataSource, appSource string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: CreateInteractionSession - createInteractionSession"))
}

// CreateInteractionEvent is the resolver for the createInteractionEvent field.
func (r *mutationResolver) CreateInteractionEvent(ctx context.Context, eventIdentifier *string, content *string, contentType *string, channel *string, interactionSession string, sentBy []*model.InteractionEventParticipantInput, sentTo []*model.InteractionEventParticipantInput, source model.DataSource, sourceOfTruth model.DataSource, appSource string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: CreateInteractionEvent - createInteractionEvent"))
}

// InteractionSession is the resolver for the interactionSession field.
func (r *queryResolver) InteractionSession(ctx context.Context, id string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSession - interactionSession"))
}

// InteractionSessionByEventIdentifier is the resolver for the interactionSession_ByEventIdentifier field.
func (r *queryResolver) InteractionSessionByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionByEventIdentifier - interactionSession_ByEventIdentifier"))
}

// InteractionEvent is the resolver for the interactionEvent field.
func (r *queryResolver) InteractionEvent(ctx context.Context, id string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEvent - interactionEvent"))
}

// InteractionEventByEventIdentifier is the resolver for the interactionEvent_ByEventIdentifier field.
func (r *queryResolver) InteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEventByEventIdentifier - interactionEvent_ByEventIdentifier"))
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
