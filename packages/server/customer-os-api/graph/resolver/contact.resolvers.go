package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

// Tags is the resolver for the tags field.
func (r *contactResolver) Tags(ctx context.Context, obj *model.Contact) ([]*model.Tag, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	tagEntities, err := dataloader.For(ctx).GetTagsForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get tags for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// JobRoles is the resolver for the jobRoles field.
func (r *contactResolver) JobRoles(ctx context.Context, obj *model.Contact) ([]*model.JobRole, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	jobRoleEntities, err := r.Services.JobRoleService.GetAllForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get job roles for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToJobRoles(jobRoleEntities), err
}

// Organizations is the resolver for the organizations field.
func (r *contactResolver) Organizations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.OrganizationService.GetOrganizationsForContact(ctx, obj.ID, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not fetch organizations for contact %s", obj.ID)
		return nil, err
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Groups is the resolver for the groups field.
func (r *contactResolver) Groups(ctx context.Context, obj *model.Contact) ([]*model.ContactGroup, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	contactGroupEntities, err := r.Services.ContactGroupService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToContactGroups(contactGroupEntities), err
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *contactResolver) PhoneNumbers(ctx context.Context, obj *model.Contact) ([]*model.PhoneNumber, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	phoneNumberEntities, err := r.Services.PhoneNumberService.FindAllForContact(ctx, obj.ID)
	return mapper.MapEntitiesToPhoneNumbers(phoneNumberEntities), err
}

// Emails is the resolver for the emails field.
func (r *contactResolver) Emails(ctx context.Context, obj *model.Contact) ([]*model.Email, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	emailEntities, err := dataloader.For(ctx).GetEmailsForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get emails for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToEmails(emailEntities), nil
}

// Locations is the resolver for the locations field.
func (r *contactResolver) Locations(ctx context.Context, obj *model.Contact) ([]*model.Location, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	locationEntities, err := dataloader.For(ctx).GetLocationsForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get locations for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToLocations(locationEntities), err
}

// CustomFields is the resolver for the customFields field.
func (r *contactResolver) CustomFields(ctx context.Context, obj *model.Contact) ([]*model.CustomField, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	var customFields []*model.CustomField
	customFieldEntities, err := r.Services.CustomFieldService.FindAllForContact(ctx, obj)
	for _, v := range mapper.MapEntitiesToCustomFields(customFieldEntities) {
		customFields = append(customFields, v)
	}
	return customFields, err
}

// FieldSets is the resolver for the fieldSets field.
func (r *contactResolver) FieldSets(ctx context.Context, obj *model.Contact) ([]*model.FieldSet, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	fieldSetEntities, err := r.Services.FieldSetService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToFieldSets(fieldSetEntities), err
}

// Template is the resolver for the template field.
func (r *contactResolver) Template(ctx context.Context, obj *model.Contact) (*model.EntityTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	templateEntity, err := r.Services.EntityTemplateService.FindLinkedWithContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact template for contact %s", obj.ID)
		return nil, err
	}
	if templateEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToEntityTemplate(templateEntity), err
}

// Owner is the resolver for the owner field.
func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	owner, err := r.Services.UserService.FindContactOwner(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get owner for contact %s", obj.ID)
		return nil, err
	}
	if owner == nil {
		return nil, nil
	}
	return mapper.MapEntityToUser(owner), err
}

