package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/grpc/events_platform"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	contractpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/contract"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestMutationResolver_ContractCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := uuid.New().String()
	calledCreateContract := false

	contractServiceCallbacks := events_platform.MockContractServiceCallbacks{
		CreateContract: func(context context.Context, contract *contractpb.CreateContractGrpcRequest) (*contractpb.ContractIdGrpcResponse, error) {
			require.Equal(t, tenantName, contract.Tenant)
			require.Equal(t, orgId, contract.OrganizationId)
			require.Equal(t, testUserId, contract.LoggedInUserId)
			require.Equal(t, string(neo4jentity.DataSourceOpenline), contract.SourceFields.Source)
			require.Equal(t, constants.AppSourceCustomerOsApi, contract.SourceFields.AppSource)
			require.Equal(t, "Contract 1", contract.Name)
			require.Equal(t, "https://contract.com", contract.ContractUrl)
			require.Equal(t, contractpb.RenewalCycle_MONTHLY_RENEWAL, contract.RenewalCycle)
			require.Equal(t, "USD", contract.Currency)
			require.Equal(t, int64(7), *contract.RenewalPeriods)
			expectedServiceStartedAt, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
			if err != nil {
				t.Fatalf("Failed to parse expected timestamp: %v", err)
			}
			require.Equal(t, timestamppb.New(expectedServiceStartedAt), contract.ServiceStartedAt)
			expectedSignedAt, err := time.Parse(time.RFC3339, "2019-02-01T00:00:00Z")
			if err != nil {
				t.Fatalf("Failed to parse expected timestamp: %v", err)
			}
			require.Equal(t, timestamppb.New(expectedSignedAt), contract.SignedAt)

			calledCreateContract = true
			neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{
				Id: contractId,
			})
			return &contractpb.ContractIdGrpcResponse{
				Id: contractId,
			}, nil
		},
	}
	events_platform.SetContractCallbacks(&contractServiceCallbacks)

	rawResponse := callGraphQL(t, "contract/create_contract", map[string]interface{}{
		"orgId": orgId,
	})

	var contractStruct struct {
		Contract_Create model.Contract
	}

	require.Nil(t, rawResponse.Errors)
	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)
	contract := contractStruct.Contract_Create
	require.Equal(t, contractId, contract.ID)

	require.True(t, calledCreateContract)
}

func TestMutationResolver_ContractCreate_DefaultValues(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := uuid.New().String()
	calledCreateContract := false

	contractServiceCallbacks := events_platform.MockContractServiceCallbacks{
		CreateContract: func(context context.Context, contract *contractpb.CreateContractGrpcRequest) (*contractpb.ContractIdGrpcResponse, error) {
			require.Equal(t, tenantName, contract.Tenant)
			require.Equal(t, orgId, contract.OrganizationId)
			require.Equal(t, testUserId, contract.LoggedInUserId)
			require.Equal(t, string(neo4jentity.DataSourceOpenline), contract.SourceFields.Source)
			require.Equal(t, constants.AppSourceCustomerOsApi, contract.SourceFields.AppSource)
			require.Equal(t, "", contract.Name)
			require.Equal(t, "", contract.ContractUrl)
			require.Equal(t, contractpb.RenewalCycle_NONE, contract.RenewalCycle)
			require.Equal(t, "", contract.Currency)
			require.False(t, contract.AutoRenew)
			require.True(t, contract.PayOnline)
			require.True(t, contract.PayAutomatically)
			require.True(t, contract.Check)
			require.True(t, contract.CanPayWithCard)
			require.True(t, contract.CanPayWithDirectDebit)
			require.True(t, contract.CanPayWithBankTransfer)
			require.Nil(t, contract.RenewalPeriods)
			require.Nil(t, contract.ServiceStartedAt)
			require.Nil(t, contract.SignedAt)
			require.Nil(t, contract.InvoicingStartDate)
			require.Equal(t, int64(0), contract.DueDays)
			calledCreateContract = true
			neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{
				Id: contractId,
			})
			return &contractpb.ContractIdGrpcResponse{
				Id: contractId,
			}, nil
		},
	}
	events_platform.SetContractCallbacks(&contractServiceCallbacks)

	rawResponse := callGraphQL(t, "contract/create_contract_default", map[string]interface{}{
		"orgId": orgId,
	})

	var contractStruct struct {
		Contract_Create model.Contract
	}

	require.Nil(t, rawResponse.Errors)
	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)
	contract := contractStruct.Contract_Create
	require.Equal(t, contractId, contract.ID)

	require.True(t, calledCreateContract)
}

