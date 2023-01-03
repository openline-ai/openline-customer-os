package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

type CompanyRepository interface {
	Create(session neo4j.Session, tenant string, company entity.CompanyEntity) (*dbtype.Node, error)
	Update(session neo4j.Session, tenant string, company entity.CompanyEntity) (*dbtype.Node, error)
	GetCompanyForRole(session neo4j.Session, tenant, roleId string) (*dbtype.Node, error)
	GetCompanyById(session neo4j.Session, tenant, companyId string) (*dbtype.Node, error)
	GetPaginatedCompanies(session neo4j.Session, tenant string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error)
	Delete(session neo4j.Session, tenant, companyId string) error
}

type companyRepository struct {
	driver *neo4j.Driver
}

func (r *companyRepository) Delete(session neo4j.Session, tenant, companyId string) error {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		_, err := tx.Run(`
			MATCH (c:Company {id:$companyId})-[:COMPANY_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			OPTIONAL MATCH (c)-[:LOCATED_AT]->(a:Address)
            DETACH DELETE a, c`,
			map[string]interface{}{
				"companyId": companyId,
				"tenant":    tenant,
			})
		return nil, err
	})
	return err
}

func NewCompanyRepository(driver *neo4j.Driver) CompanyRepository {
	return &companyRepository{
		driver: driver,
	}
}

func (r *companyRepository) Create(session neo4j.Session, tenant string, company entity.CompanyEntity) (*dbtype.Node, error) {
	query := "MATCH (t:Tenant {name:$tenant})" +
		" MERGE (t)<-[:COMPANY_BELONGS_TO_TENANT]-(c:Company {id:randomUUID()})" +
		" ON CREATE SET c.name=$name, c.description=$description, c.readonly=false," +
		" c.domain=$domain, c.website=$website, c.industry=$industry, c.isPublic=$isPublic, c:%s" +
		" RETURN c"

	if result, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		queryResult, err := tx.Run(fmt.Sprintf(query, "Company_"+tenant),
			map[string]any{
				"tenant":      tenant,
				"name":        company.Name,
				"description": company.Description,
				"readonly":    false,
				"domain":      company.Domain,
				"website":     company.Website,
				"industry":    company.Industry,
				"isPublic":    company.IsPublic,
			})
		return utils.ExtractSingleRecordFirstValueAsNode(queryResult, err)
	}); err != nil {
		return nil, err
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *companyRepository) Update(session neo4j.Session, tenant string, company entity.CompanyEntity) (*dbtype.Node, error) {
	query :=
		" MATCH (t:Tenant {name:$tenant})<-[:COMPANY_BELONGS_TO_TENANT]-(c:Company {id:$companyId})" +
			" SET c.name=$name, c.description=$description, c.domain=$domain, c.website=$website, " +
			" c.industry=$industry, c.isPublic=$isPublic " +
			" RETURN c"

	if result, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		queryResult, err := tx.Run(query,
			map[string]any{
				"tenant":      tenant,
				"companyId":   company.Id,
				"name":        company.Name,
				"description": company.Description,
				"domain":      company.Domain,
				"website":     company.Website,
				"industry":    company.Industry,
				"isPublic":    company.IsPublic,
			})
		return utils.ExtractSingleRecordFirstValueAsNode(queryResult, err)
	}); err != nil {
		return nil, err
	} else {
		return result.(*dbtype.Node), nil
	}
}

func (r *companyRepository) GetCompanyForRole(session neo4j.Session, tenant, roleId string) (*dbtype.Node, error) {
	dbRecords, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		if queryResult, err := tx.Run(`
			MATCH (r:Role {id:$roleId})-[:WORKS]->(co:Company)-[:COMPANY_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant})
			RETURN co`,
			map[string]any{
				"tenant": tenant,
				"roleId": roleId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect()
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

func (r *companyRepository) GetCompanyById(session neo4j.Session, tenant, companyId string) (*dbtype.Node, error) {
	dbRecord, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		if queryResult, err := tx.Run(`
			MATCH (co:Company {id:$companyId})-[:COMPANY_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant})
			RETURN co`,
			map[string]any{
				"tenant":    tenant,
				"companyId": companyId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Single()
		}
	})
	if err != nil {
		return nil, err
	}
	return utils.NodePtr(dbRecord.(*db.Record).Values[0].(dbtype.Node)), nil
}

func (r *companyRepository) GetPaginatedCompanies(session neo4j.Session, tenant string, skip, limit int, filter *utils.CypherFilter, sorting *utils.CypherSort) (*utils.DbNodesWithTotalCount, error) {
	dbNodesWithTotalCount := new(utils.DbNodesWithTotalCount)

	dbRecords, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		filterCypherStr, filterParams := filter.CypherFilterFragment("c")
		countParams := map[string]any{
			"tenant": tenant,
		}
		utils.MergeMapToMap(filterParams, countParams)

		queryResult, err := tx.Run(fmt.Sprintf(
			" MATCH (:Tenant {name:$tenant})<-[:COMPANY_BELONGS_TO_TENANT]-(c:Company) "+
				" %s "+
				" RETURN count(c) as count", filterCypherStr),
			countParams)
		if err != nil {
			return nil, err
		}
		count, _ := queryResult.Single()
		dbNodesWithTotalCount.Count = count.Values[0].(int64)

		params := map[string]any{
			"tenant": tenant,
			"skip":   skip,
			"limit":  limit,
		}
		utils.MergeMapToMap(filterParams, params)

		queryResult, err = tx.Run(fmt.Sprintf(
			" MATCH (:Tenant {name:$tenant})<-[:COMPANY_BELONGS_TO_TENANT]-(c:Company) "+
				" %s "+
				" RETURN c "+
				" %s "+
				" SKIP $skip LIMIT $limit", filterCypherStr, sorting.SortingCypherFragment("c")),
			params)
		return queryResult.Collect()
	})
	if err != nil {
		return nil, err
	}
	for _, v := range dbRecords.([]*neo4j.Record) {
		dbNodesWithTotalCount.Nodes = append(dbNodesWithTotalCount.Nodes, utils.NodePtr(v.Values[0].(neo4j.Node)))
	}
	return dbNodesWithTotalCount, nil
}
