package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/opentracing/opentracing-go/log"
)

// OfferingCreate is the resolver for the offering_Create field.
func (r *mutationResolver) OfferingCreate(ctx context.Context, input *model.OfferingCreateInput) (*model.Offering, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OfferingCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	offeringId, err := r.Services.OfferingService.CreateOffering(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create offering")
		return &model.Offering{Metadata: &model.Metadata{
			ID: offeringId,
		}}, err
	}

	span.LogFields(log.String("response.offeringId", offeringId))
	return mapper.MapEntityToOffering(&neo4jentity.OfferingEntity{Id: offeringId}), nil
}

// OfferingUpdate is the resolver for the offering_Update field.
func (r *mutationResolver) OfferingUpdate(ctx context.Context, input *model.OfferingUpdateInput) (*model.Offering, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OfferingUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	if input.ID == "" {
		err := errors.New("missing offering id")
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Missing offering id")
		return nil, nil
	}

	err := r.Services.OfferingService.UpdateOffering(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update offering")
		return nil, err
	}

	return mapper.MapEntityToOffering(&neo4jentity.OfferingEntity{Id: input.ID}), nil
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *offeringResolver) ExternalLinks(ctx context.Context, obj *model.Offering) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// Offerings is the resolver for the offerings field.
func (r *queryResolver) Offerings(ctx context.Context) ([]*model.Offering, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Offerings", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	entities, err := r.Services.OfferingService.GetOfferings(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get offerings")
		return nil, err
	}

	return mapper.MapEntitiesToOfferings(entities), nil
}

// Offering returns generated.OfferingResolver implementation.
func (r *Resolver) Offering() generated.OfferingResolver { return &offeringResolver{r} }

type offeringResolver struct{ *Resolver }