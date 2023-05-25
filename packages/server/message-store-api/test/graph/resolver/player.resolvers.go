package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// PlayerMerge is the resolver for the player_Merge field.
func (r *mutationResolver) PlayerMerge(ctx context.Context, input model.PlayerInput) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: PlayerMerge - player_Merge"))
}

// PlayerUpdate is the resolver for the player_Update field.
func (r *mutationResolver) PlayerUpdate(ctx context.Context, id string, update model.PlayerUpdate) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: PlayerUpdate - player_Update"))
}

// PlayerSetDefaultUser is the resolver for the player_SetDefaultUser field.
func (r *mutationResolver) PlayerSetDefaultUser(ctx context.Context, id string, userID string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: PlayerSetDefaultUser - player_SetDefaultUser"))
}

// Users is the resolver for the users field.
func (r *playerResolver) Users(ctx context.Context, obj *model.Player) ([]*model.PlayerUser, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// PlayerByAuthIDProvider is the resolver for the player_ByAuthIdProvider field.
func (r *queryResolver) PlayerByAuthIDProvider(ctx context.Context, authID string, provider string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: PlayerByAuthIDProvider - player_ByAuthIdProvider"))
}

// PlayerGetUsers is the resolver for the player_GetUsers field.
func (r *queryResolver) PlayerGetUsers(ctx context.Context) ([]*model.PlayerUser, error) {
	panic(fmt.Errorf("not implemented: PlayerGetUsers - player_GetUsers"))
}

// Player returns generated.PlayerResolver implementation.
func (r *Resolver) Player() generated.PlayerResolver { return &playerResolver{r} }

type playerResolver struct{ *Resolver }
