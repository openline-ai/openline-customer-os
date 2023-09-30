package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	commongrpc "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/common"
	usergrpc "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/user"
	"github.com/opentracing/opentracing-go/log"
)

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	userId, err := r.Services.UserService.CreateUserByEvents(ctx, *mapper.MapUserInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create user %s %s", input.FirstName, input.LastName)
		return nil, nil
	}

	if input.Player != nil {
		_, err = r.Clients.UserClient.AddPlayerInfo(ctx, &usergrpc.AddPlayerInfoGrpcRequest{
			UserId:         userId,
			Tenant:         common.GetTenantFromContext(ctx),
			LoggedInUserId: common.GetUserIdFromContext(ctx),
			AuthId:         input.Player.AuthID,
			Provider:       input.Player.Provider,
			IdentityId:     utils.IfNotNilString(input.Player.IdentityID),
			SourceFields: &commongrpc.SourceFields{
				Source:    string(entity.DataSourceOpenline),
				AppSource: utils.IfNotNilStringWithDefault(input.AppSource, constants.AppSourceCustomerOsApi),
			},
		})
		if err != nil {
			tracing.TraceErr(span, err)
			r.log.Errorf("Failed to add player info for user %s: %s", userId, err.Error())
		}
	}

	if input.Email != nil {
		emailId, err := r.Services.EmailService.CreateEmailAddressByEvents(ctx, input.Email.Email, input.AppSource)
		if err != nil {
			tracing.TraceErr(span, err)
			r.log.Errorf("Failed to create email address for user %s: %s", userId, err.Error())
		}
		if emailId != "" {
			_, err = r.Clients.UserClient.LinkEmailToUser(ctx, &usergrpc.LinkEmailToUserGrpcRequest{
				Tenant:  common.GetTenantFromContext(ctx),
				UserId:  userId,
				EmailId: emailId,
				Primary: utils.IfNotNilBool(input.Email.Primary),
				Label:   utils.IfNotNilString(input.Email.Label, func() string { return input.Email.Label.String() }),
			})
			if err != nil {
				tracing.TraceErr(span, err)
				r.log.Errorf("Failed to link email address %s to user %s: %s", emailId, userId, err.Error())
			}
		}
	}

	createdUserEntity, err := r.Services.UserService.GetById(ctx, userId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "User details not yet available. User id: %s", userId)
		return nil, nil
	}
	return mapper.MapEntityToUser(createdUserEntity), nil
}

// UserUpdate is the resolver for the user_Update field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UserUpdateInput) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", input.ID))

	if input.ID != common.GetContext(ctx).UserId {
		if !r.Services.UserService.ContainsRole(ctx, []model.Role{model.RoleAdmin, model.RoleCustomerOsPlatformOwner, model.RoleOwner}) {
			return nil, fmt.Errorf("user can not update other user")
		}
	}

	_, err := r.Clients.UserClient.UpsertUser(ctx, &usergrpc.UpsertUserGrpcRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		Id:             input.ID,
		SourceFields: &commongrpc.SourceFields{
			Source: string(entity.DataSourceOpenline),
		},
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		Name:            utils.IfNotNilString(input.Name),
		Timezone:        utils.IfNotNilString(input.Timezone),
		ProfilePhotoUrl: utils.IfNotNilString(input.ProfilePhotoURL),
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update user %s", input.ID)
		return nil, nil
	}

	updatedUserEntity, err := r.Services.UserService.GetById(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch user %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToUser(updatedUserEntity), nil
}

// UserAddRole is the resolver for the user_AddRole field.
func (r *mutationResolver) UserAddRole(ctx context.Context, id string, role model.Role) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.UserAddRole", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
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

// CustomerUserAddJobRole is the resolver for the customer_user_AddJobRole field.
func (r *mutationResolver) CustomerUserAddJobRole(ctx context.Context, id string, jobRoleInput model.JobRoleInput) (*model.CustomerUser, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.CustomerUserAddJobRole", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	role, err := r.Services.UserService.CustomerAddJobRole(ctx, &service.CustomerAddJobRoleData{
		UserId:        id,
		JobRoleEntity: mapper.MapJobRoleInputToEntity(&jobRoleInput),
	})
	return role, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.UserPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Users", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

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
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.userID", id))

	userEntity, err := r.Services.UserService.GetById(ctx, id)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
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
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.user", obj.ID))

	return obj.Roles, nil
}

// Emails is the resolver for the emails field.
func (r *userResolver) Emails(ctx context.Context, obj *model.User) ([]*model.Email, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "UserResolver.Emails", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.user", obj.ID))

	emailEntities, err := r.Services.EmailService.GetAllFor(ctx, entity.USER, obj.ID)
	return mapper.MapEntitiesToEmails(emailEntities), err
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *userResolver) PhoneNumbers(ctx context.Context, obj *model.User) ([]*model.PhoneNumber, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	phoneNumberEntities, err := dataloader.For(ctx).GetPhoneNumbersForUser(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get phone numbers for user %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get phone numbers for user %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToPhoneNumbers(phoneNumberEntities), nil
}

// JobRoles is the resolver for the jobRoles field.
func (r *userResolver) JobRoles(ctx context.Context, obj *model.User) ([]*model.JobRole, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	jobRoleEntities, err := dataloader.For(ctx).GetJobRolesForUser(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get job roles for user %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get job roles for user %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToJobRoles(jobRoleEntities), err
}

// Calendars is the resolver for the calendars field.
func (r *userResolver) Calendars(ctx context.Context, obj *model.User) ([]*model.Calendar, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	calendarsForUser, err := dataloader.For(ctx).GetCalendarsForUser(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get calendars for user %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get job roles for user %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToCalendars(calendarsForUser), err
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UserCreateInTenant(ctx context.Context, input model.UserInput, tenant string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserCreateInTenant - user_CreateInTenant"))
}
