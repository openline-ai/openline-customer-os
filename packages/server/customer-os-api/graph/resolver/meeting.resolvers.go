package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// AttendedBy is the resolver for the attendedBy field.
func (r *meetingResolver) AttendedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.AttendedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	participantEntities, err := dataloader.For(ctx).GetAttendedByParticipantsForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get participants for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToMeetingParticipants(participantEntities), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *meetingResolver) CreatedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.CreatedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	participantEntities, err := dataloader.For(ctx).GetCreatedByParticipantsForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get participants for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToMeetingParticipants(participantEntities), nil
}

// Includes is the resolver for the includes field.
func (r *meetingResolver) Includes(ctx context.Context, obj *model.Meeting) ([]*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.Includes", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	entities, err := dataloader.For(ctx).GetAttachmentsForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get attachment entities for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// DescribedBy is the resolver for the describedBy field.
func (r *meetingResolver) DescribedBy(ctx context.Context, obj *model.Meeting) ([]*model.Analysis, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.DescribedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	analysisEntities, err := dataloader.For(ctx).GetDescribedByForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get analysis for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAnalysis(analysisEntities), nil
}

// Note is the resolver for the note field.
func (r *meetingResolver) Note(ctx context.Context, obj *model.Meeting) ([]*model.Note, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.Note", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	notesForMeeting, err := dataloader.For(ctx).GetNotesForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get notes for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotes(notesForMeeting), nil
}

// Events is the resolver for the events field.
func (r *meetingResolver) Events(ctx context.Context, obj *model.Meeting) ([]*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.Events", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForMeeting(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get interaction events for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEvents(interactionEventEntities), nil
}

// Recording is the resolver for the recording field.
func (r *meetingResolver) Recording(ctx context.Context, obj *model.Meeting) (*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.Recording", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	recording := repository.INCLUDE_NATURE_RECORDING
	entities, err := r.Services.AttachmentService.GetAttachmentsForNode(ctx, repository.INCLUDED_BY_MEETING, &recording, []string{obj.ID})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get attachment entities for meeting %s", obj.ID)
		return nil, err
	}
	attachment := mapper.MapEntitiesToAttachment(entities)

	if len(attachment) == 0 {
		return nil, nil
	}

	return attachment[0], nil
}

// ExternalSystem is the resolver for the externalSystem field.
func (r *meetingResolver) ExternalSystem(ctx context.Context, obj *model.Meeting) ([]*model.ExternalSystem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MeetingResolver.ExternalSystem", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", obj.ID))

	externalSystemForMeeting, err := dataloader.For(ctx).GetExternalSystemsForEntity(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get notes for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToExternalSystems(externalSystemForMeeting), nil
}

// MeetingCreate is the resolver for the meeting_Create field.
func (r *mutationResolver) MeetingCreate(ctx context.Context, meeting model.MeetingInput) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	meetingEntity, err := r.Services.MeetingService.Create(ctx,
		&service.MeetingCreateData{
			MeetingEntity:     mapper.MapMeetingToEntity(&meeting),
			CreatedBy:         service.MapMeetingParticipantInputListToParticipant(meeting.CreatedBy),
			AttendedBy:        service.MapMeetingParticipantInputListToParticipant(meeting.AttendedBy),
			NoteInput:         meeting.Note,
			ExternalReference: mapper.MapExternalSystemReferenceInputToRelationship(meeting.ExternalSystem),
		})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to create meeting")
		return nil, err
	}
	newMeeting := mapper.MapEntityToMeeting(meetingEntity)
	return newMeeting, nil
}

// MeetingUpdate is the resolver for the meeting_Update field.
func (r *mutationResolver) MeetingUpdate(ctx context.Context, meetingID string, meeting model.MeetingUpdateInput) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID))

	input := &service.MeetingUpdateData{
		MeetingEntity: mapper.MapMeetingInputToEntity(&meeting),
		NoteEntity:    mapper.MapNoteUpdateInputToEntity(meeting.Note),
	}
	input.MeetingEntity.Id = meetingID
	meetingEntity, err := r.Services.MeetingService.Update(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update meeting")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToMeeting(meetingEntity)
	return interactionEvent, nil
}

// MeetingLinkAttendedBy is the resolver for the meeting_LinkAttendedBy field.
func (r *mutationResolver) MeetingLinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingLinkAttendedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID))

	err := r.Services.MeetingService.LinkAttendedBy(ctx, meetingID, service.MapMeetingParticipantInputToParticipant(&participant))
	if err != nil {
		return nil, err
	}

	meeting, err := r.Services.MeetingService.GetMeetingById(ctx, meetingID)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkAttendedBy is the resolver for the meeting_UnlinkAttendedBy field.
