package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/model"
)

// ContractLineItems is the resolver for the contractLineItems field.
func (r *contractResolver) ContractLineItems(ctx context.Context, obj *model.Contract) ([]*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ContractLineItems - contractLineItems"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *contractResolver) CreatedBy(ctx context.Context, obj *model.Contract) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *contractResolver) ExternalLinks(ctx context.Context, obj *model.Contract) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// Opportunities is the resolver for the opportunities field.
func (r *contractResolver) Opportunities(ctx context.Context, obj *model.Contract) ([]*model.Opportunity, error) {
	panic(fmt.Errorf("not implemented: Opportunities - opportunities"))
}

// Owner is the resolver for the owner field.
func (r *contractResolver) Owner(ctx context.Context, obj *model.Contract) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Attachments is the resolver for the attachments field.
func (r *contractResolver) Attachments(ctx context.Context, obj *model.Contract) ([]*model.Attachment, error) {
	panic(fmt.Errorf("not implemented: Attachments - attachments"))
}

// ServiceLineItems is the resolver for the serviceLineItems field.
func (r *contractResolver) ServiceLineItems(ctx context.Context, obj *model.Contract) ([]*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItems - serviceLineItems"))
}

// ContractCreate is the resolver for the contract_Create field.
func (r *mutationResolver) ContractCreate(ctx context.Context, input model.ContractInput) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: ContractCreate - contract_Create"))
}

// ContractUpdate is the resolver for the contract_Update field.
func (r *mutationResolver) ContractUpdate(ctx context.Context, input model.ContractUpdateInput) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: ContractUpdate - contract_Update"))
}

// ContractDelete is the resolver for the contract_Delete field.
func (r *mutationResolver) ContractDelete(ctx context.Context, id string) (*model.DeleteResponse, error) {
	panic(fmt.Errorf("not implemented: ContractDelete - contract_Delete"))
}

// ContractAddAttachment is the resolver for the contract_AddAttachment field.
func (r *mutationResolver) ContractAddAttachment(ctx context.Context, contractID string, attachmentID string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: ContractAddAttachment - contract_AddAttachment"))
}

// ContractRemoveAttachment is the resolver for the contract_RemoveAttachment field.
func (r *mutationResolver) ContractRemoveAttachment(ctx context.Context, contractID string, attachmentID string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: ContractRemoveAttachment - contract_RemoveAttachment"))
}

// Contract is the resolver for the contract field.
func (r *queryResolver) Contract(ctx context.Context, id string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: Contract - contract"))
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }
