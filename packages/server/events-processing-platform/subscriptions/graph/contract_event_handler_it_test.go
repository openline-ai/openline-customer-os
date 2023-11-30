package graph

import (
	"context"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contract/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contract/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contract/model"
	opportunityaggregate "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/opportunity/aggregate"
	opportunitycmdhandler "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/opportunity/command_handler"
	opportunityevent "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/opportunity/event"
	opportunitymodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/opportunity/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test"
	eventstoret "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/eventstore"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/neo4j"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContractEventHandler_OnCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// Prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{})
	userIdCreator := neo4jt.CreateUser(ctx, testDatabase.Driver, tenantName, entity.UserEntity{})
	neo4jt.CreateExternalSystem(ctx, testDatabase.Driver, tenantName, "sf")
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "User": 1, "ExternalSystem": 1, "Contract": 0})

	// Prepare the event handler
	contractEventHandler := &ContractEventHandler{
		log:                 testLogger,
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}

	// Create a ContractCreateEvent
	contractId := uuid.New().String()
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	timeNow := utils.Now()
	createEvent, err := event.NewContractCreateEvent(
		contractAggregate,
		model.ContractDataFields{
			Name:             "New Contract",
			ContractUrl:      "http://contract.url",
			OrganizationId:   orgId,
			CreatedByUserId:  userIdCreator,
			ServiceStartedAt: &timeNow,
			SignedAt:         &timeNow,
			RenewalCycle:     model.MonthlyRenewal,
			Status:           model.Live,
		},
		commonmodel.Source{
			Source:    constants.SourceOpenline,
			AppSource: constants.AppSourceEventProcessingPlatform,
		},
		commonmodel.ExternalSystem{
			ExternalSystemId: "sf",
			ExternalId:       "ext-id-1",
		},
		timeNow,
		timeNow,
	)
	require.Nil(t, err, "failed to create contract create event")

	// Execute
	err = contractEventHandler.OnCreate(context.Background(), createEvent)
	require.Nil(t, err, "failed to execute contract create event handler")

	// Assert Neo4j Node Counts
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		"Organization": 1,
		"User":         1,
		"Contract":     1, "Contract_" + tenantName: 1,
	})
	neo4jt.AssertRelationship(ctx, t, testDatabase.Driver, contractId, "CREATED_BY", userIdCreator)
	neo4jt.AssertRelationship(ctx, t, testDatabase.Driver, orgId, "HAS_CONTRACT", contractId)
	neo4jt.AssertRelationship(ctx, t, testDatabase.Driver, contractId, "IS_LINKED_WITH", "sf")

	contractDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Contract_"+tenantName, contractId)
	require.Nil(t, err)
	require.NotNil(t, contractDbNode)

	// Verify contract
	contract := graph_db.MapDbNodeToContractEntity(contractDbNode)
	require.Equal(t, contractId, contract.Id)
	require.Equal(t, "New Contract", contract.Name)
	require.Equal(t, "http://contract.url", contract.ContractUrl)
	require.Equal(t, model.Live.String(), contract.Status)
	require.Equal(t, model.MonthlyRenewal.String(), contract.RenewalCycle)
	require.True(t, timeNow.Equal(contract.CreatedAt.UTC()))
	require.True(t, timeNow.Equal(contract.UpdatedAt.UTC()))
	require.True(t, timeNow.Equal(*contract.ServiceStartedAt))
	require.True(t, timeNow.Equal(*contract.SignedAt))
	require.Nil(t, contract.EndedAt)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), contract.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, contract.AppSource)

	// Check create renewal opportunity command was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	var eventList []eventstore.Event
	for _, value := range eventsMap {
		eventList = value
	}
	require.Equal(t, 1, len(eventList))

	generatedEvent1 := eventList[0]
	require.Equal(t, opportunityevent.OpportunityCreateRenewalV1, generatedEvent1.EventType)
	var eventData1 opportunityevent.OpportunityCreateRenewalEvent
	err = generatedEvent1.GetJsonData(&eventData1)
	require.Nil(t, err)
	require.Equal(t, tenantName, eventData1.Tenant)
	require.Equal(t, contractId, eventData1.ContractId)
	require.Equal(t, constants.SourceOpenline, eventData1.Source.Source)
	test.AssertRecentTime(t, eventData1.CreatedAt)
	test.AssertRecentTime(t, eventData1.UpdatedAt)
}

