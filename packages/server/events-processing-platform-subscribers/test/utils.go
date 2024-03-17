package test

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	comlog "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/test/mocked_grpc"
	postgrest "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/test/postgres"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
	"testing"
	"time"
)

type TestDatabase struct {
	Neo4jContainer testcontainers.Container
	Driver         *neo4j.DriverWithContext
	Repositories   *repository.Repositories
	GormDB         *gorm.DB
}

func SetupTestLogger() logger.Logger {
	testLogger := logger.NewExtendedAppLogger(&comlog.Config{
		DevMode: true,
	})
	testLogger.InitLogger()
	return testLogger
}

func SetupTestDatabase() (TestDatabase, func()) {
	testDBs := TestDatabase{}

	testDBs.Neo4jContainer, testDBs.Driver = neo4jtest.InitTestNeo4jDB()

	postgresContainer, postgresGormDB, _ := postgrest.InitTestDB()
	testDBs.GormDB = postgresGormDB
	testDBs.Repositories = repository.InitRepos(testDBs.Driver, "neo4j", postgresGormDB)

	shutdown := func() {
		neo4jtest.CloseDriver(*testDBs.Driver)
		neo4jtest.Terminate(testDBs.Neo4jContainer, context.Background())
		postgrest.Terminate(postgresContainer, context.Background())
	}
	return testDBs, shutdown
}

func SetupMockedTestGrpcClient() *grpc_client.Clients {
	testDialFactory := mocked_grpc.NewMockedTestDialFactory()
	grpcConn, _ := testDialFactory.GetEventsProcessingPlatformConn()
	return grpc_client.InitGrpcClients(grpcConn)
}

func AssertRecentTime(t *testing.T, checkTime time.Time) {
	x := 2 // Set the time difference to 2 seconds

	diff := time.Since(checkTime)

	require.True(t, diff <= time.Duration(x)*time.Second, "The time is within the last %d seconds.", x)
}
