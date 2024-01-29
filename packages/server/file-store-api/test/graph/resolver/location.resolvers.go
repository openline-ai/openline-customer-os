package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// LocationRemoveFromContact is the resolver for the location_RemoveFromContact field.
func (r *mutationResolver) LocationRemoveFromContact(ctx context.Context, contactID string, locationID string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: LocationRemoveFromContact - location_RemoveFromContact"))
}

// LocationRemoveFromOrganization is the resolver for the location_RemoveFromOrganization field.
func (r *mutationResolver) LocationRemoveFromOrganization(ctx context.Context, organizationID string, locationID string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: LocationRemoveFromOrganization - location_RemoveFromOrganization"))
}

// LocationUpdate is the resolver for the location_Update field.
func (r *mutationResolver) LocationUpdate(ctx context.Context, input model.LocationUpdateInput) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: LocationUpdate - location_Update"))
}
