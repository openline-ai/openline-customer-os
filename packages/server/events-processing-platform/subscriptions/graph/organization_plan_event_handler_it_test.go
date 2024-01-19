package graph

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jmapper "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	event "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization_plan/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization_plan/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestOrganizationPlanEventHandler_OnCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	timeNow := utils.Now()
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{
		Name: "test org",
	})
	mpid := neo4jtest.CreateMasterPlan(ctx, testDatabase.Driver, tenantName, neo4jentity.MasterPlanEntity{
		Source:        constants.SourceOpenline,
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "master plan name",
		SourceOfTruth: constants.SourceOpenline,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
	})

	neo4jtest.CreateMasterPlanMilestone(ctx, testDatabase.Driver, tenantName, mpid, neo4jentity.MasterPlanMilestoneEntity{
		Source:        constants.SourceOpenline,
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "milestone name",
		SourceOfTruth: constants.SourceOpenline,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		Order:         0,
		DurationHours: 24,
		Items:         []string{"item1", "item2"},
		Optional:      false,
	})

	// Prepare the event handler
	orgPlanEventHandler := &OrganizationPlanEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	// Create an MasterPlanCreateEvent
	orgPlanId := uuid.New().String()
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	createEvent, err := event.NewOrganizationPlanCreateEvent(
		orgAggregate,
		orgPlanId,
		mpid,
		"org plan name",
		commonmodel.Source{
			Source:    constants.SourceOpenline,
			AppSource: constants.AppSourceEventProcessingPlatform,
		},
		timeNow,
	)
	require.Nil(t, err)

	// EXECUTE
	err = orgPlanEventHandler.OnCreate(context.Background(), createEvent)
	require.Nil(t, err)

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:                    1,
		neo4jutil.NodeLabelOrganizationPlan + "_" + tenantName: 1})

	orgPlanDbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlan, orgPlanId)
	require.Nil(t, err)
	require.NotNil(t, orgPlanDbNode)

	// verify org plan node
	orgPlan := neo4jmapper.MapDbNodeToOrganizationPlanEntity(orgPlanDbNode)
	require.Equal(t, orgPlanId, orgPlan.Id)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), orgPlan.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, orgPlan.AppSource)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), orgPlan.SourceOfTruth)
	require.Equal(t, timeNow, orgPlan.CreatedAt)
	require.Equal(t, timeNow, orgPlan.UpdatedAt)
	require.Equal(t, "org plan name", orgPlan.Name)
	require.Equal(t, model.NotStarted.String(), orgPlan.StatusDetails.Status)

	createdMilestones := neo4jtest.GetCountOfRelationships(ctx, testDatabase.Driver, "HAS_MILESTONE")
	// should be 2 => 1 master plan milestone + 1 org plan milestone
	require.Equal(t, 2, createdMilestones)
}

