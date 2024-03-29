package model

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"time"
)

type LogEntryData struct {
	BaseData
	Content              string                   `json:"content,omitempty"`
	ContentType          string                   `json:"contentType,omitempty"`
	StartedAtStr         string                   `json:"startedAt,omitempty"`
	StartedAt            *time.Time               `json:"startedAtTime,omitempty"`
	AuthorUser           ReferencedUser           `json:"authorUser,omitempty"`
	LoggedOrganization   ReferencedOrganization   `json:"loggedOrganization,omitempty"`
	LoggedOrganizations  []ReferencedOrganization `json:"loggedOrganizations,omitempty"`
	LoggedEntityRequired bool                     `json:"loggedEntityRequired,omitempty"`
}

func (m *LogEntryData) SetStartedAtTime() {
	if m.StartedAtStr != "" && m.StartedAt == nil {
		m.StartedAt, _ = utils.UnmarshalDateTime(m.StartedAtStr)
	}
	if m.StartedAt != nil {
		m.StartedAt = utils.TimePtr((*m.StartedAt).UTC())
	}
}

func (l *LogEntryData) Normalize() {
	l.SetTimes()
	l.BaseData.Normalize()
	l.SetStartedAtTime()
}
