// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: phone_number.proto

package phoneNumberGrpcService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePhoneNumberGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	PhoneNumber   string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	AppSource     string `protobuf:"bytes,3,opt,name=appSource,proto3" json:"appSource,omitempty"`
	Source        string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	SourceOfTruth string `protobuf:"bytes,5,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
}

func (x *CreatePhoneNumberGrpcRequest) Reset() {
	*x = CreatePhoneNumberGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneNumberGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneNumberGrpcRequest) ProtoMessage() {}

func (x *CreatePhoneNumberGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_phone_number_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneNumberGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreatePhoneNumberGrpcRequest) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePhoneNumberGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreatePhoneNumberGrpcRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreatePhoneNumberGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *CreatePhoneNumberGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreatePhoneNumberGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

type CreatePhoneNumberGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *CreatePhoneNumberGrpcResponse) Reset() {
	*x = CreatePhoneNumberGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneNumberGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneNumberGrpcResponse) ProtoMessage() {}

func (x *CreatePhoneNumberGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_phone_number_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneNumberGrpcResponse.ProtoReflect.Descriptor instead.
func (*CreatePhoneNumberGrpcResponse) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePhoneNumberGrpcResponse) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

var File_phone_number_proto protoreflect.FileDescriptor

var file_phone_number_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb4, 0x01, 0x0a,
	0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72,
	0x75, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x32, 0x9b, 0x01, 0x0a, 0x16, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x3b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phone_number_proto_rawDescOnce sync.Once
	file_phone_number_proto_rawDescData = file_phone_number_proto_rawDesc
)

func file_phone_number_proto_rawDescGZIP() []byte {
	file_phone_number_proto_rawDescOnce.Do(func() {
		file_phone_number_proto_rawDescData = protoimpl.X.CompressGZIP(file_phone_number_proto_rawDescData)
	})
	return file_phone_number_proto_rawDescData
}

var file_phone_number_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_phone_number_proto_goTypes = []interface{}{
	(*CreatePhoneNumberGrpcRequest)(nil),  // 0: phoneNumberGrpcService.CreatePhoneNumberGrpcRequest
	(*CreatePhoneNumberGrpcResponse)(nil), // 1: phoneNumberGrpcService.CreatePhoneNumberGrpcResponse
}
var file_phone_number_proto_depIdxs = []int32{
	0, // 0: phoneNumberGrpcService.phoneNumberGrpcService.CreatePhoneNumber:input_type -> phoneNumberGrpcService.CreatePhoneNumberGrpcRequest
	1, // 1: phoneNumberGrpcService.phoneNumberGrpcService.CreatePhoneNumber:output_type -> phoneNumberGrpcService.CreatePhoneNumberGrpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_phone_number_proto_init() }
func file_phone_number_proto_init() {
	if File_phone_number_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phone_number_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneNumberGrpcRequest); i {
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
		file_phone_number_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneNumberGrpcResponse); i {
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
			RawDescriptor: file_phone_number_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_phone_number_proto_goTypes,
		DependencyIndexes: file_phone_number_proto_depIdxs,
		MessageInfos:      file_phone_number_proto_msgTypes,
	}.Build()
	File_phone_number_proto = out.File
	file_phone_number_proto_rawDesc = nil
	file_phone_number_proto_goTypes = nil
	file_phone_number_proto_depIdxs = nil
}
