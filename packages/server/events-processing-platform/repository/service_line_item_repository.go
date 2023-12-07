package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type ServiceLineItemRepository interface {
	CreateForContract(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemCreateEvent, isNewVersionForExistingSLI bool, previousQuantity int64, previousPrice float64, previousBilled string) error
	GetServiceLineItemById(ctx context.Context, tenant, serviceLineItemId string) (*dbtype.Node, error)
	Update(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemUpdateEvent) error
	GetAllForContract(ctx context.Context, tenant, contractId string) ([]*neo4j.Node, error)
	Delete(ctx context.Context, tenant, serviceLineItemId string) error
	Close(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemCloseEvent) error
	GetLatestServiceLineItemByParentId(ctx context.Context, tenant, serviceLineItemParentId string, beforeDate time.Time) (*dbtype.Node, error)
}

type serviceLineItemRepository struct {
	driver   *neo4j.DriverWithContext
	database string
}

func NewServiceLineItemRepository(driver *neo4j.DriverWithContext, database string) ServiceLineItemRepository {
	return &serviceLineItemRepository{
		driver:   driver,
		database: database,
	}
}

func (r *serviceLineItemRepository) CreateForContract(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemCreateEvent, isNewVersionForExistingSLI bool, previousQuantity int64, previousPrice float64, previousBilled string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.CreateForContract")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	tracing.LogObjectAsJson(span, "event", evt)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId), log.Bool("isNewVersionForExistingSLI", isNewVersionForExistingSLI), log.Int64("previousQuantity", previousQuantity), log.Float64("previousPrice", previousPrice), log.String("previousBilled", previousBilled))

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:CONTRACT_BELONGS_TO_TENANT]-(c:Contract {id:$contractId})
							MERGE (c)-[:HAS_SERVICE]->(sli:ServiceLineItem {id:$serviceLineItemId})
							ON CREATE SET 
								sli:ServiceLineItem_%s,
								sli.createdAt=$createdAt,
								sli.updatedAt=$updatedAt,
								sli.startedAt=$startedAt,
								sli.endedAt=$endedAt,
								sli.source=$source,
								sli.sourceOfTruth=$sourceOfTruth,
								sli.appSource=$appSource,
								sli.name=$name,
								sli.price=$price,
								sli.quantity=$quantity,
								sli.billed=$billed,
								sli.parentId=$parentId
							`, tenant)
	params := map[string]any{
		"tenant":            tenant,
		"serviceLineItemId": serviceLineItemId,
		"contractId":        evt.ContractId,
		"parentId":          evt.ParentId,
		"createdAt":         evt.CreatedAt,
		"updatedAt":         evt.UpdatedAt,
		"startedAt":         evt.StartedAt,
		"endedAt":           utils.TimePtrFirstNonNilNillableAsAny(evt.EndedAt),
		"source":            helper.GetSource(evt.Source.Source),
		"sourceOfTruth":     helper.GetSourceOfTruth(evt.Source.Source),
		"appSource":         helper.GetAppSource(evt.Source.AppSource),
		"price":             evt.Price,
		"quantity":          evt.Quantity,
		"name":              evt.Name,
		"billed":            evt.Billed,
	}
	if isNewVersionForExistingSLI {
		cypher += `, sli.previousQuantity=$previousQuantity, sli.previousPrice=$previousPrice, sli.previousBilled=$previousBilled`
		params["previousQuantity"] = previousQuantity
		params["previousPrice"] = previousPrice
		params["previousBilled"] = previousBilled
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	return utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
}

func (r *serviceLineItemRepository) Update(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemUpdateEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.Update")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId), log.Object("event", evt))

	cypher := fmt.Sprintf(`MATCH (sli:ServiceLineItem {id:$serviceLineItemId})
							WHERE sli:ServiceLineItem_%s
							SET 
								sli.name = CASE WHEN sli.sourceOfTruth=$sourceOfTruth OR $overwrite=true THEN $name ELSE sli.name END,
								sli.price = CASE WHEN sli.sourceOfTruth=$sourceOfTruth OR $overwrite=true THEN $price ELSE sli.price END,
								sli.quantity = CASE WHEN sli.sourceOfTruth=$sourceOfTruth OR $overwrite=true THEN $quantity ELSE sli.quantity END,
								sli.billed = CASE WHEN sli.sourceOfTruth=$sourceOfTruth OR $overwrite=true THEN $billed ELSE sli.billed END,
								sli.sourceOfTruth = case WHEN $overwrite=true THEN $sourceOfTruth ELSE sli.sourceOfTruth END,
								sli.updatedAt=$updatedAt,
				                sli.comments=$comments
							`, tenant)
	params := map[string]any{
		"serviceLineItemId": serviceLineItemId,
		"updatedAt":         evt.UpdatedAt,
		"price":             evt.Price,
		"quantity":          evt.Quantity,
		"name":              evt.Name,
		"billed":            evt.Billed,
		"comments":          evt.Comments,
		"sourceOfTruth":     helper.GetSourceOfTruth(evt.Source.Source),
		"overwrite":         helper.GetSourceOfTruth(evt.Source.Source) == constants.SourceOpenline,
	}
	span.LogFields(log.String("cypher", cypher), log.Object("params", params))

	return utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
}

func (r *serviceLineItemRepository) GetAllForContract(ctx context.Context, tenant, contractId string) ([]*neo4j.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.GetAllForContract")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("contractId", contractId))

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:CONTRACT_BELONGS_TO_TENANT]-(c:Contract {id:$contractId})-[:HAS_SERVICE]->(sli:ServiceLineItem)
							WHERE sli:ServiceLineItem_%s
							RETURN sli`, tenant)
	params := map[string]any{
		"tenant":     tenant,
		"contractId": contractId,
	}
	span.LogFields(log.String("cypher", cypher), log.Object("params", params))

	session := utils.NewNeo4jReadSession(ctx, *r.driver, utils.WithDatabaseName(r.database))
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, cypher, params); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsFirstValueAsDbNodePtrs(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*neo4j.Node), nil
}

