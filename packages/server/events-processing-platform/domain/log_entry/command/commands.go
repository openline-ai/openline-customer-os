package command

import (
	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/log_entry/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"time"
)

type UpsertLogEntryCommand struct {
	eventstore.BaseCommand
	DataFields models.LogEntryDataFields
	Source     cmnmod.Source
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewUpsertLogEntryCommand(logEntryId, tenant, source, sourceOfTruth, appSource, userId string, dataFields models.LogEntryDataFields, createdAt, updatedAt *time.Time) *UpsertLogEntryCommand {
	return &UpsertLogEntryCommand{
		BaseCommand: eventstore.NewBaseCommand(logEntryId, tenant, userId),
		DataFields:  dataFields,
		Source: cmnmod.Source{
			Source:        source,
			SourceOfTruth: sourceOfTruth,
			AppSource:     appSource,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
