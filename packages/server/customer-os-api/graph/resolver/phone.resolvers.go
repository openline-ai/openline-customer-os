package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// PhoneNumberMergeToContact is the resolver for the phoneNumberMergeToContact field.
func (r *mutationResolver) PhoneNumberMergeToContact(ctx context.Context, contactID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	result, err := r.Services.PhoneNumberService.MergePhoneNumberToContact(ctx, contactID, mapper.MapPhoneNumberInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add phone number %s to contact %s", input.PhoneNumber, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberUpdateInContact is the resolver for the phoneNumberUpdateInContact field.
func (r *mutationResolver) PhoneNumberUpdateInContact(ctx context.Context, contactID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	result, err := r.Services.PhoneNumberService.UpdatePhoneNumberInContact(ctx, contactID, mapper.MapPhoneNumberUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update phone number %s in contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberRemoveFromContactByE164 is the resolver for the phoneNumberRemoveFromContactByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByE164(ctx context.Context, contactID string, e164 string) (*model.Result, error) {
	result, err := r.Services.PhoneNumberService.RemoveFromContactByE164(ctx, contactID, e164)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove phone number by e164 %s from contact with id %s", e164, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberRemoveFromContactByID is the resolver for the phoneNumberRemoveFromContactById field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	result, err := r.Services.PhoneNumberService.RemoveFromContactById(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove phone number by id %s from contact with id %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}
