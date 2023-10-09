package main

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/grpc_client/interceptor"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	common_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/common"
	email_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/email"
	interaction_event_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/interaction_event"
	log_entry_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/log_entry"
	organization_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/organization"
	phone_number_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/phone_number"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcApiKey = "082c1193-a5a2-42fc-87fc-e960e692fffd"

type Clients struct {
	InteractionEventClient interaction_event_grpc_service.InteractionEventGrpcServiceClient
	OrganizationClient     organization_grpc_service.OrganizationGrpcServiceClient
	EmailClient            email_grpc_service.EmailGrpcServiceClient
	PhoneNumberClient      phone_number_grpc_service.PhoneNumberGrpcServiceClient
	LogEntryClient         log_entry_grpc_service.LogEntryGrpcServiceClient
}

var clients *Clients

func main() {
	InitClients()
	//testRequestGenerateSummaryRequest()
	//testRequestGenerateActionItemsRequest()
	//testCreateOrganization()
	testUpdateOrganization()
	//testHideOrganization()
	//testShowOrganization()
	//testCreateLogEntry()
	//testUpdateLogEntry()
	//testAddCustomField()
	//testCreateEmail()
	//testCreatePhoneNumber()
	//testAddParentOrganization()
	//testRemoveParentOrganization()
}

func InitClients() {
	conn, _ := grpc.Dial("localhost:5001", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			interceptor.ApiKeyEnricher(grpcApiKey),
		))
	clients = &Clients{
		InteractionEventClient: interaction_event_grpc_service.NewInteractionEventGrpcServiceClient(conn),
		OrganizationClient:     organization_grpc_service.NewOrganizationGrpcServiceClient(conn),
		LogEntryClient:         log_entry_grpc_service.NewLogEntryGrpcServiceClient(conn),
		EmailClient:            email_grpc_service.NewEmailGrpcServiceClient(conn),
		PhoneNumberClient:      phone_number_grpc_service.NewPhoneNumberGrpcServiceClient(conn),
	}
}

func testRequestGenerateSummaryRequest() {
	tenant := "openline"
	interactionEventId := "555263fe-2e39-48f0-a8c2-c4c7a5ffb23d"

	result, _ := clients.InteractionEventClient.RequestGenerateSummary(context.TODO(), &interaction_event_grpc_service.RequestGenerateSummaryGrpcRequest{
		Tenant:             tenant,
		InteractionEventId: interactionEventId,
	})
	print(result)
}

func testRequestGenerateActionItemsRequest() {
	tenant := "openline"
	interactionEventId := "555263fe-2e39-48f0-a8c2-c4c7a5ffb23d"

	result, _ := clients.InteractionEventClient.RequestGenerateActionItems(context.TODO(), &interaction_event_grpc_service.RequestGenerateActionItensGrpcRequest{
		Tenant:             tenant,
		InteractionEventId: interactionEventId,
	})
	print(result)
}

func testCreateOrganization() {
	tenant := "openline"
	userId := "697563a8-171c-4950-a067-1aaaaf2de1d8"
	organizationId := "ccc"
	website := ""

	result, _ := clients.OrganizationClient.UpsertOrganization(context.TODO(), &organization_grpc_service.UpsertOrganizationGrpcRequest{
		Tenant:  tenant,
		Id:      organizationId,
		Website: website,
		UserId:  userId,
	})
	print(result)
}

func testUpdateOrganization() {
	tenant := "openline"
	organizationId := "cfaaf31f-ec3b-44d1-836e-4e50834632ae"
	website := "xtz.com"
	lastFoundingAmont := "1Million"
	partial := true

	result, _ := clients.OrganizationClient.UpsertOrganization(context.TODO(), &organization_grpc_service.UpsertOrganizationGrpcRequest{
		Tenant:            tenant,
		Id:                organizationId,
		Website:           website,
		LastFundingAmount: lastFoundingAmont,
		IgnoreEmptyFields: partial,
	})
	print(result)
}

func testHideOrganization() {
	tenant := "openline"
	organizationId := "ccc"

	result, _ := clients.OrganizationClient.HideOrganization(context.TODO(), &organization_grpc_service.OrganizationIdGrpcRequest{
		Tenant:         tenant,
		OrganizationId: organizationId,
	})
	print(result)
}

func testShowOrganization() {
	tenant := "openline"
	organizationId := "ccc"

	result, _ := clients.OrganizationClient.ShowOrganization(context.TODO(), &organization_grpc_service.OrganizationIdGrpcRequest{
		Tenant:         tenant,
		OrganizationId: organizationId,
	})
	print(result)
}

