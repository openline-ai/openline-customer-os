package service

import (
	"context"
	"github.com/google/uuid"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	invoiceEvents "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/invoice"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	grpcerr "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_errors"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	invoicepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/invoice"
)

type invoiceService struct {
	invoicepb.UnimplementedInvoiceServiceServer
	log            logger.Logger
	eventHandlers  *invoiceEvents.EventHandlers
	aggregateStore eventstore.AggregateStore
}

func NewInvoiceService(log logger.Logger, eventHandlers *invoiceEvents.EventHandlers, aggregateStore eventstore.AggregateStore) *invoiceService {
	return &invoiceService{
		log:            log,
		eventHandlers:  eventHandlers,
		aggregateStore: aggregateStore,
	}
}

func (s *invoiceService) NewInvoice(ctx context.Context, request *invoicepb.NewInvoiceRequest) (*invoicepb.InvoiceIdResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "InvoiceService.NewInvoice")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	invoiceId := uuid.New().String()

	baseRequest := eventstore.NewBaseRequest(invoiceId, request.Tenant, request.LoggedInUserId, commonmodel.SourceFromGrpc(request.SourceFields))

	if err := s.eventHandlers.InvoiceNew.Handle(ctx, baseRequest, request); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(InvoiceService.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
		return nil, grpcerr.ErrResponse(err)
	}

	return &invoicepb.InvoiceIdResponse{Id: invoiceId}, nil
}
