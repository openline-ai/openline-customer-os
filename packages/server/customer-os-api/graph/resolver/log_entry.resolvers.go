package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	logentrygrpc "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/log_entry"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreatedBy is the resolver for the createdBy field.
func (r *logEntryResolver) CreatedBy(ctx context.Context, obj *model.LogEntry) (*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntityNillable, err := dataloader.For(ctx).GetUserAuthorForLogEntry(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Error fetching user author for log entry %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Error fetching user author for log entry %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToUser(userEntityNillable), nil
}

// Tags is the resolver for the tags field.
func (r *logEntryResolver) Tags(ctx context.Context, obj *model.LogEntry) ([]*model.Tag, error) {
	graphql.AddErrorf(ctx, "Not implemented: LogEntry.Tags")
	return nil, nil
}

// LogEntryCreateForOrganization is the resolver for the logEntry_CreateForOrganization field.
func (r *mutationResolver) LogEntryCreateForOrganization(ctx context.Context, organizationID string, input model.LogEntryInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryCreateForOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("organizationID", organizationID), log.Object("input", input))

	organizationEntity, err := r.Services.OrganizationService.GetById(ctx, organizationID)
	if err != nil || organizationEntity == nil {
		if err == nil {
			err = fmt.Errorf("organization not found")
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Organization not found")
		return "", nil
	}

	response, err := r.Clients.LogEntryClient.UpsertLogEntry(ctx, &logentrygrpc.UpsertLogEntryGrpcRequest{
		Tenant:               common.GetTenantFromContext(ctx),
		UserId:               common.GetUserIdFromContext(ctx),
		Content:              utils.IfNotNilString(input.Content),
		ContentType:          utils.IfNotNilString(input.ContentType),
		StartedAt:            timestamppb.New(utils.IfNotNilTimeWithDefault(input.StartedAt, utils.Now())),
		AppSource:            constants.AppSourceCustomerOsApi,
		Source:               string(entity.DataSourceOpenline),
		SourceOfTruth:        string(entity.DataSourceOpenline),
		LoggedOrganizationId: utils.StringPtr(organizationID),
		AuthorUserId:         utils.StringPtr(common.GetUserIdFromContext(ctx)),
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error creating log entry")
		return "", nil
	}
	return response.Id, nil
}

// LogEntryUpdate is the resolver for the logEntry_Update field.
func (r *mutationResolver) LogEntryUpdate(ctx context.Context, id string, input model.LogEntryUpdateInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("logEntryId", id), log.Object("input", input))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil || logEntryEntity == nil {
		if err == nil {
			err = fmt.Errorf("Log entry %s not found", id)
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Log entry %s not found", id)
		return "", nil
	}
	grpcRequestMessage := logentrygrpc.UpsertLogEntryGrpcRequest{
		Id:            id,
		Tenant:        common.GetTenantFromContext(ctx),
		UserId:        common.GetUserIdFromContext(ctx),
		Content:       utils.IfNotNilString(input.Content),
		ContentType:   utils.IfNotNilString(input.ContentType),
		SourceOfTruth: string(entity.DataSourceOpenline),
	}
	if input.StartedAt != nil {
		grpcRequestMessage.StartedAt = timestamppb.New(*input.StartedAt)
	}

	response, err := r.Clients.LogEntryClient.UpsertLogEntry(ctx, &grpcRequestMessage)

	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error updating log entry")
		return "", nil
	}
	return response.Id, nil
}

// LogEntry returns generated.LogEntryResolver implementation.
func (r *Resolver) LogEntry() generated.LogEntryResolver { return &logEntryResolver{r} }

type logEntryResolver struct{ *Resolver }
