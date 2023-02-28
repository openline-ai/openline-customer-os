// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: contact/contact.proto

package contactGrpcService

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

// ContactGrpcServiceClient is the client API for ContactGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactGrpcServiceClient interface {
	CreateContact(ctx context.Context, in *CreateContactGrpcRequest, opts ...grpc.CallOption) (*CreateContactGrpcResponse, error)
}

type contactGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactGrpcServiceClient(cc grpc.ClientConnInterface) ContactGrpcServiceClient {
	return &contactGrpcServiceClient{cc}
}

func (c *contactGrpcServiceClient) CreateContact(ctx context.Context, in *CreateContactGrpcRequest, opts ...grpc.CallOption) (*CreateContactGrpcResponse, error) {
	out := new(CreateContactGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService.contactGrpcService/CreateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactGrpcServiceServer is the server API for ContactGrpcService service.
// All implementations must embed UnimplementedContactGrpcServiceServer
// for forward compatibility
type ContactGrpcServiceServer interface {
	CreateContact(context.Context, *CreateContactGrpcRequest) (*CreateContactGrpcResponse, error)
	mustEmbedUnimplementedContactGrpcServiceServer()
}

// UnimplementedContactGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContactGrpcServiceServer struct {
}

func (UnimplementedContactGrpcServiceServer) CreateContact(context.Context, *CreateContactGrpcRequest) (*CreateContactGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (UnimplementedContactGrpcServiceServer) mustEmbedUnimplementedContactGrpcServiceServer() {}

// UnsafeContactGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactGrpcServiceServer will
// result in compilation errors.
type UnsafeContactGrpcServiceServer interface {
	mustEmbedUnimplementedContactGrpcServiceServer()
}

func RegisterContactGrpcServiceServer(s grpc.ServiceRegistrar, srv ContactGrpcServiceServer) {
	s.RegisterService(&ContactGrpcService_ServiceDesc, srv)
}

func _ContactGrpcService_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactGrpcServiceServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactGrpcService.contactGrpcService/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).CreateContact(ctx, req.(*CreateContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactGrpcService_ServiceDesc is the grpc.ServiceDesc for ContactGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contactGrpcService.contactGrpcService",
	HandlerType: (*ContactGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _ContactGrpcService_CreateContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact/contact.proto",
}
