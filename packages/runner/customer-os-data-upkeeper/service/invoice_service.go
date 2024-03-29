package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/constants"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/events_processing_client"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/logger"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/repository"
	"github.com/openline-ai/openline-customer-os/packages/runner/customer-os-data-upkeeper/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/data"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	neo4jmapper "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	contractpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/contract"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type GeneratePaymentLinkEventBody struct {
	Tenant                       string `json:"tenant"`
	Currency                     string `json:"currency"`
	AmountInSmallestCurrencyUnit int64  `json:"amountInSmallestCurrencyUnit"`
	InvoiceId                    string `json:"invoiceId"`
	InvoiceDescription           string `json:"invoiceDescription"`
}

type InvoiceService interface {
	GenerateCycleInvoices()
	GenerateOffCycleInvoices()
	SendPayNotifications()
	GenerateInvoicePaymentLinks()
}

type invoiceService struct {
	cfg                    *config.Config
	log                    logger.Logger
	repositories           *repository.Repositories
	eventsProcessingClient *events_processing_client.Client
}

func NewInvoiceService(cfg *config.Config, log logger.Logger, repositories *repository.Repositories, client *events_processing_client.Client) InvoiceService {
	return &invoiceService{
		cfg:                    cfg,
		log:                    log,
		repositories:           repositories,
		eventsProcessingClient: client,
	}
}

func (s *invoiceService) GenerateCycleInvoices() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel context on exit

	span, ctx := tracing.StartTracerSpan(ctx, "InvoiceService.GenerateCycleInvoices")
	defer span.Finish()

	if s.eventsProcessingClient == nil {
		err := errors.New("eventsProcessingClient is nil")
		tracing.TraceErr(span, err)
		s.log.Error(err.Error())
		return
	}

	referenceTime := utils.Now()
	dryRun := false
	cachedTenantBaseCurrencies := make(map[string]neo4jenum.Currency)
	cachedTenantPostpaidFlags := make(map[string]bool)

	for {
		select {
		case <-ctx.Done():
			s.log.Infof("Context cancelled, stopping")
			return
		default:
			// continue as normal
		}

		records, err := s.repositories.Neo4jRepositories.ContractReadRepository.GetContractsToGenerateCycleInvoices(ctx, referenceTime)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("Error getting contracts for invoicing: %v", err)
			return
		}

		// no contracts found
		if len(records) == 0 {
			return
		}

		//process records
		for _, record := range records {
			contract := neo4jmapper.MapDbNodeToContractEntity(record.Node)
			tenant := record.Tenant

			currency := contract.Currency.String()
			if currency == "" {
				currency = s.getTenantBaseCurrency(ctx, tenant, cachedTenantBaseCurrencies).String()
			}

			isPostpaid := s.getTenantInvoicingPostpaidFlag(ctx, tenant, cachedTenantPostpaidFlags)

			var invoicePeriodStart, invoicePeriodEnd time.Time
			if contract.NextInvoiceDate != nil {
				invoicePeriodStart = *contract.NextInvoiceDate
			} else {
				invoicePeriodStart = *contract.InvoicingStartDate
			}
			invoicePeriodEnd = calculateInvoiceCycleEnd(invoicePeriodStart, contract.BillingCycle)

			readyToRequestInvoice := false
			if isPostpaid {
				readyToRequestInvoice = utils.EndOfDayInUTC(invoicePeriodEnd).Before(referenceTime)
			} else {
				readyToRequestInvoice = invoicePeriodEnd.After(invoicePeriodStart)
			}
			if readyToRequestInvoice {
				newInvoiceRequest := invoicepb.NewInvoiceForContractRequest{
					Tenant:             record.Tenant,
					ContractId:         contract.Id,
					Currency:           currency,
					InvoicePeriodStart: utils.ConvertTimeToTimestampPtr(&invoicePeriodStart),
					InvoicePeriodEnd:   utils.ConvertTimeToTimestampPtr(&invoicePeriodEnd),
					DryRun:             dryRun,
					Note:               contract.InvoiceNote,
					Postpaid:           isPostpaid,
					SourceFields: &commonpb.SourceFields{
						AppSource: constants.AppSourceDataUpkeeper,
						Source:    neo4jentity.DataSourceOpenline.String(),
					},
				}
				switch contract.BillingCycle {
				case neo4jenum.BillingCycleMonthlyBilling:
					newInvoiceRequest.BillingCycle = commonpb.BillingCycle_MONTHLY_BILLING
				case neo4jenum.BillingCycleQuarterlyBilling:
					newInvoiceRequest.BillingCycle = commonpb.BillingCycle_QUARTERLY_BILLING
				case neo4jenum.BillingCycleAnnuallyBilling:
					newInvoiceRequest.BillingCycle = commonpb.BillingCycle_ANNUALLY_BILLING
				}
				_, err = CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
					return s.eventsProcessingClient.InvoiceClient.NewInvoiceForContract(ctx, &newInvoiceRequest)
				})
				if err != nil {
					tracing.TraceErr(span, err)
					s.log.Errorf("Error generating invoice for contract %s: %s", contract.Id, err.Error())
				}

				if !dryRun && err == nil {
					nextInvoiceDate := utils.ToPtr(invoicePeriodEnd.AddDate(0, 0, 1))
					_, err = CallEventsPlatformGRPCWithRetry[*contractpb.ContractIdGrpcResponse](func() (*contractpb.ContractIdGrpcResponse, error) {
						return s.eventsProcessingClient.ContractClient.UpdateContract(ctx, &contractpb.UpdateContractGrpcRequest{
							Tenant: tenant,
							Id:     contract.Id,
							SourceFields: &commonpb.SourceFields{
								AppSource: constants.AppSourceDataUpkeeper,
							},
							NextInvoiceDate: utils.ConvertTimeToTimestampPtr(nextInvoiceDate),
							InvoiceNote:     "",
							FieldsMask: []contractpb.ContractFieldMask{
								contractpb.ContractFieldMask_CONTRACT_FIELD_INVOICE_NOTE,
								contractpb.ContractFieldMask_CONTRACT_FIELD_NEXT_INVOICE_DATE},
						})
					})
					if err != nil {
						tracing.TraceErr(span, err)
						s.log.Errorf("Error updating contract %s: %s", contract.Id, err.Error())
					}
				}
			}
			// mark invoicing started
			err = s.repositories.Neo4jRepositories.ContractWriteRepository.MarkCycleInvoicingRequested(ctx, tenant, contract.Id, utils.Now())
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error marking invoicing started for contract %s: %s", contract.Id, err.Error())
				return
			}
		}
		//sleep for async processing, then check again
		time.Sleep(10 * time.Second)
	}
}

