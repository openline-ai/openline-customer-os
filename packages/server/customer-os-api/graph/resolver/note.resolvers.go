package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// NoteCreateForContact is the resolver for the note_CreateForContact field.
func (r *mutationResolver) NoteCreateForContact(ctx context.Context, contactID string, input model.NoteInput) (*model.Note, error) {
	result, err := r.Services.NoteService.CreateNoteForContact(ctx, contactID, mapper.MapNoteInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add note %s to contact %s", input.HTML, contactID)
		return nil, err
	}
	return mapper.MapEntityToNote(result), nil
}

// NoteCreateForOrganization is the resolver for the note_CreateForOrganization field.
func (r *mutationResolver) NoteCreateForOrganization(ctx context.Context, organizationID string, input model.NoteInput) (*model.Note, error) {
	result, err := r.Services.NoteService.CreateNoteForOrganization(ctx, organizationID, mapper.MapNoteInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add note %s to organization %s", input.HTML, organizationID)
		return nil, err
	}
	return mapper.MapEntityToNote(result), nil
}

// NoteUpdate is the resolver for the note_Update field.
func (r *mutationResolver) NoteUpdate(ctx context.Context, input model.NoteUpdateInput) (*model.Note, error) {
	result, err := r.Services.NoteService.UpdateNote(ctx, mapper.MapNoteUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update note %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToNote(result), nil
}

// NoteDelete is the resolver for the note_Delete field.
func (r *mutationResolver) NoteDelete(ctx context.Context, id string) (*model.Result, error) {
	result, err := r.Services.NoteService.DeleteNote(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to delete note %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// NoteLinkAttachment is the resolver for the note_LinkAttachment field.
func (r *mutationResolver) NoteLinkAttachment(ctx context.Context, noteID string, attachmentID string) (*model.Note, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	note, err := r.Services.NoteService.NoteLinkAttachment(ctx, noteID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToNote(note), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *noteResolver) CreatedBy(ctx context.Context, obj *model.Note) (*model.User, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	creator, err := r.Services.UserService.GetNoteCreator(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get creator for note %s", obj.ID)
		return nil, err
	}
	if creator == nil {
		return nil, nil
	}
	return mapper.MapEntityToUser(creator), err
}

// Noted is the resolver for the noted field.
func (r *noteResolver) Noted(ctx context.Context, obj *model.Note) ([]model.NotedEntity, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	entities, err := dataloader.For(ctx).GetNotedEntitiesForNote(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get noted entities for note %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotedEntities(entities), nil
}

// Includes is the resolver for the includes field.
func (r *noteResolver) Includes(ctx context.Context, obj *model.Note) ([]*model.Attachment, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	entities, err := r.Services.AttachmentService.GetAttachmentsForNode(ctx, repository.INCLUDED_BY_NOTE, nil, []string{obj.ID})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get attachment entities for note %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

type noteResolver struct{ *Resolver }