func TestMutationResolver_ContractUpdate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})
	calledUpdateContract := false

	contractServiceCallbacks := events_platform.MockContractServiceCallbacks{
		UpdateContract: func(context context.Context, contract *contractpb.UpdateContractGrpcRequest) (*contractpb.ContractIdGrpcResponse, error) {
			require.Equal(t, tenantName, contract.Tenant)
			require.Equal(t, contractId, contract.Id)
			require.Equal(t, testUserId, contract.LoggedInUserId)
			require.Equal(t, string(neo4jentity.DataSourceOpenline), contract.SourceFields.Source)
			require.Equal(t, "test app source", contract.SourceFields.AppSource)
			require.Equal(t, "Updated Contract", contract.Name)
			require.Equal(t, "https://contract.com/updated", contract.ContractUrl)
			require.Equal(t, contractpb.RenewalCycle_ANNUALLY_RENEWAL, contract.RenewalCycle)
			require.Equal(t, int64(3), *contract.RenewalPeriods)
			expectedServiceStartedAt, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
			if err != nil {
				t.Fatalf("Failed to parse expected timestamp: %v", err)
			}
			require.Equal(t, timestamppb.New(expectedServiceStartedAt), contract.ServiceStartedAt)
			expectedSignedAt, err := time.Parse(time.RFC3339, "2019-02-01T00:00:00Z")
			if err != nil {
				t.Fatalf("Failed to parse expected timestamp: %v", err)
			}
			require.Equal(t, timestamppb.New(expectedSignedAt), contract.SignedAt)
			expectedEndedAt, err := time.Parse(time.RFC3339, "2019-03-01T00:00:00Z")
			if err != nil {
				t.Fatalf("Failed to parse expected timestamp: %v", err)
			}
			require.Equal(t, timestamppb.New(expectedEndedAt), contract.EndedAt)
			require.Equal(t, commonpb.BillingCycle_ANNUALLY_BILLING, contract.BillingCycle)
			require.Equal(t, "USD", contract.Currency)
			require.Equal(t, "test address line 1", contract.AddressLine1)
			require.Equal(t, "test address line 2", contract.AddressLine2)
			require.Equal(t, "test locality", contract.Locality)
			require.Equal(t, "test country", contract.Country)
			require.Equal(t, "test region", contract.Region)
			require.Equal(t, "test zip", contract.Zip)
			require.Equal(t, "test organization legal name", contract.OrganizationLegalName)
			require.Equal(t, "test invoice email", contract.InvoiceEmail)
			require.Equal(t, "test invoice note", contract.InvoiceNote)
			require.Equal(t, true, contract.CanPayWithCard)
			require.Equal(t, true, contract.CanPayWithDirectDebit)
			require.Equal(t, true, contract.CanPayWithBankTransfer)
			require.Equal(t, true, contract.PayOnline)
			require.Equal(t, true, contract.PayAutomatically)
			require.Equal(t, true, contract.AutoRenew)
			require.Equal(t, true, contract.Check)
			require.Equal(t, int64(7), contract.DueDays)
			require.Equal(t, 26, len(contract.FieldsMask))
			calledUpdateContract = true
			return &contractpb.ContractIdGrpcResponse{
				Id: contractId,
			}, nil
		},
	}
	events_platform.SetContractCallbacks(&contractServiceCallbacks)

	rawResponse := callGraphQL(t, "contract/update_contract", map[string]interface{}{
		"contractId": contractId,
	})

	var contractStruct struct {
		Contract_Update model.Contract
	}

	require.Nil(t, rawResponse.Errors)
	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)
	contract := contractStruct.Contract_Update
	require.Equal(t, contractId, contract.ID)

	require.True(t, calledUpdateContract)
}

