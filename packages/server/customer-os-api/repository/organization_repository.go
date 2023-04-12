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

const (
	Relationship_Subsidiary = "SUBSIDIARY_OF"
)

type OrganizationRepository interface {
	Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, organization entity.OrganizationEntity) (*dbtype.Node, error)
	Update(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, organization entity.OrganizationEntity) (*dbtype.Node, error)
	GetOrganizationForJobRole(ctx context.Context, session neo4j.SessionWithContext, tenant, roleId string) (*dbtype.Node, error)
	GetOrganizationById(ctx context.Context, tenant, organizationId string) (*dbtype.Node, error)
	GetPaginatedOrganizations(ctx context.Context, session neo4j.SessionWithContext, tenant string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error)
	GetPaginatedOrganizationsForContact(ctx context.Context, session neo4j.SessionWithContext, tenant, contactId string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error)
	Delete(ctx context.Context, session neo4j.SessionWithContext, tenant, organizationId string) error
	LinkWithOrganizationTypeInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId, organizationTypeId string) error
	UnlinkFromOrganizationTypesInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string) error
	LinkWithDomainsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string, domains []string) error
	UnlinkFromDomainsNotInListInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string, domains []string) error
	MergeOrganizationPropertiesInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, primaryOrganizationId, mergedOrganizationId string, sourceOfTruth entity.DataSource) error
	MergeOrganizationRelationsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, primaryOrganizationId, mergedOrganizationId string) error
	UpdateMergedOrganizationLabelsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, mergedOrganizationId string) error
	GetAllForEmails(ctx context.Context, tenant string, emailIds []string) ([]*utils.DbNodeAndId, error)
	GetAllForPhoneNumbers(ctx context.Context, tenant string, phoneNumberIds []string) ([]*utils.DbNodeAndId, error)
	GetLinkedSubOrganizations(ctx context.Context, tenant, parentOrganizationId, relationName string) ([]*utils.DbNodeAndRelation, error)
	GetLinkedParentOrganizations(ctx context.Context, tenant, organizationId, relationName string) ([]*utils.DbNodeAndRelation, error)
	LinkSubOrganization(ctx context.Context, tenant, organizationId, subOrganizationId, subOrganizationType, relationName string) error
	UnlinkSubOrganization(ctx context.Context, tenant, organizationId, subOrganizationId, relationName string) error

	GetAllCrossTenants(ctx context.Context, size int) ([]*utils.DbNodeAndId, error)
	GetAllOrganizationPhoneNumberRelationships(ctx context.Context, size int) ([]*neo4j.Record, error)
	GetAllOrganizationEmailRelationships(ctx context.Context, size int) ([]*neo4j.Record, error)
}

type organizationRepository struct {
	driver *neo4j.DriverWithContext
}

func NewOrganizationRepository(driver *neo4j.DriverWithContext) OrganizationRepository {
	return &organizationRepository{
		driver: driver,
	}
}

func (r *organizationRepository) Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, organization entity.OrganizationEntity) (*dbtype.Node, error) {
	query := "MATCH (t:Tenant {name:$tenant})" +
		" MERGE (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:randomUUID()})" +
		" ON CREATE SET org.name=$name, " +
		"				org.description=$description, " +
		"				org.readonly=false, " +
		"				org.website=$website, " +
		"				org.industry=$industry, " +
		"				org.isPublic=$isPublic, " +
		" 				org.source=$source, " +
		"				org.sourceOfTruth=$sourceOfTruth, " +
		"				org.appSource=$appSource, " +
		"				org.createdAt=$now, " +
		"				org.updatedAt=$now, " +
		" 				org:%s" +
		" RETURN org"

	queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "Organization_"+tenant),
		map[string]any{
			"tenant":        tenant,
			"name":          organization.Name,
			"description":   organization.Description,
			"readonly":      false,
			"website":       organization.Website,
			"industry":      organization.Industry,
			"isPublic":      organization.IsPublic,
			"source":        organization.Source,
			"sourceOfTruth": organization.SourceOfTruth,
			"appSource":     organization.AppSource,
			"now":           utils.Now(),
		})
	return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
}