func TestOrganizationPlanEventHandler_OnCreateMilestone(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	timeNow := utils.Now()
	mpid := neo4jtest.CreateMasterPlan(ctx, testDatabase.Driver, tenantName, neo4jentity.MasterPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "master plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
	})
	orgId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{Name: "test org"})
	opid := neo4jtest.CreateOrganizationPlan(ctx, testDatabase.Driver, tenantName, mpid, orgId, neo4jentity.OrganizationPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "org plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		StatusDetails: neo4jentity.OrganizationPlanStatusDetails{
			Status:    model.NotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan: 1,
	})

	// Prepare the event handler
	orgPlanEventHandler := &OrganizationPlanEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	// Create an OrgPlanMilestoneCreateEvent
	milestoneId := uuid.New().String()
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	createEvent, err := event.NewOrganizationPlanMilestoneCreateEvent(
		orgAggregate,
		opid,
		milestoneId,
		"milestone name",
		10,
		[]string{"item1", "item2"},
		true,
		commonmodel.Source{
			Source:    constants.SourceOpenline,
			AppSource: constants.AppSourceEventProcessingPlatform,
		},
		timeNow,
		timeNow.Add(time.Hour*24), // due date
	)
	require.Nil(t, err)

	// EXECUTE
	err = orgPlanEventHandler.OnCreateMilestone(context.Background(), createEvent)
	require.Nil(t, err)

	// verify nodes and relationships
	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:                             1,
		neo4jutil.NodeLabelOrganizationPlanMilestone:                    1,
		neo4jutil.NodeLabelOrganizationPlanMilestone + "_" + tenantName: 1})
	neo4jtest.AssertRelationship(ctx, t, testDatabase.Driver, opid, "HAS_MILESTONE", milestoneId)

	// verify org plan milestone node
	orgPlanMilestoneDbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlanMilestone, milestoneId)
	require.Nil(t, err)
	require.NotNil(t, orgPlanMilestoneDbNode)

	milestone := neo4jmapper.MapDbNodeToOrganizationPlanMilestoneEntity(orgPlanMilestoneDbNode)
	require.Equal(t, milestoneId, milestone.Id)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), milestone.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, milestone.AppSource)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), milestone.SourceOfTruth)
	require.Equal(t, timeNow, milestone.CreatedAt)
	test.AssertRecentTime(t, milestone.UpdatedAt)
	require.Equal(t, "milestone name", milestone.Name)
	require.Equal(t, int64(10), milestone.Order)
	require.Equal(t, timeNow.Add(time.Hour*24), milestone.DueDate)
	require.Equal(t, true, milestone.Optional)
	require.Equal(t, model.NotStarted.String(), milestone.StatusDetails.Status)
	for i, item := range milestone.Items {
		require.Equal(t, model.TaskNotDone.String(), item.Status)
		txt := fmt.Sprintf("item%d", i+1)
		require.Equal(t, txt, item.Text)
	}
}

func TestOrganizationPlanEventHandler_OnUpdate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	timeNow := utils.Now()
	mpid := neo4jtest.CreateMasterPlan(ctx, testDatabase.Driver, tenantName, neo4jentity.MasterPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "master plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
	})
	orgId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{Name: "test org"})
	opid := neo4jtest.CreateOrganizationPlan(ctx, testDatabase.Driver, tenantName, mpid, orgId, neo4jentity.OrganizationPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "org plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		StatusDetails: neo4jentity.OrganizationPlanStatusDetails{
			Status:    model.NotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan: 1,
	})
	// Prepare the event handler
	orgPlanEventHandler := &OrganizationPlanEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	// Create an OrgPlanUpdateEvent
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	updateTime := utils.Now()
	updateEvent, err := event.NewOrganizationPlanUpdateEvent(
		orgAggregate,
		opid,
		"org plan updated name",
		true,
		updateTime,
		[]string{event.FieldMaskName, event.FieldMaskRetired, event.FieldMaskStatusDetails},
		model.OrganizationPlanDetails{
			Status:    model.Late.String(),
			Comments:  "comments",
			UpdatedAt: updateTime,
		},
	)
	require.Nil(t, err)

	// EXECUTE
	err = orgPlanEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err)

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:                    1,
		neo4jutil.NodeLabelOrganizationPlan + "_" + tenantName: 1})

	orgPlanDbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlan, opid)
	require.Nil(t, err)
	require.NotNil(t, orgPlanDbNode)

	// verify org plan node
	orgPlan := neo4jmapper.MapDbNodeToOrganizationPlanEntity(orgPlanDbNode)
	require.Equal(t, opid, orgPlan.Id)
	require.Equal(t, updateTime, orgPlan.UpdatedAt)
	require.Equal(t, "org plan updated name", orgPlan.Name)
	require.Equal(t, true, orgPlan.Retired)
	require.Equal(t, model.Late.String(), orgPlan.StatusDetails.Status)
	require.Equal(t, "comments", orgPlan.StatusDetails.Comments)
	require.Equal(t, updateTime, orgPlan.StatusDetails.UpdatedAt)
}

