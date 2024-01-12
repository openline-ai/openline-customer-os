package invoice

import (
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
	"github.com/pkg/errors"
	"time"
)

const (
	InvoiceNewV1          = "V1_INVOICE_NEW"
	InvoiceFillV1         = "V1_INVOICE_FILL"
	InvoicePdfGeneratedV1 = "V1_INVOICE_PDF_GENERATED"
	InvoicePayV1          = "V1_INVOICE_PAY"
)

type InvoiceNewEvent struct {
	Tenant         string             `json:"tenant" validate:"required"`
	OrganizationId string             `json:"organizationId" validate:"required"`
	CreatedAt      time.Time          `json:"createdAt"`
	SourceFields   commonmodel.Source `json:"sourceFields"`
}

func NewInvoiceNewEvent(aggregate eventstore.Aggregate, organizationId string, createdAt *time.Time, sourceFields commonmodel.Source) (eventstore.Event, error) {
	eventData := InvoiceNewEvent{
		Tenant:         aggregate.GetTenant(),
		OrganizationId: organizationId,
		CreatedAt:      *createdAt,
		SourceFields:   sourceFields,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate InvoiceNewEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, InvoiceNewV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for InvoiceNewEvent")
	}

	return event, nil
}

type InvoiceFillEvent struct {
	Tenant       string             `json:"tenant" validate:"required"`
	UpdatedAt    time.Time          `json:"createdAt"`
	SourceFields commonmodel.Source `json:"sourceFields"`

	Amount float64                `json:"amount" validate:"required"`
	VAT    float64                `json:"vat" validate:"required"`
	Total  float64                `json:"total" validate:"required"`
	Lines  []InvoiceLineFillEvent `json:"invoiceLines" validate:"required"`
}

type InvoiceLineFillEvent struct {
	Tenant   string  `json:"tenant" validate:"required"`
	Index    int64   `json:"index" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Quantity int64   `json:"quantity" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	VAT      float64 `json:"vat" validate:"required"`
	Total    float64 `json:"total" validate:"required"`
}

func NewInvoiceFillEvent(aggregate eventstore.Aggregate, updatedAt *time.Time, sourceFields commonmodel.Source, request *invoicepb.FillInvoiceRequest) (eventstore.Event, error) {
	eventData := InvoiceFillEvent{
		Tenant:       aggregate.GetTenant(),
		UpdatedAt:    *updatedAt,
		SourceFields: sourceFields,

		Amount: request.Amount,
		VAT:    request.Vat,
		Total:  request.Total,
		Lines:  make([]InvoiceLineFillEvent, len(request.Lines)),
	}
	for i, line := range request.Lines {
		eventData.Lines[i] = InvoiceLineFillEvent{
			Index:    line.Index,
			Name:     line.Name,
			Price:    line.Price,
			Quantity: line.Quantity,
			Amount:   line.Amount,
			VAT:      line.Vat,
			Total:    line.Total,
		}
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate InvoiceFillEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, InvoiceFillV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for InvoiceFillEvent")
	}

	return event, nil
}

type InvoicePdfGeneratedEvent struct {
	Tenant       string             `json:"tenant" validate:"required"`
	UpdatedAt    time.Time          `json:"createdAt"`
	SourceFields commonmodel.Source `json:"sourceFields"`

	RepositoryFileId string `json:"repositoryFileId" validate:"required"`
}

func NewInvoicePdfGeneratedEvent(aggregate eventstore.Aggregate, updatedAt *time.Time, sourceFields commonmodel.Source, request *invoicepb.PdfGeneratedInvoiceRequest) (eventstore.Event, error) {
	eventData := InvoicePdfGeneratedEvent{
		Tenant:           aggregate.GetTenant(),
		UpdatedAt:        *updatedAt,
		SourceFields:     sourceFields,
		RepositoryFileId: request.RepositoryFileId,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate InvoicePdfGeneratedEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, InvoicePdfGeneratedV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for InvoicePdfGeneratedEvent")
	}

	return event, nil
}

type InvoicePayEvent struct {
	Tenant       string             `json:"tenant" validate:"required"`
	UpdatedAt    time.Time          `json:"createdAt"`
	SourceFields commonmodel.Source `json:"sourceFields"`
}

func NewInvoicePayEvent(aggregate eventstore.Aggregate, updatedAt *time.Time, sourceFields commonmodel.Source, request *invoicepb.PayInvoiceRequest) (eventstore.Event, error) {
	eventData := InvoicePayEvent{
		Tenant:       aggregate.GetTenant(),
		UpdatedAt:    *updatedAt,
		SourceFields: sourceFields,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate InvoicePayEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, InvoicePayV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for InvoicePayEvent")
	}

	return event, nil
}

type EventHandlers struct {
	InvoiceNew          InvoiceNewHandler
	InvoiceFill         InvoiceFillHandler
	InvoicePdfGenerated InvoicePdfGeneratedHandler
	InvoicePay          InvoicePayHandler
}

func NewEventHandlers(log logger.Logger, es eventstore.AggregateStore) *EventHandlers {
	return &EventHandlers{
		InvoiceNew:          NewInvoiceNewHandler(log, es),
		InvoiceFill:         NewInvoiceFillHandler(log, es),
		InvoicePdfGenerated: NewInvoicePdfGeneratedHandler(log, es),
		InvoicePay:          NewInvoicePayHandler(log, es),
	}
}
