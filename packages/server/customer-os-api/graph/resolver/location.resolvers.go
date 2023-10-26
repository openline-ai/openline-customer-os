package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// LocationRemoveFromContact is the resolver for the location_RemoveFromContact field.
func (r *mutationResolver) LocationRemoveFromContact(ctx context.Context, contactID string, locationID string) (*model.Contact, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LocationRemoveFromContactByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID), log.String("request.locationID", locationID))

	err := r.Services.LocationService.DetachFromEntity(ctx, entity.CONTACT, contactID, locationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not detach location %s from contact %s", locationID, contactID)
		return nil, nil
	}
	contactEntity, err := r.Services.ContactService.GetById(ctx, contactID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not get contact %s", contactID)
		return nil, nil
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// LocationRemoveFromOrganization is the resolver for the location_RemoveFromOrganization field.
func (r *mutationResolver) LocationRemoveFromOrganization(ctx context.Context, organizationID string, locationID string) (*model.Organization, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LocationRemoveFromOrganizationByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID), log.String("request.locationID", locationID))

	err := r.Services.LocationService.DetachFromEntity(ctx, entity.ORGANIZATION, organizationID, locationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not detach location %s from organization %s", locationID, organizationID)
		return nil, nil
	}
	organizationEntity, err := r.Services.OrganizationService.GetById(ctx, organizationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not get organization %s", organizationID)
		return nil, nil
	}
	return mapper.MapEntityToOrganization(organizationEntity), nil
}

// LocationUpdate is the resolver for the location_Update field.
func (r *mutationResolver) LocationUpdate(ctx context.Context, input model.LocationUpdateInput) (*model.Location, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LocationUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.locationID", input.ID))

	locationEntity, err := r.Services.LocationService.Update(ctx, *mapper.MapLocationUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update location")
		return nil, err
	}
	return mapper.MapEntityToLocation(locationEntity), nil
}
