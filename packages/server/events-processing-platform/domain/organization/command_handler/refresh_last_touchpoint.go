package command_handler

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"time"
)

type RefreshLastTouchpointCommandHandler interface {
	Handle(ctx context.Context, command *command.RefreshLastTouchpointCommand) error
}

type refreshLastTouchpointCommandHandler struct {
	log logger.Logger
	es  eventstore.AggregateStore
	cfg config.Utils
}

func NewRefreshLastTouchpointCommandHandler(log logger.Logger, es eventstore.AggregateStore, cfg config.Utils) RefreshLastTouchpointCommandHandler {
	return &refreshLastTouchpointCommandHandler{log: log, es: es, cfg: cfg}
}

func (h *refreshLastTouchpointCommandHandler) Handle(ctx context.Context, cmd *command.RefreshLastTouchpointCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RefreshLastTouchpointCommandHandler.Handle")
	defer span.Finish()
	tracing.SetCommandHandlerSpanTags(ctx, span, cmd.Tenant, cmd.LoggedInUserId)
	span.LogFields(log.Object("command", cmd))

	validationError, done := validator.Validate(cmd, span)
	if done {
		return validationError
	}

	for attempt := 0; attempt == 0 || attempt < h.cfg.RetriesOnOptimisticLockException; attempt++ {
		organizationAggregate, err := aggregate.LoadOrganizationAggregate(ctx, h.es, cmd.Tenant, cmd.ObjectID)
		if err != nil {
			return err
		}
		if err = organizationAggregate.HandleCommand(ctx, cmd); err != nil {
			return err
		}

		err = h.es.Save(ctx, organizationAggregate)
		if err == nil {
			return nil // Save successful
		}

		if eventstore.IsEventStoreErrorCodeWrongExpectedVersion(err) {
			// Handle concurrency error
			span.LogFields(log.Int("retryAttempt", attempt+1))
			time.Sleep(helper.BackoffDelay(attempt)) // backoffDelay is a function that increases the delay with each attempt
			continue                                 // Retry
		} else {
			// Some other error occurred
			tracing.TraceErr(span, err)
			return err
		}
	}

	err := errors.New("reached maximum number of retries")
	tracing.TraceErr(span, err)
	return err
}
