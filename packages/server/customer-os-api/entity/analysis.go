package entity

import (
	"fmt"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

type AnalysisEntity struct {
	Id        string
	CreatedAt *time.Time

	Content       string
	ContentType   string
	AnalysisType  string
	Source        neo4jentity.DataSource
	SourceOfTruth neo4jentity.DataSource
	AppSource     string

	DataloaderKey string
}

func (analysisEntity AnalysisEntity) ToString() string {
	return fmt.Sprintf("id: %s", analysisEntity.Id)
}

type AnalysisEntities []AnalysisEntity

func (AnalysisEntity) IsTimelineEvent() {
}

func (analysis *AnalysisEntity) SetDataloaderKey(key string) {
	analysis.DataloaderKey = key
}

func (analysis AnalysisEntity) GetDataloaderKey() string {
	return analysis.DataloaderKey
}

func (AnalysisEntity) TimelineEventLabel() string {
	return neo4jutil.NodeLabelAnalysis
}

func (AnalysisEntity) Labels(tenant string) []string {
	return []string{
		neo4jutil.NodeLabelAnalysis,
		neo4jutil.NodeLabelAnalysis + "_" + tenant,
	}
}
