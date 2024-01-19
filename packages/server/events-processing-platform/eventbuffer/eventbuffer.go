package eventbuffer

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	orgaggregate "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	orgevents "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository/postgres/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
)

type EventBufferWatcher struct {
	repositories  *repository.Repositories
	logger        logger.Logger
	es            eventstore.AggregateStore
	signalChannel chan os.Signal
	ticker        *time.Ticker
}

func NewEventBufferWatcher(repositories *repository.Repositories, logger logger.Logger, es eventstore.AggregateStore) *EventBufferWatcher {
	return &EventBufferWatcher{repositories: repositories, logger: logger, es: es}
}

// Start starts the EventBufferWatcher
func (eb *EventBufferWatcher) Start(ctx context.Context) {
	eb.logger.Info("EventBufferWatcher started")

	eb.ticker = time.NewTicker(time.Second * 30)
	eb.signalChannel = make(chan os.Signal, 1)
	signal.Notify(eb.signalChannel, syscall.SIGTERM, syscall.SIGINT)

	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				// Run dispatch logic every n seconds
				eb.logger.Info("EventBufferWatcher.Dispatch: dispatch buffered events")
				err := eb.Dispatch(ctx)
				if err != nil {
					eb.logger.Errorf("EventBufferWatcher.Dispatch: error dispatching events: %s", err.Error())
				}
			case <-eb.signalChannel:
				// Shutdown goroutine
				eb.logger.Info("EventBufferWatcher.Dispatch: Got signal, exiting...")
				runtime.Goexit()
			}
		}
	}(eb.ticker)
}

// Stop stops the EventBufferWatcher
func (eb *EventBufferWatcher) Stop() {
	eb.signalChannel <- syscall.SIGTERM // TODO get the signal from the caller
	eb.ticker.Stop()
	eb.logger.Info("EventBufferWatcher stopped")
	close(eb.signalChannel)
	eb.signalChannel = nil
}

func (eb *EventBufferWatcher) Park(
	ctx context.Context,
	evt eventstore.Event,
	tenant string,
	uuid string,
	expiryTimestamp time.Time,
) error {
	eventBuffer := entity.EventBuffer{
		Tenant:             tenant,
		UUID:               uuid,
		ExpiryTimestamp:    expiryTimestamp,
		EventID:            evt.EventID,
		EventType:          evt.EventType,
		EventData:          evt.Data,
		EventTimestamp:     evt.Timestamp,
		EventAggregateID:   evt.AggregateID,
		EventAggregateType: string(evt.AggregateType),
		EventVersion:       evt.Version,
		EventMetadata:      evt.Metadata,
	}
	return eb.repositories.EventBufferRepository.Upsert(eventBuffer)
}

// Dispatch dispatches all expired events from event_buffer table, and delete them after dispatching
func (eb *EventBufferWatcher) Dispatch(ctx context.Context) error {
	now := time.Now().UTC()
	eventBuffers, err := eb.repositories.EventBufferRepository.GetByExpired(now)
	if err != nil {
		return err
	}
	if len(eventBuffers) == 0 {
		return nil
	}
	for _, eventBuffer := range eventBuffers {
		err := eb.HandleEvent(ctx, eventBuffer)
		if err != nil {
			return err
		}
		err = eb.repositories.EventBufferRepository.Delete(eventBuffer)
		if err != nil {
			return err
		}
	}
	return err
}

// DispatchByUUID dispatches the event with the given uuid from event_buffer table, and delete it after dispatching
func (eb *EventBufferWatcher) DispatchByUUID(ctx context.Context, uuid string) error {
	eventBuffer, err := eb.repositories.EventBufferRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}
	err = eb.HandleEvent(ctx, eventBuffer)
	if err != nil {
		return err
	}
	err = eb.repositories.EventBufferRepository.Delete(eventBuffer)
	return err
}

// HandleEvent loads the event aggregate and applies the event to it and pushes it into event store
func (eb *EventBufferWatcher) HandleEvent(ctx context.Context, eventBuffer entity.EventBuffer) error {
	evt := eventstore.Event{
		EventID:       eventBuffer.EventID,
		EventType:     eventBuffer.EventType,
		Data:          eventBuffer.EventData,
		Timestamp:     eventBuffer.EventTimestamp.UTC(),
		AggregateType: eventstore.AggregateType(eventBuffer.EventAggregateType),
		AggregateID:   eventBuffer.EventAggregateID,
		Version:       eventBuffer.EventVersion,
		Metadata:      eventBuffer.EventMetadata,
	}
	return eb.handleEvent(ctx, evt)
}

func (eb *EventBufferWatcher) handleEvent(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EventBufferWatcher.handleEvent")
	defer span.Finish()
	switch evt.EventType {
	case orgevents.OrganizationUpdateOwnerV1:
		var data orgevents.OrganizationOwnerUpdateEvent
		json.Unmarshal(evt.Data, &data)
		organizationAggregate, err := orgaggregate.LoadOrganizationAggregate(ctx, eb.es, data.Tenant, data.OrganizationId, eventstore.LoadAggregateOptions{})
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}
		err = organizationAggregate.Apply(evt)
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}
		// Persist the changes to the event store
		if err = eb.es.Save(ctx, organizationAggregate); err != nil {
			tracing.TraceErr(span, err)
			return err
		}
		return err
	default:
		return nil
	}
}
