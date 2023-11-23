// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: contract.proto

package contract_grpc_service

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

// Enum for RenewalCycle
type RenewalCycle int32

const (
	RenewalCycle_NONE             RenewalCycle = 0
	RenewalCycle_MONTHLY_RENEWAL  RenewalCycle = 1
	RenewalCycle_ANNUALLY_RENEWAL RenewalCycle = 2
)

// Enum value maps for RenewalCycle.
var (
	RenewalCycle_name = map[int32]string{
		0: "NONE",
		1: "MONTHLY_RENEWAL",
		2: "ANNUALLY_RENEWAL",
	}
	RenewalCycle_value = map[string]int32{
		"NONE":             0,
		"MONTHLY_RENEWAL":  1,
		"ANNUALLY_RENEWAL": 2,
	}
)

func (x RenewalCycle) Enum() *RenewalCycle {
	p := new(RenewalCycle)
	*p = x
	return p
}

func (x RenewalCycle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RenewalCycle) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_proto_enumTypes[0].Descriptor()
}

func (RenewalCycle) Type() protoreflect.EnumType {
	return &file_contract_proto_enumTypes[0]
}

func (x RenewalCycle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RenewalCycle.Descriptor instead.
func (RenewalCycle) EnumDescriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

// CreateContract request message
type CreateContractGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant               string                       `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId       string                       `protobuf:"bytes,2,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	OrganizationId       string                       `protobuf:"bytes,3,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	Name                 string                       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreatedByUserId      string                       `protobuf:"bytes,5,opt,name=createdByUserId,proto3" json:"createdByUserId,omitempty"`
	CreatedAt            *timestamppb.Timestamp       `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp       `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ServiceStartedAt     *timestamppb.Timestamp       `protobuf:"bytes,8,opt,name=serviceStartedAt,proto3" json:"serviceStartedAt,omitempty"`
	SignedAt             *timestamppb.Timestamp       `protobuf:"bytes,9,opt,name=signedAt,proto3" json:"signedAt,omitempty"`
	RenewalCycle         RenewalCycle                 `protobuf:"varint,10,opt,name=renewalCycle,proto3,enum=RenewalCycle" json:"renewalCycle,omitempty"`
	SourceFields         *common.SourceFields         `protobuf:"bytes,12,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	ExternalSystemFields *common.ExternalSystemFields `protobuf:"bytes,13,opt,name=externalSystemFields,proto3" json:"externalSystemFields,omitempty"`
	ContractUrl          string                       `protobuf:"bytes,14,opt,name=contractUrl,proto3" json:"contractUrl,omitempty"`
}

func (x *CreateContractGrpcRequest) Reset() {
	*x = CreateContractGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractGrpcRequest) ProtoMessage() {}

func (x *CreateContractGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateContractGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContractGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateContractGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CreateContractGrpcRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *CreateContractGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateContractGrpcRequest) GetCreatedByUserId() string {
	if x != nil {
		return x.CreatedByUserId
	}
	return ""
}

func (x *CreateContractGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetServiceStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ServiceStartedAt
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetSignedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SignedAt
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetRenewalCycle() RenewalCycle {
	if x != nil {
		return x.RenewalCycle
	}
	return RenewalCycle_NONE
}

func (x *CreateContractGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetExternalSystemFields() *common.ExternalSystemFields {
	if x != nil {
		return x.ExternalSystemFields
	}
	return nil
}

func (x *CreateContractGrpcRequest) GetContractUrl() string {
	if x != nil {
		return x.ContractUrl
	}
	return ""
}

type UpdateContractGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant               string                       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId       string                       `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Name                 string                       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	UpdatedAt            *timestamppb.Timestamp       `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ServiceStartedAt     *timestamppb.Timestamp       `protobuf:"bytes,6,opt,name=serviceStartedAt,proto3" json:"serviceStartedAt,omitempty"`
	SignedAt             *timestamppb.Timestamp       `protobuf:"bytes,7,opt,name=signedAt,proto3" json:"signedAt,omitempty"`
	EndedAt              *timestamppb.Timestamp       `protobuf:"bytes,8,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	RenewalCycle         RenewalCycle                 `protobuf:"varint,9,opt,name=renewalCycle,proto3,enum=RenewalCycle" json:"renewalCycle,omitempty"`
	SourceFields         *common.SourceFields         `protobuf:"bytes,10,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	ExternalSystemFields *common.ExternalSystemFields `protobuf:"bytes,11,opt,name=externalSystemFields,proto3" json:"externalSystemFields,omitempty"`
	ContractUrl          string                       `protobuf:"bytes,12,opt,name=contractUrl,proto3" json:"contractUrl,omitempty"`
}

func (x *UpdateContractGrpcRequest) Reset() {
	*x = UpdateContractGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContractGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContractGrpcRequest) ProtoMessage() {}

func (x *UpdateContractGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContractGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpdateContractGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateContractGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateContractGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpdateContractGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *UpdateContractGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateContractGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetServiceStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ServiceStartedAt
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetSignedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SignedAt
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetRenewalCycle() RenewalCycle {
	if x != nil {
		return x.RenewalCycle
	}
	return RenewalCycle_NONE
}

func (x *UpdateContractGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetExternalSystemFields() *common.ExternalSystemFields {
	if x != nil {
		return x.ExternalSystemFields
	}
	return nil
}

func (x *UpdateContractGrpcRequest) GetContractUrl() string {
	if x != nil {
		return x.ContractUrl
	}
	return ""
}

type RolloutRenewalOpportunityOnExpirationGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant         string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	AppSource      string `protobuf:"bytes,4,opt,name=appSource,proto3" json:"appSource,omitempty"`
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) Reset() {
	*x = RolloutRenewalOpportunityOnExpirationGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutRenewalOpportunityOnExpirationGrpcRequest) ProtoMessage() {}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutRenewalOpportunityOnExpirationGrpcRequest.ProtoReflect.Descriptor instead.
func (*RolloutRenewalOpportunityOnExpirationGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *RolloutRenewalOpportunityOnExpirationGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

type RefreshContractStatusGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant         string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	AppSource      string `protobuf:"bytes,4,opt,name=appSource,proto3" json:"appSource,omitempty"`
}

func (x *RefreshContractStatusGrpcRequest) Reset() {
	*x = RefreshContractStatusGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshContractStatusGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshContractStatusGrpcRequest) ProtoMessage() {}

func (x *RefreshContractStatusGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshContractStatusGrpcRequest.ProtoReflect.Descriptor instead.
func (*RefreshContractStatusGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshContractStatusGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RefreshContractStatusGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *RefreshContractStatusGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *RefreshContractStatusGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

// Contract response message
type ContractIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContractIdGrpcResponse) Reset() {
	*x = ContractIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractIdGrpcResponse) ProtoMessage() {}

func (x *ContractIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*ContractIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ContractIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x05, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0c,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x52, 0x0c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x49, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x22,
	0xc2, 0x04, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x31, 0x0a, 0x0c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x14, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0xa0, 0x01, 0x0a, 0x30, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x4f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70,
	0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x2a, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x4e, 0x55, 0x41, 0x4c, 0x4c, 0x59, 0x5f,
	0x52, 0x45, 0x4e, 0x45, 0x57, 0x41, 0x4c, 0x10, 0x02, 0x32, 0xf5, 0x02, 0x0a, 0x13, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x73, 0x0a, 0x25, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x1d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3d, 0x42, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contract_proto_goTypes = []interface{}{
	(RenewalCycle)(0),                                        // 0: RenewalCycle
	(*CreateContractGrpcRequest)(nil),                        // 1: CreateContractGrpcRequest
	(*UpdateContractGrpcRequest)(nil),                        // 2: UpdateContractGrpcRequest
	(*RolloutRenewalOpportunityOnExpirationGrpcRequest)(nil), // 3: RolloutRenewalOpportunityOnExpirationGrpcRequest
	(*RefreshContractStatusGrpcRequest)(nil),                 // 4: RefreshContractStatusGrpcRequest
	(*ContractIdGrpcResponse)(nil),                           // 5: ContractIdGrpcResponse
	(*timestamppb.Timestamp)(nil),                            // 6: google.protobuf.Timestamp
	(*common.SourceFields)(nil),                              // 7: SourceFields
	(*common.ExternalSystemFields)(nil),                      // 8: ExternalSystemFields
}
var file_contract_proto_depIdxs = []int32{
	6,  // 0: CreateContractGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 1: CreateContractGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6,  // 2: CreateContractGrpcRequest.serviceStartedAt:type_name -> google.protobuf.Timestamp
	6,  // 3: CreateContractGrpcRequest.signedAt:type_name -> google.protobuf.Timestamp
	0,  // 4: CreateContractGrpcRequest.renewalCycle:type_name -> RenewalCycle
	7,  // 5: CreateContractGrpcRequest.sourceFields:type_name -> SourceFields
	8,  // 6: CreateContractGrpcRequest.externalSystemFields:type_name -> ExternalSystemFields
	6,  // 7: UpdateContractGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6,  // 8: UpdateContractGrpcRequest.serviceStartedAt:type_name -> google.protobuf.Timestamp
	6,  // 9: UpdateContractGrpcRequest.signedAt:type_name -> google.protobuf.Timestamp
	6,  // 10: UpdateContractGrpcRequest.endedAt:type_name -> google.protobuf.Timestamp
	0,  // 11: UpdateContractGrpcRequest.renewalCycle:type_name -> RenewalCycle
	7,  // 12: UpdateContractGrpcRequest.sourceFields:type_name -> SourceFields
	8,  // 13: UpdateContractGrpcRequest.externalSystemFields:type_name -> ExternalSystemFields
	1,  // 14: ContractGrpcService.CreateContract:input_type -> CreateContractGrpcRequest
	2,  // 15: ContractGrpcService.UpdateContract:input_type -> UpdateContractGrpcRequest
	3,  // 16: ContractGrpcService.RolloutRenewalOpportunityOnExpiration:input_type -> RolloutRenewalOpportunityOnExpirationGrpcRequest
	4,  // 17: ContractGrpcService.RefreshContractStatusContract:input_type -> RefreshContractStatusGrpcRequest
	5,  // 18: ContractGrpcService.CreateContract:output_type -> ContractIdGrpcResponse
	5,  // 19: ContractGrpcService.UpdateContract:output_type -> ContractIdGrpcResponse
	5,  // 20: ContractGrpcService.RolloutRenewalOpportunityOnExpiration:output_type -> ContractIdGrpcResponse
	5,  // 21: ContractGrpcService.RefreshContractStatusContract:output_type -> ContractIdGrpcResponse
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractGrpcRequest); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContractGrpcRequest); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolloutRenewalOpportunityOnExpirationGrpcRequest); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshContractStatusGrpcRequest); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractIdGrpcResponse); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		EnumInfos:         file_contract_proto_enumTypes,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
