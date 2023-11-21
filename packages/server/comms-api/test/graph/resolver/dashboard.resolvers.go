package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// Organization is the resolver for the organization field.
func (r *dashboardCustomerMapResolver) Organization(ctx context.Context, obj *model.DashboardCustomerMap) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// DashboardViewOrganizations is the resolver for the dashboardView_Organizations field.
func (r *queryResolver) DashboardViewOrganizations(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.OrganizationPage, error) {
	panic(fmt.Errorf("not implemented: DashboardViewOrganizations - dashboardView_Organizations"))
}

// DashboardCustomerMap is the resolver for the dashboard_CustomerMap field.
func (r *queryResolver) DashboardCustomerMap(ctx context.Context) ([]*model.DashboardCustomerMap, error) {
	panic(fmt.Errorf("not implemented: DashboardCustomerMap - dashboard_CustomerMap"))
}

// DashboardMRRPerCustomer is the resolver for the dashboard_MRRPerCustomer field.
func (r *queryResolver) DashboardMRRPerCustomer(ctx context.Context, year int) (*model.DashboardMRRPerCustomer, error) {
	panic(fmt.Errorf("not implemented: DashboardMRRPerCustomer - dashboard_MRRPerCustomer"))
}

// DashboardGrossRevenueRetention is the resolver for the dashboard_GrossRevenueRetention field.
func (r *queryResolver) DashboardGrossRevenueRetention(ctx context.Context, year int) (*model.DashboardGrossRevenueRetention, error) {
	panic(fmt.Errorf("not implemented: DashboardGrossRevenueRetention - dashboard_GrossRevenueRetention"))
}

// DashboardARRBreakdown is the resolver for the dashboard_ARRBreakdown field.
func (r *queryResolver) DashboardARRBreakdown(ctx context.Context, year int) (*model.DashboardARRBreakdown, error) {
	panic(fmt.Errorf("not implemented: DashboardARRBreakdown - dashboard_ARRBreakdown"))
}

// DashboardRevenueAtRisk is the resolver for the dashboard_RevenueAtRisk field.
func (r *queryResolver) DashboardRevenueAtRisk(ctx context.Context, year int) (*model.DashboardRevenueAtRisk, error) {
	panic(fmt.Errorf("not implemented: DashboardRevenueAtRisk - dashboard_RevenueAtRisk"))
}

// DashboardRetentionRate is the resolver for the dashboard_RetentionRate field.
func (r *queryResolver) DashboardRetentionRate(ctx context.Context, year int) (*model.DashboardRetentionRate, error) {
	panic(fmt.Errorf("not implemented: DashboardRetentionRate - dashboard_RetentionRate"))
}

// DashboardNewCustomers is the resolver for the dashboard_NewCustomers field.
func (r *queryResolver) DashboardNewCustomers(ctx context.Context, year int) (*model.DashboardNewCustomers, error) {
	panic(fmt.Errorf("not implemented: DashboardNewCustomers - dashboard_NewCustomers"))
}

// DashboardCustomerMap returns generated.DashboardCustomerMapResolver implementation.
func (r *Resolver) DashboardCustomerMap() generated.DashboardCustomerMapResolver {
	return &dashboardCustomerMapResolver{r}
}

type dashboardCustomerMapResolver struct{ *Resolver }
