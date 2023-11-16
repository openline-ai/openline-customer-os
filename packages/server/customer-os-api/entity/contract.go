package entity

import (
	"time"
)

type ContractEntity struct {
	Id                   string
	Name                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ServiceStartedAt     *time.Time
	SignedAt             *time.Time
	EndedAt              *time.Time
	ContractRenewalCycle ContractRenewalCycle
	ContractStatus       ContractStatus
	Source               DataSource
	SourceOfTruth        DataSource
	AppSource            string
	ContractUrl          string

	DataloaderKey string
}

type ContractEntities []ContractEntity