func TestMutationResolver_ContractUpdate_NullDateFields(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})
	calledUpdateContract := false

	contractServiceCallbacks := events_platform.MockContractServiceCallbacks{
		UpdateContract: func(context context.Context, contract *contractpb.UpdateContractGrpcRequest) (*contractpb.ContractIdGrpcResponse, error) {
			require.Equal(t, tenantName, contract.Tenant)
			require.Equal(t, contractId, contract.Id)
			require.Equal(t, testUserId, contract.LoggedInUserId)
			require.Equal(t, string(neo4jentity.DataSourceOpenline), contract.SourceFields.Source)
			require.Equal(t, "customer-os-api", contract.SourceFields.AppSource)

			require.Nil(t, contract.SignedAt)
			require.Nil(t, contract.ServiceStartedAt)
			require.Nil(t, contract.EndedAt)
			require.Nil(t, contract.InvoicingStartDate)

			require.Equal(t, 4, len(contract.FieldsMask))
			calledUpdateContract = true
			return &contractpb.ContractIdGrpcResponse{
				Id: contractId,
			}, nil
		},
	}
	events_platform.SetContractCallbacks(&contractServiceCallbacks)

	rawResponse := callGraphQL(t, "contract/update_contract_null_dates", map[string]interface{}{
		"contractId": contractId,
	})

	var contractStruct struct {
		Contract_Update model.Contract
	}

	require.Nil(t, rawResponse.Errors)
	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)
	contract := contractStruct.Contract_Update
	require.Equal(t, contractId, contract.ID)
	require.Nil(t, contract.SignedAt)
	require.Nil(t, contract.ServiceStartedAt)
	require.Nil(t, contract.EndedAt)
	require.Nil(t, contract.InvoicingStartDate)

	require.True(t, calledUpdateContract)
}

