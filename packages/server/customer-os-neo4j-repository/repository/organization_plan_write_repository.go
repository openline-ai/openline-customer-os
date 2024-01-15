package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type OrganizationPlanUpdateFields struct {
	Name          string
	Retired       bool
	UpdatedAt     time.Time
	UpdateName    bool
	UpdateRetired bool
	StatusDetails entity.OrganizationPlanStatusDetails
}

type OrganizationPlanMilestoneUpdateFields struct {
	UpdatedAt           time.Time
	Name                string
	Order               int64
	DurationHours       int64
	Items               []entity.OrganizationPlanMilestoneItem
	StatusDetails       entity.OrganizationPlanMilestoneStatusDetails
	Optional            bool
	Retired             bool
	UpdateName          bool
	UpdateOrder         bool
	UpdateItems         bool
	UpdateOptional      bool
	UpdateRetired       bool
	UpdateDurationHours bool
	UpdateStatusDetails bool
}

type OrganizationPlanWriteRepository interface {
	Create(ctx context.Context, tenant, organizationPlanId, name, source, appSource string, createdAt time.Time, statusDetails entity.OrganizationPlanStatusDetails) error
	Update(ctx context.Context, tenant, organizationPlanId string, data OrganizationPlanUpdateFields) error
	CreateMilestone(ctx context.Context, tenant, organizationPlanId, milestoneId, name, source, appSource string, order, durationHours int64, items []entity.OrganizationPlanMilestoneItem, optional bool, createdAt time.Time, statusDetails entity.OrganizationPlanMilestoneStatusDetails) error
	CreateBulkMilestones(ctx context.Context, tenant, organizationPlanId, source, appSource string, milestones []entity.OrganizationPlanMilestoneEntity, createdAt time.Time) error
	UpdateMilestone(ctx context.Context, tenant, organizationPlanId, milestoneId string, data OrganizationPlanMilestoneUpdateFields) error
	LinkWithOrganization(ctx context.Context, tenant, organizationPlanId, organizationId string, createdAt time.Time) error
	LinkWithMasterPlan(ctx context.Context, tenant, organizationPlanId, masterPlanId string, createdAt time.Time) error
}

type organizationPlanWriteRepository struct {
	driver   *neo4j.DriverWithContext
	database string
}

func NewOrganizationPlanWriteRepository(driver *neo4j.DriverWithContext, database string) OrganizationPlanWriteRepository {
	return &organizationPlanWriteRepository{
		driver:   driver,
		database: database,
	}
}

