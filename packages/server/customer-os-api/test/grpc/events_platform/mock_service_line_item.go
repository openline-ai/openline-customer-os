package events_platform

import (
	"context"
	servicelineitempb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/service_line_item"
)

type MockServiceLineItemServiceCallbacks struct {
	CreateServiceLineItem func(context.Context, *servicelineitempb.CreateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error)
	UpdateServiceLineItem func(context.Context, *servicelineitempb.UpdateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error)
	DeleteServiceLineItem func(context.Context, *servicelineitempb.DeleteServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error)
}

var serviceLineItemCallbacks = &MockServiceLineItemServiceCallbacks{}

func SetServiceLineItemCallbacks(callbacks *MockServiceLineItemServiceCallbacks) {
	serviceLineItemCallbacks = callbacks
}

type MockServiceLineItemService struct {
	servicelineitempb.UnimplementedServiceLineItemGrpcServiceServer
}

func (MockServiceLineItemService) CreateServiceLineItem(context context.Context, proto *servicelineitempb.CreateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	if serviceLineItemCallbacks.CreateServiceLineItem == nil {
		panic("serviceLineItemCallbacks.CreateServiceLineItem is not set")
	}
	return serviceLineItemCallbacks.CreateServiceLineItem(context, proto)
}

func (MockServiceLineItemService) UpdateServiceLineItem(context context.Context, proto *servicelineitempb.UpdateServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	if serviceLineItemCallbacks.UpdateServiceLineItem == nil {
		panic("serviceLineItemCallbacks.UpdateServiceLineItem is not set")
	}
	return serviceLineItemCallbacks.UpdateServiceLineItem(context, proto)
}

func (MockServiceLineItemService) DeleteServiceLineItem(context context.Context, proto *servicelineitempb.DeleteServiceLineItemGrpcRequest) (*servicelineitempb.ServiceLineItemIdGrpcResponse, error) {
	if serviceLineItemCallbacks.DeleteServiceLineItem == nil {
		panic("serviceLineItemCallbacks.DeleteServiceLineItem is not set")
	}
	return serviceLineItemCallbacks.DeleteServiceLineItem(context, proto)
}
