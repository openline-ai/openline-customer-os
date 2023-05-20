package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"time"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// TenantMerge is the resolver for the tenant_Merge field.
func (r *mutationResolver) TenantMerge(ctx context.Context, tenant model.TenantInput) (string, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	newTenant, err := r.Services.TenantService.Merge(ctx, mapper.MapTenantInputToEntity(tenant))
	if err != nil {
		return "", fmt.Errorf("TenantMerge: %w", err)
	}
	return newTenant.Name, nil
}

// Tenant is the resolver for the tenant field.
func (r *queryResolver) Tenant(ctx context.Context) (string, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	return common.GetTenantFromContext(ctx), nil
}

// TenantByWorkspace is the resolver for the tenant_ByWorkspace field.
func (r *queryResolver) TenantByWorkspace(ctx context.Context, workspace model.WorkspaceInput) (*string, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	tenant, err := r.Services.TenantService.GetTenantForWorkspace(ctx, mapper.MapWorkspaceInputToEntity(workspace))
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, nil
	}
	return &tenant.Name, nil
}