func TestContractEventHandler_OnUpdate_FrequencySet(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	contractId := neo4jt.CreateContract(ctx, testDatabase.Driver, tenantName, entity.ContractEntity{
		Name:        "test contract",
		ContractUrl: "http://contract.url",
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Contract": 1})

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	now := utils.Now()
	yesterday := now.AddDate(0, 0, -1)
	daysAgo2 := now.AddDate(0, 0, -2)
	tomorrow := now.AddDate(0, 0, 1)
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	updateEvent, err := event.NewContractUpdateEvent(contractAggregate,
		model.ContractDataFields{
			Name:             "test contract updated",
			ContractUrl:      "http://contract.url/updated",
			ServiceStartedAt: &yesterday,
			SignedAt:         &daysAgo2,
			EndedAt:          &tomorrow,
			RenewalCycle:     model.MonthlyRenewal,
			Status:           model.Live,
		},
		commonmodel.ExternalSystem{},
		constants.SourceOpenline,
		now)
	require.Nil(t, err, "failed to create event")

	// EXECUTE
	err = contractEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err, "failed to execute event handler")

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Contract": 1, "Contract_" + tenantName: 1})

	contractDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Contract_"+tenantName, contractId)
	require.Nil(t, err)
	require.NotNil(t, contractDbNode)

	// verify contract
	contract := graph_db.MapDbNodeToContractEntity(contractDbNode)
	require.Equal(t, contractId, contract.Id)
	require.Equal(t, "test contract updated", contract.Name)
	require.Equal(t, "http://contract.url/updated", contract.ContractUrl)
	require.Equal(t, model.Live.String(), contract.Status)
	require.Equal(t, model.MonthlyRenewal.String(), contract.RenewalCycle)
	require.True(t, now.Equal(contract.UpdatedAt))
	require.True(t, yesterday.Equal(*contract.ServiceStartedAt))
	require.True(t, daysAgo2.Equal(*contract.SignedAt))
	require.True(t, tomorrow.Equal(*contract.EndedAt))
	require.Equal(t, entity.DataSource(constants.SourceOpenline), contract.SourceOfTruth)

	// Check create renewal opportunity command was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	var eventList []eventstore.Event
	for _, value := range eventsMap {
		eventList = value
	}
	require.Equal(t, 1, len(eventList))

	generatedEvent1 := eventList[0]
	require.Equal(t, opportunityevent.OpportunityCreateRenewalV1, generatedEvent1.EventType)
	var eventData1 opportunityevent.OpportunityCreateRenewalEvent
	err = generatedEvent1.GetJsonData(&eventData1)
	require.Nil(t, err)
	require.Equal(t, tenantName, eventData1.Tenant)
	require.Equal(t, contractId, eventData1.ContractId)
	require.Equal(t, constants.SourceOpenline, eventData1.Source.Source)
	test.AssertRecentTime(t, eventData1.CreatedAt)
	test.AssertRecentTime(t, eventData1.UpdatedAt)
}

func TestContractEventHandler_OnUpdate_FrequencyNotChanged(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	contractId := neo4jt.CreateContract(ctx, testDatabase.Driver, tenantName, entity.ContractEntity{
		RenewalCycle: string(model.MonthlyRenewalCycleString),
	})
	opportunityId := neo4jt.CreateOpportunity(ctx, testDatabase.Driver, tenantName, entity.OpportunityEntity{
		InternalType:  string(opportunitymodel.OpportunityInternalTypeStringRenewal),
		InternalStage: string(opportunitymodel.OpportunityInternalStageStringOpen),
	})
	neo4jt.LinkContractWithOpportunity(ctx, testDatabase.Driver, contractId, opportunityId, true)

	prepareRenewalOpportunity(t, tenantName, opportunityId, aggregateStore)

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	updateEvent, err := event.NewContractUpdateEvent(contractAggregate,
		model.ContractDataFields{
			Name:         "test contract updated",
			RenewalCycle: model.MonthlyRenewal,
		},
		commonmodel.ExternalSystem{},
		constants.SourceOpenline,
		utils.Now())
	require.Nil(t, err)

	// EXECUTE
	err = contractEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err)

	// Check renew ARR event was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	var eventList []eventstore.Event
	for _, value := range eventsMap {
		eventList = value
	}
	require.Equal(t, 2, len(eventList))
	generatedEvent1 := eventList[1]
	require.Equal(t, opportunityevent.OpportunityUpdateV1, generatedEvent1.EventType)
}

