package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/utils"
	"golang.org/x/net/context"
	"time"
)

type NoteRepository interface {
	MergeNote(ctx context.Context, tenant string, syncDate time.Time, note entity.NoteData) (string, error)
	NoteLinkWithContactByExternalId(ctx context.Context, tenant, noteId, contactExternalId, externalSystem string) error
	NoteLinkWithOrganizationByExternalId(ctx context.Context, tenant, noteId, organizationExternalId, externalSystem string) error
	NoteLinkWithUserByExternalId(ctx context.Context, tenant, noteId, userExternalId, externalSystem string) error
	NoteLinkWithUserByExternalOwnerId(ctx context.Context, tenant, noteId, userExternalOwnerId, externalSystem string) error
}

type noteRepository struct {
	driver *neo4j.DriverWithContext
}

func NewNoteRepository(driver *neo4j.DriverWithContext) NoteRepository {
	return &noteRepository{
		driver: driver,
	}
}

func (r *noteRepository) MergeNote(ctx context.Context, tenant string, syncDate time.Time, note entity.NoteData) (string, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	// Create new Note if it does not exist
	// If Note exists, and sourceOfTruth is acceptable then update Note.
	//   otherwise create/update AlternateNote for incoming source, with a new relationship 'ALTERNATE'
	query := "MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem}) " +
		"MERGE (n:Note)-[r:IS_LINKED_WITH {externalId:$externalId}]->(e) " +
		"ON CREATE SET 	r.syncDate=$syncDate, " +
		"				n.id=randomUUID(), " +
		"				n.createdAt=$createdAt, " +
		"				n.updatedAt=$createdAt, " +
		"              	n.source=$source, " +
		"				n.sourceOfTruth=$sourceOfTruth, " +
		"				n.appSource=$appSource, " +
		"              	n.html=$html, " +
		"				n:%s " +
		"ON MATCH SET r.syncDate = CASE WHEN n.sourceOfTruth=$sourceOfTruth THEN $syncDate ELSE r.syncDate END, " +
		"             n.html = CASE WHEN n.sourceOfTruth=$sourceOfTruth THEN $html ELSE n.html END, " +
		"             n.updatedAt = CASE WHEN n.sourceOfTruth=$sourceOfTruth THEN $now ELSE n.updatedAt END " +
		"WITH n " +
		"FOREACH (x in CASE WHEN n.sourceOfTruth <> $sourceOfTruth THEN [n] ELSE [] END | " +
		"  MERGE (x)-[:ALTERNATE]->(alt:AlternateNote {source:$source, id:x.id}) " +
		"    SET alt.updatedAt=$now, alt.appSource=$appSource, alt.html=$html " +
		") " +
		"RETURN n.id"

	dbRecord, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "Note_"+tenant),
			map[string]interface{}{
				"tenant":         tenant,
				"source":         note.ExternalSystem,
				"sourceOfTruth":  note.ExternalSystem,
				"appSource":      note.ExternalSystem,
				"externalSystem": note.ExternalSystem,
				"externalId":     note.ExternalId,
				"syncDate":       syncDate,
				"html":           note.Html,
				"createdAt":      note.CreatedAt,
				"now":            time.Now().UTC(),
			})
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return "", err
	}
	return dbRecord.(string), nil
}

func (r *noteRepository) NoteLinkWithContactByExternalId(ctx context.Context, tenant, noteId, contactExternalId, externalSystem string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
				MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem})<-[:IS_LINKED_WITH {externalId:$contactExternalId}]-(c:Contact)
				MATCH (n:Note {id:$noteId})-[:IS_LINKED_WITH]->(e)
				MERGE (c)-[:NOTED]->(n)
				`,
			map[string]interface{}{
				"tenant":            tenant,
				"externalSystem":    externalSystem,
				"noteId":            noteId,
				"contactExternalId": contactExternalId,
			})
		return nil, err
	})
	return err
}

func (r *noteRepository) NoteLinkWithOrganizationByExternalId(ctx context.Context, tenant, noteId, organizationExternalId, externalSystem string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
				MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem})<-[:IS_LINKED_WITH {externalId:$organizationExternalId}]-(org:Organization)
				MATCH (n:Note {id:$noteId})-[:IS_LINKED_WITH]->(e)
				MERGE (org)-[:NOTED]->(n)
				`,
			map[string]interface{}{
				"tenant":                 tenant,
				"externalSystem":         externalSystem,
				"noteId":                 noteId,
				"organizationExternalId": organizationExternalId,
			})
		return nil, err
	})
	return err
}

func (r *noteRepository) NoteLinkWithUserByExternalId(ctx context.Context, tenant, noteId, userExternalId, externalSystem string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
				MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem})<-[:IS_LINKED_WITH {externalId:$userExternalId}]-(u:User)
				MATCH (n:Note {id:$noteId})-[:IS_LINKED_WITH]->(e)
				MERGE (u)-[:CREATED]->(n)
				`,
			map[string]interface{}{
				"tenant":         tenant,
				"externalSystem": externalSystem,
				"noteId":         noteId,
				"userExternalId": userExternalId,
			})
		return nil, err
	})
	return err
}

func (r *noteRepository) NoteLinkWithUserByExternalOwnerId(ctx context.Context, tenant, noteId, userExternalOwnerId, externalSystem string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
				MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem})<-[:IS_LINKED_WITH {externalOwnerId:$userExternalOwnerId}]-(u:User)
				MATCH (n:Note {id:$noteId})-[:IS_LINKED_WITH]->(e)
				MERGE (u)-[:CREATED]->(n)
				`,
			map[string]interface{}{
				"tenant":              tenant,
				"externalSystem":      externalSystem,
				"noteId":              noteId,
				"userExternalOwnerId": userExternalOwnerId,
			})
		return nil, err
	})
	return err
}
