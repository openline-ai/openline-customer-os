// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: interaction_event.proto

package interaction_event_grpc_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestGenerateSummaryGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant             string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	InteractionEventId string `protobuf:"bytes,2,opt,name=interactionEventId,proto3" json:"interactionEventId,omitempty"`
}

func (x *RequestGenerateSummaryGrpcRequest) Reset() {
	*x = RequestGenerateSummaryGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGenerateSummaryGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGenerateSummaryGrpcRequest) ProtoMessage() {}

func (x *RequestGenerateSummaryGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGenerateSummaryGrpcRequest.ProtoReflect.Descriptor instead.
func (*RequestGenerateSummaryGrpcRequest) Descriptor() ([]byte, []int) {
	return file_interaction_event_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGenerateSummaryGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *RequestGenerateSummaryGrpcRequest) GetInteractionEventId() string {
	if x != nil {
		return x.InteractionEventId
	}
	return ""
}

type InteractionEventIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InteractionEventIdGrpcResponse) Reset() {
	*x = InteractionEventIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionEventIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionEventIdGrpcResponse) ProtoMessage() {}

func (x *InteractionEventIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionEventIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*InteractionEventIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_interaction_event_proto_rawDescGZIP(), []int{1}
}

func (x *InteractionEventIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_interaction_event_proto protoreflect.FileDescriptor

var file_interaction_event_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x21, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x1e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x7c, 0x0a, 0x1b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x22, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x42, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interaction_event_proto_rawDescOnce sync.Once
	file_interaction_event_proto_rawDescData = file_interaction_event_proto_rawDesc
)

func file_interaction_event_proto_rawDescGZIP() []byte {
	file_interaction_event_proto_rawDescOnce.Do(func() {
		file_interaction_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_interaction_event_proto_rawDescData)
	})
	return file_interaction_event_proto_rawDescData
}

var file_interaction_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interaction_event_proto_goTypes = []interface{}{
	(*RequestGenerateSummaryGrpcRequest)(nil), // 0: RequestGenerateSummaryGrpcRequest
	(*InteractionEventIdGrpcResponse)(nil),    // 1: InteractionEventIdGrpcResponse
}
var file_interaction_event_proto_depIdxs = []int32{
	0, // 0: interactionEventGrpcService.RequestGenerateSummary:input_type -> RequestGenerateSummaryGrpcRequest
	1, // 1: interactionEventGrpcService.RequestGenerateSummary:output_type -> InteractionEventIdGrpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interaction_event_proto_init() }
func file_interaction_event_proto_init() {
	if File_interaction_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interaction_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGenerateSummaryGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interaction_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionEventIdGrpcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interaction_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interaction_event_proto_goTypes,
		DependencyIndexes: file_interaction_event_proto_depIdxs,
		MessageInfos:      file_interaction_event_proto_msgTypes,
	}.Build()
	File_interaction_event_proto = out.File
	file_interaction_event_proto_rawDesc = nil
	file_interaction_event_proto_goTypes = nil
	file_interaction_event_proto_depIdxs = nil
}