func TestContractEventHandler_OnUpdate_FrequencyChanged(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	contractId := neo4jt.CreateContract(ctx, testDatabase.Driver, tenantName, entity.ContractEntity{
		RenewalCycle: string(model.MonthlyRenewalCycleString),
	})
	opportunityId := neo4jt.CreateOpportunity(ctx, testDatabase.Driver, tenantName, entity.OpportunityEntity{
		InternalType:  string(opportunitymodel.OpportunityInternalTypeStringRenewal),
		InternalStage: string(opportunitymodel.OpportunityInternalStageStringOpen),
	})
	neo4jt.LinkContractWithOpportunity(ctx, testDatabase.Driver, contractId, opportunityId, true)

	prepareRenewalOpportunity(t, tenantName, opportunityId, aggregateStore)

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		log:                 testLogger,
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	updateEvent, err := event.NewContractUpdateEvent(contractAggregate,
		model.ContractDataFields{
			Name:         "test contract updated",
			RenewalCycle: model.AnnuallyRenewal,
		},
		commonmodel.ExternalSystem{},
		constants.SourceOpenline,
		utils.Now())
	require.Nil(t, err)

	// EXECUTE
	err = contractEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	var eventList []eventstore.Event
	for _, value := range eventsMap {
		eventList = value
	}
	require.Equal(t, 2, len(eventList))
	generatedEvent1 := eventList[0]
	require.Equal(t, opportunityevent.OpportunityCreateRenewalV1, generatedEvent1.EventType)
	generatedEvent2 := eventList[1]
	require.Equal(t, opportunityevent.OpportunityUpdateV1, generatedEvent2.EventType)
}

func TestContractEventHandler_OnUpdate_ServiceStartDateChanged(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	now := utils.Now()
	yesterday := now.AddDate(0, 0, -1)

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	contractId := neo4jt.CreateContract(ctx, testDatabase.Driver, tenantName, entity.ContractEntity{
		RenewalCycle:     string(model.MonthlyRenewalCycleString),
		ServiceStartedAt: &yesterday,
	})
	opportunityId := neo4jt.CreateOpportunity(ctx, testDatabase.Driver, tenantName, entity.OpportunityEntity{
		InternalType:  string(opportunitymodel.OpportunityInternalTypeStringRenewal),
		InternalStage: string(opportunitymodel.OpportunityInternalStageStringOpen),
	})
	neo4jt.LinkContractWithOpportunity(ctx, testDatabase.Driver, contractId, opportunityId, true)

	prepareRenewalOpportunity(t, tenantName, opportunityId, aggregateStore)

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		log:                 testLogger,
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	updateEvent, err := event.NewContractUpdateEvent(contractAggregate,
		model.ContractDataFields{
			Name:             "test contract updated",
			RenewalCycle:     model.MonthlyRenewal,
			ServiceStartedAt: &now,
		},
		commonmodel.ExternalSystem{},
		constants.SourceOpenline,
		utils.Now())
	require.Nil(t, err)

	// EXECUTE
	err = contractEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	var eventList []eventstore.Event
	for _, value := range eventsMap {
		eventList = value
	}
	require.Equal(t, 3, len(eventList))
	generatedEvent1 := eventList[1]
	require.Equal(t, opportunityevent.OpportunityUpdateNextCycleDateV1, generatedEvent1.EventType)
	generatedEvent2 := eventList[2]
	require.Equal(t, opportunityevent.OpportunityUpdateV1, generatedEvent2.EventType)
}

