package aggregate

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/aggregate"
	common_models "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/pkg/errors"
)

const (
	LogEntryAggregateType eventstore.AggregateType = "log_entry"
)

type LogEntryAggregate struct {
	*aggregate.CommonTenantIdAggregate
	LogEntry *models.LogEntry
}

func NewLogEntryAggregateWithTenantAndID(tenant, id string) *LogEntryAggregate {
	logEntryAggregate := LogEntryAggregate{}
	logEntryAggregate.CommonTenantIdAggregate = aggregate.NewCommonAggregateWithTenantAndId(LogEntryAggregateType, tenant, id)
	logEntryAggregate.SetWhen(logEntryAggregate.When)
	logEntryAggregate.LogEntry = &models.LogEntry{}
	logEntryAggregate.Tenant = tenant

	return &logEntryAggregate
}

func (a *LogEntryAggregate) When(event eventstore.Event) error {

	switch event.GetEventType() {
	case events.LogEntryCreateV1:
		return a.onLogEntryCreate(event)
	case events.LogEntryUpdateV1:
		return a.onLogEntryUpdate(event)
	default:
		err := eventstore.ErrInvalidEventType
		err.EventType = event.GetEventType()
		return err
	}
}

func (a *LogEntryAggregate) onLogEntryCreate(event eventstore.Event) error {
	var eventData events.LogEntryCreateEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}
	a.LogEntry.ID = a.ID
	a.LogEntry.Tenant = a.Tenant
	a.LogEntry.Content = eventData.Content
	a.LogEntry.ContentType = eventData.ContentType
	a.LogEntry.AuthorUserId = eventData.AuthorUserId
	a.LogEntry.LoggedOrganizationId = eventData.LoggedOrganizationId
	a.LogEntry.StartedAt = eventData.StartedAt
	a.LogEntry.Source = common_models.Source{
		Source:        eventData.Source,
		SourceOfTruth: eventData.SourceOfTruth,
		AppSource:     eventData.AppSource,
	}
	a.LogEntry.CreatedAt = eventData.CreatedAt
	a.LogEntry.UpdatedAt = eventData.UpdatedAt
	return nil
}

func (a *LogEntryAggregate) onLogEntryUpdate(event eventstore.Event) error {
	var eventData events.LogEntryUpdateEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}
	a.LogEntry.Content = eventData.Content
	a.LogEntry.ContentType = eventData.ContentType
	a.LogEntry.StartedAt = eventData.StartedAt
	a.LogEntry.UpdatedAt = eventData.UpdatedAt
	if eventData.SourceOfTruth != "" {
		a.LogEntry.Source.SourceOfTruth = eventData.SourceOfTruth
	}
	return nil
}
