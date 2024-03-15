package entity

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

type OrderEntity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time

	ConfirmedAt *time.Time
	PaidAt      *time.Time
	FulfilledAt *time.Time
	CancelledAt *time.Time

	Source        DataSource
	SourceOfTruth DataSource
	AppSource     string

	DataloaderKey string
}

type OrderEntities []OrderEntity

func (order OrderEntity) Labels(tenant string) []string {
	return []string{
		neo4jutil.NodeLabelOrder,
		neo4jutil.NodeLabelOrder + "_" + tenant,
	}
}
