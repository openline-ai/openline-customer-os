// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: common/billing_type.proto

package common

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

type BilledType int32

const (
	BilledType_NONE_BILLED      BilledType = 0
	BilledType_MONTHLY_BILLED   BilledType = 1
	BilledType_ANNUALLY_BILLED  BilledType = 2
	BilledType_ONCE_BILLED      BilledType = 3 // For One-Time
	BilledType_USAGE_BILLED     BilledType = 4 // For Usage
	BilledType_QUARTERLY_BILLED BilledType = 5
)

// Enum value maps for BilledType.
var (
	BilledType_name = map[int32]string{
		0: "NONE_BILLED",
		1: "MONTHLY_BILLED",
		2: "ANNUALLY_BILLED",
		3: "ONCE_BILLED",
		4: "USAGE_BILLED",
		5: "QUARTERLY_BILLED",
	}
	BilledType_value = map[string]int32{
		"NONE_BILLED":      0,
		"MONTHLY_BILLED":   1,
		"ANNUALLY_BILLED":  2,
		"ONCE_BILLED":      3,
		"USAGE_BILLED":     4,
		"QUARTERLY_BILLED": 5,
	}
)

func (x BilledType) Enum() *BilledType {
	p := new(BilledType)
	*p = x
	return p
}

func (x BilledType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BilledType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_billing_type_proto_enumTypes[0].Descriptor()
}

func (BilledType) Type() protoreflect.EnumType {
	return &file_common_billing_type_proto_enumTypes[0]
}

func (x BilledType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BilledType.Descriptor instead.
func (BilledType) EnumDescriptor() ([]byte, []int) {
	return file_common_billing_type_proto_rawDescGZIP(), []int{0}
}

var File_common_billing_type_proto protoreflect.FileDescriptor

var file_common_billing_type_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7f, 0x0a, 0x0a, 0x42,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x4e,
	0x45, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f,
	0x4e, 0x54, 0x48, 0x4c, 0x59, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x4e, 0x4e, 0x55, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x4e, 0x43, 0x45, 0x5f, 0x42, 0x49, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x49,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45,
	0x52, 0x4c, 0x59, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x42, 0x28, 0x42, 0x10,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_billing_type_proto_rawDescOnce sync.Once
	file_common_billing_type_proto_rawDescData = file_common_billing_type_proto_rawDesc
)

func file_common_billing_type_proto_rawDescGZIP() []byte {
	file_common_billing_type_proto_rawDescOnce.Do(func() {
		file_common_billing_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_billing_type_proto_rawDescData)
	})
	return file_common_billing_type_proto_rawDescData
}

var file_common_billing_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_billing_type_proto_goTypes = []interface{}{
	(BilledType)(0), // 0: BilledType
}
var file_common_billing_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_billing_type_proto_init() }
func file_common_billing_type_proto_init() {
	if File_common_billing_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_billing_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_billing_type_proto_goTypes,
		DependencyIndexes: file_common_billing_type_proto_depIdxs,
		EnumInfos:         file_common_billing_type_proto_enumTypes,
	}.Build()
	File_common_billing_type_proto = out.File
	file_common_billing_type_proto_rawDesc = nil
	file_common_billing_type_proto_goTypes = nil
	file_common_billing_type_proto_depIdxs = nil
}
