package aggregate

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func (a *ContactAggregate) CreateContact(ctx context.Context, contactDto *models.ContactDto) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ContactAggregate.CreateContact")
	defer span.Finish()
	span.LogFields(log.String("Tenant", contactDto.Tenant), log.String("AggregateID", a.GetID()))

	createdAtNotNil := utils.IfNotNilTimeWithDefault(contactDto.CreatedAt, utils.Now())
	updatedAtNotNil := utils.IfNotNilTimeWithDefault(contactDto.UpdatedAt, createdAtNotNil)
	event, err := events.NewContactCreatedEvent(a, contactDto, createdAtNotNil, updatedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewContactCreatedEvent")
	}

	if err = event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *ContactAggregate) UpdateContact(ctx context.Context, contactDto *models.ContactDto) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ContactAggregate.UpdateContact")
	defer span.Finish()
	span.LogFields(log.String("Tenant", contactDto.Tenant), log.String("AggregateID", a.GetID()))

	updatedAtNotNil := utils.IfNotNilTimeWithDefault(contactDto.UpdatedAt, utils.Now())
	if contactDto.Source.SourceOfTruth == "" {
		contactDto.Source.SourceOfTruth = a.Contact.Source.SourceOfTruth
	}

	event, err := events.NewContactUpdatedEvent(a, contactDto, updatedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewContactUpdatedEvent")
	}

	// FIXME alexb check what type of metadata should be set into event and apply it to all aggregation commands
	if err = event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
