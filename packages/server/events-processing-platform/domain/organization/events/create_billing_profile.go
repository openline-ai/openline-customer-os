package events

import (
	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/pkg/errors"
	"time"
)

type CreateBillingProfileEvent struct {
	Tenant           string        `json:"tenant" validate:"required"`
	BillingProfileId string        `json:"billingProfileId" validate:"required"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
	LegalName        string        `json:"legalName"`
	TaxId            string        `json:"taxId"`
	SourceFields     cmnmod.Source `json:"sourceFields" validate:"required"`
}

func NewCreateBillingProfileEvent(aggregate eventstore.Aggregate, billingProfileId, legalName, taxId string, sourceFields cmnmod.Source, createdAt, updatedAt time.Time) (eventstore.Event, error) {
	eventData := CreateBillingProfileEvent{
		Tenant:           aggregate.GetTenant(),
		BillingProfileId: billingProfileId,
		LegalName:        legalName,
		TaxId:            taxId,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		SourceFields:     sourceFields,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "failed to validate CreateBillingProfileEvent")
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationCreateBillingProfileV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, errors.Wrap(err, "error setting json data for CreateBillingProfileEvent")
	}
	return event, nil
}