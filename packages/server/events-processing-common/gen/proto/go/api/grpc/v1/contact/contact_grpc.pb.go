// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/contact.proto

package contact_grpc_service

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
	UpsertContact(ctx context.Context, in *UpsertContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error)
	LinkPhoneNumberToContact(ctx context.Context, in *LinkPhoneNumberToContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error)
	LinkEmailToContact(ctx context.Context, in *LinkEmailToContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error)
	UnlinkLocationFromContact(ctx context.Context, in *UnlinkLocationFromContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error)
}

type contactGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactGrpcServiceClient(cc grpc.ClientConnInterface) ContactGrpcServiceClient {
	return &contactGrpcServiceClient{cc}
}

func (c *contactGrpcServiceClient) CreateContact(ctx context.Context, in *CreateContactGrpcRequest, opts ...grpc.CallOption) (*CreateContactGrpcResponse, error) {
	out := new(CreateContactGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService/CreateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactGrpcServiceClient) UpsertContact(ctx context.Context, in *UpsertContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error) {
	out := new(ContactIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService/UpsertContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactGrpcServiceClient) LinkPhoneNumberToContact(ctx context.Context, in *LinkPhoneNumberToContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error) {
	out := new(ContactIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService/LinkPhoneNumberToContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactGrpcServiceClient) LinkEmailToContact(ctx context.Context, in *LinkEmailToContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error) {
	out := new(ContactIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService/LinkEmailToContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactGrpcServiceClient) UnlinkLocationFromContact(ctx context.Context, in *UnlinkLocationFromContactGrpcRequest, opts ...grpc.CallOption) (*ContactIdGrpcResponse, error) {
	out := new(ContactIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/contactGrpcService/UnlinkLocationFromContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactGrpcServiceServer is the server API for ContactGrpcService service.
// All implementations should embed UnimplementedContactGrpcServiceServer
// for forward compatibility
type ContactGrpcServiceServer interface {
	CreateContact(context.Context, *CreateContactGrpcRequest) (*CreateContactGrpcResponse, error)
	UpsertContact(context.Context, *UpsertContactGrpcRequest) (*ContactIdGrpcResponse, error)
	LinkPhoneNumberToContact(context.Context, *LinkPhoneNumberToContactGrpcRequest) (*ContactIdGrpcResponse, error)
	LinkEmailToContact(context.Context, *LinkEmailToContactGrpcRequest) (*ContactIdGrpcResponse, error)
	UnlinkLocationFromContact(context.Context, *UnlinkLocationFromContactGrpcRequest) (*ContactIdGrpcResponse, error)
}

// UnimplementedContactGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContactGrpcServiceServer struct {
}

func (UnimplementedContactGrpcServiceServer) CreateContact(context.Context, *CreateContactGrpcRequest) (*CreateContactGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (UnimplementedContactGrpcServiceServer) UpsertContact(context.Context, *UpsertContactGrpcRequest) (*ContactIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertContact not implemented")
}
func (UnimplementedContactGrpcServiceServer) LinkPhoneNumberToContact(context.Context, *LinkPhoneNumberToContactGrpcRequest) (*ContactIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkPhoneNumberToContact not implemented")
}
func (UnimplementedContactGrpcServiceServer) LinkEmailToContact(context.Context, *LinkEmailToContactGrpcRequest) (*ContactIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkEmailToContact not implemented")
}
func (UnimplementedContactGrpcServiceServer) UnlinkLocationFromContact(context.Context, *UnlinkLocationFromContactGrpcRequest) (*ContactIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkLocationFromContact not implemented")
}

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
		FullMethod: "/contactGrpcService/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).CreateContact(ctx, req.(*CreateContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactGrpcService_UpsertContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertContactGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactGrpcServiceServer).UpsertContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactGrpcService/UpsertContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).UpsertContact(ctx, req.(*UpsertContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactGrpcService_LinkPhoneNumberToContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkPhoneNumberToContactGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactGrpcServiceServer).LinkPhoneNumberToContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactGrpcService/LinkPhoneNumberToContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).LinkPhoneNumberToContact(ctx, req.(*LinkPhoneNumberToContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactGrpcService_LinkEmailToContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkEmailToContactGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactGrpcServiceServer).LinkEmailToContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactGrpcService/LinkEmailToContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).LinkEmailToContact(ctx, req.(*LinkEmailToContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactGrpcService_UnlinkLocationFromContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkLocationFromContactGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactGrpcServiceServer).UnlinkLocationFromContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactGrpcService/UnlinkLocationFromContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactGrpcServiceServer).UnlinkLocationFromContact(ctx, req.(*UnlinkLocationFromContactGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactGrpcService_ServiceDesc is the grpc.ServiceDesc for ContactGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contactGrpcService",
	HandlerType: (*ContactGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _ContactGrpcService_CreateContact_Handler,
		},
		{
			MethodName: "UpsertContact",
			Handler:    _ContactGrpcService_UpsertContact_Handler,
		},
		{
			MethodName: "LinkPhoneNumberToContact",
			Handler:    _ContactGrpcService_LinkPhoneNumberToContact_Handler,
		},
		{
			MethodName: "LinkEmailToContact",
			Handler:    _ContactGrpcService_LinkEmailToContact_Handler,
		},
		{
			MethodName: "UnlinkLocationFromContact",
			Handler:    _ContactGrpcService_UnlinkLocationFromContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/contact.proto",
}
