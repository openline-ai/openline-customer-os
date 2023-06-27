package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// PhoneNumberMergeToContact is the resolver for the phoneNumberMergeToContact field.
func (r *mutationResolver) PhoneNumberMergeToContact(ctx context.Context, contactID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	spanCtx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberMergeToContact", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(spanCtx, span)
	span.LogFields(log.String("request.contactID", contactID))

	result, err := r.Services.PhoneNumberService.MergePhoneNumberTo(spanCtx, entity.CONTACT, contactID, mapper.MapPhoneNumberInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(spanCtx, "Could not add phone number %s to contact %s", input.PhoneNumber, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberUpdateInContact is the resolver for the phoneNumberUpdateInContact field.
func (r *mutationResolver) PhoneNumberUpdateInContact(ctx context.Context, contactID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberUpdateInContact", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID), log.String("request.phoneNumberID", input.ID))

	result, err := r.Services.PhoneNumberService.UpdatePhoneNumberFor(ctx, entity.CONTACT, contactID, mapper.MapPhoneNumberUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update phone number %s for contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberRemoveFromContactByE164 is the resolver for the phoneNumberRemoveFromContactByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByE164(ctx context.Context, contactID string, e164 string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromContactByE164", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID))

	result, err := r.Services.PhoneNumberService.DetachFromEntityByPhoneNumber(ctx, entity.CONTACT, contactID, e164)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by e164 %s from contact with id %s", e164, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberRemoveFromContactByID is the resolver for the phoneNumberRemoveFromContactById field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromContactByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID), log.String("request.phoneNumberID", id))

	result, err := r.Services.PhoneNumberService.DetachFromEntityById(ctx, entity.CONTACT, contactID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by id %s from contact with id %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberMergeToOrganization is the resolver for the phoneNumberMergeToOrganization field.
func (r *mutationResolver) PhoneNumberMergeToOrganization(ctx context.Context, organizationID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberMergeToOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.PhoneNumberService.MergePhoneNumberTo(ctx, entity.ORGANIZATION, organizationID, mapper.MapPhoneNumberInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not add phone number %s to organization %s", input.PhoneNumber, organizationID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberUpdateInOrganization is the resolver for the phoneNumberUpdateInOrganization field.
func (r *mutationResolver) PhoneNumberUpdateInOrganization(ctx context.Context, organizationID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberUpdateInOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.PhoneNumberService.UpdatePhoneNumberFor(ctx, entity.ORGANIZATION, organizationID, mapper.MapPhoneNumberUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update phone number %s for organization %s", input.ID, organizationID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberRemoveFromOrganizationByE164 is the resolver for the phoneNumberRemoveFromOrganizationByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromOrganizationByE164(ctx context.Context, organizationID string, e164 string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromOrganizationByE164", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.PhoneNumberService.DetachFromEntityByPhoneNumber(ctx, entity.ORGANIZATION, organizationID, e164)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by e164 %s from user with id %s", e164, organizationID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberRemoveFromOrganizationByID is the resolver for the phoneNumberRemoveFromOrganizationById field.
func (r *mutationResolver) PhoneNumberRemoveFromOrganizationByID(ctx context.Context, organizationID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromOrganizationByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID), log.String("request.phoneNumberID", id))

	result, err := r.Services.PhoneNumberService.DetachFromEntityById(ctx, entity.ORGANIZATION, organizationID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by id %s from organization with id %s", id, organizationID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberMergeToUser is the resolver for the phoneNumberMergeToUser field.
func (r *mutationResolver) PhoneNumberMergeToUser(ctx context.Context, userID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberMergeToUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	result, err := r.Services.PhoneNumberService.MergePhoneNumberTo(ctx, entity.USER, userID, mapper.MapPhoneNumberInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not add phone number %s to user %s", input.PhoneNumber, userID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberUpdateInUser is the resolver for the phoneNumberUpdateInUser field.
func (r *mutationResolver) PhoneNumberUpdateInUser(ctx context.Context, userID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberUpdateInUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	result, err := r.Services.PhoneNumberService.UpdatePhoneNumberFor(ctx, entity.USER, userID, mapper.MapPhoneNumberUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update phone number %s for user %s", input.ID, userID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberRemoveFromUserByE164 is the resolver for the phoneNumberRemoveFromUserByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromUserByE164(ctx context.Context, userID string, e164 string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromUserByE164", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	result, err := r.Services.PhoneNumberService.DetachFromEntityByPhoneNumber(ctx, entity.USER, userID, e164)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by e164 %s from user with id %s", e164, userID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberRemoveFromUserByID is the resolver for the phoneNumberRemoveFromUserById field.
func (r *mutationResolver) PhoneNumberRemoveFromUserByID(ctx context.Context, userID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.PhoneNumberRemoveFromUserByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID), log.String("request.phoneNumberID", id))

	result, err := r.Services.PhoneNumberService.DetachFromEntityById(ctx, entity.USER, userID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove phone number by id %s from user with id %s", id, userID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// Users is the resolver for the users field.
func (r *phoneNumberResolver) Users(ctx context.Context, obj *model.PhoneNumber) ([]*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "PhoneNumberResolver.Users", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.phoneNumberID", obj.ID))

	userEntities, err := dataloader.For(ctx).GetUsersForPhoneNumber(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get users for phone number %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToUsers(userEntities), nil
}

// Contacts is the resolver for the contacts field.
func (r *phoneNumberResolver) Contacts(ctx context.Context, obj *model.PhoneNumber) ([]*model.Contact, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "PhoneNumberResolver.Contacts", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.phoneNumberID", obj.ID))

	contactEntities, err := dataloader.For(ctx).GetContactsForPhoneNumber(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get contacts for phone number %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToContacts(contactEntities), nil
}

// Organizations is the resolver for the organizations field.
func (r *phoneNumberResolver) Organizations(ctx context.Context, obj *model.PhoneNumber) ([]*model.Organization, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "PhoneNumberResolver.Organizations", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.phoneNumberID", obj.ID))

	organizationEntities, err := dataloader.For(ctx).GetOrganizationsForPhoneNumber(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organizations for phone number %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToOrganizations(organizationEntities), nil
}

// PhoneNumber returns generated.PhoneNumberResolver implementation.
func (r *Resolver) PhoneNumber() generated.PhoneNumberResolver { return &phoneNumberResolver{r} }

type phoneNumberResolver struct{ *Resolver }
