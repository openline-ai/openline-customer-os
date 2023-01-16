package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/model"
)

// Contacts is the resolver for the contacts field.
func (r *conversationResolver) Contacts(ctx context.Context, obj *model.Conversation) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// Users is the resolver for the users field.
func (r *conversationResolver) Users(ctx context.Context, obj *model.Conversation) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// ConversationCreate is the resolver for the conversationCreate field.
func (r *mutationResolver) ConversationCreate(ctx context.Context, input model.ConversationInput) (*model.Conversation, error) {
	if r.Resolver.ConversationCreate != nil {
		return r.Resolver.ConversationCreate(ctx, input)
	}
	panic(fmt.Errorf("not implemented: ConversationCreate - conversationCreate"))
}

// ConversationUpdate is the resolver for the conversation_Update field.
func (r *mutationResolver) ConversationUpdate(ctx context.Context, input model.ConversationUpdateInput) (*model.Conversation, error) {
	panic(fmt.Errorf("not implemented: ConversationUpdate - conversation_Update"))
}

// ConversationClose is the resolver for the conversation_Close field.
func (r *mutationResolver) ConversationClose(ctx context.Context, conversationID string) (*model.Conversation, error) {
	panic(fmt.Errorf("not implemented: ConversationClose - conversation_Close"))
}

// Conversation returns generated.ConversationResolver implementation.
func (r *Resolver) Conversation() generated.ConversationResolver { return &conversationResolver{r} }

type conversationResolver struct{ *Resolver }
