package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// PlayerMerge is the resolver for the player_Merge field.
func (r *mutationResolver) PlayerMerge(ctx context.Context, input model.PlayerInput) (*model.Player, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PlayerMerge", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	playerEntity, err := r.Services.PlayerService.Merge(ctx, mapper.MapPlayerInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to merge player")
		return nil, err
	}
	return mapper.MapEntityToPlayer(playerEntity), nil
}

// PlayerUpdate is the resolver for the player_Update field.
func (r *mutationResolver) PlayerUpdate(ctx context.Context, id string, update model.PlayerUpdate) (*model.Player, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PlayerUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.playerID", id))

	playerEntity, err := r.Services.PlayerService.Update(ctx, mapper.MapPlayerUpdateToEntity(id, &update))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update player %s", id)
		return nil, err
	}
	return mapper.MapEntityToPlayer(playerEntity), nil
}

// PlayerSetDefaultUser is the resolver for the player_SetDefaultUser field.
func (r *mutationResolver) PlayerSetDefaultUser(ctx context.Context, id string, userID string) (*model.Player, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PlayerSetDefaultUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.playerID", id), log.String("request.userID", userID))

	playerEntity, err := r.Services.PlayerService.SetDefaultUser(ctx, id, userID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to set default user for player %s", id)
		return nil, err
	}
	return mapper.MapEntityToPlayer(playerEntity), nil
}

// Users is the resolver for the users field.
func (r *playerResolver) Users(ctx context.Context, obj *model.Player) ([]*model.PlayerUser, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "PlayerResolver.Users", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.playerID", obj.ID))

	userEntities, err := dataloader.For(ctx).GetUsersForPlayer(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get users for player %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToPlayerUsers(userEntities), nil
}

// PlayerByAuthIDProvider is the resolver for the player_ByAuthIdProvider field.
func (r *queryResolver) PlayerByAuthIDProvider(ctx context.Context, authID string, provider string) (*model.Player, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "PlayerResolver.PlayerByAuthIDProvider", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.authID", authID), log.String("request.provider", provider))

	playerEntity, err := r.Services.PlayerService.GetPlayerByAuthIdProvider(ctx, authID, provider)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get player by authId and provider")
		return nil, err
	}
	return mapper.MapEntityToPlayer(playerEntity), nil
}

// PlayerGetUsers is the resolver for the player_GetUsers field.
func (r *queryResolver) PlayerGetUsers(ctx context.Context) ([]*model.PlayerUser, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.PlayerGetUsers", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	dbUsers, err := r.Services.PlayerService.GetUsers(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get users")
		return nil, err
	}
	return mapper.MapEntitiesToPlayerUsers(dbUsers), nil
}

// Player returns generated.PlayerResolver implementation.
func (r *Resolver) Player() generated.PlayerResolver { return &playerResolver{r} }

type playerResolver struct{ *Resolver }
