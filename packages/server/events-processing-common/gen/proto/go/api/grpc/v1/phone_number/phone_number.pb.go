// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: phone_number.proto

package phone_number_grpc_service

import (
	common "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/common"
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

type UpsertPhoneNumberGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant      string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	// Deprecated: Marked as deprecated in phone_number.proto.
	AppSource string `protobuf:"bytes,3,opt,name=appSource,proto3" json:"appSource,omitempty"`
	// Deprecated: Marked as deprecated in phone_number.proto.
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// Deprecated: Marked as deprecated in phone_number.proto.
	SourceOfTruth  string                 `protobuf:"bytes,5,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Id             string                 `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,9,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	SourceFields   *common.SourceFields   `protobuf:"bytes,10,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
}

func (x *UpsertPhoneNumberGrpcRequest) Reset() {
	*x = UpsertPhoneNumberGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertPhoneNumberGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertPhoneNumberGrpcRequest) ProtoMessage() {}

func (x *UpsertPhoneNumberGrpcRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpsertPhoneNumberGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertPhoneNumberGrpcRequest) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertPhoneNumberGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertPhoneNumberGrpcRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

// Deprecated: Marked as deprecated in phone_number.proto.
func (x *UpsertPhoneNumberGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

// Deprecated: Marked as deprecated in phone_number.proto.
func (x *UpsertPhoneNumberGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// Deprecated: Marked as deprecated in phone_number.proto.
func (x *UpsertPhoneNumberGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertPhoneNumberGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertPhoneNumberGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertPhoneNumberGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertPhoneNumberGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *UpsertPhoneNumberGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

type FailPhoneNumberValidationGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	PhoneNumberId  string `protobuf:"bytes,2,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	PhoneNumber    string `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	CountryCodeA2  string `protobuf:"bytes,4,opt,name=countryCodeA2,proto3" json:"countryCodeA2,omitempty"`
	ErrorMessage   string `protobuf:"bytes,5,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	AppSource      string `protobuf:"bytes,6,opt,name=appSource,proto3" json:"appSource,omitempty"`
	LoggedInUserId string `protobuf:"bytes,7,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *FailPhoneNumberValidationGrpcRequest) Reset() {
	*x = FailPhoneNumberValidationGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailPhoneNumberValidationGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailPhoneNumberValidationGrpcRequest) ProtoMessage() {}

func (x *FailPhoneNumberValidationGrpcRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FailPhoneNumberValidationGrpcRequest.ProtoReflect.Descriptor instead.
func (*FailPhoneNumberValidationGrpcRequest) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{1}
}

func (x *FailPhoneNumberValidationGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetCountryCodeA2() string {
	if x != nil {
		return x.CountryCodeA2
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *FailPhoneNumberValidationGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type PassPhoneNumberValidationGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	PhoneNumberId  string `protobuf:"bytes,2,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	PhoneNumber    string `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	E164           string `protobuf:"bytes,4,opt,name=e164,proto3" json:"e164,omitempty"`
	CountryCodeA2  string `protobuf:"bytes,5,opt,name=countryCodeA2,proto3" json:"countryCodeA2,omitempty"`
	AppSource      string `protobuf:"bytes,6,opt,name=appSource,proto3" json:"appSource,omitempty"`
	LoggedInUserId string `protobuf:"bytes,7,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *PassPhoneNumberValidationGrpcRequest) Reset() {
	*x = PassPhoneNumberValidationGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassPhoneNumberValidationGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassPhoneNumberValidationGrpcRequest) ProtoMessage() {}

func (x *PassPhoneNumberValidationGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_phone_number_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassPhoneNumberValidationGrpcRequest.ProtoReflect.Descriptor instead.
func (*PassPhoneNumberValidationGrpcRequest) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{2}
}

func (x *PassPhoneNumberValidationGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetE164() string {
	if x != nil {
		return x.E164
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetCountryCodeA2() string {
	if x != nil {
		return x.CountryCodeA2
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *PassPhoneNumberValidationGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type PhoneNumberIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PhoneNumberIdGrpcResponse) Reset() {
	*x = PhoneNumberIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumberIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumberIdGrpcResponse) ProtoMessage() {}

func (x *PhoneNumberIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_phone_number_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumberIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*PhoneNumberIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{3}
}

func (x *PhoneNumberIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_phone_number_proto protoreflect.FileDescriptor

var file_phone_number_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x03, 0x0a, 0x1c, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x61, 0x70,
	0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54,
	0x72, 0x75, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x96, 0x02, 0x0a,
	0x24, 0x46, 0x61, 0x69, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x41, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x32, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x86, 0x02, 0x0a, 0x24, 0x50, 0x61, 0x73, 0x73, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x65, 0x31, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x31,
	0x36, 0x34, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x41, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b,
	0x0a, 0x19, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa8, 0x02, 0x0a, 0x16,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x46, 0x61, 0x69, 0x6c, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x50, 0x61, 0x73, 0x73, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x42, 0x10, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x3b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_phone_number_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_phone_number_proto_goTypes = []interface{}{
	(*UpsertPhoneNumberGrpcRequest)(nil),         // 0: UpsertPhoneNumberGrpcRequest
	(*FailPhoneNumberValidationGrpcRequest)(nil), // 1: FailPhoneNumberValidationGrpcRequest
	(*PassPhoneNumberValidationGrpcRequest)(nil), // 2: PassPhoneNumberValidationGrpcRequest
	(*PhoneNumberIdGrpcResponse)(nil),            // 3: PhoneNumberIdGrpcResponse
	(*timestamppb.Timestamp)(nil),                // 4: google.protobuf.Timestamp
	(*common.SourceFields)(nil),                  // 5: SourceFields
}
var file_phone_number_proto_depIdxs = []int32{
	4, // 0: UpsertPhoneNumberGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: UpsertPhoneNumberGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	5, // 2: UpsertPhoneNumberGrpcRequest.sourceFields:type_name -> SourceFields
	0, // 3: phoneNumberGrpcService.UpsertPhoneNumber:input_type -> UpsertPhoneNumberGrpcRequest
	1, // 4: phoneNumberGrpcService.FailPhoneNumberValidation:input_type -> FailPhoneNumberValidationGrpcRequest
	2, // 5: phoneNumberGrpcService.PassPhoneNumberValidation:input_type -> PassPhoneNumberValidationGrpcRequest
	3, // 6: phoneNumberGrpcService.UpsertPhoneNumber:output_type -> PhoneNumberIdGrpcResponse
	3, // 7: phoneNumberGrpcService.FailPhoneNumberValidation:output_type -> PhoneNumberIdGrpcResponse
	3, // 8: phoneNumberGrpcService.PassPhoneNumberValidation:output_type -> PhoneNumberIdGrpcResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_phone_number_proto_init() }
func file_phone_number_proto_init() {
	if File_phone_number_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phone_number_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertPhoneNumberGrpcRequest); i {
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
			switch v := v.(*FailPhoneNumberValidationGrpcRequest); i {
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
		file_phone_number_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassPhoneNumberValidationGrpcRequest); i {
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
		file_phone_number_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumberIdGrpcResponse); i {
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
			NumMessages:   4,
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
