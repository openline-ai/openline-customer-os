package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
)

// FIXME alexb add test
// OrganizationTypeCreate is the resolver for the organizationType_Create field.
func (r *mutationResolver) OrganizationTypeCreate(ctx context.Context, input model.OrganizationTypeInput) (*model.OrganizationType, error) {
	createdOrganizationType, err := r.Services.OrganizationTypeService.Create(ctx, mapper.MapOrganizationTypeInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create organization type %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToOrganizationType(createdOrganizationType), nil
}

// FIXME alexb add test
// OrganizationTypeUpdate is the resolver for the organizationType_Update field.
func (r *mutationResolver) OrganizationTypeUpdate(ctx context.Context, input model.OrganizationTypeUpdateInput) (*model.OrganizationType, error) {
	updatedOrganizationType, err := r.Services.OrganizationTypeService.Update(ctx, mapper.MapOrganizationTypeUpdateInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update organization type %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToOrganizationType(updatedOrganizationType), nil
}

// FIXME alexb add test
// OrganizationTypeDelete is the resolver for the organizationType_Delete field.
func (r *mutationResolver) OrganizationTypeDelete(ctx context.Context, id string) (*model.Result, error) {
	result, err := r.Services.OrganizationTypeService.Delete(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to delete organization type %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// FIXME alexb add test
// OrganizationTypes is the resolver for the organizationTypes field.
func (r *queryResolver) OrganizationTypes(ctx context.Context) ([]*model.OrganizationType, error) {
	organizationTypes, err := r.Services.OrganizationTypeService.GetAll(ctx)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch organization types")
		return nil, err
	}
	return mapper.MapEntitiesToOrganizationTypes(organizationTypes), err
}
