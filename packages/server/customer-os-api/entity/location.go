package entity

import (
	"fmt"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"time"
)

type LocationEntity struct {
	Id           string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Country      string `neo4jDb:"property:country;lookupName:COUNTRY;supportCaseSensitive:true"`
	Region       string `neo4jDb:"property:region;lookupName:REGION;supportCaseSensitive:true"`
	Locality     string `neo4jDb:"property:locality;lookupName:LOCALITY;supportCaseSensitive:true"`
	Address      string
	Address2     string
	Zip          string
	AddressType  string
	HouseNumber  string
	PostalCode   string
	PlusFour     string
	Commercial   bool
	Predirection string
	District     string
	Street       string
	RawAddress   string
	Latitude     *float64
	Longitude    *float64
	TimeZone     string
	UtcOffset    int64

	SourceOfTruth neo4jentity.DataSource
	Source        neo4jentity.DataSource
	AppSource     string

	DataloaderKey string
}

func (location LocationEntity) ToString() string {
	return fmt.Sprintf("id: %s name: %s", location.Id, location.Name)
}

type LocationEntities []LocationEntity

func (location LocationEntity) Labels(tenant string) []string {
	return []string{"Location", "Location_" + tenant}
}
