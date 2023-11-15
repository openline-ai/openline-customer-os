package entity

import "time"

type ServiceLineItemEntity struct {
	ID            string
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Billed        BilledType
	Price         float64
	Quantity      int
	Source        DataSource
	SourceOfTruth DataSource
	AppSource     string

	DataloaderKey string
}

type ServiceLineItemEntities []ServiceLineItemEntity
