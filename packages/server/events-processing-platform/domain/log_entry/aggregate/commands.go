package aggregate

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/aggregate"
	cmd "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func (a *LogEntryAggregate) HandleCommand(ctx context.Context, command eventstore.Command) error {
	switch c := command.(type) {
	case *cmd.AddTagCommand:
		return a.addTag(ctx, c)
	case *cmd.RemoveTagCommand:
		return a.removeTag(ctx, c)
	default:
		return errors.New("invalid command type")
	}
}

func (a *LogEntryAggregate) CreateLogEntry(ctx context.Context, command *cmd.UpsertLogEntryCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "LogEntryAggregate.CreateLogEntry")
	defer span.Finish()
	span.LogFields(log.String("Tenant", a.Tenant), log.String("AggregateID", a.GetID()), log.Int64("AggregateVersion", a.GetVersion()))

	createdAtNotNil := utils.IfNotNilTimeWithDefault(command.CreatedAt, utils.Now())
	updatedAtNotNil := utils.IfNotNilTimeWithDefault(command.UpdatedAt, createdAtNotNil)
	startedAtNotNil := utils.IfNotNilTimeWithDefault(command.DataFields.StartedAt, createdAtNotNil)

	createEvent, err := events.NewLogEntryCreateEvent(a, command.DataFields, command.Source, createdAtNotNil, updatedAtNotNil, startedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewLogEntryCreateEvent")
	}
	aggregate.EnrichEventWithMetadata(&createEvent, &span, a.Tenant, command.UserID)

	return a.Apply(createEvent)
}

func (a *LogEntryAggregate) UpdateLogEntry(ctx context.Context, command *cmd.UpsertLogEntryCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "LogEntryAggregate.UpdateLogEntry")
	defer span.Finish()
	span.LogFields(log.String("Tenant", a.Tenant), log.String("AggregateID", a.GetID()), log.Int64("AggregateVersion", a.GetVersion()))

	updatedAtNotNil := utils.IfNotNilTimeWithDefault(command.UpdatedAt, utils.Now())
	startedAtNotNil := utils.IfNotNilTimeWithDefault(command.DataFields.StartedAt, a.LogEntry.StartedAt)
	if command.Source.SourceOfTruth == "" {
		command.Source.SourceOfTruth = a.LogEntry.Source.SourceOfTruth
	}

	event, err := events.NewLogEntryUpdateEvent(a, command.DataFields.Content, command.DataFields.ContentType,
		command.Source.SourceOfTruth, updatedAtNotNil, startedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewLogEntryUpdateEvent")
	}
	aggregate.EnrichEventWithMetadata(&event, &span, a.Tenant, command.UserID)

	return a.Apply(event)
}

func (a *LogEntryAggregate) addTag(ctx context.Context, command *cmd.AddTagCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "LogEntryAggregate.addTag")
	defer span.Finish()
	span.LogFields(log.String("Tenant", a.Tenant), log.String("AggregateID", a.GetID()), log.Int64("AggregateVersion", a.GetVersion()))

	taggedAtNotNil := utils.IfNotNilTimeWithDefault(command.TaggedAt, utils.Now())

	event, err := events.NewLogEntryAddTagEvent(a, command.TagId, taggedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewLogEntryAddTagEvent")
	}
	aggregate.EnrichEventWithMetadata(&event, &span, a.Tenant, command.UserID)

	return a.Apply(event)
}

func (a *LogEntryAggregate) removeTag(ctx context.Context, command *cmd.RemoveTagCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "LogEntryAggregate.removeTag")
	defer span.Finish()
	span.LogFields(log.String("Tenant", a.Tenant), log.String("AggregateID", a.GetID()), log.Int64("AggregateVersion", a.GetVersion()))

	event, err := events.NewLogEntryRemoveTagEvent(a, command.TagId)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewLogEntryRemoveTagEvent")
	}
	aggregate.EnrichEventWithMetadata(&event, &span, a.Tenant, command.UserID)

	return a.Apply(event)
}
