package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// ServiceLineItemCreate is the resolver for the serviceLineItemCreate field.
func (r *mutationResolver) ServiceLineItemCreate(ctx context.Context, input model.ServiceLineItemInput) (*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItemCreate - serviceLineItemCreate"))
}

// ServiceLineItemUpdate is the resolver for the serviceLineItemUpdate field.
func (r *mutationResolver) ServiceLineItemUpdate(ctx context.Context, input model.ServiceLineItemUpdateInput) (*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItemUpdate - serviceLineItemUpdate"))
}

// ServiceLineItemDelete is the resolver for the serviceLineItem_Delete field.
func (r *mutationResolver) ServiceLineItemDelete(ctx context.Context, id string) (*model.DeleteResponse, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItemDelete - serviceLineItem_Delete"))
}

// ServiceLineItemClose is the resolver for the serviceLineItem_Close field.
func (r *mutationResolver) ServiceLineItemClose(ctx context.Context, input model.ServiceLineItemCloseInput) (string, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItemClose - serviceLineItem_Close"))
}

// ServiceLineItem is the resolver for the serviceLineItem field.
func (r *queryResolver) ServiceLineItem(ctx context.Context, id string) (*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItem - serviceLineItem"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *serviceLineItemResolver) CreatedBy(ctx context.Context, obj *model.ServiceLineItem) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *serviceLineItemResolver) ExternalLinks(ctx context.Context, obj *model.ServiceLineItem) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// ServiceLineItem returns generated.ServiceLineItemResolver implementation.
func (r *Resolver) ServiceLineItem() generated.ServiceLineItemResolver {
	return &serviceLineItemResolver{r}
}

type serviceLineItemResolver struct{ *Resolver }
