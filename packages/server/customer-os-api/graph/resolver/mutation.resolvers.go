package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// EntityTemplateCreate is the resolver for the entityTemplateCreate field.
func (r *mutationResolver) EntityTemplateCreate(ctx context.Context, input model.EntityTemplateInput) (*model.EntityTemplate, error) {
	entityTemplateEntity, err := r.Services.EntityTemplateService.Create(ctx, mapper.MapEntityTemplateInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create entity template: %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToEntityTemplate(entityTemplateEntity), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
