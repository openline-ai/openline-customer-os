package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"time"
)

type OrganizationRepository interface {
	GetOrganizationWithDomain(ctx context.Context, tenant, domainId string) (*dbtype.Node, error)
	CreateOrganization(ctx context.Context, tenant, name, source, sourceOfTruth, appSource string, date time.Time) (*dbtype.Node, error)
	LinkDomainToOrganization(ctx context.Context, tenant, domainId, organizationId string) error
}

type organizationRepository struct {
	driver *neo4j.DriverWithContext
}

func NewOrganizationRepository(driver *neo4j.DriverWithContext) OrganizationRepository {
	return &organizationRepository{
		driver: driver,
	}
}

func (r *organizationRepository) GetOrganizationWithDomain(ctx context.Context, tenant, domainId string) (*dbtype.Node, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization)-[:HAS_DOMAIN]->(d:Domain{id:$domainId}) RETURN o`

	if result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, query, map[string]any{
			"tenant":   tenant,
			"domainId": domainId,
		})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil && err.Error() == "Result contains no more records" {
		return nil, nil
	} else if err != nil {
		return nil, nil
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *organizationRepository) CreateOrganization(ctx context.Context, tenant, name, source, sourceOfTruth, appSource string, date time.Time) (*dbtype.Node, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (t:Tenant {name:$tenant}) " +
		" MERGE (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:randomUUID()}) " +
		" ON CREATE SET org.createdAt=$now, " +
		"				org.updatedAt=$now, " +
		"               org.tenantOrganization=false, " +
		"               org.name=$name, " +
		"				org.source=$source, " +
		"				org.sourceOfTruth=$source, " +
		"				org.appSource=$appSource, " +
		"				org:%s " +
		" RETURN org"

	if result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, "Organization_"+tenant),
			map[string]interface{}{
				"tenant":        tenant,
				"name":          name,
				"source":        source,
				"sourceOfTruth": sourceOfTruth,
				"appSource":     appSource,
				"now":           date,
			})
		if err != nil {
			return nil, err
		}
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil {
		return nil, err
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *organizationRepository) LinkDomainToOrganization(ctx context.Context, tenant, domainId, organizationId string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$organizationId}) " +
		" MATCH (d:Domain {id:$domainId}) " +
		" MERGE (o)-[:HAS_DOMAIN]->(d)"

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"tenant":         tenant,
			"domainId":       domainId,
			"organizationId": organizationId,
		})
		return nil, err
	})
	return err
}
