package entity

import (
	"fmt"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

type NoteEntity struct {
	Id            string
	Content       string
	ContentType   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Source        neo4jentity.DataSource
	SourceOfTruth neo4jentity.DataSource
	AppSource     string

	DataloaderKey string
}

func (note NoteEntity) ToString() string {
	return fmt.Sprintf("id: %s\ncontent: %s", note.Id, note.Content)
}

type NoteEntities []NoteEntity

func (NoteEntity) IsTimelineEvent() {
}

func (NoteEntity) TimelineEventLabel() string {
	return neo4jutil.NodeLabelNote
}

func (note *NoteEntity) SetDataloaderKey(key string) {
	note.DataloaderKey = key
}

func (note NoteEntity) GetDataloaderKey() string {
	return note.DataloaderKey
}

func (note NoteEntity) Labels(tenant string) []string {
	return []string{"Note", "TimelineEvent", "Note_" + tenant, "TimelineEvent_" + tenant}
}
