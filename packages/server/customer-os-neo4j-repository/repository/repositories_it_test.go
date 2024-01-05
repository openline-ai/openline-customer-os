package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"testing"
)

var (
	neo4jContainer testcontainers.Container
	driver         *neo4j.DriverWithContext
	repositories   *Repositories
)

const tenantName = "openline"

func TestMain(m *testing.M) {
	neo4jContainer, driver = neo4jtest.InitTestNeo4jDB()
	defer func(dbContainer testcontainers.Container, driver neo4j.DriverWithContext, ctx context.Context) {
		neo4jtest.CloseDriver(driver)
		neo4jtest.Terminate(dbContainer, ctx)
	}(neo4jContainer, *driver, context.Background())

	repositories = InitNeo4jRepositories(driver, "neo4j")

	os.Exit(m.Run())
}

func tearDownTestCase(ctx context.Context) func(tb testing.TB) {
	return func(tb testing.TB) {
		tb.Logf("Teardown test %v, cleaning neo4j DB", tb.Name())
		neo4jtest.CleanupAllData(ctx, driver)
	}
}
