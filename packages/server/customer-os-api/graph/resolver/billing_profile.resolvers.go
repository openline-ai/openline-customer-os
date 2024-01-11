package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// BillingProfileCreate is the resolver for the billingProfile_Create field.
func (r *mutationResolver) BillingProfileCreate(ctx context.Context, input model.BillingProfileInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.BillingProfileCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	billingProfileId, err := r.Services.BillingProfileService.CreateBillingProfile(ctx, input.OrganizationID, utils.IfNotNilString(input.LegalName), utils.IfNotNilString(input.TaxID), input.CreatedAt)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create master plan")
	}
	return billingProfileId, err
}
