package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// TenantMerge is the resolver for the tenant_Merge field.
func (r *mutationResolver) TenantMerge(ctx context.Context, tenant model.TenantInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TenantMerge", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	newTenant, err := r.Services.TenantService.Merge(ctx, mapper.MapTenantInputToEntity(tenant))
	if err != nil {
		return "", fmt.Errorf("TenantMerge: %w", err)
	}
	return newTenant.Name, nil
}

// TenantAddBillingProfile is the resolver for the tenant_AddBillingProfile field.
func (r *mutationResolver) TenantAddBillingProfile(ctx context.Context, input model.TenantBillingProfileInput) (*model.TenantBillingProfile, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TenantAddBillingProfile", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	profileId, err := r.Services.TenantService.CreateTenantBillingProfile(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create tenant billing profile")
		return &model.TenantBillingProfile{ID: profileId}, err
	}

	createdTenantBillingProfileEntity, err := r.Services.TenantService.GetTenantBillingProfile(ctx, profileId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Tenant billing profile not yet available.")
		return &model.TenantBillingProfile{ID: profileId}, nil
	}
	span.LogFields(log.String("response.tenantBillingProfileId", profileId))
	return mapper.MapEntityToTenantBillingProfile(createdTenantBillingProfileEntity), nil
}

// TenantUpdateBillingProfile is the resolver for the tenant_UpdateBillingProfile field.
func (r *mutationResolver) TenantUpdateBillingProfile(ctx context.Context, input model.TenantBillingProfileUpdateInput) (*model.TenantBillingProfile, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TenantUpdateBillingProfile", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	if input.ID == "" {
		err := errors.New("missing tenant billing profile id")
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Missing tenant billing profile id")
		return nil, nil
	}

	err := r.Services.TenantService.UpdateTenantBillingProfile(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update master plan")
		return nil, err
	}

	updatedTenantBillingProfileEntity, err := r.Services.TenantService.GetTenantBillingProfile(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch tenant billing profile details")
		return nil, nil
	}
	return mapper.MapEntityToTenantBillingProfile(updatedTenantBillingProfileEntity), nil
}

// TenantUpdateSettings is the resolver for the tenant_UpdateSettings field.
func (r *mutationResolver) TenantUpdateSettings(ctx context.Context, input *model.TenantSettingsInput) (*model.TenantSettings, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TenantUpdateSettings", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.TenantService.UpdateTenantSettings(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update tenant settings")
		return nil, err
	}

	updatedTenantSettingsEntity, err := r.Services.TenantService.GetTenantSettings(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch tenant settings")
		return nil, nil
	}
	return mapper.MapEntityToTenantSettings(updatedTenantSettingsEntity), nil
}

// Tenant is the resolver for the tenant field.
func (r *queryResolver) Tenant(ctx context.Context) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Tenant", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	return common.GetTenantFromContext(ctx), nil
}

// TenantByWorkspace is the resolver for the tenant_ByWorkspace field.
func (r *queryResolver) TenantByWorkspace(ctx context.Context, workspace model.WorkspaceInput) (*string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TenantByWorkspace", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.workspace", workspace.Name))

	tenant, err := r.Services.TenantService.GetTenantForWorkspace(ctx, mapper.MapWorkspaceInputToEntity(workspace))
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, nil
	}
	return &tenant.Name, nil
}

// TenantByEmail is the resolver for the tenant_ByEmail field.
func (r *queryResolver) TenantByEmail(ctx context.Context, email string) (*string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TenantByEmail", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.email", email))

	tenant, err := r.Services.TenantService.GetTenantForUserEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, nil
	}
	return &tenant.Name, nil
}

// TenantBillingProfiles is the resolver for the tenantBillingProfiles field.
func (r *queryResolver) TenantBillingProfiles(ctx context.Context) ([]*model.TenantBillingProfile, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TenantBillingProfiles", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	tenantBillingProfileEntities, err := r.Services.TenantService.GetTenantBillingProfiles(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch billing profiles")
		return nil, nil
	}
	return mapper.MapEntitiesToTenantBillingProfiles(tenantBillingProfileEntities), nil
}

// TenantBillingProfile is the resolver for the tenantBillingProfile field.
func (r *queryResolver) TenantBillingProfile(ctx context.Context, id string) (*model.TenantBillingProfile, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TenantBillingProfile", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.tenantBillingProfileId", id))

	tenantBillingProfileEntity, err := r.Services.TenantService.GetTenantBillingProfile(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch tenant billing profile details")
		return nil, nil
	}
	return mapper.MapEntityToTenantBillingProfile(tenantBillingProfileEntity), nil
}

// TenantSettings is the resolver for the tenantSettings field.
func (r *queryResolver) TenantSettings(ctx context.Context) (*model.TenantSettings, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TenantSettings", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	tenantSettingsEntity, err := r.Services.TenantService.GetTenantSettings(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch tenant settings")
		return nil, nil
	}
	return mapper.MapEntityToTenantSettings(tenantSettingsEntity), nil
}
