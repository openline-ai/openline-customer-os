package resolver

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/grpc/events_platform"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestQueryResolver_Invoice(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	timeNow := utils.Now()
	neo4jtest.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, driver, tenantName, organizationId, neo4jentity.ContractEntity{})
	invoiceId := neo4jtest.CreateInvoice(ctx, driver, tenantName, contractId, neo4jentity.InvoiceEntity{
		CreatedAt:        timeNow,
		UpdatedAt:        timeNow,
		DryRun:           false,
		Number:           "1",
		Currency:         "RON",
		Date:             timeNow,
		DueDate:          timeNow,
		Amount:           100,
		Vat:              19,
		Total:            119,
		RepositoryFileId: "ABC",
	})

	neo4jtest.CreateInvoiceLine(ctx, driver, tenantName, invoiceId, neo4jentity.InvoiceLineEntity{
		CreatedAt: timeNow,
		Name:      "SLI 1",
		Price:     100,
		Quantity:  1,
		Amount:    100,
		Vat:       19,
		Total:     119,
	})

	rawResponse := callGraphQL(t, "invoice/get_invoice", map[string]interface{}{"id": invoiceId})
	require.Nil(t, rawResponse.Errors)

	var invoiceStruct struct {
		Invoice model.Invoice
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &invoiceStruct)
	require.Nil(t, err)

	invoice := invoiceStruct.Invoice
	require.Equal(t, invoiceId, invoice.ID)
	require.Equal(t, timeNow, invoice.CreatedAt)
	require.Equal(t, timeNow, invoice.UpdatedAt)
	require.Equal(t, false, invoice.DryRun)
	require.Equal(t, "1", invoice.Number)
	require.Equal(t, "RON", invoice.Currency)
	require.Equal(t, timeNow, invoice.Date)
	require.Equal(t, timeNow, invoice.DueDate)
	require.Equal(t, 100.0, invoice.Amount)
	require.Equal(t, 19.0, invoice.Vat)
	require.Equal(t, 119.0, invoice.Total)
	require.Equal(t, "ABC", invoice.RepositoryFileID)

	require.Equal(t, 1, len(invoice.InvoiceLines))
	require.Equal(t, "SLI 1", invoice.InvoiceLines[0].Name)
	require.Equal(t, 100.0, invoice.InvoiceLines[0].Price)
	require.Equal(t, 1, invoice.InvoiceLines[0].Quantity)
	require.Equal(t, 100.0, invoice.InvoiceLines[0].Amount)
	require.Equal(t, 19.0, invoice.InvoiceLines[0].Vat)
	require.Equal(t, 119.0, invoice.InvoiceLines[0].Total)
}

func TestQueryResolver_Invoices(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	timeNow := utils.Now()
	yesterday := timeNow.Add(-24 * time.Hour)
	neo4jtest.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, driver, tenantName, organizationId, neo4jentity.ContractEntity{})

	invoice1Id := neo4jtest.CreateInvoice(ctx, driver, tenantName, contractId, neo4jentity.InvoiceEntity{
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Number:    "1",
	})
	neo4jtest.CreateInvoiceLine(ctx, driver, tenantName, invoice1Id, neo4jentity.InvoiceLineEntity{
		Name: "SLI 1",
	})

	invoice2Id := neo4jtest.CreateInvoice(ctx, driver, tenantName, contractId, neo4jentity.InvoiceEntity{
		CreatedAt: yesterday,
		UpdatedAt: yesterday,
		Number:    "2",
	})
	neo4jtest.CreateInvoiceLine(ctx, driver, tenantName, invoice2Id, neo4jentity.InvoiceLineEntity{
		Name: "SLI 2",
	})

	invoice3Id := neo4jtest.CreateInvoice(ctx, driver, tenantName, contractId, neo4jentity.InvoiceEntity{
		CreatedAt: yesterday,
		UpdatedAt: yesterday,
		Number:    "11",
	})
	neo4jtest.CreateInvoiceLine(ctx, driver, tenantName, invoice3Id, neo4jentity.InvoiceLineEntity{
		Name: "SLI 3",
	})

	rawResponse := callGraphQL(t, "invoice/get_invoices", map[string]interface{}{
		"page":  0,
		"limit": 10,
	})
	require.Nil(t, rawResponse.Errors)

	var invoiceStruct struct {
		Invoices model.InvoicesPage
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &invoiceStruct)
	require.Nil(t, err)

	require.Equal(t, int64(2), invoiceStruct.Invoices.TotalElements)
	require.Equal(t, 2, len(invoiceStruct.Invoices.Content))

	require.Equal(t, invoice1Id, invoiceStruct.Invoices.Content[0].ID)
	require.Equal(t, "1", invoiceStruct.Invoices.Content[0].Number)
	require.Equal(t, "SLI 1", invoiceStruct.Invoices.Content[0].InvoiceLines[0].Name)
	require.Equal(t, "11", invoiceStruct.Invoices.Content[1].Number)
	require.Equal(t, "SLI 3", invoiceStruct.Invoices.Content[1].InvoiceLines[0].Name)
}

