// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: service_line_item.proto

package service_line_item_grpc_service

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

type CreateServiceLineItemGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,2,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Billed         common.BilledType      `protobuf:"varint,3,opt,name=billed,proto3,enum=BilledType" json:"billed,omitempty"`
	Quantity       int64                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"` // Relevant only for Subscription type
	Price          float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Name           string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	ContractId     string                 `protobuf:"bytes,7,opt,name=contractId,proto3" json:"contractId,omitempty"`
	SourceFields   *common.SourceFields   `protobuf:"bytes,8,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt        *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	VatRate        float64                `protobuf:"fixed64,13,opt,name=vatRate,proto3" json:"vatRate,omitempty"`
}

func (x *CreateServiceLineItemGrpcRequest) Reset() {
	*x = CreateServiceLineItemGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_line_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceLineItemGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceLineItemGrpcRequest) ProtoMessage() {}

func (x *CreateServiceLineItemGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_line_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceLineItemGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceLineItemGrpcRequest) Descriptor() ([]byte, []int) {
	return file_service_line_item_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceLineItemGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateServiceLineItemGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CreateServiceLineItemGrpcRequest) GetBilled() common.BilledType {
	if x != nil {
		return x.Billed
	}
	return common.BilledType(0)
}

func (x *CreateServiceLineItemGrpcRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateServiceLineItemGrpcRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateServiceLineItemGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceLineItemGrpcRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *CreateServiceLineItemGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *CreateServiceLineItemGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateServiceLineItemGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CreateServiceLineItemGrpcRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *CreateServiceLineItemGrpcRequest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *CreateServiceLineItemGrpcRequest) GetVatRate() float64 {
	if x != nil {
		return x.VatRate
	}
	return 0
}

type UpdateServiceLineItemGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant                  string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId          string                 `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Name                    string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Price                   float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity                int64                  `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Billed                  common.BilledType      `protobuf:"varint,7,opt,name=billed,proto3,enum=BilledType" json:"billed,omitempty"`
	UpdatedAt               *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	SourceFields            *common.SourceFields   `protobuf:"bytes,9,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	Comments                string                 `protobuf:"bytes,10,opt,name=comments,proto3" json:"comments,omitempty"`
	IsRetroactiveCorrection bool                   `protobuf:"varint,11,opt,name=isRetroactiveCorrection,proto3" json:"isRetroactiveCorrection,omitempty"`
	ContractId              string                 `protobuf:"bytes,12,opt,name=contractId,proto3" json:"contractId,omitempty"`
	VatRate                 float64                `protobuf:"fixed64,13,opt,name=vatRate,proto3" json:"vatRate,omitempty"`
	ParentId                string                 `protobuf:"bytes,14,opt,name=parentId,proto3" json:"parentId,omitempty"`
	StartedAt               *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
}

func (x *UpdateServiceLineItemGrpcRequest) Reset() {
	*x = UpdateServiceLineItemGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_line_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServiceLineItemGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServiceLineItemGrpcRequest) ProtoMessage() {}

func (x *UpdateServiceLineItemGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_line_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServiceLineItemGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpdateServiceLineItemGrpcRequest) Descriptor() ([]byte, []int) {
	return file_service_line_item_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateServiceLineItemGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateServiceLineItemGrpcRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *UpdateServiceLineItemGrpcRequest) GetBilled() common.BilledType {
	if x != nil {
		return x.Billed
	}
	return common.BilledType(0)
}

func (x *UpdateServiceLineItemGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpdateServiceLineItemGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *UpdateServiceLineItemGrpcRequest) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetIsRetroactiveCorrection() bool {
	if x != nil {
		return x.IsRetroactiveCorrection
	}
	return false
}

func (x *UpdateServiceLineItemGrpcRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetVatRate() float64 {
	if x != nil {
		return x.VatRate
	}
	return 0
}

func (x *UpdateServiceLineItemGrpcRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *UpdateServiceLineItemGrpcRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

// Permanently delete service line item request
type DeleteServiceLineItemGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant         string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	AppSource      string `protobuf:"bytes,4,opt,name=appSource,proto3" json:"appSource,omitempty"`
}

func (x *DeleteServiceLineItemGrpcRequest) Reset() {
	*x = DeleteServiceLineItemGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_line_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceLineItemGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceLineItemGrpcRequest) ProtoMessage() {}

func (x *DeleteServiceLineItemGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_line_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceLineItemGrpcRequest.ProtoReflect.Descriptor instead.
func (*DeleteServiceLineItemGrpcRequest) Descriptor() ([]byte, []int) {
	return file_service_line_item_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteServiceLineItemGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteServiceLineItemGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *DeleteServiceLineItemGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *DeleteServiceLineItemGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

type CloseServiceLineItemGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant         string                 `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	AppSource      string                 `protobuf:"bytes,4,opt,name=appSource,proto3" json:"appSource,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	EndedAt        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
}

func (x *CloseServiceLineItemGrpcRequest) Reset() {
	*x = CloseServiceLineItemGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_line_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseServiceLineItemGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseServiceLineItemGrpcRequest) ProtoMessage() {}

func (x *CloseServiceLineItemGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_line_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseServiceLineItemGrpcRequest.ProtoReflect.Descriptor instead.
func (*CloseServiceLineItemGrpcRequest) Descriptor() ([]byte, []int) {
	return file_service_line_item_proto_rawDescGZIP(), []int{3}
}

func (x *CloseServiceLineItemGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloseServiceLineItemGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CloseServiceLineItemGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *CloseServiceLineItemGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

func (x *CloseServiceLineItemGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CloseServiceLineItemGrpcRequest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

// Service line item response message
type ServiceLineItemIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServiceLineItemIdGrpcResponse) Reset() {
	*x = ServiceLineItemIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_line_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceLineItemIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceLineItemIdGrpcResponse) ProtoMessage() {}

func (x *ServiceLineItemIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_line_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceLineItemIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*ServiceLineItemIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_service_line_item_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceLineItemIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_service_line_item_proto protoreflect.FileDescriptor

var file_service_line_item_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x04,
	0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x22, 0xb0,
	0x04, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x62, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x69, 0x73, 0x52, 0x65, 0x74,
	0x72, 0x6f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x52, 0x65, 0x74, 0x72,
	0x6f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x76, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x90, 0x01, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x1f, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2f, 0x0a, 0x1d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8a, 0x03, 0x0a, 0x1a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x21, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5a, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x14, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x20, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x42, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_line_item_proto_rawDescOnce sync.Once
	file_service_line_item_proto_rawDescData = file_service_line_item_proto_rawDesc
)

func file_service_line_item_proto_rawDescGZIP() []byte {
	file_service_line_item_proto_rawDescOnce.Do(func() {
		file_service_line_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_line_item_proto_rawDescData)
	})
	return file_service_line_item_proto_rawDescData
}

var file_service_line_item_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_line_item_proto_goTypes = []interface{}{
	(*CreateServiceLineItemGrpcRequest)(nil), // 0: CreateServiceLineItemGrpcRequest
	(*UpdateServiceLineItemGrpcRequest)(nil), // 1: UpdateServiceLineItemGrpcRequest
	(*DeleteServiceLineItemGrpcRequest)(nil), // 2: DeleteServiceLineItemGrpcRequest
	(*CloseServiceLineItemGrpcRequest)(nil),  // 3: CloseServiceLineItemGrpcRequest
	(*ServiceLineItemIdGrpcResponse)(nil),    // 4: ServiceLineItemIdGrpcResponse
	(common.BilledType)(0),                   // 5: BilledType
	(*common.SourceFields)(nil),              // 6: SourceFields
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_service_line_item_proto_depIdxs = []int32{
	5,  // 0: CreateServiceLineItemGrpcRequest.billed:type_name -> BilledType
	6,  // 1: CreateServiceLineItemGrpcRequest.sourceFields:type_name -> SourceFields
	7,  // 2: CreateServiceLineItemGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 3: CreateServiceLineItemGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 4: CreateServiceLineItemGrpcRequest.startedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: CreateServiceLineItemGrpcRequest.endedAt:type_name -> google.protobuf.Timestamp
	5,  // 6: UpdateServiceLineItemGrpcRequest.billed:type_name -> BilledType
	7,  // 7: UpdateServiceLineItemGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6,  // 8: UpdateServiceLineItemGrpcRequest.sourceFields:type_name -> SourceFields
	7,  // 9: UpdateServiceLineItemGrpcRequest.startedAt:type_name -> google.protobuf.Timestamp
	7,  // 10: CloseServiceLineItemGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 11: CloseServiceLineItemGrpcRequest.endedAt:type_name -> google.protobuf.Timestamp
	0,  // 12: ServiceLineItemGrpcService.CreateServiceLineItem:input_type -> CreateServiceLineItemGrpcRequest
	1,  // 13: ServiceLineItemGrpcService.UpdateServiceLineItem:input_type -> UpdateServiceLineItemGrpcRequest
	2,  // 14: ServiceLineItemGrpcService.DeleteServiceLineItem:input_type -> DeleteServiceLineItemGrpcRequest
	3,  // 15: ServiceLineItemGrpcService.CloseServiceLineItem:input_type -> CloseServiceLineItemGrpcRequest
	4,  // 16: ServiceLineItemGrpcService.CreateServiceLineItem:output_type -> ServiceLineItemIdGrpcResponse
	4,  // 17: ServiceLineItemGrpcService.UpdateServiceLineItem:output_type -> ServiceLineItemIdGrpcResponse
	4,  // 18: ServiceLineItemGrpcService.DeleteServiceLineItem:output_type -> ServiceLineItemIdGrpcResponse
	4,  // 19: ServiceLineItemGrpcService.CloseServiceLineItem:output_type -> ServiceLineItemIdGrpcResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_service_line_item_proto_init() }
func file_service_line_item_proto_init() {
	if File_service_line_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_line_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceLineItemGrpcRequest); i {
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
		file_service_line_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServiceLineItemGrpcRequest); i {
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
		file_service_line_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceLineItemGrpcRequest); i {
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
		file_service_line_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseServiceLineItemGrpcRequest); i {
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
		file_service_line_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceLineItemIdGrpcResponse); i {
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
			RawDescriptor: file_service_line_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_line_item_proto_goTypes,
		DependencyIndexes: file_service_line_item_proto_depIdxs,
		MessageInfos:      file_service_line_item_proto_msgTypes,
	}.Build()
	File_service_line_item_proto = out.File
	file_service_line_item_proto_rawDesc = nil
	file_service_line_item_proto_goTypes = nil
	file_service_line_item_proto_depIdxs = nil
}
