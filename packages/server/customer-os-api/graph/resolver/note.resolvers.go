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

// NoteMergeToContact is the resolver for the note_MergeToContact field.
func (r *mutationResolver) NoteMergeToContact(ctx context.Context, contactID string, input model.NoteInput) (*model.Note, error) {
	result, err := r.Services.NoteService.MergeNoteToContact(ctx, contactID, mapper.MapNoteInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add note %s to contact %s", input.HTML, contactID)
		return nil, err
	}
	return mapper.MapEntityToNote(result), nil
}

// NoteUpdateInContact is the resolver for the note_UpdateInContact field.
func (r *mutationResolver) NoteUpdateInContact(ctx context.Context, contactID string, input model.NoteUpdateInput) (*model.Note, error) {
	result, err := r.Services.NoteService.UpdateNoteInContact(ctx, contactID, mapper.MapNoteUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update note %s in contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToNote(result), nil
}

// NoteDeleteFromContact is the resolver for the note_DeleteFromContact field.
func (r *mutationResolver) NoteDeleteFromContact(ctx context.Context, contactID string, noteID string) (*model.Result, error) {
	result, err := r.Services.NoteService.DeleteFromContact(ctx, contactID, noteID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove note %s from contact %s", noteID, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *noteResolver) CreatedBy(ctx context.Context, obj *model.Note) (*model.User, error) {
	creator, err := r.Services.UserService.FindNoteCreator(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get creator for note %s", obj.ID)
		return nil, err
	}
	if creator == nil {
		return nil, nil
	}
	return mapper.MapEntityToUser(creator), err
}

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

type noteResolver struct{ *Resolver }
