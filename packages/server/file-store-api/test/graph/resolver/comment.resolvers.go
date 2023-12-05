package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// CreatedBy is the resolver for the createdBy field.
func (r *commentResolver) CreatedBy(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *commentResolver) ExternalLinks(ctx context.Context, obj *model.Comment) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
