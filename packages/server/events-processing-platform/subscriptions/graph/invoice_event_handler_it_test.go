package graph

import (
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	neo4jmapper "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/invoice"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"testing"
)

func TestInvoiceEventHandler_OnInvoiceCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, testDatabase.Driver, tenantName, organizationId, neo4jentity.ContractEntity{})

	eventHandler := &InvoiceEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	now := utils.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	invoiceId := uuid.New().String()

	aggregate := invoice.NewInvoiceAggregateWithTenantAndID(tenantName, invoiceId)
	newEvent, err := invoice.NewInvoiceCreateEvent(
		aggregate,
		commonmodel.Source{
			Source:    constants.SourceOpenline,
			AppSource: constants.AppSourceEventProcessingPlatform,
		},
		contractId,
		"EUR",
		"INV-123",
		true,
		now,
		yesterday,
		tomorrow,
	)
	require.Nil(t, err)

	// EXECUTE
	err = eventHandler.OnInvoiceCreateV1(context.Background(), newEvent)
	require.Nil(t, err)

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelInvoice:                    1,
		neo4jutil.NodeLabelInvoice + "_" + tenantName: 1})

	dbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelInvoice, invoiceId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	// verify
	createdInvoice := neo4jmapper.MapDbNodeToInvoiceEntity(dbNode)

	require.Equal(t, invoiceId, createdInvoice.Id)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), createdInvoice.Source)
	require.Equal(t, neo4jentity.DataSource(constants.SourceOpenline), createdInvoice.SourceOfTruth)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, createdInvoice.AppSource)
	require.Equal(t, now, createdInvoice.CreatedAt)
	require.Equal(t, now, createdInvoice.UpdatedAt)
	require.Equal(t, true, createdInvoice.DryRun)
	require.Equal(t, "INV-123", createdInvoice.Number)
	require.Equal(t, yesterday, createdInvoice.PeriodStartDate)
	require.Equal(t, tomorrow, createdInvoice.PeriodEndDate)
	require.Equal(t, float64(0), createdInvoice.Amount)
	require.Equal(t, float64(0), createdInvoice.Vat)
	require.Equal(t, float64(0), createdInvoice.Amount)
	require.Equal(t, neo4jenum.CurrencyEUR, createdInvoice.Currency)
	require.Equal(t, "", createdInvoice.RepositoryFileId)
}

func TestInvoiceEventHandler_OnInvoiceFill(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, testDatabase.Driver, tenantName, organizationId, neo4jentity.ContractEntity{})
	id := neo4jtest.CreateInvoiceForContract(ctx, testDatabase.Driver, tenantName, contractId, neo4jentity.InvoiceEntity{})

	// Prepare the event handler
	eventHandler := &InvoiceEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	timeNow := utils.Now()

	aggregate := invoice.NewInvoiceAggregateWithTenantAndID(tenantName, id)
	updateEvent, err := invoice.NewInvoiceFillEvent(
		aggregate,
		&timeNow,
		commonmodel.Source{
			Source:    constants.SourceOpenline,
			AppSource: constants.AppSourceEventProcessingPlatform,
		},
		&invoicepb.FillInvoiceRequest{
			Amount: 100,
			Vat:    20,
			Total:  120,
			Lines: []*invoicepb.InvoiceLine{
				{
					Name:     "test",
					Price:    50,
					Quantity: 2,
					Amount:   100,
					Vat:      20,
					Total:    120,
				},
			},
		},
	)
	require.Nil(t, err)

	// EXECUTE
	err = eventHandler.OnInvoiceFill(context.Background(), updateEvent)
	require.Nil(t, err)

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelInvoice:                        1,
		neo4jutil.NodeLabelInvoice + "_" + tenantName:     1,
		neo4jutil.NodeLabelInvoiceLine:                    1,
		neo4jutil.NodeLabelInvoiceLine + "_" + tenantName: 1,
	})

	dbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelInvoice, id)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	// verify
	invoice := neo4jmapper.MapDbNodeToInvoiceEntity(dbNode)
	require.Equal(t, id, invoice.Id)
	require.Equal(t, timeNow, invoice.UpdatedAt)
	require.Equal(t, float64(100), invoice.Amount)
	require.Equal(t, float64(20), invoice.Vat)
	require.Equal(t, float64(120), invoice.Total)
	require.Equal(t, "", invoice.RepositoryFileId)

	// verify invoice lines
	dbNode, err = neo4jtest.GetFirstNodeByLabel(ctx, testDatabase.Driver, neo4jutil.NodeLabelInvoiceLine)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	invoiceLine := neo4jmapper.MapDbNodeToInvoiceLineEntity(dbNode)
	require.Equal(t, "test", invoiceLine.Name)
	require.Equal(t, float64(50), invoiceLine.Price)
	require.Equal(t, int64(2), invoiceLine.Quantity)
	require.Equal(t, float64(100), invoiceLine.Amount)
	require.Equal(t, float64(20), invoiceLine.Vat)
	require.Equal(t, float64(120), invoiceLine.Total)
}

func TestInvoiceEventHandler_OnInvoicePdfGenerated(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// prepare neo4j data
	neo4jtest.CreateTenant(ctx, testDatabase.Driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, testDatabase.Driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, testDatabase.Driver, tenantName, organizationId, neo4jentity.ContractEntity{})
	id := neo4jtest.CreateInvoiceForContract(ctx, testDatabase.Driver, tenantName, contractId, neo4jentity.InvoiceEntity{})

	// Prepare the event handler
	eventHandler := &InvoiceEventHandler{
		log:          testLogger,
		repositories: testDatabase.Repositories,
	}

	timeNow := utils.Now()

	aggregate := invoice.NewInvoiceAggregateWithTenantAndID(tenantName, id)
	pdfGeneratedEvent, err := invoice.NewInvoicePdfGeneratedEvent(
		aggregate,
		timeNow,
		"test",
	)
	require.Nil(t, err)

	// EXECUTE
	err = eventHandler.OnInvoicePdfGenerated(context.Background(), pdfGeneratedEvent)
	require.Nil(t, err)

	neo4jtest.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		neo4jutil.NodeLabelInvoice:                    1,
		neo4jutil.NodeLabelInvoice + "_" + tenantName: 1,
	})

	dbNode, err := neo4jtest.GetNodeById(ctx, testDatabase.Driver, neo4jutil.NodeLabelInvoice, id)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	// verify
	invoice := neo4jmapper.MapDbNodeToInvoiceEntity(dbNode)
	require.Equal(t, id, invoice.Id)
	require.Equal(t, timeNow, invoice.UpdatedAt)

	require.Equal(t, "test", invoice.RepositoryFileId)
}
