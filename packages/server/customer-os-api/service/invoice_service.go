package service

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

type InvoiceService interface {
	GetInvoices(ctx context.Context, organizationId string, page, limit int, filter *model.Filter, sortBy []*model.SortBy) (*utils.Pagination, error)
	GetById(ctx context.Context, invoiceId string) (*neo4jentity.InvoiceEntity, error)
	GetInvoiceLinesForInvoices(ctx context.Context, invoiceIds []string) (*neo4jentity.InvoiceLineEntities, error)
	SimulateInvoice(ctx context.Context, invoiceData *SimulateInvoiceData) (string, error)
	NextInvoiceDryRun(ctx context.Context, contractId string) (string, error)
	UpdateInvoice(ctx context.Context, input model.InvoiceUpdateInput) error
	PayInvoice(ctx context.Context, invoiceId string) error
	VoidInvoice(ctx context.Context, invoiceId string) error
}
type invoiceService struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
	services     *Services
}

func (s *invoiceService) UpdateInvoice(ctx context.Context, input model.InvoiceUpdateInput) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ContractService.UpdateInvoice")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	if input.ID == "" {
		err := fmt.Errorf("invoice id is missing")
		s.log.Error(err.Error())
		tracing.TraceErr(span, err)
		return err
	}

	invoiceExists, _ := s.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, common.GetTenantFromContext(ctx), input.ID, neo4jutil.NodeLabelInvoice)
	if !invoiceExists {
		err := fmt.Errorf("invoice with id {%s} not found", input.ID)
		s.log.Error(err.Error())
		tracing.TraceErr(span, err)
		return err
	}

	fieldMask := []invoicepb.InvoiceFieldMask{}
	invoiceUpdateRequest := invoicepb.UpdateInvoiceRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		InvoiceId:      input.ID,
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		AppSource:      constants.AppSourceCustomerOsApi,
	}
	// prepare invoice status
	if input.Status != nil {
		switch *input.Status {
		case model.InvoiceStatusDraft:
			invoiceUpdateRequest.Status = invoicepb.InvoiceStatus_INVOICE_STATUS_DRAFT
		case model.InvoiceStatusDue:
			invoiceUpdateRequest.Status = invoicepb.InvoiceStatus_INVOICE_STATUS_DUE
		case model.InvoiceStatusPaid:
			invoiceUpdateRequest.Status = invoicepb.InvoiceStatus_INVOICE_STATUS_PAID
		default:
			invoiceUpdateRequest.Status = invoicepb.InvoiceStatus_INVOICE_STATUS_NONE
		}
	}

	if input.Patch {
		if input.Status != nil {
			fieldMask = append(fieldMask, invoicepb.InvoiceFieldMask_INVOICE_FIELD_STATUS)
		}
		invoiceUpdateRequest.FieldsMask = fieldMask
		if len(fieldMask) == 0 {
			span.LogFields(log.String("result", "No fields to update"))
			return nil
		}
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err := CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
		return s.grpcClients.InvoiceClient.UpdateInvoice(ctx, &invoiceUpdateRequest)
	})
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return err
	}

	return nil
}

func NewInvoiceService(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients, services *Services) InvoiceService {
	return &invoiceService{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
		services:     services,
	}
}

type SimulateInvoiceData struct {
	ContractId   string
	Date         *time.Time
	InvoiceLines []SimulateInvoiceLineData
}
type SimulateInvoiceLineData struct {
	ServiceLineItemID *string
	Name              string
	Billed            enum.BilledType
	Price             float64
	Quantity          int
}

func (s *invoiceService) GetInvoices(ctx context.Context, organizationId string, page, limit int, filter *model.Filter, sortBy []*model.SortBy) (*utils.Pagination, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.GetInvoices")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("organizationId", organizationId))
	span.LogFields(log.Object("page", page))
	span.LogFields(log.Object("limit", limit))
	span.LogFields(log.Object("filter", filter))
	span.LogFields(log.Object("sortBy", sortBy))

	if len(sortBy) == 0 {
		sortBy = []*model.SortBy{
			{
				By:        "CREATED_AT",
				Direction: model.SortingDirectionDesc,
			},
		}
	}

	cypherSort, err := buildSort(sortBy, reflect.TypeOf(neo4jentity.InvoiceEntity{}))
	if err != nil {
		return nil, err
	}
	cypherFilter, err := buildFilter(filter, reflect.TypeOf(neo4jentity.InvoiceEntity{}))
	if err != nil {
		return nil, err
	}

	dbNodesWithTotalCount, err := s.repositories.Neo4jRepositories.InvoiceReadRepository.GetPaginatedInvoices(ctx, common.GetTenantFromContext(ctx), organizationId,
		page,
		limit,
		cypherFilter,
		cypherSort)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}

	var paginatedResult = utils.Pagination{
		Limit: page,
		Page:  limit,
	}

	paginatedResult.SetTotalRows(dbNodesWithTotalCount.Count)

	var invoices neo4jentity.InvoiceEntities

	for _, v := range dbNodesWithTotalCount.Nodes {
		invoices = append(invoices, *mapper.MapDbNodeToInvoiceEntity(v))
	}
	paginatedResult.SetRows(&invoices)
	return &paginatedResult, nil
}

