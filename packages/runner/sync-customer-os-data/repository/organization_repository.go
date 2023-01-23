package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/utils"
	"time"
)

type OrganizationRepository interface {
	MergeOrganization(tenant string, syncDate time.Time, organization entity.OrganizationData) (string, error)
	MergeOrganizationType(tenant, organizationId, organizationTypeName string) error
	MergeOrganizationDefaultPlace(tenant, organizationId string, organization entity.OrganizationData) error
}

type organizationRepository struct {
	driver *neo4j.Driver
}

func NewOrganizationRepository(driver *neo4j.Driver) OrganizationRepository {
	return &organizationRepository{
		driver: driver,
	}
}

func (r *organizationRepository) MergeOrganization(tenant string, syncDate time.Time, organization entity.OrganizationData) (string, error) {
	session := utils.NewNeo4jWriteSession(*r.driver)
	defer session.Close()

	// Create new Organization if it does not exist
	// If Organization exists, and sourceOfTruth is acceptable then update it.
	//   otherwise create/update AlternateOrganization for incoming source, with a new relationship 'ALTERNATE'
	// Link Organization with Tenant
	query := "MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystem}) " +
		" MERGE (org:Organization)-[r:IS_LINKED_WITH {externalId:$externalId}]->(e) " +
		" ON CREATE SET r.externalId=$externalId, r.syncDate=$syncDate, " +
		"				org.id=randomUUID(), org.createdAt=$createdAt, " +
		"               org.name=$name, org.description=$description, org.domain=$domain, " +
		"               org.website=$website, org.industry=$industry, org.isPublic=$isPublic, " +
		"				org.source=$source, org.sourceOfTruth=$sourceOfTruth, org.appSource=$appSource, " +
		"				org:%s " +
		" ON MATCH SET 	r.syncDate = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $syncDate ELSE r.syncDate END, " +
		"				org.name = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $name ELSE org.name END, " +
		"				org.description = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $description ELSE org.description END, " +
		"				org.domain = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $domain ELSE org.domain END, " +
		"				org.website = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $website ELSE org.website END, " +
		"				org.industry = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $industry ELSE org.industry END, " +
		"				org.isPublic = CASE WHEN org.sourceOfTruth=$sourceOfTruth THEN $isPublic ELSE org.isPublic END " +
		" WITH org, t " +
		" MERGE (org)-[:ORGANIZATION_BELONGS_TO_TENANT]->(t) " +
		" WITH org " +
		" FOREACH (x in CASE WHEN org.sourceOfTruth <> $sourceOfTruth THEN [org] ELSE [] END | " +
		"  MERGE (x)-[:ALTERNATE]->(alt:AlternateOrganization {source:$source, id:x.id}) " +
		"    SET alt.updatedAt=$now, alt.appSource=$appSource, " +
		" 		alt.name=$name, alt.description=$description, org.domain=$domain, org.website=$website, org.industry=$industry, org.isPublic=$isPublic " +
		") " +
		" RETURN org.id"

	dbRecord, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		queryResult, err := tx.Run(fmt.Sprintf(query, "Organization_"+tenant),
			map[string]interface{}{
				"tenant":         tenant,
				"externalSystem": organization.ExternalSystem,
				"externalId":     organization.ExternalId,
				"syncDate":       syncDate,
				"name":           organization.Name,
				"description":    organization.Description,
				"createdAt":      organization.CreatedAt,
				"domain":         organization.Domain,
				"website":        organization.Website,
				"industry":       organization.Industry,
				"isPublic":       organization.IsPublic,
				"source":         organization.ExternalSystem,
				"sourceOfTruth":  organization.ExternalSystem,
				"appSource":      organization.ExternalSystem,
				"now":            time.Now().UTC(),
			})
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single()
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

func (r *organizationRepository) MergeOrganizationType(tenant, organizationId, organizationTypeName string) error {
	session := utils.NewNeo4jWriteSession(*r.driver)
	defer session.Close()

	query := "MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant}) " +
		" MERGE (ot:OrganizationType {name:$organizationTypeName})-[:ORGANIZATION_TYPE_BELONGS_TO_TENANT]->(t) " +
		" ON CREATE SET ot.id=randomUUID() " +
		" WITH org, ot " +
		" MERGE (org)-[r:IS_OF_TYPE]->(ot) " +
		" return r"

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		queryResult, err := tx.Run(query,
			map[string]interface{}{
				"tenant":               tenant,
				"organizationId":       organizationId,
				"organizationTypeName": organizationTypeName,
			})
		if err != nil {
			return nil, err
		}
		_, err = queryResult.Single()
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (r *organizationRepository) MergeOrganizationDefaultPlace(tenant, organizationId string, organization entity.OrganizationData) error {
	session := utils.NewNeo4jWriteSession(*r.driver)
	defer session.Close()

	// Create new Place & Location if it does not exist with given source
	// If Place exists, and sourceOfTruth is acceptable then update it.
	//   otherwise create/update AlternatePlace for incoming source, with a new relationship 'ALTERNATE'
	// !!! Current assumption - there is single Location and place with source of externalSystem
	query := "MATCH (org:Organization {id:$organizationId})-[:ORGANIZATION_BELONGS_TO_TENANT]->(:Tenant {name:$tenant}) " +
		" MERGE (org)-[:ASSOCIATED_WITH]->(loc:Location {source:$source})-[:LOCATED_AT]->(p:Place {source:$source}) " +
		" ON CREATE SET p.id=randomUUID(), p.appSource=$appSource, p.sourceOfTruth=$sourceOfTruth, " +
		"				p.country=$country, p.state=$state, p.city=$city, p.address=$address, " +
		"				p.address2=$address2, p.zip=$zip, p.phone=$phone, p.createdAt=$createdAt, p.updatedAt=$createdAt, p:%s, " +
		"				loc.id=randomUUID(), loc.appSource=$appSource, loc.sourceOfTruth=$sourceOfTruth, loc.name=$locationName, " +
		"				loc.createdAt=$createdAt, loc.updatedAt=$createdAt, loc:%s " +
		" ON MATCH SET 	" +
		"             p.country = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $country ELSE p.country END, " +
		"             p.state = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $state ELSE p.state END, " +
		"             p.city = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $city ELSE p.city END, " +
		"             p.address = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $address ELSE p.address END, " +
		"             p.address2 = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $address2 ELSE p.address2 END, " +
		"             p.zip = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $zip ELSE p.zip END, " +
		"             p.phone = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $phone ELSE p.phone END, " +
		"             p.updatedAt = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $now ELSE p.updatedAt END, " +
		"             loc.updatedAt = CASE WHEN p.sourceOfTruth=$sourceOfTruth THEN $now ELSE loc.updatedAt END " +
		" WITH p " +
		" FOREACH (x in CASE WHEN p.sourceOfTruth <> $sourceOfTruth THEN [p] ELSE [] END | " +
		"  MERGE (x)-[:ALTERNATE]->(alt:AlternatePlace {source:$source, id:x.id}) " +
		"    SET alt.updatedAt=$now, alt.appSource=$appSource, " +
		" alt.country=$country, alt.state=$state, alt.city=$city, alt.address=$address, alt.address2=$address2, alt.zip=$zip, alt.phone=$phone " +
		") "

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		_, err := tx.Run(fmt.Sprintf(query, "Place_"+tenant, "Location_"+tenant),
			map[string]interface{}{
				"tenant":         tenant,
				"organizationId": organizationId,
				"country":        organization.Country,
				"state":          organization.State,
				"city":           organization.City,
				"address":        organization.Address,
				"address2":       organization.Address2,
				"zip":            organization.Zip,
				"phone":          organization.Phone,
				"source":         organization.ExternalSystem,
				"sourceOfTruth":  organization.ExternalSystem,
				"appSource":      organization.ExternalSystem,
				"locationName":   organization.DefaultLocationName,
				"createdAt":      organization.CreatedAt,
				"now":            time.Now().UTC(),
			})
		return nil, err
	})
	return err
}
