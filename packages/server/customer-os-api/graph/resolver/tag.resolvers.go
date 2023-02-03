package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// TagCreate is the resolver for the tag_Create field.
func (r *mutationResolver) TagCreate(ctx context.Context, input model.TagInput) (*model.Tag, error) {
	createdTag, err := r.Services.TagService.Merge(ctx, mapper.MapTagInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create tag %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToTag(createdTag), nil
}

// TagUpdate is the resolver for the tag_Update field.
func (r *mutationResolver) TagUpdate(ctx context.Context, input model.TagUpdateInput) (*model.Tag, error) {
	updatedTag, err := r.Services.TagService.Update(ctx, mapper.MapTagUpdateInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update tag %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToTag(updatedTag), nil
}

// TagDelete is the resolver for the tag_Delete field.
func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*model.Result, error) {
	result, err := r.Services.TagService.UnlinkAndDelete(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to delete tag %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	tags, err := r.Services.TagService.GetAll(ctx)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch tags")
		return nil, err
	}
	return mapper.MapEntitiesToTags(tags), err
}
