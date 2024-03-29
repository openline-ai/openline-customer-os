package servicet

import (
	"context"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	tenant "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/tenant"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/tenant/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test"
	eventstoret "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/eventstore"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	tenantpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/tenant"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTenantService_AddBillingProfile(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	now := utils.Now()

	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.AddBillingProfile(ctx, &tenantpb.AddBillingProfileRequest{
		Tenant: tenantName,
		SourceFields: &commonpb.SourceFields{
			AppSource: "app",
			Source:    "source",
		},
		CreatedAt:              utils.ConvertTimeToTimestampPtr(&now),
		Email:                  "email",
		Phone:                  "phone",
		AddressLine1:           "addressLine1",
		AddressLine2:           "addressLine2",
		AddressLine3:           "addressLine3",
		Locality:               "locality",
		Country:                "country",
		Region:                 "region",
		Zip:                    "zip",
		LegalName:              "legalName",
		VatNumber:              "vatNumber",
		SendInvoicesFrom:       "sendInvoicesFrom",
		CanPayWithPigeon:       true,
		CanPayWithBankTransfer: true,
		Check:                  true,
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	billingProfileId := response.Id
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantAddBillingProfileV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantBillingProfileCreateEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, now, eventData.CreatedAt)
	require.Equal(t, billingProfileId, eventData.Id)
	require.Equal(t, "phone", eventData.Phone)
	require.Equal(t, "addressLine1", eventData.AddressLine1)
	require.Equal(t, "addressLine2", eventData.AddressLine2)
	require.Equal(t, "addressLine3", eventData.AddressLine3)
	require.Equal(t, "locality", eventData.Locality)
	require.Equal(t, "country", eventData.Country)
	require.Equal(t, "region", eventData.Region)
	require.Equal(t, "zip", eventData.Zip)
	require.Equal(t, "legalName", eventData.LegalName)
	require.Equal(t, "vatNumber", eventData.VatNumber)
	require.Equal(t, "sendInvoicesFrom", eventData.SendInvoicesFrom)
	require.Equal(t, true, eventData.CanPayWithPigeon)
	require.Equal(t, true, eventData.CanPayWithBankTransfer)
	require.Equal(t, true, eventData.Check)
	require.Equal(t, "app", eventData.SourceFields.AppSource)
	require.Equal(t, "source", eventData.SourceFields.Source)
	require.Equal(t, "source", eventData.SourceFields.SourceOfTruth)
}

func TestTenantService_UpdateBillingProfile(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	billingProfileId := uuid.New().String()
	now := utils.Now()

	// setup aggregate and create initial event
	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.UpdateBillingProfile(ctx, &tenantpb.UpdateBillingProfileRequest{
		Tenant:                 tenantName,
		Id:                     billingProfileId,
		AppSource:              "test",
		UpdatedAt:              utils.ConvertTimeToTimestampPtr(&now),
		Email:                  "email",
		Phone:                  "phone",
		LegalName:              "legalName",
		AddressLine1:           "addressLine1",
		AddressLine2:           "addressLine2",
		AddressLine3:           "addressLine3",
		Locality:               "locality",
		Country:                "country",
		Region:                 "region",
		Zip:                    "zip",
		VatNumber:              "vatNumber",
		SendInvoicesFrom:       "sendInvoicesFrom",
		CanPayWithPigeon:       true,
		CanPayWithBankTransfer: true,
		Check:                  true,
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantUpdateBillingProfileV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantBillingProfileUpdateEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, now, eventData.UpdatedAt)
	require.Equal(t, billingProfileId, eventData.Id)
	require.Equal(t, "phone", eventData.Phone)
	require.Equal(t, "addressLine1", eventData.AddressLine1)
	require.Equal(t, "addressLine2", eventData.AddressLine2)
	require.Equal(t, "addressLine3", eventData.AddressLine3)
	require.Equal(t, "locality", eventData.Locality)
	require.Equal(t, "country", eventData.Country)
	require.Equal(t, "region", eventData.Region)
	require.Equal(t, "zip", eventData.Zip)
	require.Equal(t, "legalName", eventData.LegalName)
	require.Equal(t, "vatNumber", eventData.VatNumber)
	require.Equal(t, "sendInvoicesFrom", eventData.SendInvoicesFrom)
	require.Equal(t, true, eventData.CanPayWithPigeon)
	require.Equal(t, true, eventData.CanPayWithBankTransfer)
	require.Equal(t, true, eventData.Check)
	require.Equal(t, 0, len(eventData.FieldsMask))
}

func TestTenantService_UpdateTenantSettings(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	now := utils.Now()

	// setup aggregate and create initial event
	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.UpdateTenantSettings(ctx, &tenantpb.UpdateTenantSettingsRequest{
		Tenant:               tenantName,
		AppSource:            "test",
		UpdatedAt:            utils.ConvertTimeToTimestampPtr(&now),
		LogoRepositoryFileId: "logoRepositoryFileId",
		DefaultCurrency:      "USD",
		BaseCurrency:         "USD",
		InvoicingEnabled:     true,
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantUpdateSettingsV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantSettingsUpdateEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, now, eventData.UpdatedAt)
	require.Equal(t, "logoRepositoryFileId", eventData.LogoRepositoryFileId)
	require.Equal(t, "USD", eventData.BaseCurrency)
	require.Equal(t, true, eventData.InvoicingEnabled)
	require.Equal(t, 0, len(eventData.FieldsMask))
}

func TestTenantService_AddBankAccount(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	now := utils.Now()

	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.AddBankAccount(ctx, &tenantpb.AddBankAccountGrpcRequest{
		Tenant: tenantName,
		SourceFields: &commonpb.SourceFields{
			AppSource: "app",
			Source:    "source",
		},
		CreatedAt:           utils.ConvertTimeToTimestampPtr(&now),
		BankName:            "bankName",
		Currency:            "USD",
		BankTransferEnabled: true,
		AllowInternational:  true,
		AccountNumber:       "accountNumber",
		SortCode:            "sortCode",
		Iban:                "iban",
		Bic:                 "swiftBIC",
		RoutingNumber:       "routingNumber",
		OtherDetails:        "otherDetails",
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	bankAccountId := response.Id
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantAddBankAccountV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantBankAccountCreateEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, now, eventData.CreatedAt)
	require.Equal(t, bankAccountId, eventData.Id)
	require.Equal(t, "bankName", eventData.BankName)
	require.Equal(t, "USD", eventData.Currency)
	require.True(t, eventData.BankTransferEnabled)
	require.True(t, eventData.AllowInternational)
	require.Equal(t, "accountNumber", eventData.AccountNumber)
	require.Equal(t, "sortCode", eventData.SortCode)
	require.Equal(t, "iban", eventData.Iban)
	require.Equal(t, "swiftBIC", eventData.Bic)
	require.Equal(t, "routingNumber", eventData.RoutingNumber)
	require.Equal(t, "otherDetails", eventData.OtherDetails)
}

func TestTenantService_UpdateBankAccount(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	bankAccountId := uuid.New().String()
	now := utils.Now()

	// setup aggregate and create initial event
	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.UpdateBankAccount(ctx, &tenantpb.UpdateBankAccountGrpcRequest{
		Tenant:              tenantName,
		Id:                  bankAccountId,
		AppSource:           "test",
		UpdatedAt:           utils.ConvertTimeToTimestampPtr(&now),
		BankName:            "bankName",
		Currency:            "USD",
		BankTransferEnabled: true,
		AllowInternational:  true,
		AccountNumber:       "accountNumber",
		SortCode:            "sortCode",
		Iban:                "iban",
		Bic:                 "swiftBIC",
		RoutingNumber:       "routingNumber",
		OtherDetails:        "otherDetails",
		FieldsMask: []tenantpb.BankAccountFieldMask{
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_BANK_NAME,
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_CURRENCY,
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_BANK_TRANSFER_ENABLED,
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_ACCOUNT_NUMBER,
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_ROUTING_NUMBER,
			tenantpb.BankAccountFieldMask_BANK_ACCOUNT_FIELD_OTHER_DETAILS,
		},
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantUpdateBankAccountV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantBankAccountUpdateEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, now, eventData.UpdatedAt)
	require.Equal(t, bankAccountId, eventData.Id)
	require.Equal(t, "bankName", eventData.BankName)
	require.Equal(t, "USD", eventData.Currency)
	require.Equal(t, true, eventData.BankTransferEnabled)
	require.Equal(t, true, eventData.AllowInternational)
	require.Equal(t, "accountNumber", eventData.AccountNumber)
	require.Equal(t, "sortCode", eventData.SortCode)
	require.Equal(t, "iban", eventData.Iban)
	require.Equal(t, "swiftBIC", eventData.Bic)
	require.Equal(t, "routingNumber", eventData.RoutingNumber)
	require.Equal(t, "otherDetails", eventData.OtherDetails)
	require.Equal(t, 6, len(eventData.FieldsMask))
}

func TestTenantService_DeleteBankAccount(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx, testDatabase)(t)

	// setup test environment
	tenantName := "ziggy"
	bankAccountId := uuid.New().String()

	// setup aggregate and create initial event
	aggregateStore := eventstoret.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	require.Nil(t, err, "Failed to get grpc connection")
	tenantServiceClient := tenantpb.NewTenantGrpcServiceClient(grpcConnection)

	response, err := tenantServiceClient.DeleteBankAccount(ctx, &tenantpb.DeleteBankAccountGrpcRequest{
		Tenant: tenantName,
		Id:     bankAccountId,
	})
	require.Nil(t, err)
	require.NotNil(t, response)

	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))

	tenantAggregate := tenant.NewTenantAggregate(tenantName)
	eventList := eventsMap[tenantAggregate.ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, event.TenantDeleteBankAccountV1, eventList[0].GetEventType())
	require.Equal(t, string(tenant.TenantAggregateType)+"-"+tenantName, eventList[0].GetAggregateID())

	var eventData event.TenantBankAccountDeleteEvent
	err = eventList[0].GetJsonData(&eventData)
	require.Nil(t, err, "Failed to unmarshal event data")

	// Assertions to validate the contract create event data
	require.Equal(t, tenantName, eventData.Tenant)
	require.Equal(t, bankAccountId, eventData.Id)
	test.AssertRecentTime(t, eventData.DeletedAt)
}