func calculateInvoiceCycleEnd(start time.Time, cycle neo4jenum.BillingCycle) time.Time {
	var end time.Time
	switch cycle {
	case neo4jenum.BillingCycleMonthlyBilling:
		end = start.AddDate(0, 1, 0)
	case neo4jenum.BillingCycleQuarterlyBilling:
		end = start.AddDate(0, 3, 0)
	case neo4jenum.BillingCycleAnnuallyBilling:
		end = start.AddDate(1, 0, 0)
	default:
		return start
	}
	previousDay := end.AddDate(0, 0, -1)
	return previousDay
}

func (s *invoiceService) getTenantBaseCurrency(ctx context.Context, tenant string, cachedTenantBaseCurrencies map[string]neo4jenum.Currency) neo4jenum.Currency {
	if currency, ok := cachedTenantBaseCurrencies[tenant]; ok {
		return currency
	}

	dbNode, _ := s.repositories.Neo4jRepositories.TenantReadRepository.GetTenantSettings(ctx, tenant)
	tenantSettings := neo4jmapper.MapDbNodeToTenantSettingsEntity(dbNode)

	currency := tenantSettings.BaseCurrency
	cachedTenantBaseCurrencies[tenant] = currency
	return currency
}

func (s *invoiceService) getTenantInvoicingPostpaidFlag(ctx context.Context, tenant string, cachedTenantPostpaidFlags map[string]bool) bool {
	if postpaid, ok := cachedTenantPostpaidFlags[tenant]; ok {
		return postpaid
	}

	dbNode, _ := s.repositories.Neo4jRepositories.TenantReadRepository.GetTenantSettings(ctx, tenant)
	tenantSettings := neo4jmapper.MapDbNodeToTenantSettingsEntity(dbNode)

	cachedTenantPostpaidFlags[tenant] = tenantSettings.InvoicingPostpaid
	return tenantSettings.InvoicingPostpaid
}

