package graph

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	neo4jentity "github.com/openline-ai/customer-os-neo4j-repository/entity"
	neo4jmodel "github.com/openline-ai/customer-os-neo4j-repository/model"
	neo4jrepository "github.com/openline-ai/customer-os-neo4j-repository/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type ContactEventHandler struct {
	log          logger.Logger
	repositories *repository.Repositories
}

func NewContactEventHandler(log logger.Logger, repositories *repository.Repositories) *ContactEventHandler {
	return &ContactEventHandler{
		log:          log,
		repositories: repositories,
	}
}

func (h *ContactEventHandler) OnContactCreate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnContactCreate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)

	session := utils.NewNeo4jWriteSession(ctx, *h.repositories.Drivers.Neo4jDriver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		var err error

		data := neo4jrepository.ContactCreateFields{
			FirstName:       eventData.FirstName,
			LastName:        eventData.LastName,
			Prefix:          eventData.Prefix,
			Description:     eventData.Description,
			Timezone:        eventData.Timezone,
			ProfilePhotoUrl: eventData.ProfilePhotoUrl,
			Name:            eventData.Name,
			CreatedAt:       eventData.CreatedAt,
			UpdatedAt:       eventData.UpdatedAt,
			SourceFields: neo4jmodel.Source{
				Source:        helper.GetSource(eventData.Source),
				SourceOfTruth: helper.GetSourceOfTruth(eventData.SourceOfTruth),
				AppSource:     helper.GetAppSource(eventData.AppSource),
			},
		}
		err = h.repositories.Neo4jRepositories.ContactWriteRepository.CreateContactInTx(ctx, tx, eventData.Tenant, contactId, data)
		if err != nil {
			h.log.Errorf("Error while saving contact %s: %s", contactId, err.Error())
			return nil, err
		}
		if eventData.ExternalSystem.Available() {
			err = h.repositories.ExternalSystemRepository.LinkWithEntityInTx(ctx, tx, eventData.Tenant, contactId, neo4jentity.NodeLabel_Contact, eventData.ExternalSystem)
			if err != nil {
				h.log.Errorf("Error while link contact %s with external system %s: %s", contactId, eventData.ExternalSystem.ExternalSystemId, err.Error())
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	return nil
}

func (h *ContactEventHandler) OnContactUpdate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnContactUpdate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactUpdateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)
	data := neo4jrepository.ContactUpdateFields{
		FirstName:       eventData.FirstName,
		LastName:        eventData.LastName,
		Prefix:          eventData.Prefix,
		Description:     eventData.Description,
		Timezone:        eventData.Timezone,
		ProfilePhotoUrl: eventData.ProfilePhotoUrl,
		Name:            eventData.Name,
		UpdatedAt:       eventData.UpdatedAt,
		Source:          eventData.Source,
	}
	err := h.repositories.Neo4jRepositories.ContactWriteRepository.UpdateContact(ctx, eventData.Tenant, contactId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while saving contact %s: %s", contactId, err.Error())
		return err
	}

	if eventData.ExternalSystem.Available() {
		session := utils.NewNeo4jWriteSession(ctx, *h.repositories.Drivers.Neo4jDriver)
		defer session.Close(ctx)

		_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			//var err error
			if eventData.ExternalSystem.Available() {
				innerErr := h.repositories.ExternalSystemRepository.LinkWithEntityInTx(ctx, tx, eventData.Tenant, contactId, neo4jentity.NodeLabel_Contact, eventData.ExternalSystem)
				if innerErr != nil {
					h.log.Errorf("Error while link contact %s with external system %s: %s", contactId, eventData.ExternalSystem.ExternalSystemId, err.Error())
					return nil, innerErr
				}
			}
			return nil, nil
		})
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}
	}

	return nil
}

func (e *ContactEventHandler) OnPhoneNumberLinkToContact(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnPhoneNumberLinkToContact")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactLinkPhoneNumberEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)
	err := e.repositories.PhoneNumberRepository.LinkWithContact(ctx, eventData.Tenant, contactId, eventData.PhoneNumberId, eventData.Label, eventData.Primary, eventData.UpdatedAt)

	return err
}

func (h *ContactEventHandler) OnEmailLinkToContact(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnEmailLinkToContact")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactLinkEmailEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)
	err := h.repositories.EmailRepository.LinkWithContact(ctx, eventData.Tenant, contactId, eventData.EmailId, eventData.Label, eventData.Primary, eventData.UpdatedAt)

	return err
}

func (h *ContactEventHandler) OnLocationLinkToContact(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnLocationLinkToContact")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactLinkLocationEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)
	err := h.repositories.LocationRepository.LinkWithContact(ctx, eventData.Tenant, contactId, eventData.LocationId, eventData.UpdatedAt)

	return err
}

func (h *ContactEventHandler) OnContactLinkToOrganization(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContactEventHandler.OnContactLinkToOrganization")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ContactLinkWithOrganizationEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	contactId := aggregate.GetContactObjectID(evt.AggregateID, eventData.Tenant)
	err := h.repositories.JobRoleRepository.LinkContactWithOrganization(ctx, eventData.Tenant, contactId, eventData)

	return err
}