func TestQueryResolver_SimulateInvoice(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	timeNow := utils.Now()
	neo4jtest.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContract(ctx, driver, tenantName, organizationId, neo4jentity.ContractEntity{})
	invoiceId := neo4jtest.CreateInvoice(ctx, driver, tenantName, contractId, neo4jentity.InvoiceEntity{})

	calledSimulateInvoice := false
	invoiceServiceCallbacks := events_platform.MockInvoiceServiceCallbacks{
		SimulateInvoice: func(context context.Context, request *invoicepb.SimulateInvoiceRequest) (*invoicepb.InvoiceIdResponse, error) {
			require.Equal(t, tenantName, request.Tenant)
			require.Equal(t, testUserId, request.LoggedInUserId)
			require.Equal(t, contractId, request.ContractId)
			require.Equal(t, 2, len(request.DryRunServiceLineItems))

			require.Equal(t, "1", request.DryRunServiceLineItems[0].ServiceLineItemId)
			require.Equal(t, "SLI 1", request.DryRunServiceLineItems[0].Name)
			require.Equal(t, commonpb.BilledType_MONTHLY_BILLED, request.DryRunServiceLineItems[0].Billed)
			require.Equal(t, 100.0, request.DryRunServiceLineItems[0].Price)
			require.Equal(t, int64(1), request.DryRunServiceLineItems[0].Quantity)

			require.Equal(t, "", request.DryRunServiceLineItems[1].ServiceLineItemId)
			require.Equal(t, "New SLI", request.DryRunServiceLineItems[1].Name)
			require.Equal(t, commonpb.BilledType_NONE_BILLED, request.DryRunServiceLineItems[1].Billed)
			require.Equal(t, 10.0, request.DryRunServiceLineItems[1].Price)
			require.Equal(t, int64(5), request.DryRunServiceLineItems[1].Quantity)

			require.Equal(t, constants.AppSourceCustomerOsApi, request.SourceFields.AppSource)
			calledSimulateInvoice = true
			return &invoicepb.InvoiceIdResponse{
				Id: invoiceId,
			}, nil
		},
	}
	events_platform.SetInvoiceCallbacks(&invoiceServiceCallbacks)

	rawResponse := callGraphQL(t, "invoice/simulate_invoice", map[string]interface{}{
		"invoice": map[string]interface{}{
			"contractId": contractId,
			"date":       timeNow,
			"invoiceLines": []map[string]interface{}{
				{
					"serviceLineItemId": "1",
					"name":              "SLI 1",
					"billed":            "MONTHLY",
					"price":             100,
					"quantity":          1,
				},
				{
					"serviceLineItemId": "",
					"billed":            "NONE",
					"name":              "New SLI",
					"price":             10,
					"quantity":          5,
				},
			},
		},
	})
	require.Nil(t, rawResponse.Errors)

	require.True(t, calledSimulateInvoice)

	var invoiceStruct struct {
		Invoice_Simulate string
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &invoiceStruct)
	require.Nil(t, err)

	require.Equal(t, invoiceId, invoiceStruct.Invoice_Simulate)
}