func testCreateLogEntry() {
	tenant := "openline"
	organizationId := "5e72b6fb-5f20-4973-9b96-52f4543a0df3"
	userId := "development@openline.ai"
	authorId := "c61f8af2-0e46-4464-a5db-ded8e4fe242f"

	result, _ := clients.LogEntryClient.UpsertLogEntry(context.TODO(), &log_entry_grpc_service.UpsertLogEntryGrpcRequest{
		Tenant:               tenant,
		LoggedOrganizationId: utils.StringPtr(organizationId),
		SourceFields: &common_grpc_service.SourceFields{
			AppSource: "test_app",
		},
		AuthorUserId: utils.StringPtr(authorId),
		Content:      "I spoke with client",
		ContentType:  "text/plain",
		UserId:       userId,
	})
	print(result)
}

func testUpdateLogEntry() {
	tenant := "openline"
	userId := "development@openline.ai"
	logEntryId := "ccffe134-4bcd-4fa0-955f-c79b9e1a985f"

	result, _ := clients.LogEntryClient.UpsertLogEntry(context.TODO(), &log_entry_grpc_service.UpsertLogEntryGrpcRequest{
		Tenant:      tenant,
		Id:          logEntryId,
		Content:     "new content",
		ContentType: "text/plain2",
		UserId:      userId,
		StartedAt:   timestamppb.New(utils.Now()),
	})
	print(result)
}

func testAddCustomField() {
	tenant := "openline"
	organizationId := "5e72b6fb-5f20-4973-9b96-52f4543a0df3"
	userId := "development@openline.ai"
	result, _ := clients.OrganizationClient.UpsertCustomFieldToOrganization(context.TODO(), &organization_grpc_service.CustomFieldForOrganizationGrpcRequest{
		Tenant:                tenant,
		OrganizationId:        organizationId,
		UserId:                userId,
		CustomFieldTemplateId: utils.StringPtr("c70cd2fb-1c31-46fd-851c-2e47ceba508f"),
		CustomFieldName:       "CF1",
		CustomFieldDataType:   organization_grpc_service.CustomFieldDataType_TEXT,
		CustomFieldValue: &organization_grpc_service.CustomFieldValue{
			StringValue: utils.StringPtr("super secret value"),
		},
	})
	print(result)
}

func testCreateEmail() {
	tenant := "openline"
	userId := "697563a8-171c-4950-a067-1aaaaf2de1d8"
	rawEmail := "aa@test.com"

	result, _ := clients.EmailClient.UpsertEmail(context.TODO(), &email_grpc_service.UpsertEmailGrpcRequest{
		Tenant:         tenant,
		RawEmail:       rawEmail,
		LoggedInUserId: userId,
		SourceFields: &common_grpc_service.SourceFields{
			AppSource: "test_app",
		},
	})
	print(result)
}

func testCreatePhoneNumber() {
	tenant := "openline"
	userId := "697563a8-171c-4950-a067-1aaaaf2de1d8"
	rawPhoneNumber := "+1234"

	result, _ := clients.PhoneNumberClient.UpsertPhoneNumber(context.TODO(), &phone_number_grpc_service.UpsertPhoneNumberGrpcRequest{
		Tenant:         tenant,
		PhoneNumber:    rawPhoneNumber,
		LoggedInUserId: userId,
	})
	print(result)
}

func testAddParentOrganization() {
	tenant := "openline"
	orgId := "cfaaf31f-ec3b-44d1-836e-4e50834632ae"
	parentOrgId := "05f382ba-0fa9-4828-940c-efb4e2e6b84c"
	relType := "store"
	result, err := clients.OrganizationClient.AddParentOrganization(context.TODO(), &organization_grpc_service.AddParentOrganizationGrpcRequest{
		Tenant:               tenant,
		OrganizationId:       orgId,
		ParentOrganizationId: parentOrgId,
		Type:                 relType,
	})
	if err != nil {
		print(err)
	}
	print(result)
}

func testRemoveParentOrganization() {
	tenant := "openline"
	orgId := "cfaaf31f-ec3b-44d1-836e-4e50834632ae"
	parentOrgId := "05f382ba-0fa9-4828-940c-efb4e2e6b84c"
	result, err := clients.OrganizationClient.RemoveParentOrganization(context.TODO(), &organization_grpc_service.RemoveParentOrganizationGrpcRequest{
		Tenant:               tenant,
		OrganizationId:       orgId,
		ParentOrganizationId: parentOrgId,
	})
	if err != nil {
		print(err)
	}
	print(result)
}
