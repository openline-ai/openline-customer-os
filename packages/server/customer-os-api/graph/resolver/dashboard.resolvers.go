package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// DashboardViewContacts is the resolver for the dashboardView_Contacts field.
func (r *queryResolver) DashboardViewContacts(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.ContactsPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardViewContacts", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	paginatedResult, err := r.Services.QueryService.GetDashboardViewContactsData(ctx, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get contacts data")
		return nil, nil
	}
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
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
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// DashboardViewPortfolioOrganizations is the resolver for the dashboardView_PortfolioOrganizations field.
func (r *queryResolver) DashboardViewPortfolioOrganizations(ctx context.Context, ownerID string, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.OrganizationPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardViewPortfolioOrganizations", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("ownerID", ownerID))
	span.LogFields(log.Object("pagination", pagination))
	if where != nil {
		span.LogFields(log.Object("filter", *where))
	}
	if sort != nil {
		span.LogFields(log.Object("sort", *sort))
	}

	paginatedResult, err := r.Services.QueryService.GetDashboardViewOrganizationsData(ctx,
		service.DashboardViewOrganizationsRequest{
			Page:    pagination.Page,
			Limit:   pagination.Limit,
			Where:   where,
			Sort:    sort,
			OwnerId: ownerID,
		})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, nil
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// DashboardViewRelationshipOrganizations is the resolver for the dashboardView_RelationshipOrganizations field.
func (r *queryResolver) DashboardViewRelationshipOrganizations(ctx context.Context, relationships []model.OrganizationRelationship, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.OrganizationPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.DashboardViewRelationshipOrganizations", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("relationships", relationships))
	span.LogFields(log.Object("pagination", pagination))
	if where != nil {
		span.LogFields(log.Object("filter", *where))
	}
	if sort != nil {
		span.LogFields(log.Object("sort", *sort))
	}

	relationshipsStr := make([]string, 0, len(relationships))
	for _, relationship := range relationships {
		relationshipStr := mapper.MapOrgRelationshipFromModel(relationship)
		if relationshipStr.IsValid() {
			relationshipsStr = append(relationshipsStr, relationshipStr.String())
		}
	}

	paginatedResult, err := r.Services.QueryService.GetDashboardViewOrganizationsData(ctx,
		service.DashboardViewOrganizationsRequest{
			Page:          pagination.Page,
			Limit:         pagination.Limit,
			Where:         where,
			Sort:          sort,
			Relationships: relationshipsStr,
		})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organizations")
		return nil, nil
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}