func TestQueryResolver_Contract_WithServiceLineItems(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	now := utils.Now()
	tomorrow := now.Add(time.Duration(24) * time.Hour)
	yesterday := now.Add(time.Duration(-24) * time.Hour)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{
		AddressLine1:           "address line 1",
		AddressLine2:           "address line 2",
		Zip:                    "zip",
		Locality:               "locality",
		Country:                "country",
		Region:                 "region",
		OrganizationLegalName:  "organization legal name",
		InvoiceEmail:           "invoice email",
		InvoiceNote:            "invoice note",
		BillingCycle:           neo4jenum.BillingCycleMonthlyBilling,
		InvoicingStartDate:     &now,
		NextInvoiceDate:        &tomorrow,
		InvoicingEnabled:       true,
		CanPayWithCard:         true,
		CanPayWithDirectDebit:  true,
		CanPayWithBankTransfer: true,
		PayOnline:              true,
		PayAutomatically:       true,
		AutoRenew:              true,
		Check:                  true,
		DueDays:                int64(7),
	})

	serviceLineItemId1 := neo4jtest.CreateServiceLineItemForContract(ctx, driver, tenantName, contractId, neo4jentity.ServiceLineItemEntity{
		Name:      "service line item 1",
		CreatedAt: yesterday,
		UpdatedAt: yesterday,
		Billed:    neo4jenum.BilledTypeAnnually,
		Price:     13,
		Quantity:  2,
		Source:    neo4jentity.DataSourceOpenline,
		AppSource: "test1",
		VatRate:   0.1,
	})
	serviceLineItemId2 := neo4jtest.CreateServiceLineItemForContract(ctx, driver, tenantName, contractId, neo4jentity.ServiceLineItemEntity{
		Name:      "service line item 2",
		CreatedAt: now,
		UpdatedAt: now,
		Billed:    neo4jenum.BilledTypeUsage,
		Price:     255,
		Quantity:  23,
		Source:    neo4jentity.DataSourceOpenline,
		AppSource: "test2",
		VatRate:   0.2,
	})
	neo4jtest.AssertNeo4jNodeCount(ctx, t, driver, map[string]int{
		"Organization":    1,
		"Contract":        1,
		"ServiceLineItem": 2,
	})
	neo4jtest.AssertRelationship(ctx, t, driver, contractId, "HAS_SERVICE", serviceLineItemId1)
	neo4jtest.AssertRelationship(ctx, t, driver, contractId, "HAS_SERVICE", serviceLineItemId2)

	rawResponse := callGraphQL(t, "contract/get_contract_with_service_line_items",
		map[string]interface{}{"contractId": contractId})

	var contractStruct struct {
		Contract model.Contract
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)

	contract := contractStruct.Contract
	require.NotNil(t, contract)
	require.Equal(t, contractId, contract.Metadata.ID)
	require.True(t, contract.BillingEnabled)
	require.True(t, contract.AutoRenew)

	billingDetails := contract.BillingDetails
	require.Equal(t, "address line 1", *billingDetails.AddressLine1)
	require.Equal(t, "address line 2", *billingDetails.AddressLine2)
	require.Equal(t, "zip", *billingDetails.PostalCode)
	require.Equal(t, "locality", *billingDetails.Locality)
	require.Equal(t, "country", *billingDetails.Country)
	require.Equal(t, "region", *billingDetails.Region)
	require.Equal(t, "organization legal name", *billingDetails.OrganizationLegalName)
	require.Equal(t, "invoice email", *billingDetails.BillingEmail)
	require.Equal(t, "invoice note", *billingDetails.InvoiceNote)
	require.Equal(t, model.ContractBillingCycleMonthlyBilling, *billingDetails.BillingCycle)
	require.True(t, *billingDetails.CanPayWithCard)
	require.True(t, *billingDetails.CanPayWithDirectDebit)
	require.True(t, *billingDetails.CanPayWithBankTransfer)
	require.True(t, *billingDetails.Check)
	require.True(t, *billingDetails.PayOnline)
	require.True(t, *billingDetails.PayAutomatically)
	require.Equal(t, utils.StartOfDayInUTC(now), *billingDetails.InvoicingStarted)
	require.Equal(t, utils.StartOfDayInUTC(tomorrow), *billingDetails.NextInvoicing)
	require.Equal(t, int64(7), *billingDetails.DueDays)

	require.Equal(t, 2, len(contract.ContractLineItems))

	firstContractLineItem := contract.ContractLineItems[0]
	require.Equal(t, serviceLineItemId1, firstContractLineItem.Metadata.ID)
	require.Equal(t, "service line item 1", firstContractLineItem.Description)
	require.Equal(t, yesterday, firstContractLineItem.Metadata.Created)
	require.Equal(t, yesterday, firstContractLineItem.Metadata.LastUpdated)
	require.Equal(t, model.BilledTypeAnnually, firstContractLineItem.BillingCycle)
	require.Equal(t, float64(13), firstContractLineItem.Price)
	require.Equal(t, int64(2), firstContractLineItem.Quantity)
	require.Equal(t, model.DataSourceOpenline, firstContractLineItem.Metadata.Source)
	require.Equal(t, "test1", firstContractLineItem.Metadata.AppSource)
	require.Equal(t, 0.1, firstContractLineItem.Tax.TaxRate)

	secondContractLineItem := contract.ContractLineItems[1]
	require.Equal(t, serviceLineItemId2, secondContractLineItem.Metadata.ID)
	require.Equal(t, "service line item 2", secondContractLineItem.Description)
	require.Equal(t, now, secondContractLineItem.Metadata.Created)
	require.Equal(t, now, secondContractLineItem.Metadata.LastUpdated)
	require.Equal(t, model.BilledTypeUsage, secondContractLineItem.BillingCycle)
	require.Equal(t, float64(255), secondContractLineItem.Price)
	require.Equal(t, int64(23), secondContractLineItem.Quantity)
	require.Equal(t, model.DataSourceOpenline, secondContractLineItem.Metadata.Source)
	require.Equal(t, "test2", secondContractLineItem.Metadata.AppSource)
	require.Equal(t, 0.2, secondContractLineItem.Tax.TaxRate)
}

