package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/sirupsen/logrus"
)

// PhoneNumberUpsertInEventStore is the resolver for the phoneNumberUpsertInEventStore field.
func (r *mutationResolver) PhoneNumberUpsertInEventStore(ctx context.Context, size int) (int, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, _, err := r.Services.PhoneNumberService.UpsertInEventStore(ctx, size)
	if err != nil {
		logrus.Errorf("Failed to call method: %v", err)
		graphql.AddErrorf(ctx, "Failed to upsert phone numbers to event store")
	}

	return result, nil
}

// ContactUpsertInEventStore is the resolver for the contactUpsertInEventStore field.
func (r *mutationResolver) ContactUpsertInEventStore(ctx context.Context, size int) (int, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, _, err := r.Services.ContactService.UpsertInEventStore(ctx, size)
	if err != nil {
		logrus.Errorf("Failed to call method: %v", err)
		graphql.AddErrorf(ctx, "Failed to upsert contacts to event store")
	}

	return result, nil
}

// ContactPhoneNumberRelationUpsertInEventStore is the resolver for the contactPhoneNumberRelationUpsertInEventStore field.
func (r *mutationResolver) ContactPhoneNumberRelationUpsertInEventStore(ctx context.Context, size int) (int, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, _, err := r.Services.ContactService.UpsertPhoneNumberRelationInEventStore(ctx, size)
	if err != nil {
		logrus.Errorf("Failed to call method: %v", err)
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
	}

	return result, nil
}

// UpsertInEventStore is the resolver for the UpsertInEventStore field.
func (r *mutationResolver) UpsertInEventStore(ctx context.Context, size int) (*model.UpsertToEventStoreResult, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	output := model.UpsertToEventStoreResult{}

	processedPhoneNumbers, failedPhoneNumbers, err := r.Services.PhoneNumberService.UpsertInEventStore(ctx, size)
	output.PhoneNumberCount = processedPhoneNumbers
	output.PhoneNumberCountFailed = failedPhoneNumbers
	if err != nil || failedPhoneNumbers > 0 {
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	processedContacts, failedContacts, err := r.Services.ContactService.UpsertInEventStore(ctx, size)
	output.ContactCount = processedContacts
	output.ContactCountFailed = failedContacts
	if err != nil || failedContacts > 0 {
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	processedEmails, failedEmails, err := r.Services.EmailService.UpsertInEventStore(ctx, size)
	output.EmailCount = processedEmails
	output.EmailCountFailed = failedEmails
	if err != nil || failedEmails > 0 {
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	if processedPhoneNumbers < size && processedContacts < size {
		processedCount, failedCount, err := r.Services.ContactService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.ContactPhoneNumberRelationCount = processedCount
		output.ContactPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	if processedEmails < size && processedContacts < size {
		processedCount, failedCount, err := r.Services.ContactService.UpsertEmailRelationInEventStore(ctx, size)
		output.ContactEmailRelationCount = processedCount
		output.ContactEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	processedOrgs, failedOrgs, err := r.Services.OrganizationService.UpsertInEventStore(ctx, size)
	output.OrganizationCount = processedOrgs
	output.OrganizationCountFailed = failedOrgs
	if err != nil || failedOrgs > 0 {
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	if processedPhoneNumbers < size && processedOrgs < size {
		processedCount, failedCount, err := r.Services.OrganizationService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.OrganizationPhoneNumberRelationCount = processedCount
		output.OrganizationPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	if processedEmails < size && processedOrgs < size {
		processedCount, failedCount, err := r.Services.OrganizationService.UpsertEmailRelationInEventStore(ctx, size)
		output.OrganizationEmailRelationCount = processedCount
		output.OrganizationEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	processedUsers, failedUsers, err := r.Services.UserService.UpsertInEventStore(ctx, size)
	output.UserCount = processedUsers
	output.UserCountFailed = failedUsers
	if err != nil || failedUsers > 0 {
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	if processedPhoneNumbers < size && processedUsers < size {
		processedCount, failedCount, err := r.Services.UserService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.UserPhoneNumberRelationCount = processedCount
		output.UserPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	if processedEmails < size && processedUsers < size {
		processedCount, failedCount, err := r.Services.UserService.UpsertEmailRelationInEventStore(ctx, size)
		output.UserEmailRelationCount = processedCount
		output.UserEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	return &output, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
