package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// DashboardViewOrganizations is the resolver for the dashboardView_Organizations field.
func (r *queryResolver) DashboardViewOrganizations(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.OrganizationPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardViewOrganizations", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("pagination", pagination))
	if where != nil {
		span.LogFields(log.Object("filter", *where))
	}
	if sort != nil {
		span.LogFields(log.Object("sort", *sort))
	}

	paginatedResult, err := r.Services.QueryService.GetDashboardViewOrganizationsData(ctx, service.DashboardViewOrganizationsRequest{
		Page:  pagination.Page,
		Limit: pagination.Limit,
		Where: where,
		Sort:  sort,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, nil
	}
	countOrganizations, err := r.Services.OrganizationService.CountOrganizations(ctx, common.GetTenantFromContext(ctx))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, nil
	}

	return &model.OrganizationPage{
		Content:        mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:     paginatedResult.TotalPages,
		TotalElements:  paginatedResult.TotalRows,
		TotalAvailable: countOrganizations,
	}, err
}

// DashboardNewCustomers is the resolver for the dashboard_NewCustomers field.
func (r *queryResolver) DashboardNewCustomers(ctx context.Context, year int) (*model.DashboardNewCustomers, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardNewCustomers", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("year", year))

	newCustomersData, err := r.Services.QueryService.GetDashboardNewCustomersData(ctx, year)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get new customers data for year %d", year)
		return nil, nil
	}

	return mapper.MapDashboardNewCustomersData(newCustomersData), nil
}

// DashboardRetentionRate is the resolver for the dashboard_RetentionRate field.
func (r *queryResolver) DashboardRetentionRate(ctx context.Context, year int) (*model.DashboardRetentionRate, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardRetentionRate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("year", year))

	newCustomersData, err := r.Services.QueryService.GetDashboardRetentionRateData(ctx, year)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get new customers data for year %d", year)
		return nil, nil
	}

	return mapper.MapDashboardRetentionRateData(newCustomersData), nil
}
