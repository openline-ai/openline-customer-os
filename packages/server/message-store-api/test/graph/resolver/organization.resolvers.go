package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// OrganizationCreate is the resolver for the organization_Create field.
func (r *mutationResolver) OrganizationCreate(ctx context.Context, input model.OrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationCreate - organization_Create"))
}

// OrganizationUpdate is the resolver for the organization_Update field.
func (r *mutationResolver) OrganizationUpdate(ctx context.Context, input model.OrganizationUpdateInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationUpdate - organization_Update"))
}

// OrganizationDelete is the resolver for the organization_Delete field.
func (r *mutationResolver) OrganizationDelete(ctx context.Context, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: OrganizationDelete - organization_Delete"))
}

// OrganizationMerge is the resolver for the organization_Merge field.
func (r *mutationResolver) OrganizationMerge(ctx context.Context, primaryOrganizationID string, mergedOrganizationIds []string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationMerge - organization_Merge"))
}

// Domains is the resolver for the domains field.
func (r *organizationResolver) Domains(ctx context.Context, obj *model.Organization) ([]string, error) {
	panic(fmt.Errorf("not implemented: Domains - domains"))
}

// OrganizationType is the resolver for the organizationType field.
func (r *organizationResolver) OrganizationType(ctx context.Context, obj *model.Organization) (*model.OrganizationType, error) {
	panic(fmt.Errorf("not implemented: OrganizationType - organizationType"))
}

// Locations is the resolver for the locations field.
func (r *organizationResolver) Locations(ctx context.Context, obj *model.Organization) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented: Locations - locations"))
}

// Contacts is the resolver for the contacts field.
func (r *organizationResolver) Contacts(ctx context.Context, obj *model.Organization, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// JobRoles is the resolver for the jobRoles field.
func (r *organizationResolver) JobRoles(ctx context.Context, obj *model.Organization) ([]*model.JobRole, error) {
	panic(fmt.Errorf("not implemented: JobRoles - jobRoles"))
}

// Notes is the resolver for the notes field.
func (r *organizationResolver) Notes(ctx context.Context, obj *model.Organization, pagination *model.Pagination) (*model.NotePage, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Tags is the resolver for the tags field.
func (r *organizationResolver) Tags(ctx context.Context, obj *model.Organization) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

type organizationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *organizationResolver) Addresses(ctx context.Context, obj *model.Organization) ([]*model.Place, error) {
	panic(fmt.Errorf("not implemented: Addresses - addresses"))
}