func TestQueryResolver_Contract_WithOpportunities(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	now := utils.Now()
	yesterday := now.Add(time.Duration(-24) * time.Hour)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})

	opportunityId1 := neo4jt.CreateOpportunityForContract(ctx, driver, tenantName, contractId, entity.OpportunityEntity{
		Name:          "oppo 1",
		CreatedAt:     now,
		UpdatedAt:     now,
		Amount:        49,
		InternalType:  entity.InternalTypeUpsell,
		InternalStage: entity.InternalStageOpen,
		Source:        neo4jentity.DataSourceOpenline,
		GeneralNotes:  "test notes 1",
		Comments:      "test comments 1",
		AppSource:     "test1",
	})
	opportunityId2 := neo4jt.CreateOpportunityForContract(ctx, driver, tenantName, contractId, entity.OpportunityEntity{
		Name:          "oppo 2",
		CreatedAt:     yesterday,
		UpdatedAt:     yesterday,
		Amount:        1239,
		InternalType:  entity.InternalTypeNbo,
		InternalStage: entity.InternalStageEvaluating,
		Source:        neo4jentity.DataSourceOpenline,
		GeneralNotes:  "test notes 2",
		Comments:      "test comments 2",
		AppSource:     "test2",
	})
	neo4jtest.AssertNeo4jNodeCount(ctx, t, driver, map[string]int{
		"Organization": 1,
		"Contract":     1,
		"Opportunity":  2,
	})
	neo4jtest.AssertRelationship(ctx, t, driver, contractId, "HAS_OPPORTUNITY", opportunityId1)

	rawResponse := callGraphQL(t, "contract/get_contract_with_opportunities",
		map[string]interface{}{"contractId": contractId})

	var contractStruct struct {
		Contract model.Contract
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &contractStruct)
	require.Nil(t, err)

	contract := contractStruct.Contract
	require.NotNil(t, contract)
	require.Equal(t, 2, len(contract.Opportunities))

	firstOpportunity := contract.Opportunities[0]
	require.Equal(t, opportunityId1, firstOpportunity.ID)
	require.Equal(t, "oppo 1", firstOpportunity.Name)
	require.Equal(t, now, firstOpportunity.CreatedAt)
	require.Equal(t, now, firstOpportunity.UpdatedAt)
	require.Equal(t, float64(49), firstOpportunity.Amount)
	require.Equal(t, model.InternalStageOpen, firstOpportunity.InternalStage)
	require.Equal(t, model.InternalTypeUpsell, firstOpportunity.InternalType)
	require.Equal(t, model.DataSourceOpenline, firstOpportunity.Source)
	require.Equal(t, "test notes 1", firstOpportunity.GeneralNotes)
	require.Equal(t, "test comments 1", firstOpportunity.Comments)
	require.Equal(t, "test1", firstOpportunity.AppSource)

	secondOpportunity := contract.Opportunities[1]
	require.Equal(t, opportunityId2, secondOpportunity.ID)
	require.Equal(t, "oppo 2", secondOpportunity.Name)
	require.Equal(t, yesterday, secondOpportunity.CreatedAt)
	require.Equal(t, yesterday, secondOpportunity.UpdatedAt)
	require.Equal(t, float64(1239), secondOpportunity.Amount)
	require.Equal(t, model.InternalStageEvaluating, secondOpportunity.InternalStage)
	require.Equal(t, model.InternalTypeNbo, secondOpportunity.InternalType)
	require.Equal(t, model.DataSourceOpenline, secondOpportunity.Source)
	require.Equal(t, "test notes 2", secondOpportunity.GeneralNotes)
	require.Equal(t, "test comments 2", secondOpportunity.Comments)
	require.Equal(t, "test2", secondOpportunity.AppSource)
}

func TestMutationResolver_ContractDelete(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})

	calledDeleteContractEvent := false

	contractCallbacks := events_platform.MockContractServiceCallbacks{
		SoftDeleteContract: func(context context.Context, contract *contractpb.SoftDeleteContractGrpcRequest) (*emptypb.Empty, error) {
			require.Equal(t, tenantName, contract.Tenant)
			require.Equal(t, contractId, contract.Id)
			require.Equal(t, testUserId, contract.LoggedInUserId)
			require.Equal(t, constants.AppSourceCustomerOsApi, constants.AppSourceCustomerOsApi)
			calledDeleteContractEvent = true
			return &emptypb.Empty{}, nil
		},
	}
	events_platform.SetContractCallbacks(&contractCallbacks)

	rawResponse := callGraphQL(t, "contract/delete_contract", map[string]interface{}{
		"contractId": contractId,
	})

	var response struct {
		Contract_Delete model.DeleteResponse
	}

	require.Nil(t, rawResponse.Errors)
	err := decode.Decode(rawResponse.Data.(map[string]any), &response)
	require.Nil(t, err)
	require.True(t, response.Contract_Delete.Accepted)
	require.False(t, response.Contract_Delete.Completed)
	require.True(t, calledDeleteContractEvent)
}

