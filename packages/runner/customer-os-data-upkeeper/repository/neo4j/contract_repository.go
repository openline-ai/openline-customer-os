package neo4j

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type TenantAndContractId struct {
	Tenant     string
	ContractId string
}

type ContractRepository interface {
	GetContractsForStatusRenewal(ctx context.Context, referenceTime time.Time) ([]TenantAndContractId, error)
	GetContractsForRenewalRollout(ctx context.Context, referenceTime time.Time) ([]TenantAndContractId, error)
	// Deprecated, use neo4j module instead
	MarkPayNotificationRequested(ctx context.Context, tenant, invoiceId string, requestedAt time.Time) error
}

type contractRepository struct {
	driver *neo4j.DriverWithContext
}

func NewContractRepository(driver *neo4j.DriverWithContext) ContractRepository {
	return &contractRepository{
		driver: driver,
	}
}

func (r *contractRepository) GetContractsForStatusRenewal(ctx context.Context, referenceTime time.Time) ([]TenantAndContractId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContractRepository.GetContractsForStatusRenewal")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(span)
	span.LogFields(log.Object("referenceTime", referenceTime))

	cypher := `MATCH (t:Tenant)<-[:CONTRACT_BELONGS_TO_TENANT]-(c:Contract)
				WHERE ((c.status <> $endedStatus AND c.endedAt < $referenceTime) OR
						((c.endedAt is null OR c.endedAt > $referenceTime) AND c.status <> $liveStatus AND c.serviceStartedAt < $referenceTime))
					AND (c.techStatusRenewalRequestedAt IS NULL OR c.techStatusRenewalRequestedAt + duration({hours: 2}) < $referenceTime)
				RETURN t.name, c.id LIMIT 100`
	params := map[string]any{
		"referenceTime": referenceTime,
		"endedStatus":   "ENDED",
		"liveStatus":    "LIVE",
	}
	span.LogFields(log.String("query", cypher), log.Object("params", params))

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)

	})
	if err != nil {
		return nil, err
	}
	output := make([]TenantAndContractId, 0)
	for _, v := range records.([]*neo4j.Record) {
		output = append(output,
			TenantAndContractId{
				Tenant:     v.Values[0].(string),
				ContractId: v.Values[1].(string),
			})
	}
	span.LogFields(log.Int("output - length", len(output)))
	return output, nil
}

func (r *contractRepository) GetContractsForRenewalRollout(ctx context.Context, referenceTime time.Time) ([]TenantAndContractId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContractRepository.GetContractsForStatusRenewal")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(span)
	span.LogFields(log.Object("referenceTime", referenceTime))

	cypher := `MATCH (t:Tenant)<-[:CONTRACT_BELONGS_TO_TENANT]-(c:Contract),
				(c)-[:ACTIVE_RENEWAL]->(op:RenewalOpportunity)
				WHERE c.status = $liveStatus 
					AND op.renewedAt < $referenceTime
					AND (c.techRolloutRenewalRequestedAt IS NULL OR c.techRolloutRenewalRequestedAt + duration({hours: 2}) < $referenceTime)
				RETURN t.name, c.id LIMIT 100`
	params := map[string]any{
		"referenceTime": referenceTime,
		"liveStatus":    "LIVE",
	}
	span.LogFields(log.String("query", cypher), log.Object("params", params))

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)

	})
	if err != nil {
		return nil, err
	}
	output := make([]TenantAndContractId, 0)
	for _, v := range records.([]*neo4j.Record) {
		output = append(output,
			TenantAndContractId{
				Tenant:     v.Values[0].(string),
				ContractId: v.Values[1].(string),
			})
	}
	span.LogFields(log.Int("output - length", len(output)))
	return output, nil
}

func (r *contractRepository) MarkPayNotificationRequested(ctx context.Context, tenant, invoiceId string, requestedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContractRepository.MarkPayNotificationRequested")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(span)
	span.LogFields(log.String("invoiceId", invoiceId), log.Object("requestedAt", requestedAt))

	cypher := `MATCH (:Tenant {name:$tenant})<-[:INVOICE_BELONGS_TO_TENANT]-(i:Invoice {id:$invoiceId})
				SET i.techPayNotificationRequestedAt=$requestedAt`
	params := map[string]any{
		"tenant":      tenant,
		"invoiceId":   invoiceId,
		"requestedAt": requestedAt,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}
