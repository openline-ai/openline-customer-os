// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: interaction_event.proto

package interaction_event_grpc_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InteractionEventGrpcServiceClient is the client API for InteractionEventGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionEventGrpcServiceClient interface {
	RequestGenerateSummary(ctx context.Context, in *RequestGenerateSummaryGrpcRequest, opts ...grpc.CallOption) (*InteractionEventIdGrpcResponse, error)
}

type interactionEventGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionEventGrpcServiceClient(cc grpc.ClientConnInterface) InteractionEventGrpcServiceClient {
	return &interactionEventGrpcServiceClient{cc}
}

func (c *interactionEventGrpcServiceClient) RequestGenerateSummary(ctx context.Context, in *RequestGenerateSummaryGrpcRequest, opts ...grpc.CallOption) (*InteractionEventIdGrpcResponse, error) {
	out := new(InteractionEventIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/interactionEventGrpcService/RequestGenerateSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionEventGrpcServiceServer is the server API for InteractionEventGrpcService service.
// All implementations should embed UnimplementedInteractionEventGrpcServiceServer
// for forward compatibility
type InteractionEventGrpcServiceServer interface {
	RequestGenerateSummary(context.Context, *RequestGenerateSummaryGrpcRequest) (*InteractionEventIdGrpcResponse, error)
}

// UnimplementedInteractionEventGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInteractionEventGrpcServiceServer struct {
}

func (UnimplementedInteractionEventGrpcServiceServer) RequestGenerateSummary(context.Context, *RequestGenerateSummaryGrpcRequest) (*InteractionEventIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestGenerateSummary not implemented")
}

// UnsafeInteractionEventGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionEventGrpcServiceServer will
// result in compilation errors.
type UnsafeInteractionEventGrpcServiceServer interface {
	mustEmbedUnimplementedInteractionEventGrpcServiceServer()
}

func RegisterInteractionEventGrpcServiceServer(s grpc.ServiceRegistrar, srv InteractionEventGrpcServiceServer) {
	s.RegisterService(&InteractionEventGrpcService_ServiceDesc, srv)
}

func _InteractionEventGrpcService_RequestGenerateSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGenerateSummaryGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionEventGrpcServiceServer).RequestGenerateSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionEventGrpcService/RequestGenerateSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionEventGrpcServiceServer).RequestGenerateSummary(ctx, req.(*RequestGenerateSummaryGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionEventGrpcService_ServiceDesc is the grpc.ServiceDesc for InteractionEventGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionEventGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactionEventGrpcService",
	HandlerType: (*InteractionEventGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestGenerateSummary",
			Handler:    _InteractionEventGrpcService_RequestGenerateSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction_event.proto",
}