func TestMutationResolver_AddAttachmentToContract(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})

	attachmentId := neo4jt.CreateAttachment(ctx, driver, tenantName, entity.AttachmentEntity{
		MimeType: "text/plain",
		FileName: "readme.txt",
	})

	rawResponse, err := c.RawPost(getQuery("contract/contract_add_attachment"),
		client.Var("contractId", contractId),
		client.Var("attachmentId", attachmentId))
	assertRawResponseSuccess(t, rawResponse, err)

	require.Equal(t, 1, neo4jtest.GetCountOfNodes(ctx, driver, "Contract"))
	require.Equal(t, 1, neo4jtest.GetCountOfNodes(ctx, driver, "Attachment"))
	require.Equal(t, 1, neo4jtest.GetCountOfRelationships(ctx, driver, "INCLUDES"))

	var meeting struct {
		Contract_AddAttachment model.Contract
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &meeting)
	require.Nil(t, err)

	require.NotNil(t, meeting.Contract_AddAttachment.ID)
	require.Len(t, meeting.Contract_AddAttachment.Attachments, 1)
	require.Equal(t, meeting.Contract_AddAttachment.Attachments[0].ID, attachmentId)
}

func TestMutationResolver_RemoveAttachmentFromContract(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	contractId := neo4jtest.CreateContractForOrganization(ctx, driver, tenantName, orgId, neo4jentity.ContractEntity{})

	attachmentId1 := neo4jt.CreateAttachment(ctx, driver, tenantName, entity.AttachmentEntity{
		MimeType: "text/plain",
		FileName: "readme1.txt",
	})

	attachmentId2 := neo4jt.CreateAttachment(ctx, driver, tenantName, entity.AttachmentEntity{
		MimeType: "text/plain",
		FileName: "readme2.txt",
	})

	addAttachment1Response, err := c.RawPost(getQuery("contract/contract_add_attachment"),
		client.Var("contractId", contractId),
		client.Var("attachmentId", attachmentId1))
	assertRawResponseSuccess(t, addAttachment1Response, err)

	addAttachment2Response, err := c.RawPost(getQuery("contract/contract_add_attachment"),
		client.Var("contractId", contractId),
		client.Var("attachmentId", attachmentId2))
	assertRawResponseSuccess(t, addAttachment2Response, err)

	require.Equal(t, 1, neo4jtest.GetCountOfNodes(ctx, driver, "Contract"))
	require.Equal(t, 2, neo4jtest.GetCountOfNodes(ctx, driver, "Attachment"))
	require.Equal(t, 2, neo4jtest.GetCountOfRelationships(ctx, driver, "INCLUDES"))

	removeAttachmentResponse, err := c.RawPost(getQuery("contract/contract_remove_attachment"),
		client.Var("contractId", contractId),
		client.Var("attachmentId", attachmentId2))
	assertRawResponseSuccess(t, removeAttachmentResponse, err)

	require.Equal(t, 1, neo4jtest.GetCountOfNodes(ctx, driver, "Contract"))
	require.Equal(t, 2, neo4jtest.GetCountOfNodes(ctx, driver, "Attachment"))
	require.Equal(t, 1, neo4jtest.GetCountOfRelationships(ctx, driver, "INCLUDES"))

	var meeting struct {
		Contract_RemoveAttachment model.Contract
	}

	err = decode.Decode(removeAttachmentResponse.Data.(map[string]any), &meeting)
	require.Nil(t, err)

	require.NotNil(t, meeting.Contract_RemoveAttachment.ID)
	require.Len(t, meeting.Contract_RemoveAttachment.Attachments, 1)
	require.Equal(t, meeting.Contract_RemoveAttachment.Attachments[0].ID, attachmentId1)
}
