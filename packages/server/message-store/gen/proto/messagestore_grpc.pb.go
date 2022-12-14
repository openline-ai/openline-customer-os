// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: messagestore.proto

package proto

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

// MessageStoreServiceClient is the client API for MessageStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageStoreServiceClient interface {
	SaveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*MessagePagedResponse, error)
	GetMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error)
	GetFeeds(ctx context.Context, in *GetFeedsPagedRequest, opts ...grpc.CallOption) (*FeedItemPagedResponse, error)
	GetFeed(ctx context.Context, in *Id, opts ...grpc.CallOption) (*FeedItem, error)
}

type messageStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageStoreServiceClient(cc grpc.ClientConnInterface) MessageStoreServiceClient {
	return &messageStoreServiceClient{cc}
}

func (c *messageStoreServiceClient) SaveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/MessageStoreService/saveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*MessagePagedResponse, error) {
	out := new(MessagePagedResponse)
	err := c.cc.Invoke(ctx, "/MessageStoreService/getMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/MessageStoreService/getMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetFeeds(ctx context.Context, in *GetFeedsPagedRequest, opts ...grpc.CallOption) (*FeedItemPagedResponse, error) {
	out := new(FeedItemPagedResponse)
	err := c.cc.Invoke(ctx, "/MessageStoreService/getFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetFeed(ctx context.Context, in *Id, opts ...grpc.CallOption) (*FeedItem, error) {
	out := new(FeedItem)
	err := c.cc.Invoke(ctx, "/MessageStoreService/getFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStoreServiceServer is the server API for MessageStoreService service.
// All implementations must embed UnimplementedMessageStoreServiceServer
// for forward compatibility
type MessageStoreServiceServer interface {
	SaveMessage(context.Context, *Message) (*Message, error)
	GetMessages(context.Context, *GetMessagesRequest) (*MessagePagedResponse, error)
	GetMessage(context.Context, *Id) (*Message, error)
	GetFeeds(context.Context, *GetFeedsPagedRequest) (*FeedItemPagedResponse, error)
	GetFeed(context.Context, *Id) (*FeedItem, error)
	mustEmbedUnimplementedMessageStoreServiceServer()
}

// UnimplementedMessageStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageStoreServiceServer struct {
}

func (UnimplementedMessageStoreServiceServer) SaveMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*MessagePagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetMessage(context.Context, *Id) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetFeeds(context.Context, *GetFeedsPagedRequest) (*FeedItemPagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetFeed(context.Context, *Id) (*FeedItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedMessageStoreServiceServer) mustEmbedUnimplementedMessageStoreServiceServer() {}

// UnsafeMessageStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageStoreServiceServer will
// result in compilation errors.
type UnsafeMessageStoreServiceServer interface {
	mustEmbedUnimplementedMessageStoreServiceServer()
}

func RegisterMessageStoreServiceServer(s grpc.ServiceRegistrar, srv MessageStoreServiceServer) {
	s.RegisterService(&MessageStoreService_ServiceDesc, srv)
}

func _MessageStoreService_SaveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).SaveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStoreService/saveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).SaveMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStoreService/getMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStoreService/getMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetMessage(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedsPagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStoreService/getFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetFeeds(ctx, req.(*GetFeedsPagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStoreService/getFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetFeed(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageStoreService_ServiceDesc is the grpc.ServiceDesc for MessageStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageStoreService",
	HandlerType: (*MessageStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveMessage",
			Handler:    _MessageStoreService_SaveMessage_Handler,
		},
		{
			MethodName: "getMessages",
			Handler:    _MessageStoreService_GetMessages_Handler,
		},
		{
			MethodName: "getMessage",
			Handler:    _MessageStoreService_GetMessage_Handler,
		},
		{
			MethodName: "getFeeds",
			Handler:    _MessageStoreService_GetFeeds_Handler,
		},
		{
			MethodName: "getFeed",
			Handler:    _MessageStoreService_GetFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messagestore.proto",
}
