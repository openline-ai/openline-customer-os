package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// ContactPhoneNumberRelationUpsertInEventStore is the resolver for the contactPhoneNumberRelationUpsertInEventStore field.
func (r *mutationResolver) ContactPhoneNumberRelationUpsertInEventStore(ctx context.Context, size int) (int, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContactPhoneNumberRelationUpsertInEventStore", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	result, _, err := r.Services.ContactService.UpsertPhoneNumberRelationInEventStore(ctx, size)
	if err != nil {
		r.log.Errorf("%s - Failed to call method: {%v}", utils.GetFunctionName(), err.Error())
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
	}

	return result, nil
}

// UpsertInEventStore is the resolver for the UpsertInEventStore field.
func (r *mutationResolver) UpsertInEventStore(ctx context.Context, size int) (*model.UpsertToEventStoreResult, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UpsertInEventStore", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	output := model.UpsertToEventStoreResult{}

	return &output, nil

	{
		processedCount, failedCount, err := r.Services.ContactService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.ContactPhoneNumberRelationCount = processedCount
		output.ContactPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	{
		processedCount, failedCount, err := r.Services.ContactService.UpsertEmailRelationInEventStore(ctx, size)
		output.ContactEmailRelationCount = processedCount
		output.ContactEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	processedOrgs, failedOrgs, err := r.Services.OrganizationService.UpsertInEventStore(ctx, size)
	output.OrganizationCount = processedOrgs
	output.OrganizationCountFailed = failedOrgs
	if err != nil || failedOrgs > 0 {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed: {%s}", err)
		return &output, err
	}

	if processedOrgs < size {
		processedCount, failedCount, err := r.Services.OrganizationService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.OrganizationPhoneNumberRelationCount = processedCount
		output.OrganizationPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	if processedOrgs < size {
		processedCount, failedCount, err := r.Services.OrganizationService.UpsertEmailRelationInEventStore(ctx, size)
		output.OrganizationEmailRelationCount = processedCount
		output.OrganizationEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	{
		processedCount, failedCount, err := r.Services.UserService.UpsertPhoneNumberRelationInEventStore(ctx, size)
		output.UserPhoneNumberRelationCount = processedCount
		output.UserPhoneNumberRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	{
		processedCount, failedCount, err := r.Services.UserService.UpsertEmailRelationInEventStore(ctx, size)
		output.UserEmailRelationCount = processedCount
		output.UserEmailRelationCountFailed = failedCount
		if err != nil || failedCount > 0 {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed: {%s}", err)
			return &output, err
		}
	}

	return &output, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
