package entity

import "time"

type OrganizationData struct {
	Id             string
	Name           string
	Description    string
	Domain         string
	Website        string
	Industry       string
	IsPublic       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExternalId     string
	ExternalSystem string
	Country        string
	State          string
	City           string
	Address        string
	Address2       string
	Zip            string
	Phone          string

	OrganizationTypeName string
}
