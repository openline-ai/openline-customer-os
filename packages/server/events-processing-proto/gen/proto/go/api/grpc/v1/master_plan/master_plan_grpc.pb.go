// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: master_plan.proto

package master_plan_grpc_service

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

// MasterPlanGrpcServiceClient is the client API for MasterPlanGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterPlanGrpcServiceClient interface {
	CreateMasterPlan(ctx context.Context, in *CreateMasterPlanGrpcRequest, opts ...grpc.CallOption) (*MasterPlanIdGrpcResponse, error)
}

type masterPlanGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterPlanGrpcServiceClient(cc grpc.ClientConnInterface) MasterPlanGrpcServiceClient {
	return &masterPlanGrpcServiceClient{cc}
}

func (c *masterPlanGrpcServiceClient) CreateMasterPlan(ctx context.Context, in *CreateMasterPlanGrpcRequest, opts ...grpc.CallOption) (*MasterPlanIdGrpcResponse, error) {
	out := new(MasterPlanIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/MasterPlanGrpcService/CreateMasterPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterPlanGrpcServiceServer is the server API for MasterPlanGrpcService service.
// All implementations should embed UnimplementedMasterPlanGrpcServiceServer
// for forward compatibility
type MasterPlanGrpcServiceServer interface {
	CreateMasterPlan(context.Context, *CreateMasterPlanGrpcRequest) (*MasterPlanIdGrpcResponse, error)
}

// UnimplementedMasterPlanGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMasterPlanGrpcServiceServer struct {
}

func (UnimplementedMasterPlanGrpcServiceServer) CreateMasterPlan(context.Context, *CreateMasterPlanGrpcRequest) (*MasterPlanIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMasterPlan not implemented")
}

// UnsafeMasterPlanGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterPlanGrpcServiceServer will
// result in compilation errors.
type UnsafeMasterPlanGrpcServiceServer interface {
	mustEmbedUnimplementedMasterPlanGrpcServiceServer()
}

func RegisterMasterPlanGrpcServiceServer(s grpc.ServiceRegistrar, srv MasterPlanGrpcServiceServer) {
	s.RegisterService(&MasterPlanGrpcService_ServiceDesc, srv)
}

func _MasterPlanGrpcService_CreateMasterPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMasterPlanGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterPlanGrpcServiceServer).CreateMasterPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MasterPlanGrpcService/CreateMasterPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterPlanGrpcServiceServer).CreateMasterPlan(ctx, req.(*CreateMasterPlanGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterPlanGrpcService_ServiceDesc is the grpc.ServiceDesc for MasterPlanGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterPlanGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MasterPlanGrpcService",
	HandlerType: (*MasterPlanGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMasterPlan",
			Handler:    _MasterPlanGrpcService_CreateMasterPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_plan.proto",
}