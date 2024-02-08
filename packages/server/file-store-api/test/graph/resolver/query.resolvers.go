package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// EntityTemplates is the resolver for the entityTemplates field.
func (r *queryResolver) EntityTemplates(ctx context.Context, extends *model.EntityTemplateExtension) ([]*model.EntityTemplate, error) {
	panic(fmt.Errorf("not implemented: EntityTemplates - entityTemplates"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
