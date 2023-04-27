package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// InteractionSession is the resolver for the interactionSession field.
func (r *interactionEventResolver) InteractionSession(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSession - interactionSession"))
}

// Meeting is the resolver for the meeting field.
func (r *interactionEventResolver) Meeting(ctx context.Context, obj *model.InteractionEvent) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: Meeting - meeting"))
}

// SentBy is the resolver for the sentBy field.
func (r *interactionEventResolver) SentBy(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	panic(fmt.Errorf("not implemented: SentBy - sentBy"))
}

// SentTo is the resolver for the sentTo field.
func (r *interactionEventResolver) SentTo(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	panic(fmt.Errorf("not implemented: SentTo - sentTo"))
}

// RepliesTo is the resolver for the repliesTo field.
func (r *interactionEventResolver) RepliesTo(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: RepliesTo - repliesTo"))
}

// Includes is the resolver for the includes field.
func (r *interactionEventResolver) Includes(ctx context.Context, obj *model.InteractionEvent) ([]*model.Attachment, error) {
	panic(fmt.Errorf("not implemented: Includes - includes"))
}

// Events is the resolver for the events field.
func (r *interactionSessionResolver) Events(ctx context.Context, obj *model.InteractionSession) ([]*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: Events - events"))
}

// AttendedBy is the resolver for the attendedBy field.
func (r *interactionSessionResolver) AttendedBy(ctx context.Context, obj *model.InteractionSession) ([]model.InteractionSessionParticipant, error) {
	panic(fmt.Errorf("not implemented: AttendedBy - attendedBy"))
}

// Includes is the resolver for the includes field.
func (r *interactionSessionResolver) Includes(ctx context.Context, obj *model.InteractionSession) ([]*model.Attachment, error) {
	panic(fmt.Errorf("not implemented: Includes - includes"))
}

// InteractionSessionCreate is the resolver for the interactionSession_Create field.
func (r *mutationResolver) InteractionSessionCreate(ctx context.Context, session model.InteractionSessionInput) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionCreate - interactionSession_Create"))
}

// InteractionSessionLinkAttachment is the resolver for the interactionSession_LinkAttachment field.
func (r *mutationResolver) InteractionSessionLinkAttachment(ctx context.Context, sessionID string, attachmentID string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionLinkAttachment - interactionSession_LinkAttachment"))
}

// InteractionEventCreate is the resolver for the interactionEvent_Create field.
func (r *mutationResolver) InteractionEventCreate(ctx context.Context, event model.InteractionEventInput) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEventCreate - interactionEvent_Create"))
}

// InteractionEventLinkAttachment is the resolver for the interactionEvent_LinkAttachment field.
func (r *mutationResolver) InteractionEventLinkAttachment(ctx context.Context, eventID string, attachmentID string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEventLinkAttachment - interactionEvent_LinkAttachment"))
}

// InteractionSession is the resolver for the interactionSession field.
func (r *queryResolver) InteractionSession(ctx context.Context, id string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSession - interactionSession"))
}

// InteractionSessionBySessionIdentifier is the resolver for the interactionSession_BySessionIdentifier field.
func (r *queryResolver) InteractionSessionBySessionIdentifier(ctx context.Context, sessionIdentifier string) (*model.InteractionSession, error) {
	panic(fmt.Errorf("not implemented: InteractionSessionBySessionIdentifier - interactionSession_BySessionIdentifier"))
}

// InteractionEvent is the resolver for the interactionEvent field.
func (r *queryResolver) InteractionEvent(ctx context.Context, id string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEvent - interactionEvent"))
}

// InteractionEventByEventIdentifier is the resolver for the interactionEvent_ByEventIdentifier field.
func (r *queryResolver) InteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEventByEventIdentifier - interactionEvent_ByEventIdentifier"))
}

// InteractionEvent returns generated.InteractionEventResolver implementation.
func (r *Resolver) InteractionEvent() generated.InteractionEventResolver {
	return &interactionEventResolver{r}
}

// InteractionSession returns generated.InteractionSessionResolver implementation.
func (r *Resolver) InteractionSession() generated.InteractionSessionResolver {
	return &interactionSessionResolver{r}
}

type interactionEventResolver struct{ *Resolver }
type interactionSessionResolver struct{ *Resolver }