func TestOrganizationPlanEventHandler_OnUpdateMilestone(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	timeNow := utils.Now()
	mpid := neo4jtest.CreateMasterPlan(ctx, testDatabase.Driver, tenantName, neo4jentity.MasterPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "master plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
	})
	orgId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{Name: "test org"})
	opid := neo4jtest.CreateOrganizationPlan(ctx, testDatabase.Driver, tenantName, mpid, orgId, neo4jentity.OrganizationPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "org plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		StatusDetails: neo4jentity.OrganizationPlanStatusDetails{
			Status:    model.NotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})

	milestoneId := neo4jtest.CreateOrganizationPlanMilestone(ctx, testDatabase.Driver, tenantName, opid, neo4jentity.OrganizationPlanMilestoneEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "milestone name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		Order:         0,
		DueDate:       timeNow.Add(time.Hour * 24),
		Items:         []neo4jentity.OrganizationPlanMilestoneItem{{Text: "item1", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}, {Text: "item2", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}},
		Optional:      false,
		StatusDetails: neo4jentity.OrganizationPlanMilestoneStatusDetails{
			Status:    model.MilestoneNotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:                             1,
		neo4jutil.NodeLabelOrganizationPlanMilestone:                    1,
		neo4jutil.NodeLabelOrganizationPlanMilestone + "_" + tenantName: 1,
	})

	// Prepare the event handler
	orgPlanEventHandler := &OrganizationPlanEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	// Create an MasterPlanMilestoneCreateEvent
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	updateTime := utils.Now()
	updateEvent, err := event.NewOrganizationPlanMilestoneUpdateEvent(
		orgAggregate,
		opid,
		milestoneId,
		"new name",
		10,
		[]model.OrganizationPlanMilestoneItem{{Text: "item1", Status: model.TaskDone.String(), UpdatedAt: updateTime}, {Text: "item2Change", Status: model.TaskNotDone.String(), UpdatedAt: updateTime}},
		[]string{event.FieldMaskName, event.FieldMaskOptional, event.FieldMaskItems, event.FieldMaskDueDate, event.FieldMaskOrder, event.FieldMaskStatusDetails},
		true,
		true,
		updateTime,
		timeNow.Add(time.Hour*48), // due date
		model.OrganizationPlanDetails{
			Status:    model.MilestoneStarted.String(),
			Comments:  "comments",
			UpdatedAt: updateTime,
		},
	)
	require.Nil(t, err)

	// EXECUTE
	err = orgPlanEventHandler.OnUpdateMilestone(context.Background(), updateEvent)
	require.Nil(t, err)

	// verify nodes and relationships
	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:          1,
		neo4jutil.NodeLabelOrganizationPlanMilestone: 1})

	// verify master plan milestone node
	orgPlanMilestoneDbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlanMilestone, milestoneId)
	require.Nil(t, err)
	require.NotNil(t, orgPlanMilestoneDbNode)

	milestone := neo4jmapper.MapDbNodeToOrganizationPlanMilestoneEntity(orgPlanMilestoneDbNode)
	require.Equal(t, milestoneId, milestone.Id)
	require.Equal(t, updateTime, milestone.UpdatedAt)
	require.Equal(t, "new name", milestone.Name)
	require.Equal(t, int64(10), milestone.Order)
	require.Equal(t, timeNow.Add(time.Hour*48), milestone.DueDate)
	require.Equal(t, true, milestone.Optional)
	require.Equal(t, false, milestone.Retired) // mask not passed so we ignore this field update
	require.Equal(t, model.MilestoneStarted.String(), milestone.StatusDetails.Status)
	require.Equal(t, "comments", milestone.StatusDetails.Comments)
	require.Equal(t, updateTime, milestone.StatusDetails.UpdatedAt)
	for i, item := range milestone.Items {
		if i == 0 {
			require.Equal(t, model.TaskDone.String(), item.Status)
			require.Equal(t, "item1", item.Text)
		} else {
			require.Equal(t, model.TaskNotDone.String(), item.Status)
			require.Equal(t, "item2Change", item.Text)
		}
	}
}

