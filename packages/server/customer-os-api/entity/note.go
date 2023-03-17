package entity

import (
	"fmt"
	"time"
)

type NoteEntity struct {
	Id            string
	Html          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Source        DataSource
	SourceOfTruth DataSource
	AppSource     string

	DataloaderKey string
}

func (note NoteEntity) ToString() string {
	return fmt.Sprintf("id: %s\nhtml: %s", note.Id, note.Html)
}

type NoteEntities []NoteEntity

func (NoteEntity) IsTimelineEvent() {
}

func (NoteEntity) TimelineEventLabel() string {
	return NodeLabel_Note
}

func (note NoteEntity) Labels(tenant string) []string {
	return []string{"Note", "Action", "Note_" + tenant, "Action_" + tenant}
}
