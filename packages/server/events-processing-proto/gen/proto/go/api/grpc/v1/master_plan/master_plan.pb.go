// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: master_plan.proto

package master_plan_grpc_service

import (
	common "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
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

type MasterPlanFieldMask int32

const (
	MasterPlanFieldMask_MASTER_PLAN_PROPERTY_NONE    MasterPlanFieldMask = 0 // No property, ignored by the server
	MasterPlanFieldMask_MASTER_PLAN_PROPERTY_ALL     MasterPlanFieldMask = 1 // All properties, if present in the list all other properties are ignored
	MasterPlanFieldMask_MASTER_PLAN_PROPERTY_NAME    MasterPlanFieldMask = 2
	MasterPlanFieldMask_MASTER_PLAN_PROPERTY_RETIRED MasterPlanFieldMask = 3
)

// Enum value maps for MasterPlanFieldMask.
var (
	MasterPlanFieldMask_name = map[int32]string{
		0: "MASTER_PLAN_PROPERTY_NONE",
		1: "MASTER_PLAN_PROPERTY_ALL",
		2: "MASTER_PLAN_PROPERTY_NAME",
		3: "MASTER_PLAN_PROPERTY_RETIRED",
	}
	MasterPlanFieldMask_value = map[string]int32{
		"MASTER_PLAN_PROPERTY_NONE":    0,
		"MASTER_PLAN_PROPERTY_ALL":     1,
		"MASTER_PLAN_PROPERTY_NAME":    2,
		"MASTER_PLAN_PROPERTY_RETIRED": 3,
	}
)

func (x MasterPlanFieldMask) Enum() *MasterPlanFieldMask {
	p := new(MasterPlanFieldMask)
	*p = x
	return p
}

func (x MasterPlanFieldMask) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MasterPlanFieldMask) Descriptor() protoreflect.EnumDescriptor {
	return file_master_plan_proto_enumTypes[0].Descriptor()
}

func (MasterPlanFieldMask) Type() protoreflect.EnumType {
	return &file_master_plan_proto_enumTypes[0]
}

func (x MasterPlanFieldMask) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MasterPlanFieldMask.Descriptor instead.
func (MasterPlanFieldMask) EnumDescriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{0}
}

type CreateMasterPlanGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,2,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	SourceFields   *common.SourceFields   `protobuf:"bytes,5,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
}

func (x *CreateMasterPlanGrpcRequest) Reset() {
	*x = CreateMasterPlanGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMasterPlanGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMasterPlanGrpcRequest) ProtoMessage() {}

func (x *CreateMasterPlanGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMasterPlanGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateMasterPlanGrpcRequest) Descriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMasterPlanGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateMasterPlanGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CreateMasterPlanGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMasterPlanGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateMasterPlanGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

type CreateMasterPlanMilestoneGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	MasterPlanId   string                 `protobuf:"bytes,2,opt,name=masterPlanId,proto3" json:"masterPlanId,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Name           string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	SourceFields   *common.SourceFields   `protobuf:"bytes,6,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	Optional       bool                   `protobuf:"varint,7,opt,name=optional,proto3" json:"optional,omitempty"`
	DurationHours  int64                  `protobuf:"varint,8,opt,name=durationHours,proto3" json:"durationHours,omitempty"`
	Items          []string               `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
	Order          int64                  `protobuf:"varint,10,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *CreateMasterPlanMilestoneGrpcRequest) Reset() {
	*x = CreateMasterPlanMilestoneGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMasterPlanMilestoneGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMasterPlanMilestoneGrpcRequest) ProtoMessage() {}

func (x *CreateMasterPlanMilestoneGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMasterPlanMilestoneGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateMasterPlanMilestoneGrpcRequest) Descriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetMasterPlanId() string {
	if x != nil {
		return x.MasterPlanId
	}
	return ""
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetDurationHours() int64 {
	if x != nil {
		return x.DurationHours
	}
	return 0
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateMasterPlanMilestoneGrpcRequest) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

type UpdateMasterPlanGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterPlanId   string                 `protobuf:"bytes,1,opt,name=masterPlanId,proto3" json:"masterPlanId,omitempty"`
	Tenant         string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Name           string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Retired        bool                   `protobuf:"varint,5,opt,name=retired,proto3" json:"retired,omitempty"`
	FieldsMask     []MasterPlanFieldMask  `protobuf:"varint,6,rep,packed,name=fieldsMask,proto3,enum=MasterPlanFieldMask" json:"fieldsMask,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppSource      string                 `protobuf:"bytes,8,opt,name=appSource,proto3" json:"appSource,omitempty"`
}

func (x *UpdateMasterPlanGrpcRequest) Reset() {
	*x = UpdateMasterPlanGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMasterPlanGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMasterPlanGrpcRequest) ProtoMessage() {}

func (x *UpdateMasterPlanGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMasterPlanGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpdateMasterPlanGrpcRequest) Descriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMasterPlanGrpcRequest) GetMasterPlanId() string {
	if x != nil {
		return x.MasterPlanId
	}
	return ""
}

func (x *UpdateMasterPlanGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpdateMasterPlanGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *UpdateMasterPlanGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMasterPlanGrpcRequest) GetRetired() bool {
	if x != nil {
		return x.Retired
	}
	return false
}

func (x *UpdateMasterPlanGrpcRequest) GetFieldsMask() []MasterPlanFieldMask {
	if x != nil {
		return x.FieldsMask
	}
	return nil
}

func (x *UpdateMasterPlanGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpdateMasterPlanGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

type MasterPlanIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MasterPlanIdGrpcResponse) Reset() {
	*x = MasterPlanIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterPlanIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterPlanIdGrpcResponse) ProtoMessage() {}

func (x *MasterPlanIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterPlanIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*MasterPlanIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{3}
}

func (x *MasterPlanIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MasterPlanMilestoneIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MasterPlanMilestoneIdGrpcResponse) Reset() {
	*x = MasterPlanMilestoneIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterPlanMilestoneIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterPlanMilestoneIdGrpcResponse) ProtoMessage() {}

func (x *MasterPlanMilestoneIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterPlanMilestoneIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*MasterPlanMilestoneIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_master_plan_proto_rawDescGZIP(), []int{4}
}

func (x *MasterPlanMilestoneIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_master_plan_proto protoreflect.FileDescriptor

var file_master_plan_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xf9, 0x02, 0x0a, 0x24, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x4d,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xbd, 0x02, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x38,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x33, 0x0a, 0x21, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e,
	0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x93, 0x01, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19,
	0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x50,
	0x45, 0x52, 0x54, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4d,
	0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45,
	0x52, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x54, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x32, 0x99, 0x02,
	0x0a, 0x15, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x66, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x6c, 0x61, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x25,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c,
	0x61, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45, 0x42, 0x0f, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x3b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_plan_proto_rawDescOnce sync.Once
	file_master_plan_proto_rawDescData = file_master_plan_proto_rawDesc
)

func file_master_plan_proto_rawDescGZIP() []byte {
	file_master_plan_proto_rawDescOnce.Do(func() {
		file_master_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_plan_proto_rawDescData)
	})
	return file_master_plan_proto_rawDescData
}

var file_master_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_master_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_master_plan_proto_goTypes = []interface{}{
	(MasterPlanFieldMask)(0),                     // 0: MasterPlanFieldMask
	(*CreateMasterPlanGrpcRequest)(nil),          // 1: CreateMasterPlanGrpcRequest
	(*CreateMasterPlanMilestoneGrpcRequest)(nil), // 2: CreateMasterPlanMilestoneGrpcRequest
	(*UpdateMasterPlanGrpcRequest)(nil),          // 3: UpdateMasterPlanGrpcRequest
	(*MasterPlanIdGrpcResponse)(nil),             // 4: MasterPlanIdGrpcResponse
	(*MasterPlanMilestoneIdGrpcResponse)(nil),    // 5: MasterPlanMilestoneIdGrpcResponse
	(*timestamppb.Timestamp)(nil),                // 6: google.protobuf.Timestamp
	(*common.SourceFields)(nil),                  // 7: SourceFields
}
var file_master_plan_proto_depIdxs = []int32{
	6, // 0: CreateMasterPlanGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	7, // 1: CreateMasterPlanGrpcRequest.sourceFields:type_name -> SourceFields
	6, // 2: CreateMasterPlanMilestoneGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	7, // 3: CreateMasterPlanMilestoneGrpcRequest.sourceFields:type_name -> SourceFields
	0, // 4: UpdateMasterPlanGrpcRequest.fieldsMask:type_name -> MasterPlanFieldMask
	6, // 5: UpdateMasterPlanGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	1, // 6: MasterPlanGrpcService.CreateMasterPlan:input_type -> CreateMasterPlanGrpcRequest
	3, // 7: MasterPlanGrpcService.UpdateMasterPlan:input_type -> UpdateMasterPlanGrpcRequest
	2, // 8: MasterPlanGrpcService.CreateMasterPlanMilestone:input_type -> CreateMasterPlanMilestoneGrpcRequest
	4, // 9: MasterPlanGrpcService.CreateMasterPlan:output_type -> MasterPlanIdGrpcResponse
	4, // 10: MasterPlanGrpcService.UpdateMasterPlan:output_type -> MasterPlanIdGrpcResponse
	5, // 11: MasterPlanGrpcService.CreateMasterPlanMilestone:output_type -> MasterPlanMilestoneIdGrpcResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_master_plan_proto_init() }
func file_master_plan_proto_init() {
	if File_master_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMasterPlanGrpcRequest); i {
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
		file_master_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMasterPlanMilestoneGrpcRequest); i {
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
		file_master_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMasterPlanGrpcRequest); i {
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
		file_master_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterPlanIdGrpcResponse); i {
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
		file_master_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterPlanMilestoneIdGrpcResponse); i {
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
			RawDescriptor: file_master_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_plan_proto_goTypes,
		DependencyIndexes: file_master_plan_proto_depIdxs,
		EnumInfos:         file_master_plan_proto_enumTypes,
		MessageInfos:      file_master_plan_proto_msgTypes,
	}.Build()
	File_master_plan_proto = out.File
	file_master_plan_proto_rawDesc = nil
	file_master_plan_proto_goTypes = nil
	file_master_plan_proto_depIdxs = nil
}
