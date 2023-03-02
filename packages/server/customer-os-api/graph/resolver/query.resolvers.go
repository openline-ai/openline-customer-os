package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// EntityTemplates is the resolver for the entityTemplates field.
func (r *queryResolver) EntityTemplates(ctx context.Context, extends *model.EntityTemplateExtension) ([]*model.EntityTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.EntityTemplateService.FindAll(ctx, utils.StringPtr(extends.String()))
	return mapper.MapEntitiesToEntityTemplates(result), err
}

// DashboardView is the resolver for the dashboardView field.
func (r *queryResolver) DashboardView(ctx context.Context, pagination model.Pagination, searchTerm *string) (*model.DashboardViewItemPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	paginatedResult, err := r.Services.QueryService.GetDashboardViewData(ctx, pagination.Page, pagination.Limit, searchTerm)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, err
	}
	return &model.DashboardViewItemPage{
		Content:       mapper.MapEntitiesToDashboardViewItems(paginatedResult.Rows.([]*entity.DashboardViewResultEntity)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
