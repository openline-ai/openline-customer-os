package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/model"
)

// TenantMerge is the resolver for the tenant_Merge field.
func (r *mutationResolver) TenantMerge(ctx context.Context, tenant model.TenantInput) (string, error) {
	panic(fmt.Errorf("not implemented: TenantMerge - tenant_Merge"))
}

// TenantAddBillingProfile is the resolver for the tenant_AddBillingProfile field.
func (r *mutationResolver) TenantAddBillingProfile(ctx context.Context, input model.TenantBillingProfileInput) (*model.TenantBillingProfile, error) {
	panic(fmt.Errorf("not implemented: TenantAddBillingProfile - tenant_AddBillingProfile"))
}

// TenantUpdateBillingProfile is the resolver for the tenant_UpdateBillingProfile field.
func (r *mutationResolver) TenantUpdateBillingProfile(ctx context.Context, input model.TenantBillingProfileUpdateInput) (*model.TenantBillingProfile, error) {
	panic(fmt.Errorf("not implemented: TenantUpdateBillingProfile - tenant_UpdateBillingProfile"))
}

// TenantUpdateSettings is the resolver for the tenant_UpdateSettings field.
func (r *mutationResolver) TenantUpdateSettings(ctx context.Context, input *model.TenantSettingsInput) (*model.TenantSettings, error) {
	panic(fmt.Errorf("not implemented: TenantUpdateSettings - tenant_UpdateSettings"))
}

// Tenant is the resolver for the tenant field.
func (r *queryResolver) Tenant(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Tenant - tenant"))
}

// TenantByWorkspace is the resolver for the tenant_ByWorkspace field.
func (r *queryResolver) TenantByWorkspace(ctx context.Context, workspace model.WorkspaceInput) (*string, error) {
	panic(fmt.Errorf("not implemented: TenantByWorkspace - tenant_ByWorkspace"))
}

// TenantByEmail is the resolver for the tenant_ByEmail field.
func (r *queryResolver) TenantByEmail(ctx context.Context, email string) (*string, error) {
	panic(fmt.Errorf("not implemented: TenantByEmail - tenant_ByEmail"))
}

// TenantBillingProfiles is the resolver for the tenantBillingProfiles field.
func (r *queryResolver) TenantBillingProfiles(ctx context.Context) ([]*model.TenantBillingProfile, error) {
	panic(fmt.Errorf("not implemented: TenantBillingProfiles - tenantBillingProfiles"))
}

// TenantBillingProfile is the resolver for the tenantBillingProfile field.
func (r *queryResolver) TenantBillingProfile(ctx context.Context, id string) (*model.TenantBillingProfile, error) {
	panic(fmt.Errorf("not implemented: TenantBillingProfile - tenantBillingProfile"))
}

// TenantSettings is the resolver for the tenantSettings field.
func (r *queryResolver) TenantSettings(ctx context.Context) (*model.TenantSettings, error) {
	panic(fmt.Errorf("not implemented: TenantSettings - tenantSettings"))
}
