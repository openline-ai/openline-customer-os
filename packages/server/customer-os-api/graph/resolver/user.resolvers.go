package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))

	createdUserEntity, err := r.Services.UserService.Create(ctx, &service.UserCreateData{
		UserEntity:   mapper.MapUserInputToEntity(input),
		EmailEntity:  mapper.MapEmailInputToEntity(input.Email),
		PlayerEntity: mapper.MapPlayerInputToEntity(input.Player),
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create user %s %s", input.FirstName, input.LastName)
		return nil, err
	}
	return mapper.MapEntityToUser(createdUserEntity), nil
}

// UserCreateInTenant is the resolver for the user_CreateInTenant field.
func (r *mutationResolver) UserCreateInTenant(ctx context.Context, input model.UserInput, tenant string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserCreateInTenant - user_CreateInTenant"))
}

// UserUpdate is the resolver for the user_Update field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UserUpdateInput) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.userID", input.ID))

	updatedUserEntity, err := r.Services.UserService.Update(ctx, mapper.MapUserUpdateInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update user %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToUser(updatedUserEntity), nil
}

// UserAddRole is the resolver for the user_AddRole field.
func (r *mutationResolver) UserAddRole(ctx context.Context, id string, role model.Role) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserAddRole", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.userID", id))

	userResult, err := r.Services.UserService.AddRole(ctx, id, role)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to add role %s to user %s", role, id)
		return nil, err
	}
	return mapper.MapEntityToUser(userResult), nil
}

// UserRemoveRole is the resolver for the user_RemoveRole field.
func (r *mutationResolver) UserRemoveRole(ctx context.Context, id string, role model.Role) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserRemoveRole", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.userID", id))

	userResult, err := r.Services.UserService.DeleteRole(ctx, id, role)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to remove role %s from user %s", role, id)
		return nil, err
	}
	return mapper.MapEntityToUser(userResult), nil
}

// UserAddRoleInTenant is the resolver for the user_AddRoleInTenant field.
func (r *mutationResolver) UserAddRoleInTenant(ctx context.Context, id string, tenant string, role model.Role) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserAddRoleInTenant", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, tenant)
	span.LogFields(log.String("request.userID", id))

	userResult, err := r.Services.UserService.AddRoleInTenant(ctx, id, tenant, role)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to add role %s to user %s in tenant %s", role, id, tenant)
		return nil, err
	}
	return mapper.MapEntityToUser(userResult), nil
}

// UserRemoveRoleInTenant is the resolver for the user_RemoveRoleInTenant field.
func (r *mutationResolver) UserRemoveRoleInTenant(ctx context.Context, id string, tenant string, role model.Role) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserRemoveRoleInTenant", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, tenant)
	span.LogFields(log.String("request.userID", id))

	userResult, err := r.Services.UserService.DeleteRoleInTenant(ctx, id, tenant, role)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to remove role %s from user %s in tenant %s", role, id, tenant)
		return nil, err
	}
	return mapper.MapEntityToUser(userResult), nil
}

// UserDelete is the resolver for the user_Delete field.
func (r *mutationResolver) UserDelete(ctx context.Context, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: UserDelete - user_Delete"))
}

// UserDeleteInTenant is the resolver for the user_DeleteInTenant field.
func (r *mutationResolver) UserDeleteInTenant(ctx context.Context, id string, tenant string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: UserDeleteInTenant - user_DeleteInTenant"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.UserPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Users", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	span.LogFields(log.Int("request.page", pagination.Page), log.Int("request.limit", pagination.Limit))
	paginatedResult, err := r.Services.UserService.GetAll(ctx, pagination.Page, pagination.Limit, where, sort)
	return &model.UserPage{
		Content:       mapper.MapEntitiesToUsers(paginatedResult.Rows.(*entity.UserEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.User", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.userID", id))

	userEntity, err := r.Services.UserService.FindUserById(ctx, id)
	if err != nil || userEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "User with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToUser(userEntity), nil
}

// UserByEmail is the resolver for the user_ByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.UserByEmail", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.email", email))

	userEntity, err := r.Services.UserService.FindUserByEmail(ctx, email)
	if err != nil || userEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "User with email %s not identified", email)
		return nil, err
	}
	return mapper.MapEntityToUser(userEntity), nil
}

// Player is the resolver for the player field.
func (r *userResolver) Player(ctx context.Context, obj *model.User) (*model.Player, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.Player", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.user", obj.ID))

	playerEntity, err := r.Services.PlayerService.GetPlayerForUser(ctx, common.GetContext(ctx).Tenant, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get player for user %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntityToPlayer(playerEntity), nil
}

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]model.Role, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.Roles", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.user", obj.ID))

	return obj.Roles, nil
}

// Emails is the resolver for the emails field.
func (r *userResolver) Emails(ctx context.Context, obj *model.User) ([]*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.Emails", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.user", obj.ID))

	emailEntities, err := r.Services.EmailService.GetAllFor(ctx, entity.USER, obj.ID)
	return mapper.MapEntitiesToEmails(emailEntities), err
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *userResolver) PhoneNumbers(ctx context.Context, obj *model.User) ([]*model.PhoneNumber, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.PhoneNumbers", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.user", obj.ID))

	phoneNumberEntities, err := dataloader.For(ctx).GetPhoneNumbersForUser(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get phone numbers for user %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToPhoneNumbers(phoneNumberEntities), nil
}

// Conversations is the resolver for the conversations field.
func (r *userResolver) Conversations(ctx context.Context, obj *model.User, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.Conversations", graphql.GetOperationContext(ctx))
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, common.GetTenantFromContext(ctx))
	span.LogFields(log.String("request.user", obj.ID))

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ConversationService.GetConversationsForUser(ctx, obj.ID, pagination.Page, pagination.Limit, sort)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get user %s conversations", obj.ID)
		return nil, err
	}
	return &model.ConversationPage{
		Content:       mapper.MapEntitiesToConversations(paginatedResult.Rows.(*entity.ConversationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
