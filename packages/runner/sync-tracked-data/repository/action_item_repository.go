package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/entity"
)

type ActionRepository interface {
	CreatePageViewAction(contactId string, pv entity.PageView) error
}

type actionRepository struct {
	driver *neo4j.Driver
}

func NewActionRepository(driver *neo4j.Driver) ActionRepository {
	return &actionRepository{
		driver: driver,
	}
}

func (r *actionRepository) CreatePageViewAction(contactId string, pv entity.PageView) error {
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
		"application":    pv.Application,
		"sessionId":      pv.SessionID,
		"trackerName":    pv.TrackerName,
		"pageUrl":        pv.Url,
		"pageTitle":      pv.Title,
		"orderInSession": pv.OrderInSession,
		"engagedTime":    pv.EngagedTime,
		"source":         "openline",
		"sourceOfTruth":  "openline",
		"appSource":      pv.Application,
	}
	query := "MATCH (c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant}) " +
		" MERGE (c)-[:HAS_ACTION]->(a:Action:PageView {id:$pvId, trackerName:$trackerName})" +
		" ON CREATE SET " +
		" a.startedAt=$start, " +
		" a.endedAt=$end, " +
		" a.application=$application, " +
		" a.pageUrl=$pageUrl, " +
		" a.pageTitle=$pageTitle, " +
		" a.sessionId=$sessionId, " +
		" a.orderInSession=$orderInSession, " +
		" a.engagedTime=$engagedTime, " +
		" a.source=$source, " +
		" a.sourceOfTruth=$sourceOfTruth, " +
		" a.appSource=$appSource "

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(query, params)
		return nil, err
	})

	return err
}
