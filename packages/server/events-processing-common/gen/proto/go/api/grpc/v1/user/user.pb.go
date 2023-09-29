// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: user.proto

package user_grpc_service

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

// Request to upsert a user
type UpsertUserGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant    string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Marked as deprecated in user.proto.
	AppSource string `protobuf:"bytes,6,opt,name=appSource,proto3" json:"appSource,omitempty"`
	// Deprecated: Marked as deprecated in user.proto.
	Source string `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	// Deprecated: Marked as deprecated in user.proto.
	SourceOfTruth        string                       `protobuf:"bytes,8,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt            *timestamppb.Timestamp       `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp       `protobuf:"bytes,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Internal             bool                         `protobuf:"varint,11,opt,name=internal,proto3" json:"internal,omitempty"`
	ProfilePhotoUrl      string                       `protobuf:"bytes,12,opt,name=profilePhotoUrl,proto3" json:"profilePhotoUrl,omitempty"`
	Timezone             string                       `protobuf:"bytes,13,opt,name=timezone,proto3" json:"timezone,omitempty"`
	SourceFields         *common.SourceFields         `protobuf:"bytes,14,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	ExternalSystemFields *common.ExternalSystemFields `protobuf:"bytes,15,opt,name=externalSystemFields,proto3" json:"externalSystemFields,omitempty"`
	UserId               string                       `protobuf:"bytes,16,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpsertUserGrpcRequest) Reset() {
	*x = UpsertUserGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertUserGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserGrpcRequest) ProtoMessage() {}

func (x *UpsertUserGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserGrpcRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertUserGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Marked as deprecated in user.proto.
func (x *UpsertUserGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

// Deprecated: Marked as deprecated in user.proto.
func (x *UpsertUserGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// Deprecated: Marked as deprecated in user.proto.
func (x *UpsertUserGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertUserGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertUserGrpcRequest) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *UpsertUserGrpcRequest) GetProfilePhotoUrl() string {
	if x != nil {
		return x.ProfilePhotoUrl
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UpsertUserGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *UpsertUserGrpcRequest) GetExternalSystemFields() *common.ExternalSystemFields {
	if x != nil {
		return x.ExternalSystemFields
	}
	return nil
}

func (x *UpsertUserGrpcRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LinkJobRoleToUserGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant    string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	JobRoleId string `protobuf:"bytes,3,opt,name=jobRoleId,proto3" json:"jobRoleId,omitempty"`
}

func (x *LinkJobRoleToUserGrpcRequest) Reset() {
	*x = LinkJobRoleToUserGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkJobRoleToUserGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkJobRoleToUserGrpcRequest) ProtoMessage() {}

func (x *LinkJobRoleToUserGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkJobRoleToUserGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkJobRoleToUserGrpcRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LinkJobRoleToUserGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkJobRoleToUserGrpcRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LinkJobRoleToUserGrpcRequest) GetJobRoleId() string {
	if x != nil {
		return x.JobRoleId
	}
	return ""
}

type LinkPhoneNumberToUserGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant        string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PhoneNumberId string `protobuf:"bytes,3,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	Primary       bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label         string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *LinkPhoneNumberToUserGrpcRequest) Reset() {
	*x = LinkPhoneNumberToUserGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkPhoneNumberToUserGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkPhoneNumberToUserGrpcRequest) ProtoMessage() {}

func (x *LinkPhoneNumberToUserGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkPhoneNumberToUserGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkPhoneNumberToUserGrpcRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *LinkPhoneNumberToUserGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkPhoneNumberToUserGrpcRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LinkPhoneNumberToUserGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *LinkPhoneNumberToUserGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkPhoneNumberToUserGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type LinkEmailToUserGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant  string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	EmailId string `protobuf:"bytes,3,opt,name=emailId,proto3" json:"emailId,omitempty"`
	Primary bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label   string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *LinkEmailToUserGrpcRequest) Reset() {
	*x = LinkEmailToUserGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkEmailToUserGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkEmailToUserGrpcRequest) ProtoMessage() {}

func (x *LinkEmailToUserGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkEmailToUserGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkEmailToUserGrpcRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LinkEmailToUserGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkEmailToUserGrpcRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LinkEmailToUserGrpcRequest) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *LinkEmailToUserGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkEmailToUserGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type UserIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdGrpcResponse) Reset() {
	*x = UserIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdGrpcResponse) ProtoMessage() {}

func (x *UserIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*UserIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe1, 0x04, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x31,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x49, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x1c, 0x4c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x52,
	0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x20, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x96, 0x01,
	0x0a, 0x1a, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xab, 0x02, 0x0a,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x4c,
	0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x42, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_proto_goTypes = []interface{}{
	(*UpsertUserGrpcRequest)(nil),            // 0: UpsertUserGrpcRequest
	(*LinkJobRoleToUserGrpcRequest)(nil),     // 1: LinkJobRoleToUserGrpcRequest
	(*LinkPhoneNumberToUserGrpcRequest)(nil), // 2: LinkPhoneNumberToUserGrpcRequest
	(*LinkEmailToUserGrpcRequest)(nil),       // 3: LinkEmailToUserGrpcRequest
	(*UserIdGrpcResponse)(nil),               // 4: UserIdGrpcResponse
	(*timestamppb.Timestamp)(nil),            // 5: google.protobuf.Timestamp
	(*common.SourceFields)(nil),              // 6: SourceFields
	(*common.ExternalSystemFields)(nil),      // 7: ExternalSystemFields
}
var file_user_proto_depIdxs = []int32{
	5, // 0: UpsertUserGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	5, // 1: UpsertUserGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6, // 2: UpsertUserGrpcRequest.sourceFields:type_name -> SourceFields
	7, // 3: UpsertUserGrpcRequest.externalSystemFields:type_name -> ExternalSystemFields
	0, // 4: userGrpcService.UpsertUser:input_type -> UpsertUserGrpcRequest
	1, // 5: userGrpcService.LinkJobRoleToUser:input_type -> LinkJobRoleToUserGrpcRequest
	2, // 6: userGrpcService.LinkPhoneNumberToUser:input_type -> LinkPhoneNumberToUserGrpcRequest
	3, // 7: userGrpcService.LinkEmailToUser:input_type -> LinkEmailToUserGrpcRequest
	4, // 8: userGrpcService.UpsertUser:output_type -> UserIdGrpcResponse
	4, // 9: userGrpcService.LinkJobRoleToUser:output_type -> UserIdGrpcResponse
	4, // 10: userGrpcService.LinkPhoneNumberToUser:output_type -> UserIdGrpcResponse
	4, // 11: userGrpcService.LinkEmailToUser:output_type -> UserIdGrpcResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertUserGrpcRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkJobRoleToUserGrpcRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkPhoneNumberToUserGrpcRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkEmailToUserGrpcRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdGrpcResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
