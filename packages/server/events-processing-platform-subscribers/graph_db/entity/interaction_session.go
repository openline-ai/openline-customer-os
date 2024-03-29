package entity

import (
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

// Deprecated
type InteractionSessionEntity struct {
	Id            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Identifier    string
	Name          string
	Status        string
	Type          string
	Channel       string
	ChannelData   string
	AppSource     string
	Source        neo4jentity.DataSource
	SourceOfTruth neo4jentity.DataSource
}

func (InteractionSessionEntity) IsTimelineEvent() {
}

func (InteractionSessionEntity) TimelineEventLabel() string {
	return neo4jutil.NodeLabelInteractionSession
}
