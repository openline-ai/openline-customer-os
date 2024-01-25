// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tenant.proto

package tenant_grpc_service

import (
	context "context"
	common "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TenantGrpcServiceClient is the client API for TenantGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantGrpcServiceClient interface {
	AddBillingProfile(ctx context.Context, in *AddBillingProfileRequest, opts ...grpc.CallOption) (*common.IdResponse, error)
	UpdateBillingProfile(ctx context.Context, in *UpdateBillingProfileRequest, opts ...grpc.CallOption) (*common.IdResponse, error)
}

type tenantGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantGrpcServiceClient(cc grpc.ClientConnInterface) TenantGrpcServiceClient {
	return &tenantGrpcServiceClient{cc}
}

func (c *tenantGrpcServiceClient) AddBillingProfile(ctx context.Context, in *AddBillingProfileRequest, opts ...grpc.CallOption) (*common.IdResponse, error) {
	out := new(common.IdResponse)
	err := c.cc.Invoke(ctx, "/tenantGrpcService/AddBillingProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantGrpcServiceClient) UpdateBillingProfile(ctx context.Context, in *UpdateBillingProfileRequest, opts ...grpc.CallOption) (*common.IdResponse, error) {
	out := new(common.IdResponse)
	err := c.cc.Invoke(ctx, "/tenantGrpcService/UpdateBillingProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantGrpcServiceServer is the server API for TenantGrpcService service.
// All implementations should embed UnimplementedTenantGrpcServiceServer
// for forward compatibility
type TenantGrpcServiceServer interface {
	AddBillingProfile(context.Context, *AddBillingProfileRequest) (*common.IdResponse, error)
	UpdateBillingProfile(context.Context, *UpdateBillingProfileRequest) (*common.IdResponse, error)
}

// UnimplementedTenantGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantGrpcServiceServer struct {
}

func (UnimplementedTenantGrpcServiceServer) AddBillingProfile(context.Context, *AddBillingProfileRequest) (*common.IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBillingProfile not implemented")
}
func (UnimplementedTenantGrpcServiceServer) UpdateBillingProfile(context.Context, *UpdateBillingProfileRequest) (*common.IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBillingProfile not implemented")
}

// UnsafeTenantGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantGrpcServiceServer will
// result in compilation errors.
type UnsafeTenantGrpcServiceServer interface {
	mustEmbedUnimplementedTenantGrpcServiceServer()
}

func RegisterTenantGrpcServiceServer(s grpc.ServiceRegistrar, srv TenantGrpcServiceServer) {
	s.RegisterService(&TenantGrpcService_ServiceDesc, srv)
}

func _TenantGrpcService_AddBillingProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBillingProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantGrpcServiceServer).AddBillingProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenantGrpcService/AddBillingProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantGrpcServiceServer).AddBillingProfile(ctx, req.(*AddBillingProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantGrpcService_UpdateBillingProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBillingProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantGrpcServiceServer).UpdateBillingProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenantGrpcService/UpdateBillingProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantGrpcServiceServer).UpdateBillingProfile(ctx, req.(*UpdateBillingProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantGrpcService_ServiceDesc is the grpc.ServiceDesc for TenantGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenantGrpcService",
	HandlerType: (*TenantGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBillingProfile",
			Handler:    _TenantGrpcService_AddBillingProfile_Handler,
		},
		{
			MethodName: "UpdateBillingProfile",
			Handler:    _TenantGrpcService_UpdateBillingProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}