func (r *organizationRepository) Update(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, organization entity.OrganizationEntity) (*dbtype.Node, error) {
	query :=
		" MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:$organizationId})" +
			" SET 	org.name=$name, " +
			"		org.description=$description, " +
			"		org.website=$website, " +
			" 		org.industry=$industry, " +
			"		org.isPublic=$isPublic, " +
			"		org.sourceOfTruth=$sourceOfTruth," +
			"		org.updatedAt=datetime({timezone: 'UTC'}) " +
			" RETURN org"

	queryResult, err := tx.Run(ctx, query,
		map[string]any{
			"tenant":         tenant,
			"organizationId": organization.ID,
			"name":           organization.Name,
			"description":    organization.Description,
			"website":        organization.Website,
			"industry":       organization.Industry,
			"isPublic":       organization.IsPublic,
			"sourceOfTruth":  organization.SourceOfTruth,
		})
	return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
}

func (r *organizationRepository) Delete(ctx context.Context, session neo4j.SessionWithContext, tenant, organizationId string) error {
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			OPTIONAL MATCH (org)-[:ASSOCIATED_WITH]->(l:Location)
			OPTIONAL MATCH (l)-[:LOCATED_AT]->(p:Place)	
            DETACH DELETE p, l, org`,
			map[string]interface{}{
				"organizationId": organizationId,
				"tenant":         tenant,
			})
		return nil, err
	})
	return err
}

func (r *organizationRepository) GetOrganizationForJobRole(ctx context.Context, session neo4j.SessionWithContext, tenant, roleId string) (*dbtype.Node, error) {
	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (:JobRole {id:$roleId})-[:ROLE_IN]->(org:Organization)-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			RETURN org`,
			map[string]any{
				"tenant": tenant,
				"roleId": roleId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect(ctx)
		}
	})
	if err != nil {
		return nil, err
	}
	if len(dbRecords.([]*neo4j.Record)) == 0 {
		return nil, nil
	}
	return utils.NodePtr(dbRecords.([]*neo4j.Record)[0].Values[0].(dbtype.Node)), nil
}

