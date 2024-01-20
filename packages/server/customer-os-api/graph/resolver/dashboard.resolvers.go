package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// Organization is the resolver for the organization field.
func (r *dashboardCustomerMapResolver) Organization(ctx context.Context, obj *model.DashboardCustomerMap) (*model.Organization, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "DashboardCustomerMapResolver.Organization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", obj.OrganizationID))

	if obj.OrganizationID == "" {
		tracing.TraceErr(span, errors.New("missing organization input id"))
		graphql.AddErrorf(ctx, "Missing organization input id")
		return nil, nil
	}

	organizationEntityPtr, err := dataloader.For(ctx).GetOrganization(ctx, obj.OrganizationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organization by id %s", obj.OrganizationID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntityPtr), nil
}

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

// DashboardViewRenewals is the resolver for the dashboardView_Renewals field.
func (r *queryResolver) DashboardViewRenewals(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.RenewalsPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardViewRenewals", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("pagination", pagination))
	if where != nil {
		span.LogFields(log.Object("filter", *where))
	}
	if sort != nil {
		span.LogFields(log.Object("sort", *sort))
	}

	paginatedResult, err := r.Services.QueryService.GetDashboardViewRenewalsData(ctx, service.DashboardViewRenewalsRequest{
		Page:  pagination.Page,
		Limit: pagination.Limit,
		Where: where,
		Sort:  sort,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get renewals data")
		return nil, nil
	}
	countContracts, err := r.Services.ContractService.CountContracts(ctx, common.GetTenantFromContext(ctx))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get contracts count")
		return nil, nil
	}

	return &model.RenewalsPage{
		Content:        mapper.MapEntitiesToRenewalRecords(paginatedResult.Rows.(*entity.RenewalsRecordEntities)),
		TotalPages:     paginatedResult.TotalPages,
		TotalElements:  paginatedResult.TotalRows,
		TotalAvailable: countContracts,
	}, err
}

// DashboardCustomerMap is the resolver for the dashboard_CustomerMap field.
func (r *queryResolver) DashboardCustomerMap(ctx context.Context) ([]*model.DashboardCustomerMap, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardCustomerMap", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	data, err := r.Services.QueryService.GetDashboardCustomerMapData(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the customer map")
		return nil, nil
	}

	return mapper.MapDashboardCustomerMapDataList(data), nil
}

// DashboardMRRPerCustomer is the resolver for the dashboard_MRRPerCustomer field.
func (r *queryResolver) DashboardMRRPerCustomer(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardMRRPerCustomer, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardMRRPerCustomer", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardMRRPerCustomerData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the MRR per customer data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return mapper.MapDashboardMRRPerCustomerData(data), nil
}

// DashboardGrossRevenueRetention is the resolver for the dashboard_GrossRevenueRetention field.
func (r *queryResolver) DashboardGrossRevenueRetention(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardGrossRevenueRetention, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardGrossRevenueRetention", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardGrossRevenueRetentionData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the gross revenue retention data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return mapper.MapDashboardGrossRevenueRetentionData(data), nil
}

// DashboardARRBreakdown is the resolver for the dashboard_ARRBreakdown field.
func (r *queryResolver) DashboardARRBreakdown(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardARRBreakdown, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardARRBreakdown", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardARRBreakdownData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the arr breakdown data for period %s - %s", startTime, endTime)
		return nil, nil
	}

	return mapper.MapDashboardARRBreakdownData(data), nil
}

// DashboardRevenueAtRisk is the resolver for the dashboard_RevenueAtRisk field.
func (r *queryResolver) DashboardRevenueAtRisk(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardRevenueAtRisk, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardRevenueAtRisk", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardRevenueAtRiskData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the revenue at risk data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return mapper.MapDashboardRevenueAtRiskData(data), nil
}

// DashboardRetentionRate is the resolver for the dashboard_RetentionRate field.
func (r *queryResolver) DashboardRetentionRate(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardRetentionRate, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardRetentionRate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardRetentionRateData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the retention rate data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return mapper.MapDashboardRetentionRateData(data), nil
}

// DashboardNewCustomers is the resolver for the dashboard_NewCustomers field.
func (r *queryResolver) DashboardNewCustomers(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardNewCustomers, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardNewCustomers", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardNewCustomersData(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get new customers data for period %s - %s", period.Start.String(), period.End.String())
		return nil, nil
	}

	return mapper.MapDashboardNewCustomersData(data), nil
}

// DashboardTimeToOnboard is the resolver for the dashboard_TimeToOnboard field.
func (r *queryResolver) DashboardTimeToOnboard(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardTimeToOnboard, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardTimeToOnboard", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardAverageTimeToOnboardPerMonth(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get time to onboard data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return data, nil
}

// DashboardOnboardingCompletion is the resolver for the dashboard_OnboardingCompletion field.
func (r *queryResolver) DashboardOnboardingCompletion(ctx context.Context, period *model.DashboardPeriodInput) (*model.DashboardOnboardingCompletion, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardOnboardingCompletion", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("period", period))

	startTime, endTime, err := getPeriod(period)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get the data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	data, err := r.Services.QueryService.GetDashboardOnboardingCompletionPerMonth(ctx, startTime, endTime)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get onboarding completion data for period %s - %s", startTime.String(), endTime.String())
		return nil, nil
	}

	return data, nil
}

// DashboardCustomerMap returns generated.DashboardCustomerMapResolver implementation.
func (r *Resolver) DashboardCustomerMap() generated.DashboardCustomerMapResolver {
	return &dashboardCustomerMapResolver{r}
}

type dashboardCustomerMapResolver struct{ *Resolver }
