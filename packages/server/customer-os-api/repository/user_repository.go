package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, entity entity.UserEntity) (*dbtype.Node, error)
	Update(ctx context.Context, session neo4j.SessionWithContext, tenant string, entity entity.UserEntity) (*dbtype.Node, error)
	FindUserByEmail(ctx context.Context, session neo4j.SessionWithContext, tenant string, email string) (*dbtype.Node, error)
	GetOwnerForContact(ctx context.Context, tx neo4j.ManagedTransaction, tenant, contactId string) (*dbtype.Node, error)
	GetCreatorForNote(ctx context.Context, tx neo4j.ManagedTransaction, tenant, noteId string) (*dbtype.Node, error)
	GetPaginatedUsers(ctx context.Context, session neo4j.SessionWithContext, tenant string, skip, limit int, filter *utils.CypherFilter, sort *utils.CypherSort) (*utils.DbNodesWithTotalCount, error)
	GetById(ctx context.Context, session neo4j.SessionWithContext, tenant, userId string) (*dbtype.Node, error)
	GetAllForConversation(ctx context.Context, session neo4j.SessionWithContext, tenant, conversationId string) ([]*dbtype.Node, error)
	GetAllForEmails(ctx context.Context, tenant string, emailIds []string) ([]*utils.DbNodeAndId, error)
	GetAllForPhoneNumbers(ctx context.Context, tenant string, phoneNumberIds []string) ([]*utils.DbNodeAndId, error)
}

type userRepository struct {
	driver *neo4j.DriverWithContext
}

func NewUserRepository(driver *neo4j.DriverWithContext) UserRepository {
	return &userRepository{
		driver: driver,
	}
}

func (r *userRepository) Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, entity entity.UserEntity) (*dbtype.Node, error) {
	query := "MATCH (t:Tenant {name:$tenant}) " +
		" MERGE (u:User {id: randomUUID()})-[:USER_BELONGS_TO_TENANT]->(t) " +
		" ON CREATE SET u.firstName=$firstName, " +
		"				u.lastName=$lastName, " +
		"				u.createdAt=$now, " +
		"				u.updatedAt=$now, " +
		" 				u.source=$source, " +
		"				u.sourceOfTruth=$sourceOfTruth, " +
		"				u:%s" +
		" RETURN u"

	queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "User_"+tenant),
		map[string]any{
			"tenant":        tenant,
			"firstName":     entity.FirstName,
			"lastName":      entity.LastName,
			"source":        entity.Source,
			"sourceOfTruth": entity.SourceOfTruth,
			"now":           utils.Now(),
		})
	return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
}

func (r *userRepository) Update(ctx context.Context, session neo4j.SessionWithContext, tenant string, entity entity.UserEntity) (*dbtype.Node, error) {
	query := "MATCH (u:User {id:$userId})-[:USER_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant}) " +
		" SET 	u.firstName=$firstName, " +
		"		u.lastName=$lastName, " +
		"		u.updatedAt=datetime({timezone: 'UTC'}), " +
		"		u.sourceOfTruth=$sourceOfTruth " +
		" RETURN u"

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		queryResult, err := tx.Run(ctx, query,
			map[string]any{
				"userId":        entity.Id,
				"tenant":        tenant,
				"firstName":     entity.FirstName,
				"lastName":      entity.LastName,
				"sourceOfTruth": entity.SourceOfTruth,
			})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	})
	if err != nil {
		return nil, err
	}
	return result.(*dbtype.Node), nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, session neo4j.SessionWithContext, tenant string, email string) (*dbtype.Node, error) {
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		queryResult, err := tx.Run(ctx, `
			MATCH (:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:HAS]->(e:Email) 
			WHERE e.email=$email OR e.rawEmail=$email
			RETURN DISTINCT u limit 1`,
			map[string]any{
				"tenant": tenant,
				"email":  email,
			})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	})
	if err != nil {
		return nil, err
	}
	return result.(*dbtype.Node), nil
}

