package models

import (
	commonModels "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"time"
)

type OrganizationCoreFields struct {
	Name             string
	Description      string
	Website          string
	Industry         string
	SubIndustry      string
	IndustryGroup    string
	TargetAudience   string
	ValueProposition string
	IsPublic         bool
	Employees        int64
	Market           string
}

type OrganizationDto struct {
	ID                     string
	Tenant                 string
	OrganizationCoreFields OrganizationCoreFields
	Source                 commonModels.Source
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
}
