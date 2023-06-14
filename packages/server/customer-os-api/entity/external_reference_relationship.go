package entity

import "time"

type ExternalSystemId string

const (
	Hubspot        ExternalSystemId = "hubspot"
	ZendeskSupport ExternalSystemId = "zendesk_support"
)

type ExternalSystemEntity struct {
	ExternalSystemId ExternalSystemId
	Relationship     struct {
		ExternalId  string
		SyncDate    *time.Time
		ExternalUrl *string
	}
	DataloaderKey string
}

type ExternalSystemEntities []ExternalSystemEntity

func ExternalSystemTypeFromString(input string) ExternalSystemId {
	for _, v := range []ExternalSystemId{
		Hubspot, ZendeskSupport,
	} {
		if string(v) == input {
			return v
		}
	}
	// Return a default value or handle the case when the input string doesn't match any ExternalSystemId
	return ""
}
