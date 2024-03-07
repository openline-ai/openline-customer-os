// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: country.proto

package country_grpc_service

import (
	common "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
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

type CreateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoggedInUserId string               `protobuf:"bytes,1,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	SourceFields   *common.SourceFields `protobuf:"bytes,2,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	Name           string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CodeA2         string               `protobuf:"bytes,4,opt,name=codeA2,proto3" json:"codeA2,omitempty"`
	CodeA3         string               `protobuf:"bytes,5,opt,name=codeA3,proto3" json:"codeA3,omitempty"`
	PhoneCode      string               `protobuf:"bytes,6,opt,name=phoneCode,proto3" json:"phoneCode,omitempty"`
}

func (x *CreateCountryRequest) Reset() {
	*x = CreateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryRequest) ProtoMessage() {}

func (x *CreateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryRequest.ProtoReflect.Descriptor instead.
func (*CreateCountryRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCountryRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CreateCountryRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *CreateCountryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCountryRequest) GetCodeA2() string {
	if x != nil {
		return x.CodeA2
	}
	return ""
}

func (x *CreateCountryRequest) GetCodeA3() string {
	if x != nil {
		return x.CodeA3
	}
	return ""
}

func (x *CreateCountryRequest) GetPhoneCode() string {
	if x != nil {
		return x.PhoneCode
	}
	return ""
}

type CountryIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountryIdGrpcResponse) Reset() {
	*x = CountryIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryIdGrpcResponse) ProtoMessage() {}

func (x *CountryIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*CountryIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{1}
}

func (x *CountryIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_country_proto protoreflect.FileDescriptor

var file_country_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x64, 0x65, 0x41, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x33, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x33, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x54, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x42, 0x0c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x3b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_country_proto_rawDescOnce sync.Once
	file_country_proto_rawDescData = file_country_proto_rawDesc
)

func file_country_proto_rawDescGZIP() []byte {
	file_country_proto_rawDescOnce.Do(func() {
		file_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_country_proto_rawDescData)
	})
	return file_country_proto_rawDescData
}

var file_country_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_country_proto_goTypes = []interface{}{
	(*CreateCountryRequest)(nil),  // 0: CreateCountryRequest
	(*CountryIdGrpcResponse)(nil), // 1: CountryIdGrpcResponse
	(*common.SourceFields)(nil),   // 2: SourceFields
}
var file_country_proto_depIdxs = []int32{
	2, // 0: CreateCountryRequest.sourceFields:type_name -> SourceFields
	0, // 1: CountryGrpcService.CreateCountry:input_type -> CreateCountryRequest
	1, // 2: CountryGrpcService.CreateCountry:output_type -> CountryIdGrpcResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_country_proto_init() }
func file_country_proto_init() {
	if File_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryRequest); i {
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
		file_country_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountryIdGrpcResponse); i {
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
			RawDescriptor: file_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_country_proto_goTypes,
		DependencyIndexes: file_country_proto_depIdxs,
		MessageInfos:      file_country_proto_msgTypes,
	}.Build()
	File_country_proto = out.File
	file_country_proto_rawDesc = nil
	file_country_proto_goTypes = nil
	file_country_proto_depIdxs = nil
}