func (s *invoiceService) SendPayNotifications() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel context on exit

	span, ctx := tracing.StartTracerSpan(ctx, "InvoiceService.SendPayNotifications")
	defer span.Finish()

	if s.eventsProcessingClient == nil {
		err := errors.New("eventsProcessingClient is nil")
		tracing.TraceErr(span, err)
		s.log.Error(err.Error())
		return
	}

	referenceTime := utils.Now()

	for {
		select {
		case <-ctx.Done():
			s.log.Infof("Context cancelled, stopping")
			return
		default:
			// continue as normal
		}

		records, err := s.repositories.Neo4jRepositories.InvoiceReadRepository.GetInvoicesForPayNotifications(
			ctx, s.cfg.ProcessConfig.DelaySendPayInvoiceNotificationInMinutes, s.cfg.ProcessConfig.RetrySendPayInvoiceNotificationDays, referenceTime)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("Error getting invoices for pay notifications: %v", err)
			return
		}

		// no invoices found
		if len(records) == 0 {
			return
		}

		//process records
		for _, record := range records {
			invoice := neo4jmapper.MapDbNodeToInvoiceEntity(record.Node)
			tenant := record.Tenant

			grpcRequest := invoicepb.PayInvoiceNotificationRequest{
				Tenant:    record.Tenant,
				AppSource: constants.AppSourceDataUpkeeper,
				InvoiceId: invoice.Id,
			}
			_, err = CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
				return s.eventsProcessingClient.InvoiceClient.PayInvoiceNotification(ctx, &grpcRequest)
			})
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error sending pay notification for invoice %s: %s", invoice.Id, err.Error())
			}

			// mark invoicing started
			err = s.repositories.Neo4jRepositories.InvoiceWriteRepository.MarkPayNotificationRequested(ctx, tenant, invoice.Id, utils.Now())
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error marking pay notification requested for invoice %s: %s", invoice.Id, err.Error())
			}
		}
		//sleep for async processing, then check again
		time.Sleep(5 * time.Second)
	}
}

func (s *invoiceService) GenerateOffCycleInvoices() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel context on exit

	span, ctx := tracing.StartTracerSpan(ctx, "InvoiceService.GenerateOffCycleInvoices")
	defer span.Finish()

	if s.eventsProcessingClient == nil {
		err := errors.New("eventsProcessingClient is nil")
		tracing.TraceErr(span, err)
		s.log.Error(err.Error())
		return
	}

	referenceTime := utils.Now()
	dryRun := false
	cachedTenantBaseCurrencies := make(map[string]neo4jenum.Currency)

	for {
		select {
		case <-ctx.Done():
			s.log.Infof("Context cancelled, stopping")
			return
		default:
			// continue as normal
		}

		records, err := s.repositories.Neo4jRepositories.ContractReadRepository.GetContractsToGenerateOffCycleInvoices(ctx, referenceTime)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("Error getting contracts for off-cycle invoicing: %v", err)
			return
		}

		// no contracts found
		if len(records) == 0 {
			return
		}

		//process records
		for _, record := range records {
			contract := neo4jmapper.MapDbNodeToContractEntity(record.Node)
			tenant := record.Tenant

			currency := contract.Currency.String()
			if currency == "" {
				currency = s.getTenantBaseCurrency(ctx, tenant, cachedTenantBaseCurrencies).String()
			}

			invoicePeriodStart := utils.StartOfDayInUTC(referenceTime)
			invoicePeriodEnd := utils.StartOfDayInUTC(utils.IfNotNilTimeWithDefault(contract.NextInvoiceDate, referenceTime).AddDate(0, 0, -1))

			readyToRequestInvoice := !invoicePeriodEnd.Before(invoicePeriodStart)
			if readyToRequestInvoice {
				newInvoiceRequest := invoicepb.NewInvoiceForContractRequest{
					Tenant:             record.Tenant,
					ContractId:         contract.Id,
					Currency:           currency,
					InvoicePeriodStart: utils.ConvertTimeToTimestampPtr(&invoicePeriodStart),
					InvoicePeriodEnd:   utils.ConvertTimeToTimestampPtr(&invoicePeriodEnd),
					DryRun:             dryRun,
					OffCycle:           true,
					SourceFields: &commonpb.SourceFields{
						AppSource: constants.AppSourceDataUpkeeper,
						Source:    neo4jentity.DataSourceOpenline.String(),
					},
				}
				_, err = CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
					return s.eventsProcessingClient.InvoiceClient.NewInvoiceForContract(ctx, &newInvoiceRequest)
				})

				if err != nil {
					tracing.TraceErr(span, err)
					s.log.Errorf("Error generating off-cycle invoice for contract %s: %s", contract.Id, err.Error())
				}
			}
			// mark invoicing started
			err = s.repositories.Neo4jRepositories.ContractWriteRepository.MarkCycleInvoicingRequested(ctx, tenant, contract.Id, utils.Now())
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error marking invoicing started for contract %s: %s", contract.Id, err.Error())
			}
		}
		//sleep for async processing, then check again
		time.Sleep(10 * time.Second)
	}
}

