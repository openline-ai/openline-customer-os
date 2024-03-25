package event

const (
	ContractCreateV1                    = "V1_CONTRACT_CREATE"
	ContractUpdateV1                    = "V1_CONTRACT_UPDATE"
	ContractUpdateStatusV1              = "V1_CONTRACT_UPDATE_STATUS"
	ContractRolloutRenewalOpportunityV1 = "V1_CONTRACT_ROLLOUT_RENEWAL_OPPORTUNITY"
	ContractDeleteV1                    = "V1_CONTRACT_DELETE"
)

const (
	FieldMaskName                   = "name"
	FieldMaskContractURL            = "contractURL"
	FieldMaskSignedAt               = "signedAt"
	FieldMaskEndedAt                = "endedAt"
	FieldMaskServiceStartedAt       = "serviceStartedAt"
	FieldMaskInvoicingStartDate     = "invoicingStartDate"
	FieldMaskRenewalCycle           = "renewalCycle"
	FieldMaskRenewalPeriods         = "renewalPeriods"
	FieldMaskBillingCycle           = "billingCycle"
	FieldMaskCurrency               = "currency"
	FieldMaskAddressLine1           = "addressLine1"
	FieldMaskAddressLine2           = "addressLine2"
	FieldMaskZip                    = "zip"
	FieldMaskCountry                = "country"
	FieldMaskLocality               = "locality"
	FieldMaskOrganizationLegalName  = "organizationLegalName"
	FieldMaskInvoiceEmail           = "invoiceEmail"
	FieldMaskStatus                 = "status"
	FieldMaskInvoiceNote            = "invoiceNote"
	FieldMaskNextInvoiceDate        = "nextInvoiceDate"
	FieldMaskCanPayWithCard         = "canPayWithCard"
	FieldMaskCanPayWithDirectDebit  = "canPayWithDirectDebit"
	FieldMaskCanPayWithBankTransfer = "canPayWithBankTransfer"
	FieldMaskPayOnline              = "payOnline"
	FieldMaskPayAutomatically       = "payAutomatically"
	FieldMaskInvoicingEnabled       = "invoicingEnabled"
	FieldMaskAutoRenew              = "autoRenew"
)
