package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/settings-api/repository/entity"
)

// TODO the state should come from the actual running service
func MapTenantSettingsEntityToDTO(tenantSettings *entity.TenantSettings) *map[string]interface{} {
	responseMap := make(map[string]interface{})

	if tenantSettings == nil {
		return &responseMap
	}

	if tenantSettings.HubspotPrivateAppKey != nil {
		responseMap["hubspot"] = make(map[string]interface{})
		responseMap["hubspot"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.ZendeskAPIKey != nil && tenantSettings.ZendeskSubdomain != nil && tenantSettings.ZendeskAdminEmail != nil {
		responseMap["zendesk"] = make(map[string]interface{})
		responseMap["zendesk"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.SmartSheetId != nil && tenantSettings.SmartSheetAccessToken != nil {
		responseMap["smartsheet"] = make(map[string]interface{})
		responseMap["smartsheet"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.JiraAPIToken != nil && tenantSettings.JiraDomain != nil && tenantSettings.JiraEmail != nil {
		responseMap["jira"] = make(map[string]interface{})
		responseMap["jira"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.TrelloAPIToken != nil && tenantSettings.TrelloAPIKey != nil {
		responseMap["trello"] = make(map[string]interface{})
		responseMap["trello"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.AhaAPIUrl != nil && tenantSettings.AhaAPIKey != nil {
		responseMap["aha"] = make(map[string]interface{})
		responseMap["aha"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.AirtablePersonalAccessToken != nil {
		responseMap["airtable"] = make(map[string]interface{})
		responseMap["airtable"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.AmplitudeSecretKey != nil && tenantSettings.AmplitudeAPIKey != nil {
		responseMap["amplitude"] = make(map[string]interface{})
		responseMap["amplitude"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.AsanaAccessToken != nil {
		responseMap["asana"] = make(map[string]interface{})
		responseMap["asana"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.BatonAPIKey != nil {
		responseMap["baton"] = make(map[string]interface{})
		responseMap["baton"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.BabelforceRegionEnvironment != nil && tenantSettings.BabelforceAccessKeyId != nil && tenantSettings.BabelforceAccessToken != nil {
		responseMap["babelforce"] = make(map[string]interface{})
		responseMap["babelforce"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.BigQueryServiceAccountKey != nil {
		responseMap["bigquery"] = make(map[string]interface{})
		responseMap["bigquery"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.BraintreeEnvironment != nil && tenantSettings.BraintreeMerchantId != nil && tenantSettings.BraintreePublicKey != nil && tenantSettings.BraintreePrivateKey != nil {
		responseMap["braintree"] = make(map[string]interface{})
		responseMap["braintree"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.CallRailAccount != nil && tenantSettings.CallRailApiToken != nil {
		responseMap["callrail"] = make(map[string]interface{})
		responseMap["callrail"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.ChargebeeProductCatalog != nil && tenantSettings.ChargebeeApiKey != nil {
		responseMap["chargebee"] = make(map[string]interface{})
		responseMap["chargebee"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.ChargifyApiKey != nil && tenantSettings.ChargifyDomain != nil {
		responseMap["chargify"] = make(map[string]interface{})
		responseMap["chargify"].(map[string]interface{})["state"] = "ACTIVE"
	}

	if tenantSettings != nil && tenantSettings.ClickUpApiKey != nil {
		responseMap["clickup"] = make(map[string]interface{})
		responseMap["clickup"].(map[string]interface{})["state"] = "ACTIVE"
	}

	return &responseMap
}
