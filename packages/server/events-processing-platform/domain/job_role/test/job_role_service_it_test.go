package test

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/grpc"
	"testing"
)

var testDatabase *test.TestDatabase
var dialFactory *grpc.TestDialFactoryImpl

func TestMain(m *testing.M) {
	//myDatabase, shutdown := test.SetupTestDatabase()
	//testDatabase = &myDatabase
	//
	//dialFactory = &grpc.TestDialFactoryImpl{}
	//defer shutdown()
	//
	//os.Exit(m.Run())
}

func tearDownTestCase(ctx context.Context, database *test.TestDatabase) func(tb testing.TB) {
	return func(tb testing.TB) {
		//tb.Logf("Teardown test %v, cleaning neo4j DB", tb.Name())
		//neo4jt.CleanupAllData(ctx, database.Driver)
	}
}

//func TestJobRoleService_CreateJobRole(t *testing.T) {
//	ctx := context.TODO()
//	defer tearDownTestCase(ctx, testDatabase)(t)
//
//	aggregateStore := eventstore.NewTestAggregateStore()
//	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
//	if err != nil {
//		t.Fatalf("Failed to connect to events processing platform: %v", err)
//	}
//	jobRoleClient := job_role_grpc_service.NewJobRoleGrpcServiceClient(grpcConnection)
//	timeNow := time.Now().UTC()
//	timeStarted := time.Now().UTC().AddDate(0, -6, 0)
//	timeEnded := time.Now().UTC().AddDate(0, 6, 0)
//	description := "I clean things"
//	response, err := jobRoleClient.CreateJobRole(ctx, &job_role_grpc_service.CreateJobRoleGrpcRequest{
//		Tenant:        "ziggy",
//		JobTitle:      "Chief Janitor",
//		Description:   &description,
//		Source:        "N/A",
//		SourceOfTruth: "N/A",
//		AppSource:     "unit-test",
//		CreatedAt:     timestamppb.New(timeNow),
//		StartedAt:     timestamppb.New(timeStarted),
//		EndedAt:       timestamppb.New(timeEnded),
//	})
//	if err != nil {
//		t.Fatalf("Failed to create job role: %v", err)
//	}
//	require.Nil(t, err)
//	require.NotNil(t, response)
//	eventsMap := aggregateStore.GetEventMap()
//	require.Equal(t, 1, len(eventsMap))
//	eventList := eventsMap[aggregate.NewJobRoleAggregateWithTenantAndID("ziggy", response.Id).ID]
//	require.Equal(t, 1, len(eventList))
//	require.Equal(t, events.JobRoleCreateV1, eventList[0].GetEventType())
//	var eventData events.JobRoleCreateEvent
//	if err := eventList[0].GetJsonData(&eventData); err != nil {
//		t.Errorf("Failed to unmarshal event data: %v", err)
//	}
//	require.Equal(t, "Chief Janitor", eventData.JobTitle)
//	require.Equal(t, "I clean things", *eventData.Description)
//	require.Equal(t, "N/A", eventData.Source)
//	require.Equal(t, "N/A", eventData.SourceOfTruth)
//	require.Equal(t, "unit-test", eventData.AppSource)
//	require.Equal(t, timeNow, eventData.CreatedAt)
//	require.Equal(t, timeNow, eventData.UpdatedAt)
//	require.Equal(t, timeStarted, *eventData.StartedAt)
//	require.Equal(t, timeEnded, *eventData.EndedAt)
//	require.Equal(t, "ziggy", eventData.Tenant)
//
//}
