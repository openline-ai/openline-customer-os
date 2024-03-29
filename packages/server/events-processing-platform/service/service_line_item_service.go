package service

import (
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contract/aggregate"
	sliaggregate "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/command_handler"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	grpcerr "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_errors"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	servicelineitempb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/service_line_item"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serviceLineItemService struct {
	servicelineitempb.UnimplementedServiceLineItemGrpcServiceServer
	log                            logger.Logger
	serviceLineItemCommandHandlers *command_handler.CommandHandlers
	aggregateStore                 eventstore.AggregateStore
}

func NewServiceLineItemService(log logger.Logger, commandHandlers *command_handler.CommandHandlers, aggregateStore eventstore.AggregateStore) *serviceLineItemService {
	return &serviceLineItemService{
		log:                            log,
		serviceLineItemCommandHandlers: commandHandlers,
		aggregateStore:                 aggregateStore,
	}
}

func (s *serviceLineItemService) CreateServiceLineItem(ctx context.Context, request *servicelineitempb.CreateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ServiceLineItemService.CreateServiceLineItem")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.Object("request", request))

	// Validate contract ID
	if request.ContractId == "" {
		return nil, grpcerr.ErrResponse(grpcerr.ErrMissingField("contractId"))
	}
	// Check if the contract aggregate exists
	contractExists, err := s.checkContractExists(ctx, request.Tenant, request.ContractId)
	if err != nil {
		s.log.Error(err, "error checking contract existence")
		return nil, status.Errorf(codes.Internal, "error checking contract existence: %v", err)
	}
	if !contractExists {
		return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", request.ContractId)
	}

	serviceLineItemId := uuid.New().String()

	source := commonmodel.Source{}
	source.FromGrpc(request.SourceFields)

	createdAt, updatedAt := convertCreateAndUpdateProtoTimestampsToTime(request.CreatedAt, request.UpdatedAt)

	createCommand := command.NewCreateServiceLineItemCommand(
		serviceLineItemId,
		request.Tenant,
		request.LoggedInUserId,
		"",
		model.ServiceLineItemDataFields{
			Billed:     model.BilledType(request.Billed),
			Quantity:   request.Quantity,
			Price:      request.Price,
			Name:       request.Name,
			ContractId: request.ContractId,
			ParentId:   serviceLineItemId,
			VatRate:    request.VatRate,
		},
		source,
		createdAt,
		updatedAt,
	)
	createCommand.StartedAt = utils.TimestampProtoToTimePtr(request.StartedAt)
	createCommand.EndedAt = utils.TimestampProtoToTimePtr(request.EndedAt)

	if err = s.serviceLineItemCommandHandlers.CreateServiceLineItem.Handle(ctx, createCommand); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(CreateServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
		return nil, grpcerr.ErrResponse(err)
	}

	// Return the ID of the newly created service line item
	return &servicelineitempb.ServiceLineItemIdGrpcResponse{Id: serviceLineItemId}, nil
}

