// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user.proto

package user_grpc_service

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

// UserGrpcServiceClient is the client API for UserGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcServiceClient interface {
	UpsertUser(ctx context.Context, in *UpsertUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	AddPlayerInfo(ctx context.Context, in *AddPlayerInfoGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	LinkJobRoleToUser(ctx context.Context, in *LinkJobRoleToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	LinkPhoneNumberToUser(ctx context.Context, in *LinkPhoneNumberToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	LinkEmailToUser(ctx context.Context, in *LinkEmailToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	AddRole(ctx context.Context, in *AddRoleGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
	RemoveRole(ctx context.Context, in *RemoveRoleGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error)
}

type userGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcServiceClient(cc grpc.ClientConnInterface) UserGrpcServiceClient {
	return &userGrpcServiceClient{cc}
}

func (c *userGrpcServiceClient) UpsertUser(ctx context.Context, in *UpsertUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/UpsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) AddPlayerInfo(ctx context.Context, in *AddPlayerInfoGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/AddPlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) LinkJobRoleToUser(ctx context.Context, in *LinkJobRoleToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/LinkJobRoleToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) LinkPhoneNumberToUser(ctx context.Context, in *LinkPhoneNumberToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/LinkPhoneNumberToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) LinkEmailToUser(ctx context.Context, in *LinkEmailToUserGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/LinkEmailToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) AddRole(ctx context.Context, in *AddRoleGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/AddRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) RemoveRole(ctx context.Context, in *RemoveRoleGrpcRequest, opts ...grpc.CallOption) (*UserIdGrpcResponse, error) {
	out := new(UserIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/userGrpcService/RemoveRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServiceServer is the server API for UserGrpcService service.
// All implementations should embed UnimplementedUserGrpcServiceServer
// for forward compatibility
type UserGrpcServiceServer interface {
	UpsertUser(context.Context, *UpsertUserGrpcRequest) (*UserIdGrpcResponse, error)
	AddPlayerInfo(context.Context, *AddPlayerInfoGrpcRequest) (*UserIdGrpcResponse, error)
	LinkJobRoleToUser(context.Context, *LinkJobRoleToUserGrpcRequest) (*UserIdGrpcResponse, error)
	LinkPhoneNumberToUser(context.Context, *LinkPhoneNumberToUserGrpcRequest) (*UserIdGrpcResponse, error)
	LinkEmailToUser(context.Context, *LinkEmailToUserGrpcRequest) (*UserIdGrpcResponse, error)
	AddRole(context.Context, *AddRoleGrpcRequest) (*UserIdGrpcResponse, error)
	RemoveRole(context.Context, *RemoveRoleGrpcRequest) (*UserIdGrpcResponse, error)
}

// UnimplementedUserGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserGrpcServiceServer struct {
}

func (UnimplementedUserGrpcServiceServer) UpsertUser(context.Context, *UpsertUserGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUser not implemented")
}
func (UnimplementedUserGrpcServiceServer) AddPlayerInfo(context.Context, *AddPlayerInfoGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlayerInfo not implemented")
}
func (UnimplementedUserGrpcServiceServer) LinkJobRoleToUser(context.Context, *LinkJobRoleToUserGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkJobRoleToUser not implemented")
}
func (UnimplementedUserGrpcServiceServer) LinkPhoneNumberToUser(context.Context, *LinkPhoneNumberToUserGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkPhoneNumberToUser not implemented")
}
func (UnimplementedUserGrpcServiceServer) LinkEmailToUser(context.Context, *LinkEmailToUserGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkEmailToUser not implemented")
}
func (UnimplementedUserGrpcServiceServer) AddRole(context.Context, *AddRoleGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedUserGrpcServiceServer) RemoveRole(context.Context, *RemoveRoleGrpcRequest) (*UserIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRole not implemented")
}

// UnsafeUserGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServiceServer will
// result in compilation errors.
type UnsafeUserGrpcServiceServer interface {
	mustEmbedUnimplementedUserGrpcServiceServer()
}

func RegisterUserGrpcServiceServer(s grpc.ServiceRegistrar, srv UserGrpcServiceServer) {
	s.RegisterService(&UserGrpcService_ServiceDesc, srv)
}

func _UserGrpcService_UpsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).UpsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/UpsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).UpsertUser(ctx, req.(*UpsertUserGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_AddPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlayerInfoGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).AddPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/AddPlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).AddPlayerInfo(ctx, req.(*AddPlayerInfoGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_LinkJobRoleToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkJobRoleToUserGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).LinkJobRoleToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/LinkJobRoleToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).LinkJobRoleToUser(ctx, req.(*LinkJobRoleToUserGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_LinkPhoneNumberToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkPhoneNumberToUserGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).LinkPhoneNumberToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/LinkPhoneNumberToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).LinkPhoneNumberToUser(ctx, req.(*LinkPhoneNumberToUserGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_LinkEmailToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkEmailToUserGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).LinkEmailToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/LinkEmailToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).LinkEmailToUser(ctx, req.(*LinkEmailToUserGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/AddRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).AddRole(ctx, req.(*AddRoleGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userGrpcService/RemoveRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).RemoveRole(ctx, req.(*RemoveRoleGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpcService_ServiceDesc is the grpc.ServiceDesc for UserGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userGrpcService",
	HandlerType: (*UserGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertUser",
			Handler:    _UserGrpcService_UpsertUser_Handler,
		},
		{
			MethodName: "AddPlayerInfo",
			Handler:    _UserGrpcService_AddPlayerInfo_Handler,
		},
		{
			MethodName: "LinkJobRoleToUser",
			Handler:    _UserGrpcService_LinkJobRoleToUser_Handler,
		},
		{
			MethodName: "LinkPhoneNumberToUser",
			Handler:    _UserGrpcService_LinkPhoneNumberToUser_Handler,
		},
		{
			MethodName: "LinkEmailToUser",
			Handler:    _UserGrpcService_LinkEmailToUser_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _UserGrpcService_AddRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _UserGrpcService_RemoveRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
