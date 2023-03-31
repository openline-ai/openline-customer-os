// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: phone_number.proto

package phoneNumberGrpcService

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

// PhoneNumberGrpcServiceClient is the client API for PhoneNumberGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhoneNumberGrpcServiceClient interface {
	CreatePhoneNumber(ctx context.Context, in *CreatePhoneNumberGrpcRequest, opts ...grpc.CallOption) (*CreatePhoneNumberGrpcResponse, error)
}

type phoneNumberGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhoneNumberGrpcServiceClient(cc grpc.ClientConnInterface) PhoneNumberGrpcServiceClient {
	return &phoneNumberGrpcServiceClient{cc}
}

func (c *phoneNumberGrpcServiceClient) CreatePhoneNumber(ctx context.Context, in *CreatePhoneNumberGrpcRequest, opts ...grpc.CallOption) (*CreatePhoneNumberGrpcResponse, error) {
	out := new(CreatePhoneNumberGrpcResponse)
	err := c.cc.Invoke(ctx, "/phoneNumberGrpcService.phoneNumberGrpcService/CreatePhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneNumberGrpcServiceServer is the server API for PhoneNumberGrpcService service.
// All implementations must embed UnimplementedPhoneNumberGrpcServiceServer
// for forward compatibility
type PhoneNumberGrpcServiceServer interface {
	CreatePhoneNumber(context.Context, *CreatePhoneNumberGrpcRequest) (*CreatePhoneNumberGrpcResponse, error)
	mustEmbedUnimplementedPhoneNumberGrpcServiceServer()
}

// UnimplementedPhoneNumberGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPhoneNumberGrpcServiceServer struct {
}

func (UnimplementedPhoneNumberGrpcServiceServer) CreatePhoneNumber(context.Context, *CreatePhoneNumberGrpcRequest) (*CreatePhoneNumberGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePhoneNumber not implemented")
}
func (UnimplementedPhoneNumberGrpcServiceServer) mustEmbedUnimplementedPhoneNumberGrpcServiceServer() {
}

// UnsafePhoneNumberGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhoneNumberGrpcServiceServer will
// result in compilation errors.
type UnsafePhoneNumberGrpcServiceServer interface {
	mustEmbedUnimplementedPhoneNumberGrpcServiceServer()
}

func RegisterPhoneNumberGrpcServiceServer(s grpc.ServiceRegistrar, srv PhoneNumberGrpcServiceServer) {
	s.RegisterService(&PhoneNumberGrpcService_ServiceDesc, srv)
}

func _PhoneNumberGrpcService_CreatePhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhoneNumberGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneNumberGrpcServiceServer).CreatePhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phoneNumberGrpcService.phoneNumberGrpcService/CreatePhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneNumberGrpcServiceServer).CreatePhoneNumber(ctx, req.(*CreatePhoneNumberGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhoneNumberGrpcService_ServiceDesc is the grpc.ServiceDesc for PhoneNumberGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhoneNumberGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phoneNumberGrpcService.phoneNumberGrpcService",
	HandlerType: (*PhoneNumberGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePhoneNumber",
			Handler:    _PhoneNumberGrpcService_CreatePhoneNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phone_number.proto",
}