func (s *serviceLineItemService) UpdateServiceLineItem(ctx context.Context, request *servicelineitempb.UpdateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ServiceLineItemService.UpdateServiceLineItem")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.Object("request", request))

	// Validate service line item ID
	if request.Id == "" {
		return nil, grpcerr.ErrResponse(grpcerr.ErrMissingField("id"))
	}

	updatedAt := utils.TimestampProtoToTimePtr(request.UpdatedAt)

	source := commonmodel.Source{}
	source.FromGrpc(request.SourceFields)

	if request.IsRetroactiveCorrection {
		updateServiceLineItemCommand := command.NewUpdateServiceLineItemCommand(
			request.Id,
			request.Tenant,
			request.LoggedInUserId,
			model.ServiceLineItemDataFields{
				Billed:   model.BilledType(request.Billed),
				Quantity: request.Quantity,
				Price:    request.Price,
				Name:     request.Name,
				Comments: request.Comments,
				VatRate:  request.VatRate,
			},
			source,
			updatedAt,
		)
		if request.StartedAt != nil {
			updateServiceLineItemCommand.StartedAt = utils.TimestampProtoToTimePtr(request.StartedAt)
		}

		if err := s.serviceLineItemCommandHandlers.UpdateServiceLineItem.Handle(ctx, updateServiceLineItemCommand); err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("(UpdateServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
			return nil, grpcerr.ErrResponse(err)
		}
		// Return the ID of the updated service line item
		return &servicelineitempb.ServiceLineItemIdGrpcResponse{Id: request.Id}, nil

	} else {
		versionDate := utils.NowPtr()
		if request.StartedAt != nil {
			versionDate = utils.TimestampProtoToTimePtr(request.StartedAt)
		}

		// Validate contract ID
		if request.ContractId == "" {
			return nil, grpcerr.ErrResponse(grpcerr.ErrMissingField("contractId"))
		}

		// Check if the contract aggregate exists prior to closing the service line item
		contractExists, err := s.checkContractExists(ctx, request.Tenant, request.ContractId)
		if err != nil {
			s.log.Error(err, "error checking contract existence")
			return nil, status.Errorf(codes.Internal, "error checking contract existence: %v", err)
		}
		if !contractExists {
			return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", request.ContractId)
		}

		// Check if current SLI is not ended
		sliEnded, err := s.checkSLINotEnded(ctx, request.Tenant, request.Id)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Error(err, "error checking service line item end status")
			return nil, status.Errorf(codes.Internal, "error checking service line item end status: %v", err)
		}
		if sliEnded {
			return nil, status.Errorf(codes.FailedPrecondition, "service line item with ID %s is already ended", request.Id)
		}

		//Close service line item
		closeSliCommand := command.NewCloseServiceLineItemCommand(request.Id, request.Tenant, request.LoggedInUserId, source.AppSource, false,
			versionDate, utils.TimestampProtoToTimePtr(request.UpdatedAt))

		if err := s.serviceLineItemCommandHandlers.CloseServiceLineItem.Handle(ctx, closeSliCommand); err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("(CloseServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
			return nil, grpcerr.ErrResponse(err)
		}

		//Create new service line item
		serviceLineItemId := uuid.New().String()

		createCommand := command.NewCreateServiceLineItemCommand(
			serviceLineItemId,
			request.Tenant,
			request.LoggedInUserId,
			request.Id,
			model.ServiceLineItemDataFields{
				Billed:     model.BilledType(request.Billed),
				Quantity:   request.Quantity,
				Price:      float64(request.Price),
				Name:       request.Name,
				ContractId: request.ContractId,
				ParentId:   request.ParentId,
				Comments:   request.Comments,
			},
			source,
			utils.NowPtr(),
			updatedAt,
		)
		createCommand.StartedAt = versionDate

		if err = s.serviceLineItemCommandHandlers.CreateServiceLineItem.Handle(ctx, createCommand); err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("(CreateServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
			return nil, grpcerr.ErrResponse(err)
		}
		return &servicelineitempb.ServiceLineItemIdGrpcResponse{Id: serviceLineItemId}, nil
	}
}

func (s *serviceLineItemService) DeleteServiceLineItem(ctx context.Context, request *servicelineitempb.DeleteServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ServiceLineItemService.DeleteServiceLineItem")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	// Validate service line item ID
	if request.Id == "" {
		return nil, grpcerr.ErrResponse(grpcerr.ErrMissingField("id"))
	}

	deleteSliCommand := command.NewDeleteServiceLineItemCommand(request.Id, request.Tenant, request.LoggedInUserId, request.AppSource)

	if err := s.serviceLineItemCommandHandlers.DeleteServiceLineItem.Handle(ctx, deleteSliCommand); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(DeleteServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
		return nil, grpcerr.ErrResponse(err)
	}

	// Return the ID of the updated service line item
	return &servicelineitempb.ServiceLineItemIdGrpcResponse{Id: request.Id}, nil
}

func (s *serviceLineItemService) CloseServiceLineItem(ctx context.Context, request *servicelineitempb.CloseServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ServiceLineItemService.CloseServiceLineItem")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.Object("request", request))

	// Validate service line item ID
	if request.Id == "" {
		return nil, grpcerr.ErrResponse(grpcerr.ErrMissingField("id"))
	}

	closeSliCommand := command.NewCloseServiceLineItemCommand(request.Id, request.Tenant, request.LoggedInUserId, request.AppSource, true,
		utils.TimestampProtoToTimePtr(request.EndedAt), utils.TimestampProtoToTimePtr(request.UpdatedAt))

	if err := s.serviceLineItemCommandHandlers.CloseServiceLineItem.Handle(ctx, closeSliCommand); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(CloseServiceLineItem.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
		return nil, grpcerr.ErrResponse(err)
	}

	// Return the ID of the updated service line item
	return &servicelineitempb.ServiceLineItemIdGrpcResponse{Id: request.Id}, nil
}

func (s *serviceLineItemService) checkContractExists(ctx context.Context, tenant, contractId string) (bool, error) {
	contractAggregate := aggregate.NewContractAggregateWithTenantAndID(tenant, contractId)
	err := s.aggregateStore.Exists(ctx, contractAggregate.GetID())
	if err != nil {
		if errors.Is(err, eventstore.ErrAggregateNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil // The contract exists
}

func (s *serviceLineItemService) checkSLINotEnded(ctx context.Context, tenant, id string) (bool, error) {
	sliAggregate := sliaggregate.NewServiceLineItemAggregateWithTenantAndID(tenant, id)
	err := s.aggregateStore.Exists(ctx, sliAggregate.GetID())
	if err != nil {
		if errors.Is(err, eventstore.ErrAggregateNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	return sliAggregate.ServiceLineItem.EndedAt != nil, nil
}
