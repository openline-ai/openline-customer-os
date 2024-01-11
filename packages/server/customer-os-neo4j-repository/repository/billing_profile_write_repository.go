package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type BillingProfileCreateFields struct {
	OrganizationId string       `json:"organizationId"`
	Name           string       `json:"name"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	SourceFields   model.Source `json:"sourceFields"`
}

type BillingProfileWriteRepository interface {
	Create(ctx context.Context, tenant, billingProfileId string, data BillingProfileCreateFields) error
}

type billingProfileWriteRepository struct {
	driver   *neo4j.DriverWithContext
	database string
}

func NewBillingProfileWriteRepository(driver *neo4j.DriverWithContext, database string) BillingProfileWriteRepository {
	return &billingProfileWriteRepository{
		driver:   driver,
		database: database,
	}
}

func (r *billingProfileWriteRepository) Create(ctx context.Context, tenant, billingProfileId string, data BillingProfileCreateFields) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BillingProfileWriteRepository.Create")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, billingProfileId)
	tracing.LogObjectAsJson(span, "data", data)

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:$orgId})
							MERGE (bp:BillingProfile {id:$billingProfileId})<-[:HAS_BILLING_PROFILE]-(org)
							ON CREATE SET 
								bp:BillingProfile_%s,
								bp.createdAt=$createdAt,
								bp.updatedAt=$updatedAt,
								bp.source=$source,
								bp.sourceOfTruth=$sourceOfTruth,
								bp.appSource=$appSource,
								bp.name=$name`, tenant)
	params := map[string]any{
		"tenant":           tenant,
		"billingProfileId": billingProfileId,
		"orgId":            data.OrganizationId,
		"createdAt":        data.CreatedAt,
		"updatedAt":        data.UpdatedAt,
		"source":           data.SourceFields.Source,
		"sourceOfTruth":    data.SourceFields.Source,
		"appSource":        data.SourceFields.AppSource,
		"name":             data.Name,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}