func (r *organizationPlanWriteRepository) Create(ctx context.Context, tenant, organizationPlanId, name, source, appSource string, createdAt time.Time, statusDetails entity.OrganizationPlanStatusDetails) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.Create")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationPlanId)

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})
							MERGE (t)<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							ON CREATE SET 
								op:OrganizationPlan_%s,
								op.createdAt=$createdAt,
								op.updatedAt=$updatedAt,
								op.source=$source,
								op.sourceOfTruth=$sourceOfTruth,
								op.appSource=$appSource,
								op.name=$name
								op.statusDetails=$statusDetails
							`, tenant)
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"createdAt":          createdAt,
		"updatedAt":          createdAt,
		"source":             source,
		"sourceOfTruth":      source,
		"appSource":          appSource,
		"name":               name,
		"statusDetails":      statusDetails,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) CreateMilestone(ctx context.Context, tenant, organizationPlanId, milestoneId, name, source, appSource string, order, durationHours int64, items []entity.OrganizationPlanMilestoneItem, optional bool, createdAt time.Time, statusDetails entity.OrganizationPlanMilestoneStatusDetails) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.CreateMilestone")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, milestoneId)

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							MERGE (op)-[:HAS_MILESTONE]->(m:OrganizationPlanMilestone {id:$milestoneId})
							ON CREATE SET 
								m:OrganizationPlanMilestone_%s,
								m.createdAt=$createdAt,
								m.updatedAt=$updatedAt,
								m.source=$source,
								m.sourceOfTruth=$sourceOfTruth,
								m.appSource=$appSource,
								m.name=$name,
								m.order=$order,
								m.durationHours=$durationHours,
								m.optional=$optional,
								m.items=$items
								m.statusDetails=$statusDetails
							`, tenant)
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"milestoneId":        milestoneId,
		"createdAt":          createdAt,
		"updatedAt":          createdAt,
		"source":             source,
		"sourceOfTruth":      source,
		"appSource":          appSource,
		"name":               name,
		"order":              order,
		"durationHours":      durationHours,
		"optional":           optional,
		"items":              items,
		"statusDetails":      statusDetails,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) CreateBulkMilestones(ctx context.Context, tenant, organizationPlanId, source, appSource string, milestones []entity.OrganizationPlanMilestoneEntity, createdAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.CreateBulkMilestones")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	tracing.LogObjectAsJson(span, "milestones", milestones)

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							UNWIND $milestones as milestone
							MERGE (op)-[:HAS_MILESTONE]->(m:OrganizationPlanMilestone {id:milestone.id})
							ON CREATE SET 
								m:OrganizationPlanMilestone_%s,
								m.createdAt=$createdAt,
								m.updatedAt=$updatedAt,
								m.source=$source,
								m.sourceOfTruth=$sourceOfTruth,
								m.appSource=$appSource,
								m.name=milestone.name,
								m.order=milestone.order,
								m.durationHours=milestone.durationHours,
								m.optional=milestone.optional,
								m.items=milestone.items,
								m.statusDetails=milestone.statusDetails
							`, tenant)
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"milestones":         milestones,
		"createdAt":          createdAt,
		"updatedAt":          createdAt,
		"source":             source,
		"sourceOfTruth":      source,
		"appSource":          appSource,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) Update(ctx context.Context, tenant, organizationPlanId string, data OrganizationPlanUpdateFields) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.Update")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationPlanId)

	cypher := `MATCH (:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							SET op.updatedAt=$updatedAt`
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"updatedAt":          data.UpdatedAt,
	}
	if data.UpdateName {
		cypher += ", op.name=$name"
		params["name"] = data.Name
	}
	if data.UpdateRetired {
		cypher += ", op.retired=$retired"
		params["retired"] = data.Retired
	}

	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) UpdateMilestone(ctx context.Context, tenant, organizationPlanId, milestoneId string, data OrganizationPlanMilestoneUpdateFields) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.UpdateMilestone")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, milestoneId)
	tracing.LogObjectAsJson(span, "data", data)

	cypher := `MATCH (:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId})-[:HAS_MILESTONE]->(m:OrganizationPlanMilestone {id:$milestoneId}) 
							SET m.updatedAt=$updatedAt`
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"milestoneId":        milestoneId,
		"updatedAt":          data.UpdatedAt,
	}
	if data.UpdateName {
		cypher += ", m.name=$name"
		params["name"] = data.Name
	}
	if data.UpdateOrder {
		cypher += ", m.order=$order"
		params["order"] = data.Order
	}
	if data.UpdateDurationHours {
		cypher += ", m.durationHours=$durationHours"
		params["durationHours"] = data.DurationHours
	}
	if data.UpdateItems {
		cypher += ", m.items=$items"
		params["items"] = data.Items
	}
	if data.UpdateOptional {
		cypher += ", m.optional=$optional"
		params["optional"] = data.Optional
	}
	if data.UpdateRetired {
		cypher += ", m.retired=$retired"
		params["retired"] = data.Retired
	}

	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) LinkWithOrganization(ctx context.Context, tenant, organizationPlanId, organizationId string, createdAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.LinkWithOrganization")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationPlanId)

	cypher := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							MATCH (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$organizationId}) 
							MERGE (op)-[:ORGANIZATION_PLAN_BELONGS_TO_ORGANIZATION]->(o)`
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"organizationId":     organizationId,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}

func (r *organizationPlanWriteRepository) LinkWithMasterPlan(ctx context.Context, tenant, organizationPlanId, masterPlanId string, createdAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationPlanWriteRepository.LinkWithMasterPlan")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(span, tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationPlanId)

	cypher := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_PLAN_BELONGS_TO_TENANT]-(op:OrganizationPlan {id:$organizationPlanId}) 
							MATCH (t)<-[:MASTER_PLAN_BELONGS_TO_TENANT]-(m:MasterPlan {id:$masterPlanId}) 
							MERGE (op)-[:ORGANIZATION_PLAN_BELONGS_TO_MASTER_PLAN]->(m)`
	params := map[string]any{
		"tenant":             tenant,
		"organizationPlanId": organizationPlanId,
		"masterPlanId":       masterPlanId,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	err := utils.ExecuteWriteQuery(ctx, *r.driver, cypher, params)
	if err != nil {
		tracing.TraceErr(span, err)
	}
	return err
}
