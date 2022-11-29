package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/entity"
)

type ActionItemRepository interface {
	CreatePageViewActionItem(contactId string, pv entity.PageView) error
}

type actionItemRepository struct {
	driver *neo4j.Driver
}

func NewActionItemRepository(driver *neo4j.Driver) ActionItemRepository {
	return &actionItemRepository{
		driver: driver,
	}
}

func (r *actionItemRepository) CreatePageViewActionItem(contactId string, pv entity.PageView) error {
	session := (*r.driver).NewSession(
		neo4j.SessionConfig{
			AccessMode: neo4j.AccessModeWrite,
			BoltLogger: neo4j.ConsoleBoltLogger()})
	defer session.Close()

	params := map[string]interface{}{
		"tenant":         pv.Tenant,
		"contactId":      contactId,
		"pvId":           pv.ID,
		"start":          pv.Start,
		"end":            pv.End,
		"appId":          pv.AppId,
		"trackerName":    pv.TrackerName,
		"pageUrl":        pv.Url,
		"pageTitle":      pv.Title,
		"orderInSession": pv.OrderInSession,
		"engagedTime":    pv.EngagedTime,
	}
	query := "MATCH (c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant}) " +
		" MERGE (c)-[:HAS_ACTION]->(a:ActionItem:PageViewAction {id:$pvId, trackerName:$trackerName})" +
		" ON CREATE SET " +
		" a.startedAt=$start, " +
		" a.endedAt=$end, " +
		" a.appId=$appId, " +
		" a.pageUrl=$pageUrl, " +
		" a.pageTitle=$pageTitle, " +
		" a.orderInSession=$orderInSession, " +
		" a.engagedTime=$engagedTime"

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(query, params)
		return nil, err
	})

	return err
}