func (r *mutationResolver) MeetingUnlinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingUnlinkAttendedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID))

	err := r.Services.MeetingService.UnlinkAttendedBy(ctx, meetingID, service.MapMeetingParticipantInputToParticipant(&participant))
	if err != nil {
		return nil, err
	}

	meeting, err := r.Services.MeetingService.GetMeetingById(ctx, meetingID)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingLinkAttachment is the resolver for the meeting_LinkAttachment field.
func (r *mutationResolver) MeetingLinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingLinkAttachment", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID), log.String("request.attachmentID", attachmentID))

	meeting, err := r.Services.MeetingService.LinkAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkAttachment is the resolver for the meeting_UnlinkAttachment field.
func (r *mutationResolver) MeetingUnlinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingUnlinkAttachment", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID), log.String("request.attachmentID", attachmentID))

	meeting, err := r.Services.MeetingService.UnlinkAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingLinkRecording is the resolver for the meeting_LinkRecording field.
func (r *mutationResolver) MeetingLinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingLinkRecording", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID), log.String("request.attachmentID", attachmentID))

	meeting, err := r.Services.MeetingService.LinkRecordingAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkRecording is the resolver for the meeting_UnlinkRecording field.
func (r *mutationResolver) MeetingUnlinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingUnlinkRecording", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID), log.String("request.attachmentID", attachmentID))

	meeting, err := r.Services.MeetingService.UnlinkRecordingAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingAddNewLocation is the resolver for the meeting_AddNewLocation field.
func (r *mutationResolver) MeetingAddNewLocation(ctx context.Context, meetingID string) (*model.Location, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MeetingAddNewLocation", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", meetingID))

	locationEntity, err := r.Services.LocationService.CreateLocationForEntity(ctx, entity.MEETING, meetingID, entity.SourceFields{
		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
		AppSource:     constants.AppSourceCustomerOsApi,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error creating location for meeting %s", meetingID)
		return nil, err
	}
	return mapper.MapEntityToLocation(locationEntity), nil
}

// ExternalMeetingUpdate is the resolver for the externalMeeting_Update field.
func (r *mutationResolver) ExternalMeetingUpdate(ctx context.Context, externalSystemID string, externalID string, meeting model.MeetingUpdateInput) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ExternalMeetingUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.externalSystemID", externalSystemID))
	span.LogFields(log.String("request.externalID", externalID))
	response, findErr := r.Services.MeetingService.FindAll(ctx, externalSystemID, &externalID, 1, 1, nil, nil)
	meetings := mapper.MapEntitiesToMeetings(response.Rows.(*entity.MeetingEntities))
	if findErr != nil {
		tracing.TraceErr(span, findErr)
		graphql.AddErrorf(ctx, "Error find external meeting %s %s", externalSystemID, externalID)
		return nil, findErr
	}

	input := &service.MeetingUpdateData{
		MeetingEntity: mapper.MapMeetingInputToEntity(&meeting),
		NoteEntity:    mapper.MapNoteUpdateInputToEntity(meeting.Note),
	}
	input.MeetingEntity.Id = meetings[0].ID
	meetingEntity, err := r.Services.MeetingService.Update(ctx, input)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update meeting")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToMeeting(meetingEntity)
	return interactionEvent, nil
}

// Meeting is the resolver for the meeting field.
func (r *queryResolver) Meeting(ctx context.Context, id string) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Meeting", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.meetingID", id))

	meetingEntity, err := r.Services.MeetingService.GetMeetingById(ctx, id)
	if err != nil || meetingEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Meeting with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToMeeting(meetingEntity), nil
}

// ExternalMeetings is the resolver for the externalMeetings field.
func (r *queryResolver) ExternalMeetings(ctx context.Context, externalSystemID string, externalID *string, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.MeetingsPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Meetings", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	span.LogFields(log.Int("request.page", pagination.Page), log.Int("request.limit", pagination.Limit))
	paginatedResult, err := r.Services.MeetingService.FindAll(ctx, externalSystemID, externalID, pagination.Page, pagination.Limit, where, sort)
	return &model.MeetingsPage{
		Content:       mapper.MapEntitiesToMeetings(paginatedResult.Rows.(*entity.MeetingEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Meeting returns generated.MeetingResolver implementation.
func (r *Resolver) Meeting() generated.MeetingResolver { return &meetingResolver{r} }

type meetingResolver struct{ *Resolver }