func (r *userRepository) GetOwnerForContact(ctx context.Context, tx neo4j.ManagedTransaction, tenant, contactId string) (*dbtype.Node, error) {
	if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$contactId})<-[:OWNS]-(u:User)
			RETURN u`,
		map[string]any{
			"tenant":    tenant,
			"contactId": contactId,
		}); err != nil {
		return nil, err
	} else {
		dbRecords, err := queryResult.Collect(ctx)
		if err != nil {
			return nil, err
		} else if len(dbRecords) == 0 {
			return nil, nil
		} else {
			return utils.NodePtr(dbRecords[0].Values[0].(dbtype.Node)), nil
		}
	}
}

func (r *userRepository) GetCreatorForNote(ctx context.Context, tx neo4j.ManagedTransaction, tenant, noteId string) (*dbtype.Node, error) {
	if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:CREATED]->(n:Note {id:$noteId})
			RETURN u`,
		map[string]any{
			"tenant": tenant,
			"noteId": noteId,
		}); err != nil {
		return nil, err
	} else {
		dbRecords, err := queryResult.Collect(ctx)
		if err != nil {
			return nil, err
		} else if len(dbRecords) == 0 {
			return nil, nil
		} else {
			return utils.NodePtr(dbRecords[0].Values[0].(dbtype.Node)), nil
		}
	}
}

func (r *userRepository) GetPaginatedUsers(ctx context.Context, session neo4j.SessionWithContext, tenant string, skip, limit int, filter *utils.CypherFilter, sort *utils.CypherSort) (*utils.DbNodesWithTotalCount, error) {
	dbNodesWithTotalCount := new(utils.DbNodesWithTotalCount)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		filterCypherStr, filterParams := filter.CypherFilterFragment("u")
		countParams := map[string]any{
			"tenant": tenant,
		}
		utils.MergeMapToMap(filterParams, countParams)
		queryResult, err := tx.Run(ctx, fmt.Sprintf("MATCH (:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User) %s RETURN count(u) as count", filterCypherStr),
			countParams)
		if err != nil {
			return nil, err
		}
		count, _ := queryResult.Single(ctx)
		dbNodesWithTotalCount.Count = count.Values[0].(int64)

		params := map[string]any{
			"tenant": tenant,
			"skip":   skip,
			"limit":  limit,
		}
		utils.MergeMapToMap(filterParams, params)

		queryResult, err = tx.Run(ctx, fmt.Sprintf(
			"MATCH (:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User) "+
				" %s "+
				" RETURN u "+
				" %s "+
				" SKIP $skip LIMIT $limit", filterCypherStr, sort.SortingCypherFragment("u")),
			params)
		return queryResult.Collect(ctx)
	})
	if err != nil {
		return nil, err
	}
	for _, v := range dbRecords.([]*neo4j.Record) {
		dbNodesWithTotalCount.Nodes = append(dbNodesWithTotalCount.Nodes, utils.NodePtr(v.Values[0].(neo4j.Node)))
	}
	return dbNodesWithTotalCount, nil
}

func (r *userRepository) GetById(ctx context.Context, session neo4j.SessionWithContext, tenant, userId string) (*dbtype.Node, error) {
	dbRecord, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User {id:$userId})
			RETURN u`,
			map[string]any{
				"tenant": tenant,
				"userId": userId,
			})
		if err != nil {
			return nil, err
		}
		return queryResult.Single(ctx)
	})
	return utils.NodePtr(dbRecord.(*db.Record).Values[0].(dbtype.Node)), err
}

func (r *userRepository) GetAllForConversation(ctx context.Context, session neo4j.SessionWithContext, tenant, conversationId string) ([]*dbtype.Node, error) {
	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:PARTICIPATES]->(o:Conversation {id:$conversationId})
			RETURN u`,
			map[string]any{
				"tenant":         tenant,
				"conversationId": conversationId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect(ctx)
		}
	})
	if err != nil {
		return nil, err
	}
	dbNodes := []*dbtype.Node{}
	for _, v := range dbRecords.([]*neo4j.Record) {
		if v.Values[0] != nil {
			dbNodes = append(dbNodes, utils.NodePtr(v.Values[0].(dbtype.Node)))
		}
	}
	return dbNodes, err
}

func (r *userRepository) GetAllForEmails(ctx context.Context, tenant string, emailIds []string) ([]*utils.DbNodeAndId, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:HAS]->(e:Email)-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]->(t)
			WHERE e.id IN $emailIds
			RETURN u, e.id as emailId ORDER BY u.firstName, u.lastName`,
			map[string]any{
				"tenant":   tenant,
				"emailIds": emailIds,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeAndId), err
}

func (r *userRepository) GetAllForPhoneNumbers(ctx context.Context, tenant string, phoneNumberIds []string) ([]*utils.DbNodeAndId, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:HAS]->(p:PhoneNumber)-[:PHONE_NUMBER_BELONGS_TO_TENANT]->(t)
			WHERE p.id IN $phoneNumberIds
			RETURN u, p.id as phoneNumberId ORDER BY u.firstName, u.lastName`,
			map[string]any{
				"tenant":         tenant,
				"phoneNumberIds": phoneNumberIds,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeAndId), err
}
