package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// SocialUpdate is the resolver for the social_Update field.
func (r *mutationResolver) SocialUpdate(ctx context.Context, input model.SocialUpdateInput) (*model.Social, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.SocialUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.socialID", input.ID))

	socialEntity, err := r.Services.SocialService.Update(ctx, *mapper.MapSocialUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update social")
		return nil, err
	}
	return mapper.MapEntityToSocial(socialEntity), nil
}

// SocialRemove is the resolver for the social_Remove field.
func (r *mutationResolver) SocialRemove(ctx context.Context, socialID string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.SocialRemove", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.socialID", socialID))

	err := r.Services.SocialService.Remove(ctx, socialID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to remove social")
		return &model.Result{Result: false}, nil
	}
	return &model.Result{Result: true}, nil
}
