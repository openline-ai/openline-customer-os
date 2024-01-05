package graph

import (
	"context"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jmodel "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/model"
	neo4jrepository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/interaction_session/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/interaction_session/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type InteractionSessionEventHandler struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
}

func NewInteractionSessionEventHandler(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients) *InteractionSessionEventHandler {
	return &InteractionSessionEventHandler{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
	}
}

func (h *InteractionSessionEventHandler) OnCreate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionSessionEventHandler.OnCreate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.InteractionSessionCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	span.SetTag(tracing.SpanTagTenant, eventData.Tenant)

	interactionSessionId := aggregate.GetInteractionSessionObjectID(evt.AggregateID, eventData.Tenant)
	span.SetTag(tracing.SpanTagEntityId, interactionSessionId)

	data := neo4jrepository.InteractionSessionCreateFields{
		CreatedAt: eventData.CreatedAt,
		UpdatedAt: eventData.UpdatedAt,
		SourceFields: neo4jmodel.Source{
			Source:    helper.GetSource(eventData.Source),
			AppSource: helper.GetAppSource(eventData.AppSource),
		},
		Channel:     eventData.Channel,
		ChannelData: eventData.ChannelData,
		Identifier:  eventData.Identifier,
		Type:        eventData.Type,
		Status:      eventData.Status,
		Name:        eventData.Name,
	}
	err := h.repositories.Neo4jRepositories.InteractionSessionWriteRepository.Create(ctx, eventData.Tenant, interactionSessionId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while saving interaction session %s: %s", interactionSessionId, err.Error())
		return err
	}

	if eventData.ExternalSystem.Available() {
		err = h.repositories.ExternalSystemRepository.LinkWithEntity(ctx, eventData.Tenant, interactionSessionId, neo4jentity.NodeLabelInteractionSession, eventData.ExternalSystem)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Error while link interaction session %s with external system %s: %s", interactionSessionId, eventData.ExternalSystem.ExternalSystemId, err.Error())
			return err
		}
	}

	return nil
}