func (s *invoiceService) GenerateInvoicePaymentLinks() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel context on exit

	span, ctx := tracing.StartTracerSpan(ctx, "InvoiceService.GenerateInvoicePaymentLinks")
	defer span.Finish()

	if s.eventsProcessingClient == nil {
		err := errors.New("eventsProcessingClient is nil")
		tracing.TraceErr(span, err)
		s.log.Error(err.Error())
		return
	}

	if s.cfg.EventNotifications.EndPoints.GeneratePaymentLinkUrl == "" {
		err := errors.New("GeneratePaymentLinkUrl is not configured")
		tracing.TraceErr(span, err)
		s.log.Error(err.Error())
		return
	}

	referenceTime := utils.Now()

	for {
		select {
		case <-ctx.Done():
			s.log.Infof("Context cancelled, stopping")
			return
		default:
			// continue as normal
		}

		records, err := s.repositories.Neo4jRepositories.InvoiceReadRepository.GetInvoicesForPaymentLinkRequest(
			ctx, s.cfg.ProcessConfig.DelayRequestPaymentLinkInMinutes, s.cfg.ProcessConfig.RequestPaymentLinkLookBackWindowInDays, referenceTime)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("Error getting invoices for payment links generation: %v", err)
			return
		}

		// no invoices found
		if len(records) == 0 {
			return
		}

		//process records
		for _, record := range records {
			invoice := neo4jmapper.MapDbNodeToInvoiceEntity(record.Node)
			tenant := record.Tenant

			// convert amount to the smallest currency unit
			amountInSmallestCurrencyUnit, err := data.InSmallestCurrencyUnit(invoice.Currency.String(), invoice.TotalAmount)
			if err != nil {
				tracing.TraceErr(span, err)
			}

			// mark payment link request first, before sending the event
			err = s.repositories.Neo4jRepositories.InvoiceWriteRepository.MarkPaymentLinkRequested(ctx, tenant, invoice.Id, utils.Now())
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error marking payment link requested for invoice %s: %s", invoice.Id, err.Error())
			}

			requestBody := GeneratePaymentLinkEventBody{
				Tenant:                       tenant,
				Currency:                     invoice.Currency.String(),
				AmountInSmallestCurrencyUnit: amountInSmallestCurrencyUnit,
				InvoiceId:                    invoice.Id,
				InvoiceDescription:           fmt.Sprintf("Invoice %s", invoice.Number),
			}

			// Convert the request body to JSON
			requestBodyJSON, err := json.Marshal(requestBody)
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("error encoding JSON: %v", err)
				continue
			}

			// Create an HTTP client
			client := &http.Client{}

			// Create a POST request with headers and body
			req, err := http.NewRequest("POST", s.cfg.EventNotifications.EndPoints.GeneratePaymentLinkUrl, bytes.NewBuffer(requestBodyJSON))
			if err != nil {
				tracing.TraceErr(span, err)
				continue
			}

			// Set the content type header
			req.Header.Set("Content-Type", "application/json")

			// Send the POST request
			resp, err := client.Do(req)
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("error sending request: %v", err)
				continue
			}
			defer resp.Body.Close()

			// Check the response status code
			if resp.StatusCode != http.StatusOK {
				tracing.TraceErr(span, fmt.Errorf("request failed with status code: %s", resp.Status))
				s.log.Errorf("request failed with status code: %s", resp.Status)
			}
		}
	}
}
