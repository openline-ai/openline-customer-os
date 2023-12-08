package entity

import "time"

type ServiceLineItemEntity struct {
	ID               string
	Name             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	StartedAt        time.Time
	EndedAt          *time.Time
	IsCanceled       bool
	Billed           BilledType
	Price            float64
	Quantity         int64
	PreviousBilled   BilledType
	PreviousPrice    float64
	PreviousQuantity int64
	Comments         string
	Source           DataSource
	SourceOfTruth    DataSource
	AppSource        string
	ParentID         string

	DataloaderKey string
}

type ServiceLineItemEntities []ServiceLineItemEntity