func TestContractEventHandler_OnUpdate_CurrentSourceOpenline_UpdateSourceNonOpenline(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	now := utils.Now()
	contractId := neo4jt.CreateContract(ctx, testDatabase.Driver, tenantName, entity.ContractEntity{
		Name:             "test contract",
		ContractUrl:      "http://contract.url",
		Status:           "DRAFT",
		RenewalCycle:     "ANNUALLY",
		ServiceStartedAt: &now,
		SignedAt:         &now,
		EndedAt:          &now,
		SourceOfTruth:    constants.SourceOpenline,
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Contract": 1})

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	yesterday := now.AddDate(0, 0, -1)
	daysAgo2 := now.AddDate(0, 0, -2)
	tomorrow := now.AddDate(0, 0, 1)
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	updateEvent, err := event.NewContractUpdateEvent(contractAggregate,
		model.ContractDataFields{
			Name:             "test contract updated",
			ContractUrl:      "http://contract.url/updated",
			ServiceStartedAt: &yesterday,
			SignedAt:         &daysAgo2,
			EndedAt:          &tomorrow,
			RenewalCycle:     model.MonthlyRenewal,
			Status:           model.Live,
		},
		commonmodel.ExternalSystem{},
		"hubspot",
		now)
	require.Nil(t, err, "failed to create event")

	// EXECUTE
	err = contractEventHandler.OnUpdate(context.Background(), updateEvent)
	require.Nil(t, err, "failed to execute event handler")

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Contract": 1, "Contract_" + tenantName: 1})

	contractDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Contract_"+tenantName, contractId)
	require.Nil(t, err)
	require.NotNil(t, contractDbNode)

	// verify contract
	contract := graph_db.MapDbNodeToContractEntity(contractDbNode)
	require.Equal(t, contractId, contract.Id)
	require.Equal(t, "test contract", contract.Name)
	require.Equal(t, "http://contract.url", contract.ContractUrl)
	require.Equal(t, model.Draft.String(), contract.Status)
	require.Equal(t, model.AnnuallyRenewal.String(), contract.RenewalCycle)
	require.True(t, now.Equal(contract.UpdatedAt))
	require.True(t, now.Equal(*contract.ServiceStartedAt))
	require.True(t, now.Equal(*contract.SignedAt))
	require.True(t, now.Equal(*contract.EndedAt))
	require.Equal(t, entity.DataSource(constants.SourceOpenline), contract.SourceOfTruth)
}

func prepareRenewalOpportunity(t *testing.T, tenant, opportunityId string, aggregateStore *eventstoret.TestAggregateStore) {
	// prepare aggregate
	opportunityAggregate := opportunityaggregate.NewOpportunityAggregateWithTenantAndID(tenant, opportunityId)
	createEvent := eventstore.NewBaseEvent(opportunityAggregate, opportunityevent.OpportunityCreateRenewalV1)
	preconfiguredEventData := opportunityevent.OpportunityCreateRenewalEvent{
		Tenant:        tenant,
		InternalType:  string(opportunitymodel.OpportunityInternalTypeStringRenewal),
		InternalStage: string(opportunitymodel.OpportunityInternalStageStringOpen),
	}
	err := createEvent.SetJsonData(&preconfiguredEventData)
	require.Nil(t, err)
	opportunityAggregate.UncommittedEvents = []eventstore.Event{
		createEvent,
	}
	err = aggregateStore.Save(context.Background(), opportunityAggregate)
	require.Nil(t, err)
}

func TestContractEventHandler_OnUpdateStatusEnded(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// Prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{})
	contractId := neo4jt.CreateContractForOrganization(ctx, testDatabase.Driver, tenantName, orgId, entity.ContractEntity{
		Name:   "test contract",
		Status: string(model.ContractStatusStringDraft),
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"Contract": 1, "Contract_" + tenantName: 1, "Action": 0, "TimelineEvent": 0})

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		log:                 testLogger,
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	now := utils.Now()
	event, err := event.NewContractUpdateStatusEvent(contractAggregate, string(model.ContractStatusStringEnded), &now, nil)
	require.Nil(t, err)

	// EXECUTE
	err = contractEventHandler.OnUpdateStatus(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"Contract": 1, "Contract_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})

	contractDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Contract_"+tenantName, contractId)
	require.Nil(t, err)
	require.NotNil(t, contractDbNode)

	// verify contract
	contract := graph_db.MapDbNodeToContractEntity(contractDbNode)
	require.Equal(t, contractId, contract.Id)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, entity.ActionContractStatusUpdated, action.Type)
	require.Equal(t, "test contract has ended", action.Content)
	require.Equal(t, `{"status":"ENDED"}`, action.Metadata)

	// Check request was not generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 0, len(eventsMap))
}

func TestContractEventHandler_OnUpdateStatusLive(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstoret.NewTestAggregateStore()

	// Prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{})
	contractId := neo4jt.CreateContractForOrganization(ctx, testDatabase.Driver, tenantName, orgId, entity.ContractEntity{
		Name:   "test contract",
		Status: string(model.ContractStatusStringDraft),
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"Contract": 1, "Contract_" + tenantName: 1, "Action": 0, "TimelineEvent": 0})

	// prepare event handler
	contractEventHandler := &ContractEventHandler{
		log:                 testLogger,
		repositories:        testDatabase.Repositories,
		opportunityCommands: opportunitycmdhandler.NewCommandHandlers(testLogger, &config.Config{}, aggregateStore),
	}
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenantName, contractId)
	now := utils.Now()
	event, err := event.NewContractUpdateStatusEvent(contractAggregate, string(model.ContractStatusStringLive), &now, nil)
	require.Nil(t, err)

	// EXECUTE
	err = contractEventHandler.OnUpdateStatus(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"Contract": 1, "Contract_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})

	contractDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Contract_"+tenantName, contractId)
	require.Nil(t, err)
	require.NotNil(t, contractDbNode)

	// verify contract
	contract := graph_db.MapDbNodeToContractEntity(contractDbNode)
	require.Equal(t, contractId, contract.Id)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, entity.ActionContractStatusUpdated, action.Type)
	require.Equal(t, "test contract is nowt live", action.Content)
	require.Equal(t, `{"status":"LIVE"}`, action.Metadata)

	// Check request was not generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 0, len(eventsMap))
}