func (s *invoiceService) GetById(ctx context.Context, invoiceId string) (*neo4jentity.InvoiceEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.GetById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("invoiceId", invoiceId))

	if invoiceDbNode, err := s.repositories.Neo4jRepositories.InvoiceReadRepository.GetInvoiceById(ctx, common.GetContext(ctx).Tenant, invoiceId); err != nil {
		tracing.TraceErr(span, err)
		wrappedErr := errors.Wrap(err, fmt.Sprintf("Invoice with id {%s} not found", invoiceId))
		return nil, wrappedErr
	} else {
		return mapper.MapDbNodeToInvoiceEntity(invoiceDbNode), nil
	}
}

func (s *invoiceService) GetInvoiceLinesForInvoices(ctx context.Context, invoiceIds []string) (*neo4jentity.InvoiceLineEntities, error) {
	invoiceLines, err := s.repositories.Neo4jRepositories.InvoiceLineReadRepository.GetAllForInvoices(ctx, common.GetTenantFromContext(ctx), invoiceIds)
	if err != nil {
		return nil, err
	}
	invoiceLineEntities := make(neo4jentity.InvoiceLineEntities, 0, len(invoiceLines))
	for _, v := range invoiceLines {
		invoiceLineEntity := mapper.MapDbNodeToInvoiceLineEntity(v.Node)
		invoiceLineEntity.DataloaderKey = v.LinkedNodeId
		invoiceLineEntities = append(invoiceLineEntities, *invoiceLineEntity)
	}
	return &invoiceLineEntities, nil
}

func (s *invoiceService) SimulateInvoice(ctx context.Context, invoiceData *SimulateInvoiceData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.SimulateInvoice")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("invoiceData", invoiceData))

	if invoiceData.InvoiceLines == nil {
		err := fmt.Errorf("no invoice lines to simulate")
		tracing.TraceErr(span, err)
		return "", err
	}

	now := time.Now()
	simulateInvoiceRequest := invoicepb.SimulateInvoiceRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		ContractId:     invoiceData.ContractId,
		CreatedAt:      utils.ConvertTimeToTimestampPtr(&now),
		SourceFields: &commonpb.SourceFields{
			Source:    neo4jentity.DataSourceOpenline.String(),
			AppSource: constants.AppSourceCustomerOsApi,
		},
		Date:                   utils.ConvertTimeToTimestampPtr(invoiceData.Date),
		DryRunServiceLineItems: make([]*invoicepb.DryRunServiceLineItem, 0, len(invoiceData.InvoiceLines)),
	}
	for _, invoiceLine := range invoiceData.InvoiceLines {
		dryRunServiceLineItem := invoicepb.DryRunServiceLineItem{
			ServiceLineItemId: utils.IfNotNilStringWithDefault(invoiceLine.ServiceLineItemID, ""),
			Name:              invoiceLine.Name,
			Price:             invoiceLine.Price,
			Quantity:          int64(invoiceLine.Quantity),
		}

		switch invoiceLine.Billed {
		case enum.BilledTypeMonthly:
			dryRunServiceLineItem.Billed = commonpb.BilledType_MONTHLY_BILLED
		case enum.BilledTypeQuarterly:
			dryRunServiceLineItem.Billed = commonpb.BilledType_QUARTERLY_BILLED
		case enum.BilledTypeAnnually:
			dryRunServiceLineItem.Billed = commonpb.BilledType_ANNUALLY_BILLED
		case enum.BilledTypeOnce:
			dryRunServiceLineItem.Billed = commonpb.BilledType_ONCE_BILLED
		case enum.BilledTypeUsage:
			dryRunServiceLineItem.Billed = commonpb.BilledType_USAGE_BILLED
		case enum.BilledTypeNone:
			dryRunServiceLineItem.Billed = commonpb.BilledType_NONE_BILLED
		}

		simulateInvoiceRequest.DryRunServiceLineItems = append(simulateInvoiceRequest.DryRunServiceLineItems, &dryRunServiceLineItem)
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	response, err := CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
		return s.grpcClients.InvoiceClient.SimulateInvoice(ctx, &simulateInvoiceRequest)
	})
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return "", err
	}

	WaitForNodeCreatedInNeo4j(ctx, s.repositories, response.Id, neo4jutil.NodeLabelInvoice, span)

	span.LogFields(log.String("output - createdInvoiceId", response.Id))
	return response.Id, nil
}

