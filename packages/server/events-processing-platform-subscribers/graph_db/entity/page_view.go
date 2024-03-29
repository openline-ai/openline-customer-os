package entity

import (
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

// Deprecated
type PageViewEntity struct {
	Id             string
	Application    string
	TrackerName    string
	SessionId      string
	PageUrl        string
	PageTitle      string
	OrderInSession int64
	EngagedTime    int64
	StartedAt      time.Time
	EndedAt        time.Time
	Source         neo4jentity.DataSource
	SourceOfTruth  neo4jentity.DataSource
	AppSource      string

	DataloaderKey string
}

func (PageViewEntity) IsTimelineEvent() {
}

func (PageViewEntity) TimelineEventLabel() string {
	return neo4jutil.NodeLabelPageView
}
