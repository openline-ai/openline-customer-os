// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: log_entry.proto

package log_entry_grpc_service

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

// LogEntryGrpcServiceClient is the client API for LogEntryGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogEntryGrpcServiceClient interface {
	UpsertLogEntry(ctx context.Context, in *UpsertLogEntryGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error)
	AddTag(ctx context.Context, in *AddTagGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error)
	RemoveTag(ctx context.Context, in *RemoveTagGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error)
}

type logEntryGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogEntryGrpcServiceClient(cc grpc.ClientConnInterface) LogEntryGrpcServiceClient {
	return &logEntryGrpcServiceClient{cc}
}

func (c *logEntryGrpcServiceClient) UpsertLogEntry(ctx context.Context, in *UpsertLogEntryGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error) {
	out := new(LogEntryIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/logEntryGrpcService/UpsertLogEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logEntryGrpcServiceClient) AddTag(ctx context.Context, in *AddTagGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error) {
	out := new(LogEntryIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/logEntryGrpcService/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logEntryGrpcServiceClient) RemoveTag(ctx context.Context, in *RemoveTagGrpcRequest, opts ...grpc.CallOption) (*LogEntryIdGrpcResponse, error) {
	out := new(LogEntryIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/logEntryGrpcService/RemoveTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogEntryGrpcServiceServer is the server API for LogEntryGrpcService service.
// All implementations should embed UnimplementedLogEntryGrpcServiceServer
// for forward compatibility
type LogEntryGrpcServiceServer interface {
	UpsertLogEntry(context.Context, *UpsertLogEntryGrpcRequest) (*LogEntryIdGrpcResponse, error)
	AddTag(context.Context, *AddTagGrpcRequest) (*LogEntryIdGrpcResponse, error)
	RemoveTag(context.Context, *RemoveTagGrpcRequest) (*LogEntryIdGrpcResponse, error)
}

// UnimplementedLogEntryGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLogEntryGrpcServiceServer struct {
}

func (UnimplementedLogEntryGrpcServiceServer) UpsertLogEntry(context.Context, *UpsertLogEntryGrpcRequest) (*LogEntryIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertLogEntry not implemented")
}
func (UnimplementedLogEntryGrpcServiceServer) AddTag(context.Context, *AddTagGrpcRequest) (*LogEntryIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedLogEntryGrpcServiceServer) RemoveTag(context.Context, *RemoveTagGrpcRequest) (*LogEntryIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTag not implemented")
}

// UnsafeLogEntryGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogEntryGrpcServiceServer will
// result in compilation errors.
type UnsafeLogEntryGrpcServiceServer interface {
	mustEmbedUnimplementedLogEntryGrpcServiceServer()
}

func RegisterLogEntryGrpcServiceServer(s grpc.ServiceRegistrar, srv LogEntryGrpcServiceServer) {
	s.RegisterService(&LogEntryGrpcService_ServiceDesc, srv)
}

func _LogEntryGrpcService_UpsertLogEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertLogEntryGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogEntryGrpcServiceServer).UpsertLogEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logEntryGrpcService/UpsertLogEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogEntryGrpcServiceServer).UpsertLogEntry(ctx, req.(*UpsertLogEntryGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogEntryGrpcService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogEntryGrpcServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logEntryGrpcService/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogEntryGrpcServiceServer).AddTag(ctx, req.(*AddTagGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogEntryGrpcService_RemoveTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTagGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogEntryGrpcServiceServer).RemoveTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logEntryGrpcService/RemoveTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogEntryGrpcServiceServer).RemoveTag(ctx, req.(*RemoveTagGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogEntryGrpcService_ServiceDesc is the grpc.ServiceDesc for LogEntryGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogEntryGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logEntryGrpcService",
	HandlerType: (*LogEntryGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertLogEntry",
			Handler:    _LogEntryGrpcService_UpsertLogEntry_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _LogEntryGrpcService_AddTag_Handler,
		},
		{
			MethodName: "RemoveTag",
			Handler:    _LogEntryGrpcService_RemoveTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log_entry.proto",
}
