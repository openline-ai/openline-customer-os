package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
)

// ExternalSystemInstances is the resolver for the externalSystemInstances field.
func (r *queryResolver) ExternalSystemInstances(ctx context.Context) ([]*model.ExternalSystemInstance, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.ExternalSystemInstances", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	externalSystemEntities, err := r.Services.ExternalSystemService.GetAllExternalSystemInstances(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get external system instances")
		return nil, err
	}
	return mapper.MapExternalSystemEntitiesToExternalSystemInstances(externalSystemEntities), nil
}
