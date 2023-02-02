package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
)

// OrganizationCreate is the resolver for the organization_Create field.
func (r *mutationResolver) OrganizationCreate(ctx context.Context, input model.OrganizationInput) (*model.Organization, error) {
	createdOrganizationEntity, err := r.Services.OrganizationService.Create(ctx,
		&service.OrganizationCreateData{
			OrganizationEntity: mapper.MapOrganizationInputToEntity(&input),
			OrganizationTypeID: input.OrganizationTypeID,
		})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create organization %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToOrganization(createdOrganizationEntity), nil
}

// OrganizationUpdate is the resolver for the organization_Update field.
func (r *mutationResolver) OrganizationUpdate(ctx context.Context, input model.OrganizationUpdateInput) (*model.Organization, error) {
	organization := mapper.MapOrganizationUpdateInputToEntity(&input)

	updatedOrganizationEntity, err := r.Services.OrganizationService.Update(ctx,
		&service.OrganizationUpdateData{
			OrganizationEntity: organization,
			OrganizationTypeID: input.OrganizationTypeID,
		})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update organization %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(updatedOrganizationEntity), nil
}

// OrganizationDelete is the resolver for the organization_Delete field.
func (r *mutationResolver) OrganizationDelete(ctx context.Context, id string) (*model.Result, error) {
	result, err := r.Services.OrganizationService.PermanentDelete(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to delete organization %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// OrganizationType is the resolver for the organizationType field.
func (r *organizationResolver) OrganizationType(ctx context.Context, obj *model.Organization) (*model.OrganizationType, error) {
	organizationTypeEntity, err := r.Services.OrganizationTypeService.FindOrganizationTypeForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization type for organization %s", obj.ID)
		return nil, err
	}
	if organizationTypeEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToOrganizationType(organizationTypeEntity), nil
}

// Addresses is the resolver for the addresses field.
func (r *organizationResolver) Addresses(ctx context.Context, obj *model.Organization) ([]*model.Place, error) {
	addressEntities, err := r.Services.PlaceService.FindAllForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get addresses for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToPlaces(addressEntities), err
}

// Contacts is the resolver for the contacts field.
func (r *organizationResolver) Contacts(ctx context.Context, obj *model.Organization, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ContactService.GetContactsForOrganization(ctx, obj.ID, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not fetch contacts for organization %s", obj.ID)
		return nil, err
	}
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// ContactRoles is the resolver for the contactRoles field.
func (r *organizationResolver) ContactRoles(ctx context.Context, obj *model.Organization) ([]*model.ContactRole, error) {
	contactRoleEntities, err := r.Services.ContactRoleService.FindAllForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get roles for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToContactRoles(contactRoleEntities), err
}

// Notes is the resolver for the notes field.
func (r *organizationResolver) Notes(ctx context.Context, obj *model.Organization, pagination *model.Pagination) (*model.NotePage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.NoteService.GetNotesForOrganization(ctx, obj.ID, pagination.Page, pagination.Limit)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization %s notes", obj.ID)
		return nil, err
	}
	return &model.NotePage{
		Content:       mapper.MapEntitiesToNotes(paginatedResult.Rows.(*entity.NoteEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.OrganizationService.FindAll(ctx, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not fetch organizations")
		return nil, err
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	organizationEntityPtr, err := r.Services.OrganizationService.GetOrganizationById(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization by id %s", id)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntityPtr), nil
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

type organizationResolver struct{ *Resolver }
