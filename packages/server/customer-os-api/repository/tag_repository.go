package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
	"golang.org/x/net/context"
)

type TagRepository interface {
	Merge(ctx context.Context, tenant string, tag entity.TagEntity) (*dbtype.Node, error)
	Update(ctx context.Context, tenant string, tag entity.TagEntity) (*dbtype.Node, error)
	UnlinkAndDelete(ctx context.Context, tenant string, tagId string) error
	GetAll(ctx context.Context, tenant string) ([]*dbtype.Node, error)
	GetForContact(ctx context.Context, tenant, contactId string) ([]*dbtype.Node, error)
}

type tagRepository struct {
	driver *neo4j.DriverWithContext
}

func NewTagRepository(driver *neo4j.DriverWithContext) TagRepository {
	return &tagRepository{
		driver: driver,
	}
}

func (r *tagRepository) Merge(ctx context.Context, tenant string, tag entity.TagEntity) (*dbtype.Node, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (t:Tenant {name:$tenant}) " +
		" MERGE (t)<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag {name:$name}) " +
		" ON CREATE SET " +
		"  tag.id=randomUUID(), " +
		"  tag.createdAt=$now, " +
		"  tag.updatedAt=$now, " +
		"  tag.source=$source, " +
		"  tag.appSource=$appSource, " +
		"  tag:%s" +
		" RETURN tag"

	if result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "Tag_"+tenant),
			map[string]any{
				"tenant":    tenant,
				"name":      tag.Name,
				"source":    tag.Source,
				"appSource": tag.AppSource,
				"now":       utils.Now(),
			})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil {
		return nil, err
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *tagRepository) Update(ctx context.Context, tenant string, tag entity.TagEntity) (*dbtype.Node, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	if result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag {id:$id})
			SET tag.name=$name, tag.updatedAt=$now
			RETURN tag`,
			map[string]any{
				"tenant": tenant,
				"id":     tag.Id,
				"name":   tag.Name,
				"now":    utils.Now(),
			})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil {
		return nil, err
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *tagRepository) UnlinkAndDelete(ctx context.Context, tenant string, tagId string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag {id:$id})
			DETACH DELETE tag`,
			map[string]any{
				"tenant": tenant,
				"id":     tagId,
			})
		return nil, err
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func (r *tagRepository) GetAll(ctx context.Context, tenant string) ([]*dbtype.Node, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag)
			RETURN tag ORDER BY tag.name`,
			map[string]any{
				"tenant": tenant,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsFirstValueAsNodePtrs(ctx, queryResult, err)
		}
	})
	return result.([]*dbtype.Node), err
}

func (r *tagRepository) GetForContact(ctx context.Context, tenant, contactId string) ([]*dbtype.Node, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$contactId})-[:TAGGED]->(tag:Tag)
			RETURN tag ORDER BY tag.name`,
			map[string]any{
				"tenant":    tenant,
				"contactId": contactId,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsFirstValueAsNodePtrs(ctx, queryResult, err)
		}
	})
	return result.([]*dbtype.Node), err
}