func (r *organizationRepository) GetOrganizationById(ctx context.Context, tenant, organizationId string) (*dbtype.Node, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	dbRecord, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			RETURN org`,
			map[string]any{
				"tenant":         tenant,
				"organizationId": organizationId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Single(ctx)
		}
	})
	if err != nil {
		return nil, err
	}
	return utils.NodePtr(dbRecord.(*db.Record).Values[0].(dbtype.Node)), nil
}

func (r *organizationRepository) GetPaginatedOrganizations(ctx context.Context, session neo4j.SessionWithContext, tenant string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error) {
	dbNodesWithTotalCount := new(utils.DbNodesWithTotalCount)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		filterCypherStr, filterParams := filter.CypherFilterFragment("org")
		countParams := map[string]any{
			"tenant": tenant,
		}
		utils.MergeMapToMap(filterParams, countParams)

		queryResult, err := tx.Run(ctx, fmt.Sprintf(
			" MATCH (:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization) "+
				" %s "+
				" RETURN count(org) as count", filterCypherStr),
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
			" MATCH (:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization) "+
				" %s "+
				" RETURN org "+
				" %s "+
				" SKIP $skip LIMIT $limit", filterCypherStr, sorting.SortingCypherFragment("org")),
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

func (r *organizationRepository) GetPaginatedOrganizationsForContact(ctx context.Context, session neo4j.SessionWithContext, tenant, contactId string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error) {
	dbNodesWithTotalCount := new(utils.DbNodesWithTotalCount)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		filterCypherStr, filterParams := filter.CypherFilterFragment("org")
		countParams := map[string]any{
			"tenant":    tenant,
			"contactId": contactId,
		}
		utils.MergeMapToMap(filterParams, countParams)

		queryResult, err := tx.Run(ctx, fmt.Sprintf(
			" MATCH (:Tenant {name:$tenant})<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$contactId})--(org:Organization) "+
				" %s "+
				" RETURN count(org) as count", filterCypherStr),
			countParams)
		if err != nil {
			return nil, err
		}
		count, _ := queryResult.Single(ctx)
		dbNodesWithTotalCount.Count = count.Values[0].(int64)

		params := map[string]any{
			"tenant":    tenant,
			"contactId": contactId,
			"skip":      skip,
			"limit":     limit,
		}
		utils.MergeMapToMap(filterParams, params)

		queryResult, err = tx.Run(ctx, fmt.Sprintf(
			" MATCH (:Tenant {name:$tenant})<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$contactId})--(org:Organization) "+
				" %s "+
				" RETURN org "+
				" %s "+
				" SKIP $skip LIMIT $limit", filterCypherStr, sorting.SortingCypherFragment("org")),
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

func (r *organizationRepository) LinkWithOrganizationTypeInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId, organizationTypeId string) error {
	queryResult, err := tx.Run(ctx, `
			MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})<-[:ORGANIZATION_TYPE_BELONGS_TO_TENANT]-(ot:OrganizationType {id:$organizationTypeId})
			MERGE (org)-[r:IS_OF_TYPE]->(ot)
			RETURN r`,
		map[string]any{
			"tenant":             tenant,
			"organizationId":     organizationId,
			"organizationTypeId": organizationTypeId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *organizationRepository) UnlinkFromOrganizationTypesInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string) error {
	if _, err := tx.Run(ctx, `
			MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}),
				(org)-[r:IS_OF_TYPE]->(:OrganizationType)
			DELETE r`,
		map[string]any{
			"tenant":         tenant,
			"organizationId": organizationId,
		}); err != nil {
		return err
	}
	return nil
}

func (r *organizationRepository) LinkWithDomainsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string, domains []string) error {
	_, err := tx.Run(ctx, `
			MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}),
			(d:Domain) WHERE d.domain IN $domains
			MERGE (org)-[:HAS_DOMAIN]->(d)`,
		map[string]any{
			"tenant":         tenant,
			"organizationId": organizationId,
			"domains":        domains,
		})
	return err
}

func (r *organizationRepository) UnlinkFromDomainsNotInListInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, organizationId string, domains []string) error {
	_, err := tx.Run(ctx, `
			MATCH (:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:$organizationId})-[rel:HAS_DOMAIN]->(d:Domain)
			WHERE NOT d.domain IN $domains
			DELETE rel`,
		map[string]any{
			"tenant":         tenant,
			"organizationId": organizationId,
			"domains":        domains,
		})
	return err
}

func (r *organizationRepository) MergeOrganizationPropertiesInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, primaryOrganizationId, mergedOrganizationId string, sourceOfTruth entity.DataSource) error {
	_, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(primary:Organization {id:$primaryOrganizationId}),
			(t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(merged:Organization {id:$mergedOrganizationId})
			SET primary.website = CASE WHEN primary.website is null OR primary.website = '' THEN merged.website ELSE primary.website END, 
				primary.industry = CASE WHEN primary.industry is null OR primary.industry = '' THEN merged.industry ELSE primary.industry END, 
				primary.name = CASE WHEN primary.name is null OR primary.name = '' THEN merged.name ELSE primary.name END, 
				primary.description = CASE WHEN primary.description is null OR primary.description = '' THEN merged.description ELSE primary.description END, 
				primary.isPublic = CASE WHEN primary.isPublic is null THEN merged.isPublic ELSE primary.isPublic END, 
				primary.sourceOfTruth=$sourceOfTruth,
				primary.updatedAt = $now
			`,
		map[string]any{
			"tenant":                tenant,
			"primaryOrganizationId": primaryOrganizationId,
			"mergedOrganizationId":  mergedOrganizationId,
			"sourceOfTruth":         string(sourceOfTruth),
			"now":                   utils.Now(),
		})
	return err
}