func TestOrganizationPlanEventHandler_OnReorderMilestones(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	timeNow := utils.Now()
	mpid := neo4jtest.CreateMasterPlan(ctx, testDatabase.Driver, tenantName, neo4jentity.MasterPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "master plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
	})
	orgId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{Name: "test org"})
	opid := neo4jtest.CreateOrganizationPlan(ctx, testDatabase.Driver, tenantName, mpid, orgId, neo4jentity.OrganizationPlanEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "org plan name",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		StatusDetails: neo4jentity.OrganizationPlanStatusDetails{
			Status:    model.NotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})

	milestoneId1 := neo4jtest.CreateOrganizationPlanMilestone(ctx, testDatabase.Driver, tenantName, opid, neo4jentity.OrganizationPlanMilestoneEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "milestone name 1",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		Order:         0,
		DueDate:       timeNow.Add(time.Hour * 24),
		Items:         []neo4jentity.OrganizationPlanMilestoneItem{{Text: "item1", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}, {Text: "item2", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}},
		Optional:      false,
		StatusDetails: neo4jentity.OrganizationPlanMilestoneStatusDetails{
			Status:    model.MilestoneNotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})
	milestoneId2 := neo4jtest.CreateOrganizationPlanMilestone(ctx, testDatabase.Driver, tenantName, opid, neo4jentity.OrganizationPlanMilestoneEntity{
		Source:        neo4jentity.DataSource(constants.SourceOpenline),
		AppSource:     constants.AppSourceEventProcessingPlatform,
		Name:          "milestone name 2",
		SourceOfTruth: neo4jentity.DataSource(constants.SourceOpenline),
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		Retired:       false,
		Order:         1,
		DueDate:       timeNow.Add(time.Hour * 24),
		Items:         []neo4jentity.OrganizationPlanMilestoneItem{{Text: "item1", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}, {Text: "item2", Status: model.TaskNotDone.String(), UpdatedAt: timeNow}},
		Optional:      false,
		StatusDetails: neo4jentity.OrganizationPlanMilestoneStatusDetails{
			Status:    model.MilestoneNotStarted.String(),
			Comments:  "",
			UpdatedAt: timeNow,
		},
	})
	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:          1,
		neo4jutil.NodeLabelOrganizationPlanMilestone: 2,
	})

	// Prepare the event handler
	orgPlanEventHandler := &OrganizationPlanEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	// Create an MasterPlanMilestoneCreateEvent
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	reorderEvent, err := event.NewOrganizationPlanMilestoneReorderEvent(
		orgAggregate,
		opid,
		[]string{milestoneId2, milestoneId1},
		timeNow,
	)
	require.Nil(t, err)

	// EXECUTE
	err = orgPlanEventHandler.OnReorderMilestones(context.Background(), reorderEvent)
	require.Nil(t, err)

	// verify nodes and relationships
	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelOrganizationPlan:          1,
		neo4jutil.NodeLabelOrganizationPlanMilestone: 2})

	// verify master plan milestone nodes
	orgPlanMilestoneDbNode1, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlanMilestone, milestoneId1)
	require.Nil(t, err)
	require.NotNil(t, orgPlanMilestoneDbNode1)
	milestone1 := neo4jmapper.MapDbNodeToOrganizationPlanMilestoneEntity(orgPlanMilestoneDbNode1)
	require.Equal(t, int64(1), milestone1.Order)

	orgPlanMilestoneDbNode2, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelOrganizationPlanMilestone, milestoneId2)
	require.Nil(t, err)
	require.NotNil(t, orgPlanMilestoneDbNode2)
	milestone2 := neo4jmapper.MapDbNodeToOrganizationPlanMilestoneEntity(orgPlanMilestoneDbNode2)
	require.Equal(t, int64(0), milestone2.Order)
}