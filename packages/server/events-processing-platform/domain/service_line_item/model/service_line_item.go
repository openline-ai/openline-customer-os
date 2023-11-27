package model

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"time"
)

// ServiceLineItem represents the state of a service line item aggregate.
type ServiceLineItem struct {
	ID         string             `json:"id"`
	ContractId string             `json:"contractId"`
	ParentId   string             `json:"parentId"`
	Billed     string             `json:"billed"`
	Quantity   int64              `json:"quantity"` // Relevant only for Subscription type
	Price      float64            `json:"price"`
	Name       string             `json:"name"`
	Comments   string             `json:"comments,omitempty"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
	StartedAt  time.Time          `json:"startedAt"`
	EndedAt    *time.Time         `json:"endedAt,omitempty"`
	Source     commonmodel.Source `json:"source"`
	IsDeleted  bool               `json:"isDeleted"`
}

// ServiceLineItemDataFields contains all the fields that may be used to create or update a service line item.
type ServiceLineItemDataFields struct {
	Billed     BilledType `json:"billed"`
	Quantity   int64      `json:"quantity"` // Relevant only for Subscription type
	Price      float64    `json:"price"`
	Name       string     `json:"name"`
	ContractId string     `json:"contractId"`
	ParentId   string     `json:"parentId"`
	Comments   string     `json:"comments,omitempty"`
}

func (sli ServiceLineItem) IsEnded() bool {
	return sli.EndedAt != nil && (*sli.EndedAt).Before(utils.Now())
}

// BilledType enum represents the billing type for a service line item.
type BilledType int32

const (
	MonthlyBilled BilledType = iota
	AnnuallyBilled
	OnceBilled  // For One-Time
	UsageBilled // For Usage-Based
)

func (bt BilledType) String() string {
	return [...]string{string(MonthlyBilledString), string(AnnuallyBilledString), string(OnceBilledString), string(UsageBilledString)}[bt]
}

func (bt BilledType) IsOneTime() bool {
	return bt == OnceBilled
}

func (bt BilledType) IsUsage() bool {
	return bt == UsageBilled
}

func (bt BilledType) IsRecurrent() bool {
	return bt == MonthlyBilled || bt == AnnuallyBilled
}

type BilledString string

const (
	MonthlyBilledString  BilledString = "MONTHLY"
	AnnuallyBilledString BilledString = "ANNUALLY"
	OnceBilledString     BilledString = "ONCE"
	UsageBilledString    BilledString = "USAGE"
)
