package resolver

import (
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/grpc/events_platform"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jtest "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/test"
	organizationpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/organization"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"testing"
)

func TestMutationResolver_BillingProfileCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	billingProfileId := uuid.New().String()
	calledCreateBillingProfile := false

	organizationServiceCallbacks := events_platform.MockOrganizationServiceCallbacks{
		CreateBillingProfile: func(context context.Context, billingProfile *organizationpb.CreateBillingProfileGrpcRequest) (*organizationpb.BillingProfileIdGrpcResponse, error) {
			require.Equal(t, tenantName, billingProfile.Tenant)
			require.Equal(t, testUserId, billingProfile.LoggedInUserId)
			require.Equal(t, neo4jentity.DataSourceOpenline.String(), billingProfile.SourceFields.Source)
			require.Equal(t, constants.AppSourceCustomerOsApi, billingProfile.SourceFields.AppSource)
			require.Equal(t, "New profile", billingProfile.LegalName)
			require.Equal(t, "123456789", billingProfile.TaxId)
			require.Nil(t, billingProfile.CreatedAt)
			require.Equal(t, orgId, billingProfile.OrganizationId)
			calledCreateBillingProfile = true
			return &organizationpb.BillingProfileIdGrpcResponse{
				Id: billingProfileId,
			}, nil
		},
	}
	events_platform.SetOrganizationCallbacks(&organizationServiceCallbacks)

	rawResponse := callGraphQL(t, "billing_profile/create_billing_profile", map[string]interface{}{"organizationId": orgId})
	require.Nil(t, rawResponse.Errors)

	var billingProfileStruct struct {
		BillingProfile_Create string
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &billingProfileStruct)
	require.Nil(t, err)

	// Verify
	require.Equal(t, billingProfileId, billingProfileStruct.BillingProfile_Create)

	require.True(t, calledCreateBillingProfile)
}

func TestMutationResolver_BillingProfileUpdate(t *testing.T) {
	ctx := context.Background()
	defer tearDownTestCase(ctx)(t)

	neo4jtest.CreateTenant(ctx, driver, tenantName)
	neo4jtest.CreateUserWithId(ctx, driver, tenantName, testUserId)
	orgId := neo4jtest.CreateOrganization(ctx, driver, tenantName, neo4jentity.OrganizationEntity{})
	billingProfileId := neo4jtest.CreateBillingProfileForOrganization(ctx, driver, tenantName, orgId, neo4jentity.BillingProfileEntity{
		LegalName: "Old profile",
		TaxId:     "987654321",
	})
	calledUpdateBillingProfile := false

	organizationServiceCallbacks := events_platform.MockOrganizationServiceCallbacks{
		UpdateBillingProfile: func(context context.Context, billingProfile *organizationpb.UpdateBillingProfileGrpcRequest) (*organizationpb.BillingProfileIdGrpcResponse, error) {
			require.Equal(t, tenantName, billingProfile.Tenant)
			require.Equal(t, testUserId, billingProfile.LoggedInUserId)
			require.Equal(t, "New name", billingProfile.LegalName)
			require.Equal(t, "", billingProfile.TaxId)
			require.Equal(t, []organizationpb.BillingProfileFieldMask{organizationpb.BillingProfileFieldMask_BILLING_PROFILE_PROPERTY_LEGAL_NAME}, billingProfile.FieldsMask)
			require.Nil(t, billingProfile.UpdatedAt)
			require.Equal(t, orgId, billingProfile.OrganizationId)
			calledUpdateBillingProfile = true
			return &organizationpb.BillingProfileIdGrpcResponse{
				Id: billingProfileId,
			}, nil
		},
	}
	events_platform.SetOrganizationCallbacks(&organizationServiceCallbacks)

	rawResponse := callGraphQL(t, "billing_profile/update_billing_profile", map[string]interface{}{"organizationId": orgId, "billingProfileId": billingProfileId})
	require.Nil(t, rawResponse.Errors)

	var billingProfileStruct struct {
		BillingProfile_Update string
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &billingProfileStruct)
	require.Nil(t, err)

	// Verify
	require.Equal(t, billingProfileId, billingProfileStruct.BillingProfile_Update)

	require.True(t, calledUpdateBillingProfile)
}