func (r *organizationRepository) MergeOrganizationRelationsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, primaryOrganizationId, mergedOrganizationId string) error {
	matchQuery := "MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(primary:Organization {id:$primaryOrganizationId}), " +
		"(t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(merged:Organization {id:$mergedOrganizationId})"

	params := map[string]any{
		"tenant":                tenant,
		"primaryOrganizationId": primaryOrganizationId,
		"mergedOrganizationId":  mergedOrganizationId,
		"now":                   utils.Now(),
	}

	if _, err := tx.Run(ctx, matchQuery+" "+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:HAS_DOMAIN]->(d:Domain) "+
		" MERGE (primary)-[newRel:HAS_DOMAIN]->(d)"+
		" ON CREATE SET	newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged"+
		" MATCH (merged)<-[rel:ROLE_IN]-(jb:JobRole) "+
		" MERGE (primary)<-[newRel:ROLE_IN]-(jb) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged"+
		" MATCH (merged)-[rel:IS_OF_TYPE]->(ot:OrganizationType) "+
		" MERGE (primary)-[newRel:IS_OF_TYPE]->(ot) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:ASSOCIATED_WITH]->(loc:Location) "+
		" MERGE (primary)-[newRel:ASSOCIATED_WITH]->(loc) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:NOTED]->(n:Note) "+
		" MERGE (primary)-[newRel:NOTED]->(n) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:CREATED]->(n:Note) "+
		" MERGE (primary)-[newRel:CREATED]->(n) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:HAS]->(e:Email) "+
		" MERGE (primary)-[newRel:HAS]->(e) "+
		" ON CREATE SET newRel.primary=rel.primary, "+
		"				newRel.label=rel.label, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:HAS]->(p:PhoneNumber) "+
		" MERGE (primary)-[newRel:HAS]->(p) "+
		" ON CREATE SET newRel.primary=rel.primary, "+
		"				newRel.label=rel.label, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:TAGGED]->(t:Tag) "+
		" MERGE (primary)-[newRel:TAGGED]->(t) "+
		" ON CREATE SET newRel.taggedAt=rel.taggedAt, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)<-[rel:REPORTED_BY]-(i:Issue) "+
		" MERGE (primary)<-[newRel:REPORTED_BY]-(i) "+
		" ON CREATE SET newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:IS_LINKED_WITH]->(ext:ExternalSystem) "+
		" MERGE (primary)-[newRel:IS_LINKED_WITH {externalId:rel.externalId}]->(ext) "+
		" ON CREATE SET newRel.syncDate=rel.syncDate, "+
		"				newRel.externalUrl=rel.externalUrl, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)-[rel:SUBSIDIARY_OF]->(org:Organization) "+
		" WHERE org.id <> primary.id "+
		" MERGE (primary)-[newRel:SUBSIDIARY_OF]->(org) "+
		" ON CREATE SET newRel.type = rel.type, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MATCH (merged)<-[rel:SUBSIDIARY_OF]-(org:Organization) "+
		" WHERE org.id <> primary.id "+
		" MERGE (primary)<-[newRel:SUBSIDIARY_OF]-(org) "+
		" ON CREATE SET newRel.type = rel.type, "+
		"				newRel.mergedFrom = $mergedOrganizationId, "+
		"				newRel.createdAt = $now "+
		"			SET	rel.merged=true", params); err != nil {
		return err
	}

	if _, err := tx.Run(ctx, matchQuery+
		" WITH primary, merged "+
		" MERGE (merged)-[:IS_MERGED_INTO]->(primary)", params); err != nil {
		return err
	}

	return nil
}

func (r *organizationRepository) UpdateMergedOrganizationLabelsInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, mergedOrganizationId string) error {
	query := "MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:$organizationId}) " +
		" SET org:MergedOrganization:%s " +
		" REMOVE org:Organization:%s"

	_, err := tx.Run(ctx, fmt.Sprintf(query, "MergedOrganization_"+tenant, "Organization_"+tenant),
		map[string]any{
			"tenant":         tenant,
			"organizationId": mergedOrganizationId,
		})
	return err
}

func (r *organizationRepository) GetAllForEmails(ctx context.Context, tenant string, emailIds []string) ([]*utils.DbNodeAndId, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization)-[:HAS]->(e:Email)-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]->(t)
			WHERE e.id IN $emailIds
			RETURN o, e.id as emailId ORDER BY o.name`,
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

func (r *organizationRepository) GetAllForPhoneNumbers(ctx context.Context, tenant string, phoneNumberIds []string) ([]*utils.DbNodeAndId, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization)-[:HAS]->(p:PhoneNumber)-[:PHONE_NUMBER_BELONGS_TO_TENANT]->(t)
			WHERE p.id IN $phoneNumberIds
			RETURN o, p.id as phoneNumberId ORDER BY o.name`,
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

