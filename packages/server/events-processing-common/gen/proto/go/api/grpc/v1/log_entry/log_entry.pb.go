// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: log_entry.proto

package log_entry_grpc_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpsertLogEntryGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant               string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UserId               string                 `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Content              string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ContentType          string                 `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	StartedAt            *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	AppSource            string                 `protobuf:"bytes,9,opt,name=appSource,proto3" json:"appSource,omitempty"`
	Source               string                 `protobuf:"bytes,10,opt,name=source,proto3" json:"source,omitempty"`
	SourceOfTruth        string                 `protobuf:"bytes,11,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	AuthorUserId         *string                `protobuf:"bytes,12,opt,name=authorUserId,proto3,oneof" json:"authorUserId,omitempty"`
	LoggedOrganizationId *string                `protobuf:"bytes,13,opt,name=loggedOrganizationId,proto3,oneof" json:"loggedOrganizationId,omitempty"`
}

func (x *UpsertLogEntryGrpcRequest) Reset() {
	*x = UpsertLogEntryGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertLogEntryGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertLogEntryGrpcRequest) ProtoMessage() {}

func (x *UpsertLogEntryGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertLogEntryGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertLogEntryGrpcRequest) Descriptor() ([]byte, []int) {
	return file_log_entry_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertLogEntryGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertLogEntryGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertLogEntryGrpcRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *UpsertLogEntryGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetAuthorUserId() string {
	if x != nil && x.AuthorUserId != nil {
		return *x.AuthorUserId
	}
	return ""
}

func (x *UpsertLogEntryGrpcRequest) GetLoggedOrganizationId() string {
	if x != nil && x.LoggedOrganizationId != nil {
		return *x.LoggedOrganizationId
	}
	return ""
}

type LogEntryIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LogEntryIdGrpcResponse) Reset() {
	*x = LogEntryIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryIdGrpcResponse) ProtoMessage() {}

func (x *LogEntryIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*LogEntryIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_log_entry_proto_rawDescGZIP(), []int{1}
}

func (x *LogEntryIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_log_entry_proto protoreflect.FileDescriptor

var file_log_entry_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x19, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70,
	0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f,
	0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x37, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x14, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x5c, 0x0a, 0x13,
	0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x42, 0x0d, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x3b, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_log_entry_proto_rawDescOnce sync.Once
	file_log_entry_proto_rawDescData = file_log_entry_proto_rawDesc
)

func file_log_entry_proto_rawDescGZIP() []byte {
	file_log_entry_proto_rawDescOnce.Do(func() {
		file_log_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_entry_proto_rawDescData)
	})
	return file_log_entry_proto_rawDescData
}

var file_log_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_log_entry_proto_goTypes = []interface{}{
	(*UpsertLogEntryGrpcRequest)(nil), // 0: UpsertLogEntryGrpcRequest
	(*LogEntryIdGrpcResponse)(nil),    // 1: LogEntryIdGrpcResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_log_entry_proto_depIdxs = []int32{
	2, // 0: UpsertLogEntryGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	2, // 1: UpsertLogEntryGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 2: UpsertLogEntryGrpcRequest.startedAt:type_name -> google.protobuf.Timestamp
	0, // 3: logEntryGrpcService.UpsertLogEntry:input_type -> UpsertLogEntryGrpcRequest
	1, // 4: logEntryGrpcService.UpsertLogEntry:output_type -> LogEntryIdGrpcResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_log_entry_proto_init() }
func file_log_entry_proto_init() {
	if File_log_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertLogEntryGrpcRequest); i {
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
		file_log_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryIdGrpcResponse); i {
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
	file_log_entry_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_log_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_entry_proto_goTypes,
		DependencyIndexes: file_log_entry_proto_depIdxs,
		MessageInfos:      file_log_entry_proto_msgTypes,
	}.Build()
	File_log_entry_proto = out.File
	file_log_entry_proto_rawDesc = nil
	file_log_entry_proto_goTypes = nil
	file_log_entry_proto_depIdxs = nil
}
