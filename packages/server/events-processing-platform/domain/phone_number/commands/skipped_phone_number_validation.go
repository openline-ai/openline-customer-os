package commands

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type SkippedPhoneNumberValidationCommandHandler interface {
	Handle(ctx context.Context, command *SkippedPhoneNumberValidationCommand) error
}

type skippedPhoneNumberValidationCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  eventstore.AggregateStore
}

func NewSkippedPhoneNumberValidationCommandHandler(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) SkippedPhoneNumberValidationCommandHandler {
	return &skippedPhoneNumberValidationCommandHandler{log: log, cfg: cfg, es: es}
}

func (h *skippedPhoneNumberValidationCommandHandler) Handle(ctx context.Context, command *SkippedPhoneNumberValidationCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SkippedPhoneNumberValidationCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("Tenant", command.Tenant), log.String("ObjectID", command.ObjectID))

	phoneNumberAggregate, err := aggregate.LoadPhoneNumberAggregate(ctx, h.es, command.Tenant, command.ObjectID)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	if err = phoneNumberAggregate.SkippedPhoneNumberValidation(ctx, command.Tenant, command.RawPhoneNumber, command.CountryCodeA2, command.ValidationSkipReason); err != nil {
		return err
	}
	return h.es.Save(ctx, phoneNumberAggregate)
}
