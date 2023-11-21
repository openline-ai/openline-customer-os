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

// CreatedBy is the resolver for the createdBy field.
func (r *opportunityResolver) CreatedBy(ctx context.Context, obj *model.Opportunity) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// Owner is the resolver for the owner field.
func (r *opportunityResolver) Owner(ctx context.Context, obj *model.Opportunity) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *opportunityResolver) ExternalLinks(ctx context.Context, obj *model.Opportunity) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// Opportunity is the resolver for the opportunity field.
func (r *queryResolver) Opportunity(ctx context.Context, id string) (*model.Opportunity, error) {
	panic(fmt.Errorf("not implemented: Opportunity - opportunity"))
}

// Opportunity returns generated.OpportunityResolver implementation.
func (r *Resolver) Opportunity() generated.OpportunityResolver { return &opportunityResolver{r} }

type opportunityResolver struct{ *Resolver }