// Notes is the resolver for the notes field.
func (r *contactResolver) Notes(ctx context.Context, obj *model.Contact, pagination *model.Pagination) (*model.NotePage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.NoteService.GetNotesForContact(ctx, obj.ID, pagination.Page, pagination.Limit)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact %s notes", obj.ID)
		return nil, err
	}
	return &model.NotePage{
		Content:       mapper.MapEntitiesToNotes(paginatedResult.Rows.(*entity.NoteEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Conversations is the resolver for the conversations field.
func (r *contactResolver) Conversations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ConversationService.GetConversationsForContact(ctx, obj.ID, pagination.Page, pagination.Limit, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact %s conversations", obj.ID)
		return nil, err
	}
	return &model.ConversationPage{
		Content:       mapper.MapEntitiesToConversations(paginatedResult.Rows.(*entity.ConversationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Actions is the resolver for the actions field.
func (r *contactResolver) Actions(ctx context.Context, obj *model.Contact, from time.Time, to time.Time, actionTypes []model.ActionType) ([]model.Action, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	actions, err := r.Services.ActionsService.GetContactActions(ctx, obj.ID, from, to, actionTypes)
	if err != nil {
		graphql.AddErrorf(ctx, "failed to get actions for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToActions(actions), nil
}

// ContactCreate is the resolver for the contact_Create field.
func (r *mutationResolver) ContactCreate(ctx context.Context, input model.ContactInput) (*model.Contact, error) {
	contactNodeCreated, err := r.Services.ContactService.Create(ctx, &service.ContactCreateData{
		ContactEntity:     mapper.MapContactInputToEntity(input),
		CustomFields:      mapper.MapCustomFieldInputsToEntities(input.CustomFields),
		FieldSets:         mapper.MapFieldSetInputsToEntities(input.FieldSets),
		PhoneNumberEntity: mapper.MapPhoneNumberInputToEntity(input.PhoneNumber),
		EmailEntity:       mapper.MapEmailInputToEntity(input.Email),
		ExternalReference: mapper.MapExternalSystemReferenceInputToRelationship(input.ExternalReference),
		TemplateId:        input.TemplateID,
		OwnerUserId:       input.OwnerID,
		Source:            entity.DataSourceOpenline,
		SourceOfTruth:     entity.DataSourceOpenline,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create contact %s %s", *input.FirstName, *input.LastName)
		return nil, err
	}
	return mapper.MapEntityToContact(contactNodeCreated), nil
}

// ContactUpdate is the resolver for the contact_Update field.
func (r *mutationResolver) ContactUpdate(ctx context.Context, input model.ContactUpdateInput) (*model.Contact, error) {
	updatedContact, err := r.Services.ContactService.Update(ctx, &service.ContactUpdateData{
		ContactEntity: mapper.MapContactUpdateInputToEntity(input),
		OwnerUserId:   input.OwnerID,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update contact %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToContact(updatedContact), nil
}

// ContactHardDelete is the resolver for the contact_HardDelete field.
func (r *mutationResolver) ContactHardDelete(ctx context.Context, contactID string) (*model.Result, error) {
	result, err := r.Services.ContactService.PermanentDelete(ctx, contactID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not hard delete contact %s", contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// ContactSoftDelete is the resolver for the contact_SoftDelete field.
func (r *mutationResolver) ContactSoftDelete(ctx context.Context, contactID string) (*model.Result, error) {
	result, err := r.Services.ContactService.SoftDelete(ctx, contactID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not soft delete contact %s", contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// ContactAddTagByID is the resolver for the contact_AddTagById field.
func (r *mutationResolver) ContactAddTagByID(ctx context.Context, input *model.ContactTagInput) (*model.Contact, error) {
	updatedContact, err := r.Services.ContactService.AddTag(ctx, input.ContactID, input.TagID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to add tag %s to contact %s", input.TagID, input.ContactID)
		return nil, err
	}
	return mapper.MapEntityToContact(updatedContact), nil
}

// ContactRemoveTagByID is the resolver for the contact_RemoveTagById field.
func (r *mutationResolver) ContactRemoveTagByID(ctx context.Context, input *model.ContactTagInput) (*model.Contact, error) {
	updatedContact, err := r.Services.ContactService.RemoveTag(ctx, input.ContactID, input.TagID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to remove tag %s from contact %s", input.TagID, input.ContactID)
		return nil, err
	}
	return mapper.MapEntityToContact(updatedContact), nil
}

// Contact is the resolver for the contact field.
func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	contactEntity, err := r.Services.ContactService.GetContactById(ctx, id)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ContactService.FindAll(ctx, pagination.Page, pagination.Limit, where, sort)
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// ContactByEmail is the resolver for the contactByEmail field.
func (r *queryResolver) ContactByEmail(ctx context.Context, email string) (*model.Contact, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	contactEntity, err := r.Services.ContactService.FindContactByEmail(ctx, email)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with email %s not identified", email)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// ContactByPhone is the resolver for the contactByPhone field.
func (r *queryResolver) ContactByPhone(ctx context.Context, e164 string) (*model.Contact, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	contactEntity, err := r.Services.ContactService.FindContactByPhoneNumber(ctx, e164)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with phone number %s not identified", e164)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

type contactResolver struct{ *Resolver }
