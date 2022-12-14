package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/model"
)

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserCreate - userCreate"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.UserPage, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Conversations is the resolver for the conversations field.
func (r *userResolver) Conversations(ctx context.Context, obj *model.User, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	panic(fmt.Errorf("not implemented: Conversations - conversations"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