func (r *serviceLineItemRepository) Delete(ctx context.Context, tenant, serviceLineItemId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.Delete")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	cypher := fmt.Sprintf(`MATCH (sli:ServiceLineItem {id:$serviceLineItemId})
							WHERE sli:ServiceLineItem_%s
							DETACH DELETE sli`, tenant)
	params := map[string]any{
		"serviceLineItemId": serviceLineItemId,
	}
	span.LogFields(log.String("cypher", cypher), log.Object("params", params))

	return utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
}

func (r *serviceLineItemRepository) Close(ctx context.Context, tenant, serviceLineItemId string, evt event.ServiceLineItemCloseEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.Close")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	params := map[string]any{
		"serviceLineItemId": serviceLineItemId,
		"updatedAt":         evt.UpdatedAt,
		"endedAt":           evt.EndedAt,
	}
	cypher := fmt.Sprintf(`MATCH (sli:ServiceLineItem {id:$serviceLineItemId})
							WHERE sli:ServiceLineItem_%s SET
							sli.endedAt = $endedAt,
							sli.updatedAt = $updatedAt`, tenant)
	if evt.IsCanceled {
		params["isCanceled"] = evt.IsCanceled
		cypher += `, sli.isCanceled = $isCanceled`
	}
	span.LogFields(log.String("cypher", cypher), log.Object("params", params))

	return utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
}

func (r *serviceLineItemRepository) GetServiceLineItemById(ctx context.Context, tenant, serviceLineItemId string) (*dbtype.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.GetServiceLineItemById")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	cypher := `MATCH(sli:ServiceLineItem {id:$id}) RETURN sli`
	params := map[string]any{
		"tenant": tenant,
		"id":     serviceLineItemId,
	}
	span.LogFields(log.String("query", cypher), log.Object("params", params))

	session := utils.NewNeo4jReadSession(ctx, *r.driver, utils.WithDatabaseName(r.database))
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, cypher, params); err != nil {
			return nil, err
		} else {
			return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.(*dbtype.Node), nil
}

func (r *serviceLineItemRepository) GetLatestServiceLineItemByParentId(ctx context.Context, tenant, serviceLineItemParentId string, beforeDate time.Time) (*dbtype.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemRepository.GetLatestServiceLineItemByParentId")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("serviceLineItemParentId", serviceLineItemParentId), log.Object("beforeDate", beforeDate))

	cypher := `MATCH (sli:ServiceLineItem {parentId:$parentId}) WHERE sli.startedAt < $before RETURN sli ORDER BY sli.startedAt DESC LIMIT 1`
	params := map[string]any{
		"tenant":   tenant,
		"parentId": serviceLineItemParentId,
		"before":   beforeDate.Add(time.Millisecond * 1),
	}
	span.LogFields(log.String("query", cypher), log.Object("params", params))

	session := utils.NewNeo4jReadSession(ctx, *r.driver, utils.WithDatabaseName(r.database))
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, cypher, params); err != nil {
			return nil, err
		} else {
			return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.(*dbtype.Node), nil
}
