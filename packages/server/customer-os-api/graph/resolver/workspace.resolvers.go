package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
)

// WorkspaceMergeToTenant is the resolver for the workspace_MergeToTenant field.
func (r *mutationResolver) WorkspaceMergeToTenant(ctx context.Context, workspace model.WorkspaceInput, tenant string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.WorkspaceMergeToTenant", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	result, err := r.Services.WorkspaceService.MergeToTenant(ctx, mapper.MapWorkspaceInputToEntity(workspace), tenant)
	if err != nil {
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// WorkspaceMerge is the resolver for the workspace_Merge field.
func (r *mutationResolver) WorkspaceMerge(ctx context.Context, workspace model.WorkspaceInput) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.WorkspaceMerge", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	result, err := r.Services.WorkspaceService.MergeToTenant(ctx, mapper.MapWorkspaceInputToEntity(workspace), common.GetContext(ctx).Tenant)
	if err != nil {
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}
