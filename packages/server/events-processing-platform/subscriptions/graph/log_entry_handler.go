package graph

import (
	"context"
	neo4jentity "github.com/openline-ai/customer-os-neo4j-repository/entity"
	neo4jmodel "github.com/openline-ai/customer-os-neo4j-repository/model"
	neo4jrepository "github.com/openline-ai/customer-os-neo4j-repository/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	organizationpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/organization"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type LogEntryEventHandler struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
}

func NewLogEntryEventHandler(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients) *LogEntryEventHandler {
	return &LogEntryEventHandler{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
	}
}

func (h *LogEntryEventHandler) OnCreate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogEntryEventHandler.OnCreate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.LogEntryCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	logEntryId := aggregate.GetLogEntryObjectID(evt.AggregateID, eventData.Tenant)
	data := neo4jrepository.LogEntryCreateFields{
		Content:              eventData.Content,
		ContentType:          eventData.ContentType,
		StartedAt:            eventData.StartedAt,
		AuthorUserId:         eventData.AuthorUserId,
		LoggedOrganizationId: eventData.LoggedOrganizationId,
		SourceFields: neo4jmodel.Source{
			Source:        helper.GetSource(eventData.Source),
			SourceOfTruth: helper.GetSourceOfTruth(eventData.SourceOfTruth),
			AppSource:     helper.GetAppSource(eventData.AppSource),
		},
		CreatedAt: eventData.CreatedAt,
		UpdatedAt: eventData.UpdatedAt,
	}
	err := h.repositories.Neo4jRepositories.LogEntryWriteRepository.Create(ctx, eventData.Tenant, logEntryId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while saving log entry %s: %s", logEntryId, err.Error())
		return err
	}

	if eventData.ExternalSystem.Available() {
		err = h.repositories.ExternalSystemRepository.LinkWithEntity(ctx, eventData.Tenant, logEntryId, neo4jentity.NodeLabel_LogEntry, eventData.ExternalSystem)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Error while link log entry %s with external system %s: %s", logEntryId, eventData.ExternalSystem.ExternalSystemId, err.Error())
			return err
		}
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err = h.grpcClients.OrganizationClient.RefreshLastTouchpoint(ctx, &organizationpb.OrganizationIdGrpcRequest{
		Tenant:         eventData.Tenant,
		OrganizationId: eventData.LoggedOrganizationId,
		AppSource:      constants.AppSourceEventProcessingPlatform,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while refreshing last touchpoint for organization %s: %s", eventData.LoggedOrganizationId, err.Error())
	}

	return nil
}

func (h *LogEntryEventHandler) OnUpdate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogEntryEventHandler.OnUpdate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.LogEntryUpdateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	logEntryId := aggregate.GetLogEntryObjectID(evt.AggregateID, eventData.Tenant)
	data := neo4jrepository.LogEntryUpdateFields{
		Content:              eventData.Content,
		ContentType:          eventData.ContentType,
		StartedAt:            eventData.StartedAt,
		LoggedOrganizationId: eventData.LoggedOrganizationId,
		UpdatedAt:            eventData.UpdatedAt,
		Source:               helper.GetSource(eventData.SourceOfTruth),
	}
	err := h.repositories.Neo4jRepositories.LogEntryWriteRepository.Update(ctx, eventData.Tenant, logEntryId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while saving log entry %s: %s", logEntryId, err.Error())
	}

	return err
}

func (h *LogEntryEventHandler) OnAddTag(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogEntryEventHandler.OnAddTag")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.LogEntryAddTagEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	logEntryId := aggregate.GetLogEntryObjectID(evt.AggregateID, eventData.Tenant)
	err := h.repositories.Neo4jRepositories.TagWriteRepository.LinkTagByIdToEntity(ctx, eventData.Tenant, eventData.TagId, logEntryId, "LogEntry", eventData.TaggedAt)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while adding tag %s to log entry %s: %s", eventData.TagId, logEntryId, err.Error())
	}

	return err
}

func (h *LogEntryEventHandler) OnRemoveTag(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogEntryEventHandler.OnRemoveTag")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.LogEntryRemoveTagEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	logEntryId := aggregate.GetLogEntryObjectID(evt.AggregateID, eventData.Tenant)
	err := h.repositories.Neo4jRepositories.TagWriteRepository.UnlinkTagByIdFromEntity(ctx, eventData.Tenant, eventData.TagId, logEntryId, "LogEntry")
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while removing tag %s to log entry %s: %s", eventData.TagId, logEntryId, err.Error())
	}

	return err
}
