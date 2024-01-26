package invoice

import (
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/pkg/errors"
	"time"
)

type InvoiceForContractCreateEvent struct {
	Tenant          string             `json:"tenant" validate:"required"`
	ContractId      string             `json:"organizationId" validate:"required"`
	CreatedAt       time.Time          `json:"createdAt"`
	DueDate         time.Time          `json:"dueDate"`
	SourceFields    commonmodel.Source `json:"sourceFields"`
	DryRun          bool               `json:"dryRun"`
	InvoiceNumber   string             `json:"invoiceNumber"`
	Currency        string             `json:"currency"`
	PeriodStartDate time.Time          `json:"periodStartDate"`
	PeriodEndDate   time.Time          `json:"periodEndDate"`
	BillingCycle    string             `json:"billingCycle"`
	Note            string             `json:"note"`
}

func NewInvoiceForContractCreateEvent(aggregate eventstore.Aggregate, sourceFields commonmodel.Source, contractId, currency, invoiceNumber, billingCycle, note string, dryRun bool, createdAt, periodStartDate, periodEndDate time.Time) (eventstore.Event, error) {
	eventData := InvoiceForContractCreateEvent{
		Tenant:          aggregate.GetTenant(),
		ContractId:      contractId,
		CreatedAt:       createdAt,
		DueDate:         createdAt,
		SourceFields:    sourceFields,
		Currency:        currency,
		InvoiceNumber:   invoiceNumber,
		DryRun:          dryRun,
		PeriodStartDate: periodStartDate,
		PeriodEndDate:   periodEndDate,
		BillingCycle:    billingCycle,
		Note:            note,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate InvoiceCreateEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, InvoiceCreateForContractV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for InvoiceCreateEvent")
	}

	return event, nil
}