func (s *invoiceService) NextInvoiceDryRun(ctx context.Context, contractId string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.NextInvoiceDryRun")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("contractId", contractId))

	tenant := common.GetTenantFromContext(ctx)
	now := time.Now()

	contract, err := s.services.ContractService.GetById(ctx, contractId)
	if err != nil {
		tracing.TraceErr(span, err)
		return "", err
	}

	var invoicePeriodStart, invoicePeriodEnd time.Time
	if contract.NextInvoiceDate != nil {
		invoicePeriodStart = *contract.NextInvoiceDate
	} else if contract.InvoicingStartDate != nil {
		invoicePeriodStart = *contract.InvoicingStartDate
	} else {
		err = fmt.Errorf("contract has no next invoice date or invoicing start date")
		tracing.TraceErr(span, err)
		return "", err
	}
	invoicePeriodEnd = calculateInvoiceCycleEnd(invoicePeriodStart, contract.BillingCycle)

	tenantSettings, err := s.services.TenantService.GetTenantSettings(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		return "", err
	}

	currency := contract.Currency.String()
	if currency == "" {
		currency = tenantSettings.BaseCurrency.String()
	}

	dryRunInvoiceRequest := invoicepb.NewInvoiceForContractRequest{
		Tenant:             tenant,
		LoggedInUserId:     common.GetUserIdFromContext(ctx),
		ContractId:         contractId,
		DryRun:             true,
		CreatedAt:          utils.ConvertTimeToTimestampPtr(&now),
		InvoicePeriodStart: utils.ConvertTimeToTimestampPtr(&invoicePeriodStart),
		InvoicePeriodEnd:   utils.ConvertTimeToTimestampPtr(&invoicePeriodEnd),
		Currency:           currency,
		Note:               contract.InvoiceNote,
		Postpaid:           tenantSettings.InvoicingPostpaid,
		SourceFields: &commonpb.SourceFields{
			Source:    neo4jentity.DataSourceOpenline.String(),
			AppSource: constants.AppSourceCustomerOsApi,
		},
	}

	switch contract.BillingCycle {
	case enum.BillingCycleMonthlyBilling:
		dryRunInvoiceRequest.BillingCycle = commonpb.BillingCycle_MONTHLY_BILLING
	case enum.BillingCycleQuarterlyBilling:
		dryRunInvoiceRequest.BillingCycle = commonpb.BillingCycle_QUARTERLY_BILLING
	case enum.BillingCycleAnnuallyBilling:
		dryRunInvoiceRequest.BillingCycle = commonpb.BillingCycle_ANNUALLY_BILLING
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	response, err := CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
		return s.grpcClients.InvoiceClient.NewInvoiceForContract(ctx, &dryRunInvoiceRequest)
	})
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return "", err
	}

	WaitForNodeCreatedInNeo4j(ctx, s.repositories, response.Id, neo4jutil.NodeLabelInvoice, span)

	span.LogFields(log.String("output - createdInvoiceId", response.Id))
	return response.Id, nil
}

func (s *invoiceService) PayInvoice(ctx context.Context, invoiceId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.PayInvoice")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("invoiceId", invoiceId))

	tenant := common.GetTenantFromContext(ctx)

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	response, err := CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
		return s.grpcClients.InvoiceClient.PayInvoice(ctx, &invoicepb.PayInvoiceRequest{
			Tenant:         tenant,
			InvoiceId:      invoiceId,
			LoggedInUserId: common.GetUserIdFromContext(ctx),
			SourceFields: &commonpb.SourceFields{
				AppSource: constants.AppSourceCustomerOsApi,
			},
		})
	})

	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return err
	}

	span.LogFields(log.String("output - payInvoiceId", response.Id))
	return nil
}

func (s *invoiceService) VoidInvoice(ctx context.Context, invoiceId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InvoiceService.VoidInvoice")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("invoiceId", invoiceId))

	tenant := common.GetTenantFromContext(ctx)

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	response, err := CallEventsPlatformGRPCWithRetry[*invoicepb.InvoiceIdResponse](func() (*invoicepb.InvoiceIdResponse, error) {
		return s.grpcClients.InvoiceClient.VoidInvoice(ctx, &invoicepb.VoidInvoiceRequest{
			Tenant:         tenant,
			InvoiceId:      invoiceId,
			LoggedInUserId: common.GetUserIdFromContext(ctx),
			AppSource:      constants.AppSourceCustomerOsApi,
		})
	})

	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return err
	}

	span.LogFields(log.String("output - voidInvoiceId", response.Id))
	return nil
}

func calculateInvoiceCycleEnd(start time.Time, cycle enum.BillingCycle) time.Time {
	var end time.Time
	switch cycle {
	case enum.BillingCycleMonthlyBilling:
		end = start.AddDate(0, 1, 0)
	case enum.BillingCycleQuarterlyBilling:
		end = start.AddDate(0, 3, 0)
	case enum.BillingCycleAnnuallyBilling:
		end = start.AddDate(1, 0, 0)
	default:
		return start
	}
	previousDay := end.AddDate(0, 0, -1)
	return previousDay
}
