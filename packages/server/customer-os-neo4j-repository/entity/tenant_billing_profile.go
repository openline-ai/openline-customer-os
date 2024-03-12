package entity

import (
	"time"
)

type TenantBillingProfileEntity struct {
	Id                                string
	CreatedAt                         time.Time
	UpdatedAt                         time.Time
	LegalName                         string
	Phone                             string
	AddressLine1                      string
	AddressLine2                      string
	AddressLine3                      string
	Locality                          string
	Country                           string
	Zip                               string
	DomesticPaymentsBankInfo          string
	InternationalPaymentsBankInfo     string
	DomesticPaymentsBankName          string
	DomesticPaymentsAccountNumber     string
	DomesticPaymentsSortCode          string
	InternationalPaymentsSwiftBic     string
	InternationalPaymentsBankName     string
	InternationalPaymentsBankAddress  string
	InternationalPaymentsInstructions string
	VatNumber                         string
	SendInvoicesFrom                  string
	SendInvoicesBcc                   string
	CanPayWithCard                    bool //Deprecated
	CanPayWithDirectDebitSEPA         bool //Deprecated
	CanPayWithDirectDebitACH          bool //Deprecated
	CanPayWithDirectDebitBacs         bool //Deprecated
	CanPayWithPigeon                  bool
	CanPayWithBankTransfer            bool
	Source                            DataSource
	SourceOfTruth                     DataSource
	AppSource                         string
}

type TenantBillingProfileEntities []TenantBillingProfileEntity
