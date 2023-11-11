package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
)

// Owner is the resolver for the owner field.
func (r *contractResolver) Owner(ctx context.Context, obj *model.Contract) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *contractResolver) ExternalLinks(ctx context.Context, obj *model.Contract) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// ContractCreate is the resolver for the contract_Create field.
func (r *mutationResolver) ContractCreate(ctx context.Context, input model.ContractInput) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: ContractCreate - contract_Create"))
}

// Contract is the resolver for the contract field.
func (r *queryResolver) Contract(ctx context.Context, id string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: Contract - contract"))
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *contractResolver) ServiceStartedAtTimelineEvent(ctx context.Context, obj *model.Contract) (model.TimelineEvent, error) {
	panic(fmt.Errorf("not implemented: ServiceStartedAtTimelineEvent - serviceStartedAtTimelineEvent"))
}
func (r *contractResolver) ServiceSignedAtTimelineEvent(ctx context.Context, obj *model.Contract) (model.TimelineEvent, error) {
	panic(fmt.Errorf("not implemented: ServiceSignedAtTimelineEvent - serviceSignedAtTimelineEvent"))
}
func (r *contractResolver) SignedAtTimelineEvent(ctx context.Context, obj *model.Contract) (model.TimelineEvent, error) {
	panic(fmt.Errorf("not implemented: SignedAtTimelineEvent - signedAtTimelineEvent"))
}
