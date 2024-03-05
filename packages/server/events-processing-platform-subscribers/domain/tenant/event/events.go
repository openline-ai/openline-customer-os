package event

const (
	TenantAddBillingProfileV1    = "V1_TENANT_BILLING_PROFILE_NEW"
	TenantUpdateBillingProfileV1 = "V1_TENANT_BILLING_PROFILE_UPDATE"
	TenantUpdateSettingsV1       = "V1_TENANT_SETTINGS_UPDATE"
)

const (
	FieldMaskPhone                         = "phone"
	FieldMaskAddressLine1                  = "addressLine1"
	FieldMaskAddressLine2                  = "addressLine2"
	FieldMaskAddressLine3                  = "addressLine3"
	FieldMaskZip                           = "zip"
	FieldMaskCountry                       = "country"
	FieldMaskLocality                      = "locality"
	FieldMaskLegalName                     = "legalName"
	FieldMaskDomesticPaymentsBankInfo      = "domesticPaymentsBankInfo"
	FieldMaskInternationalPaymentsBankInfo = "internationalPaymentsBankInfo"
	FieldMaskVatNumber                     = "vatNumber"
	FieldMaskSendInvoicesFrom              = "sendInvoicesFrom"
	FieldMaskSendInvoicesBcc               = "sendInvoicesBcc"
	FieldMaskCanPayWithCard                = "canPayWithCard"
	FieldMaskCanPayWithDirectDebitSEPA     = "canPayWithDirectDebitSEPA"
	FieldMaskCanPayWithDirectDebitACH      = "canPayWithDirectDebitACH"
	FieldMaskCanPayWithDirectDebitBacs     = "canPayWithDirectDebitBacs"
	FieldMaskCanPayWithPigeon              = "canPayWithPigeon"

	FieldMaskLogoRepositoryFileId = "logoRepositoryFileId"
	FieldMaskBaseCurrency         = "baseCurrency"
	FieldMaskInvoicingEnabled     = "invoicingEnabled"
	FieldMaskInvoicingPostpaid    = "invoicingPostpaid"
)
