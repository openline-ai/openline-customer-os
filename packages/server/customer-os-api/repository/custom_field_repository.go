package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/customer-os-api/utils"
)

type CustomFieldRepository interface {
	AddTextCustomFieldToContactInTx(contactId string, input entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error)
	MergeTextCustomFieldToContactInTx(tenant, contactId string, entity *entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error)
	MergeTextCustomFieldToFieldSetInTx(tenant, contactId, fieldSet string, entity *entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error)

	LinkWithCustomFieldDefinitionForContactInTx(fieldId, contactId, definitionId string, tx neo4j.Transaction) error
	LinkWithCustomFieldDefinitionForFieldSetInTx(fieldId, fieldSetId, definitionId string, tx neo4j.Transaction) error
}

type customFieldRepository struct {
	driver *neo4j.Driver
	repos  *RepositoryContainer
}

func NewCustomFieldRepository(driver *neo4j.Driver, repos *RepositoryContainer) CustomFieldRepository {
	return &customFieldRepository{
		driver: driver,
		repos:  repos,
	}
}

func (r *customFieldRepository) AddTextCustomFieldToContactInTx(contactId string, input entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error) {
	queryResult, err := tx.Run(`
			MATCH (c:Contact {id:$contactId})
			MERGE (f:TextCustomField {
				id: randomUUID(),
				name: $name,
				value: $value
			})<-[:HAS_TEXT_PROPERTY]-(c)
			RETURN f`,
		map[string]any{
			"contactId": contactId,
			"name":      input.Name,
			"value":     input.Value,
		})
	return utils.ExtractSingleRecordFirstValue(queryResult, err)
}

func (r *customFieldRepository) MergeTextCustomFieldToContactInTx(tenant, contactId string, entity *entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error) {
	queryResult, err := tx.Run(`
			MATCH (c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			MERGE (f:TextCustomField {name: $name})<-[:HAS_TEXT_PROPERTY]-(c)
            ON CREATE SET f.value=$value, f.id=randomUUID()
            ON MATCH SET f.value=$value
			RETURN f`,
		map[string]any{
			"tenant":    tenant,
			"contactId": contactId,
			"name":      entity.Name,
			"value":     entity.Value,
		})
	return utils.ExtractSingleRecordFirstValue(queryResult, err)
}

func (r *customFieldRepository) MergeTextCustomFieldToFieldSetInTx(tenant, contactId, fieldSetId string, entity *entity.TextCustomFieldEntity, tx neo4j.Transaction) (any, error) {
	queryResult, err := tx.Run(`
			MATCH (s:FieldSet {id:$fieldSetId})<-[:HAS_COMPLEX_PROPERTY]-(c:Contact {id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(:Tenant {name:$tenant})
			MERGE (f:TextCustomField {name: $name})<-[:HAS_TEXT_PROPERTY]-(s)
            ON CREATE SET f.value=$value, f.id=randomUUID()
            ON MATCH SET f.value=$value
			RETURN f`,
		map[string]any{
			"tenant":     tenant,
			"contactId":  contactId,
			"fieldSetId": fieldSetId,
			"name":       entity.Name,
			"value":      entity.Value,
		})
	return utils.ExtractSingleRecordFirstValue(queryResult, err)
}

func (r *customFieldRepository) LinkWithCustomFieldDefinitionForContactInTx(fieldId, contactId, definitionId string, tx neo4j.Transaction) error {
	queryResult, err := tx.Run(`
			MATCH (f:TextCustomField {id:$fieldId})<-[:HAS_TEXT_PROPERTY]-(c:Contact {id:$contactId}),
				  (c)-[:IS_DEFINED_BY]->(e:EntityDefinition),
				  (e)-[:CONTAINS]->(d:CustomFieldDefinition {id:$customFieldDefinitionId})
			MERGE (f)-[r:IS_DEFINED_BY]->(d)
			RETURN r`,
		map[string]any{
			"customFieldDefinitionId": definitionId,
			"fieldId":                 fieldId,
			"contactId":               contactId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single()
	return err
}

func (r *customFieldRepository) LinkWithCustomFieldDefinitionForFieldSetInTx(fieldId, fieldSetId, definitionId string, tx neo4j.Transaction) error {
	queryResult, err := tx.Run(`
			MATCH (f:TextCustomField {id:$fieldId})<-[:HAS_TEXT_PROPERTY]-(s:FieldSet {id:$fieldSetId}),
				  (s)-[:IS_DEFINED_BY]->(e:FieldSetDefinition),
				  (e)-[:CONTAINS]->(d:CustomFieldDefinition {id:$customFieldDefinitionId})
			MERGE (f)-[r:IS_DEFINED_BY]->(d)
			RETURN r`,
		map[string]any{
			"customFieldDefinitionId": definitionId,
			"fieldId":                 fieldId,
			"fieldSetId":              fieldSetId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single()
	return err
}
