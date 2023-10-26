package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	usergrpc "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/user"
	"github.com/opentracing/opentracing-go/log"
)

// Users is the resolver for the users field.
func (r *emailResolver) Users(ctx context.Context, obj *model.Email) ([]*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntities, err := dataloader.For(ctx).GetUsersForEmail(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get users for email %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get users for email %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToUsers(userEntities), nil
}

// Contacts is the resolver for the contacts field.
func (r *emailResolver) Contacts(ctx context.Context, obj *model.Email) ([]*model.Contact, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	contactEntities, err := dataloader.For(ctx).GetContactsForEmail(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get contacts for email %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get contacts for email %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToContacts(contactEntities), nil
}

// Organizations is the resolver for the organizations field.
func (r *emailResolver) Organizations(ctx context.Context, obj *model.Email) ([]*model.Organization, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	organizationEntities, err := dataloader.For(ctx).GetOrganizationsForEmail(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get organizations for email %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get organizations for email %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToOrganizations(organizationEntities), nil
}

// EmailMergeToContact is the resolver for the emailMergeToContact field.
func (r *mutationResolver) EmailMergeToContact(ctx context.Context, contactID string, input model.EmailInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailMergeToContact", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID))

	result, err := r.Services.EmailService.MergeEmailTo(ctx, entity.CONTACT, contactID, mapper.MapEmailInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not add email %s to contact %s", input.Email, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailUpdateInContact is the resolver for the emailUpdateInContact field.
func (r *mutationResolver) EmailUpdateInContact(ctx context.Context, contactID string, input model.EmailUpdateInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailUpdateInContact", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID))

	result, err := r.Services.EmailService.UpdateEmailFor(ctx, entity.CONTACT, contactID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update email %s for contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromContact is the resolver for the EmailRemoveFromContact field.
func (r *mutationResolver) EmailRemoveFromContact(ctx context.Context, contactID string, email string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromContact", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID))

	result, err := r.Services.EmailService.DetachFromEntity(ctx, entity.CONTACT, contactID, email)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from contact %s", email, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailRemoveFromContactByID is the resolver for the emailRemoveFromContactById field.
func (r *mutationResolver) EmailRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromContactByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contactID", contactID), log.String("request.emailID", id))

	result, err := r.Services.EmailService.DetachFromEntityById(ctx, entity.CONTACT, contactID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from contact %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailMergeToUser is the resolver for the emailMergeToUser field.
func (r *mutationResolver) EmailMergeToUser(ctx context.Context, userID string, input model.EmailInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailMergeToUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	inputEmail := strings.TrimSpace(input.Email)

	if inputEmail == "" {
		graphql.AddErrorf(ctx, "Email address is required")
		return nil, nil
	}

	emailId, err := r.Services.EmailService.CreateEmailAddressByEvents(ctx, inputEmail, utils.IfNotNilString(input.AppSource))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Operation failed to create email %s", inputEmail)
		return nil, err
	}

	_, err = r.Clients.UserClient.LinkEmailToUser(ctx, &usergrpc.LinkEmailToUserGrpcRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		UserId:         userID,
		EmailId:        emailId,
		Primary:        utils.IfNotNilBool(input.Primary),
		Label:          utils.IfNotNilString(input.Label, func() string { return input.Label.String() }),
		LoggedInUserId: common.GetUserIdFromContext(ctx),
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not add email %s to user %s", inputEmail, userID)
		return nil, err
	}

	emailEntity, err := r.Services.EmailService.GetById(ctx, emailId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch email details %s", inputEmail)
		return nil, nil
	}

	return mapper.MapEntityToEmail(emailEntity), nil
}

// EmailUpdateInUser is the resolver for the emailUpdateInUser field.
func (r *mutationResolver) EmailUpdateInUser(ctx context.Context, userID string, input model.EmailUpdateInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailUpdateInUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	result, err := r.Services.EmailService.UpdateEmailFor(ctx, entity.USER, userID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update email %s for user %s", input.ID, userID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromUser is the resolver for the emailRemoveFromUser field.
func (r *mutationResolver) EmailRemoveFromUser(ctx context.Context, userID string, email string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromUser", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID))

	result, err := r.Services.EmailService.DetachFromEntity(ctx, entity.USER, userID, email)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from user %s", email, userID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailRemoveFromUserByID is the resolver for the emailRemoveFromUserById field.
func (r *mutationResolver) EmailRemoveFromUserByID(ctx context.Context, userID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromUserByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", userID), log.String("request.emailID", id))

	result, err := r.Services.EmailService.DetachFromEntityById(ctx, entity.USER, userID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from user %s", id, userID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailMergeToOrganization is the resolver for the emailMergeToOrganization field.
func (r *mutationResolver) EmailMergeToOrganization(ctx context.Context, organizationID string, input model.EmailInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailMergeToOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.EmailService.MergeEmailTo(ctx, entity.ORGANIZATION, organizationID, mapper.MapEmailInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not add email %s to organization %s", input.Email, organizationID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailUpdateInOrganization is the resolver for the emailUpdateInOrganization field.
func (r *mutationResolver) EmailUpdateInOrganization(ctx context.Context, organizationID string, input model.EmailUpdateInput) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailUpdateInOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.EmailService.UpdateEmailFor(ctx, entity.ORGANIZATION, organizationID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not update email %s for organization %s", input.ID, organizationID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromOrganization is the resolver for the emailRemoveFromOrganization field.
func (r *mutationResolver) EmailRemoveFromOrganization(ctx context.Context, organizationID string, email string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID))

	result, err := r.Services.EmailService.DetachFromEntity(ctx, entity.ORGANIZATION, organizationID, email)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from organization %s", email, organizationID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailRemoveFromOrganizationByID is the resolver for the emailRemoveFromOrganizationById field.
func (r *mutationResolver) EmailRemoveFromOrganizationByID(ctx context.Context, organizationID string, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailRemoveFromOrganizationByID", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.organizationID", organizationID), log.String("request.emailID", id))

	result, err := r.Services.EmailService.DetachFromEntityById(ctx, entity.ORGANIZATION, organizationID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s from organization %s", id, organizationID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailDelete is the resolver for the emailDelete field.
func (r *mutationResolver) EmailDelete(ctx context.Context, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.EmailDelete", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.emailID", id))

	result, err := r.Services.EmailService.DeleteById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Could not remove email %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// Email is the resolver for the email field.
func (r *queryResolver) Email(ctx context.Context, id string) (*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Email", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.emailID", id))

	emailEntity, err := r.Services.EmailService.GetById(ctx, id)
	if err != nil || emailEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed getting email with id %s", id)
		return nil, nil
	}
	return mapper.MapEntityToEmail(emailEntity), nil
}

// Email returns generated.EmailResolver implementation.
func (r *Resolver) Email() generated.EmailResolver { return &emailResolver{r} }

type emailResolver struct{ *Resolver }
