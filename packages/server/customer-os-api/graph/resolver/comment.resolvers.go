package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	opentracing "github.com/opentracing/opentracing-go"
)

// CreatedBy is the resolver for the createdBy field.
func (r *commentResolver) CreatedBy(ctx context.Context, obj *model.Comment) (*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntityNillable, err := dataloader.For(ctx).GetUserAuthorForComment(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("error fetching user author for comment %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "error fetching user author for comment %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToUser(userEntityNillable), nil
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *commentResolver) ExternalLinks(ctx context.Context, obj *model.Comment) ([]*model.ExternalSystem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetExternalSystemsForComment(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("failed to get external system for comment %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "failed to get external system for comment %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToExternalSystems(entities), nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
