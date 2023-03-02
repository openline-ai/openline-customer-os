package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// Organization is the resolver for the organization field.
func (r *jobRoleResolver) Organization(ctx context.Context, obj *model.JobRole) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	organizationEntity, err := r.Services.OrganizationService.GetOrganizationForJobRole(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization for job role %s", obj.ID)
		return nil, err
	}
	if organizationEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToOrganization(organizationEntity), nil
}

// Contact is the resolver for the contact field.
func (r *jobRoleResolver) Contact(ctx context.Context, obj *model.JobRole) (*model.Contact, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	contactEntity, err := r.Services.ContactService.GetContactForRole(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact job for role %s", obj.ID)
		return nil, err
	}
	if contactEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// JobRoleDelete is the resolver for the jobRole_Delete field.
func (r *mutationResolver) JobRoleDelete(ctx context.Context, contactID string, roleID string) (*model.Result, error) {
	result, err := r.Services.JobRoleService.DeleteJobRole(ctx, contactID, roleID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed remove job role %s from contact %s", roleID, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// JobRoleCreate is the resolver for the jobRole_Create field.
func (r *mutationResolver) JobRoleCreate(ctx context.Context, contactID string, input model.JobRoleInput) (*model.JobRole, error) {
	result, err := r.Services.JobRoleService.CreateJobRole(ctx, contactID, input.OrganizationID, mapper.MapJobRoleInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed add job role to contact %s", contactID)
		return nil, err
	}
	return mapper.MapEntityToJobRole(result), nil
}

// JobRoleUpdate is the resolver for the jobRole_Update field.
func (r *mutationResolver) JobRoleUpdate(ctx context.Context, contactID string, input model.JobRoleUpdateInput) (*model.JobRole, error) {
	result, err := r.Services.JobRoleService.UpdateJobRole(ctx, contactID, input.OrganizationID, mapper.MapJobRoleUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed update role %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToJobRole(result), nil
}

// JobRole returns generated.JobRoleResolver implementation.
func (r *Resolver) JobRole() generated.JobRoleResolver { return &jobRoleResolver{r} }

type jobRoleResolver struct{ *Resolver }
