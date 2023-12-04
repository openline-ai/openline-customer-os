package models

import (
	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"time"
)

type OrganizationDataFields struct {
	Name              string
	Hide              bool
	Description       string
	Website           string
	Industry          string
	SubIndustry       string
	IndustryGroup     string
	TargetAudience    string
	ValueProposition  string
	IsPublic          bool
	IsCustomer        bool
	Employees         int64
	Market            string
	LastFundingRound  string
	LastFundingAmount string
	ReferenceId       string
	Note              string
}

type OrganizationFields struct {
	ID                     string
	Tenant                 string
	IgnoreEmptyFields      bool
	OrganizationDataFields OrganizationDataFields
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
	Source                 cmnmod.Source
	ExternalSystem         cmnmod.ExternalSystem
}
