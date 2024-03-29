package model

import (
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"time"

	neo4jmodel "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/model"

	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	orgplanmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization_plan/model"
)

const (
	FieldMaskName               = "name"
	FieldMaskTargetAudience     = "targetAudience"
	FieldMaskValueProposition   = "valueProposition"
	FieldMaskIndustry           = "industry"
	FieldMaskSubIndustry        = "subIndustry"
	FieldMaskIndustryGroup      = "industryGroup"
	FieldMaskMarket             = "market"
	FieldMaskHide               = "hide"
	FieldMaskDescription        = "description"
	FieldMaskNote               = "note"
	FieldMaskIsPublic           = "isPublic"
	FieldMaskIsCustomer         = "isCustomer"
	FieldMaskEmployees          = "employees"
	FieldMaskLastFundingRound   = "lastFundingRound"
	FieldMaskLastFundingAmount  = "lastFundingAmount"
	FieldMaskReferenceId        = "referenceId"
	FieldMaskWebsite            = "website"
	FieldMaskYearFounded        = "yearFounded"
	FieldMaskHeadquarters       = "headquarters"
	FieldMaskLogoUrl            = "logoUrl"
	FieldMaskEmployeeGrowthRate = "employeeGrowthRate"
	FieldMaskSlackChannelId     = "slackChannelId"
)

type CustomFieldDataType string

const (
	CustomFieldDataTypeText     CustomFieldDataType = "TEXT"
	CustomFieldDataTypeBool     CustomFieldDataType = "BOOL"
	CustomFieldDataTypeDatetime CustomFieldDataType = "DATETIME"
	CustomFieldDataTypeInteger  CustomFieldDataType = "INTEGER"
	CustomFieldDataTypeDecimal  CustomFieldDataType = "DECIMAL"
)

type Social struct {
	PlatformName string `json:"platformName"`
	Url          string `json:"url"`
}

type CustomField struct {
	Id                  string                      `json:"id"`
	Name                string                      `json:"name"`
	TemplateId          *string                     `json:"templateId,omitempty"`
	CustomFieldValue    neo4jmodel.CustomFieldValue `json:"customFieldValue"`
	CustomFieldDataType CustomFieldDataType         `json:"customFieldDataType"`
	Source              cmnmod.Source               `json:"source"`
	CreatedAt           time.Time                   `json:"createdAt,omitempty"`
	UpdatedAt           time.Time                   `json:"updatedAt,omitempty"`
}

type Organization struct {
	ID                  string                                   `json:"id"`
	Name                string                                   `json:"name"`
	Hide                bool                                     `json:"hide"`
	Description         string                                   `json:"description"`
	Website             string                                   `json:"website"`
	Industry            string                                   `json:"industry"`
	SubIndustry         string                                   `json:"subIndustry"`
	IndustryGroup       string                                   `json:"industryGroup"`
	TargetAudience      string                                   `json:"targetAudience"`
	ValueProposition    string                                   `json:"valueProposition"`
	IsPublic            bool                                     `json:"isPublic"`
	IsCustomer          bool                                     `json:"isCustomer"`
	Employees           int64                                    `json:"employees"`
	Market              string                                   `json:"market"`
	LastFundingRound    string                                   `json:"lastFundingRound"`
	LastFundingAmount   string                                   `json:"lastFundingAmount"`
	ReferenceId         string                                   `json:"referenceId"`
	Note                string                                   `json:"note"`
	Source              cmnmod.Source                            `json:"source"`
	CreatedAt           time.Time                                `json:"createdAt,omitempty"`
	UpdatedAt           time.Time                                `json:"updatedAt,omitempty"`
	PhoneNumbers        map[string]OrganizationPhoneNumber       `json:"phoneNumbers"`
	Emails              map[string]OrganizationEmail             `json:"emails"`
	Locations           []string                                 `json:"locations,omitempty"`
	Domains             []string                                 `json:"domains,omitempty"`
	Socials             map[string]Social                        `json:"socials,omitempty"`
	CustomFields        map[string]CustomField                   `json:"customFields,omitempty"`
	ExternalSystems     []cmnmod.ExternalSystem                  `json:"externalSystems"`
	ParentOrganizations map[string]ParentOrganization            `json:"parentOrganizations,omitempty"`
	LogoUrl             string                                   `json:"logoUrl,omitempty"`
	YearFounded         *int64                                   `json:"yearFounded,omitempty"`
	Headquarters        string                                   `json:"headquarters,omitempty"`
	EmployeeGrowthRate  string                                   `json:"employeeGrowthRate,omitempty"`
	SlackChannelId      string                                   `json:"slackChannelId,omitempty"`
	OnboardingDetails   OnboardingDetails                        `json:"onboardingDetails,omitempty"`
	BillingProfiles     map[string]BillingProfile                `json:"billingProfiles,omitempty"`
	OrganizationPlans   map[string]orgplanmodel.OrganizationPlan `json:"organizationPlans,omitempty"`
}

type BillingProfile struct {
	Id             string        `json:"id"`
	LegalName      string        `json:"legalName"`
	TaxId          string        `json:"taxId"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	SourceFields   cmnmod.Source `json:"sourceFields"`
	PrimaryEmailId string        `json:"primaryEmailId"`
	EmailIds       []string      `json:"emailIds"`
	LocationIds    []string      `json:"locationIds"`
}

type OnboardingDetails struct {
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
	Comments  string    `json:"comments"`
}

type OrganizationPhoneNumber struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

type OrganizationEmail struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

type ParentOrganization struct {
	OrganizationId string `json:"organizationId"`
	Type           string `json:"type"`
}

func (o *Organization) String() string {
	return fmt.Sprintf("Organization{ID: %s, Name: %s, Description: %s, Website: %s, Industry: %s, IsPublic: %t, Source: %s, CreatedAt: %s, UpdatedAt: %s}", o.ID, o.Name, o.Description, o.Website, o.Industry, o.IsPublic, o.Source, o.CreatedAt, o.UpdatedAt)
}

func (o *Organization) GetSocialIdForUrl(url string) string {
	if o.Socials == nil {
		return ""
	}
	for key, social := range o.Socials {
		if social.Url == url {
			return key
		}
	}
	return ""
}

func (o *Organization) ContainsExternalSystem(externalSystem string) bool {
	for _, es := range o.ExternalSystems {
		if es.ExternalSystemId == externalSystem {
			return true
		}
	}
	return false
}

func (o *Organization) SkipUpdate(fields *OrganizationFields) bool {
	if fields.ExternalSystem.Available() && !o.ContainsExternalSystem(fields.ExternalSystem.ExternalSystemId) {
		return false
	}
	if o.Source.SourceOfTruth == constants.SourceOpenline && fields.ExternalSystem.Available() {
		return true
	}
	return false
}