func (r *organizationRepository) GetLinkedSubOrganizations(ctx context.Context, tenant, parentOrganizationId, relationName string) ([]*utils.DbNodeAndRelation, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(:Organization {id:$parentOrganizationId})<-[rel:%s]-(org:Organization)-[:ORGANIZATION_BELONGS_TO_TENANT]->(t)
			RETURN org, rel ORDER BY org.name`

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, relationName),
			map[string]any{
				"tenant":               tenant,
				"parentOrganizationId": parentOrganizationId,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeAndRelation(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeAndRelation), err
}

func (r *organizationRepository) GetLinkedParentOrganizations(ctx context.Context, tenant, organizationId, relationName string) ([]*utils.DbNodeAndRelation, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(:Organization {id:$organizationId})-[rel:%s]->(org:Organization)-[:ORGANIZATION_BELONGS_TO_TENANT]->(t)
			RETURN org, rel ORDER BY org.name`

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, relationName),
			map[string]any{
				"tenant":         tenant,
				"organizationId": organizationId,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeAndRelation(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeAndRelation), err
}

func (r *organizationRepository) LinkSubOrganization(ctx context.Context, tenant, organizationId, subOrganizationId, subOrganizationType, relationName string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(parent:Organization {id:$organizationId})," +
		" (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(sub:Organization {id:$subOrganizationId}) " +
		" MERGE (parent)<-[rel:%s]-(sub) " +
		" ON CREATE SET rel.type=$type " +
		" ON MATCH SET rel.type=$type " +
		" return parent.id"

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, relationName),
			map[string]interface{}{
				"tenant":            tenant,
				"organizationId":    organizationId,
				"subOrganizationId": subOrganizationId,
				"type":              subOrganizationType,
				"now":               utils.Now(),
			})
		if err != nil {
			return nil, err
		}
		_, err = queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (r *organizationRepository) UnlinkSubOrganization(ctx context.Context, tenant, organizationId, subOrganizationId, relationName string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(parent:Organization {id:$organizationId})<-[rel:%s]-(sub:Organization {id:$subOrganizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(t)" +
		" DELETE rel "

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, relationName),
			map[string]interface{}{
				"tenant":            tenant,
				"organizationId":    organizationId,
				"subOrganizationId": subOrganizationId,
				"now":               utils.Now(),
			})
		return nil, err
	})
	return err
}

func (r *organizationRepository) GetAllCrossTenants(ctx context.Context, size int) ([]*utils.DbNodeAndId, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (org:Organization)-[:ORGANIZATION_BELONGS_TO_TENANT]->(t:Tenant)
 			WHERE (org.syncedWithEventStore is null or org.syncedWithEventStore=false)
			RETURN org, t.name limit $size`,
			map[string]any{
				"size": size,
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

func (r *organizationRepository) GetAllOrganizationPhoneNumberRelationships(ctx context.Context, size int) ([]*neo4j.Record, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization)-[rel:HAS]->(p:PhoneNumber)
 			WHERE (rel.syncedWithEventStore is null or rel.syncedWithEventStore=false)
			RETURN rel, org.id, p.id, t.name limit $size`,
			map[string]any{
				"size": size,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect(ctx)
		}
	})
	if err != nil {
		return nil, err
	}
	return dbRecords.([]*neo4j.Record), err
}

func (r *organizationRepository) GetAllOrganizationEmailRelationships(ctx context.Context, size int) ([]*neo4j.Record, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `
			MATCH (t:Tenant)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization)-[rel:HAS]->(e:Email)
 			WHERE (rel.syncedWithEventStore is null or rel.syncedWithEventStore=false)
			RETURN rel, org.id, e.id, t.name limit $size`,
			map[string]any{
				"size": size,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect(ctx)
		}
	})
	if err != nil {
		return nil, err
	}
	return dbRecords.([]*neo4j.Record), err
}
