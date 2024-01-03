package entity

import (
	"fmt"
	neo4jentity "github.com/openline-ai/customer-os-neo4j-repository/entity"
	"time"
)

type TagEntity struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Source    neo4jentity.DataSource
	AppSource string
	TaggedAt  time.Time

	DataloaderKey string
}

func (tag TagEntity) ToString() string {
	return fmt.Sprintf("id: %s\nname: %s", tag.Id, tag.Name)
}

type TagEntities []TagEntity

func (TagEntity) Labels(tenant string) []string {
	return []string{
		NodeLabel_Tag,
		NodeLabel_Tag + "_" + tenant,
	}
}
