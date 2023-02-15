package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"golang.org/x/net/context"
)

type ExternalSystemRepository interface {
	LinkContactWithExternalSystemInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, contactId string, relationship entity.ExternalReferenceRelationship) error
}

type externalSystemRepository struct {
	driver *neo4j.DriverWithContext
}

func NewExternalSystemRepository(driver *neo4j.DriverWithContext) ExternalSystemRepository {
	return &externalSystemRepository{
		driver: driver,
	}
}

func (e *externalSystemRepository) LinkContactWithExternalSystemInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, contactId string, relationship entity.ExternalReferenceRelationship) error {
	query := "MATCH (e:ExternalSystem {id:$externalSystemId})-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})," +
		" (c:Contact {id:$contactId}) " +
		" MERGE (c)-[r:IS_LINKED_WITH {id:$referenceId}]->(e) " +
		" ON CREATE SET e:%s, " +
		"				r.syncDate=$syncDate, " +
		"				e.createdAt=datetime({timezone: 'UTC'}) " +
		" ON MATCH SET r.syncDate=$syncDate " +
		" RETURN r"

	queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "ExternalSystem_"+tenant),
		map[string]any{
			"contactId":        contactId,
			"tenant":           tenant,
			"syncDate":         relationship.SyncDate,
			"referenceId":      relationship.Id,
			"externalSystemId": relationship.ExternalSystemId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}
