package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
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
