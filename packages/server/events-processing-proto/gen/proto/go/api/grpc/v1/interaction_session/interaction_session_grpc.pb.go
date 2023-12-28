// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: interaction_session.proto

package interaction_session_grpc_service

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

// InteractionSessionGrpcServiceClient is the client API for InteractionSessionGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionSessionGrpcServiceClient interface {
	UpsertInteractionSession(ctx context.Context, in *UpsertInteractionSessionGrpcRequest, opts ...grpc.CallOption) (*InteractionSessionIdGrpcResponse, error)
}

type interactionSessionGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionSessionGrpcServiceClient(cc grpc.ClientConnInterface) InteractionSessionGrpcServiceClient {
	return &interactionSessionGrpcServiceClient{cc}
}

func (c *interactionSessionGrpcServiceClient) UpsertInteractionSession(ctx context.Context, in *UpsertInteractionSessionGrpcRequest, opts ...grpc.CallOption) (*InteractionSessionIdGrpcResponse, error) {
	out := new(InteractionSessionIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/interactionSessionGrpcService/UpsertInteractionSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionSessionGrpcServiceServer is the server API for InteractionSessionGrpcService service.
// All implementations should embed UnimplementedInteractionSessionGrpcServiceServer
// for forward compatibility
type InteractionSessionGrpcServiceServer interface {
	UpsertInteractionSession(context.Context, *UpsertInteractionSessionGrpcRequest) (*InteractionSessionIdGrpcResponse, error)
}

// UnimplementedInteractionSessionGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInteractionSessionGrpcServiceServer struct {
}

func (UnimplementedInteractionSessionGrpcServiceServer) UpsertInteractionSession(context.Context, *UpsertInteractionSessionGrpcRequest) (*InteractionSessionIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertInteractionSession not implemented")
}

// UnsafeInteractionSessionGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionSessionGrpcServiceServer will
// result in compilation errors.
type UnsafeInteractionSessionGrpcServiceServer interface {
	mustEmbedUnimplementedInteractionSessionGrpcServiceServer()
}

func RegisterInteractionSessionGrpcServiceServer(s grpc.ServiceRegistrar, srv InteractionSessionGrpcServiceServer) {
	s.RegisterService(&InteractionSessionGrpcService_ServiceDesc, srv)
}

func _InteractionSessionGrpcService_UpsertInteractionSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertInteractionSessionGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionSessionGrpcServiceServer).UpsertInteractionSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionSessionGrpcService/UpsertInteractionSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionSessionGrpcServiceServer).UpsertInteractionSession(ctx, req.(*UpsertInteractionSessionGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionSessionGrpcService_ServiceDesc is the grpc.ServiceDesc for InteractionSessionGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionSessionGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactionSessionGrpcService",
	HandlerType: (*InteractionSessionGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertInteractionSession",
			Handler:    _InteractionSessionGrpcService_UpsertInteractionSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction_session.proto",
}