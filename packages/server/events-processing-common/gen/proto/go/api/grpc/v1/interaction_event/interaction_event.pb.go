// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: interaction_event.proto

package interaction_event_grpc_service

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

type UpsertInteractionEventGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant               string                       `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	LoggedInUserId       string                       `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	Content              string                       `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ContentType          string                       `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Identifier           string                       `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Channel              string                       `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	ChannelData          string                       `protobuf:"bytes,8,opt,name=channelData,proto3" json:"channelData,omitempty"`
	EventType            string                       `protobuf:"bytes,9,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Hide                 string                       `protobuf:"bytes,10,opt,name=hide,proto3" json:"hide,omitempty"`
	CreatedAt            *timestamppb.Timestamp       `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp       `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	PartOfIssueId        *string                      `protobuf:"bytes,13,opt,name=partOfIssueId,proto3,oneof" json:"partOfIssueId,omitempty"`
	PartOfSessionId      *string                      `protobuf:"bytes,14,opt,name=partOfSessionId,proto3,oneof" json:"partOfSessionId,omitempty"`
	SourceFields         *common.SourceFields         `protobuf:"bytes,15,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	ExternalSystemFields *common.ExternalSystemFields `protobuf:"bytes,16,opt,name=externalSystemFields,proto3" json:"externalSystemFields,omitempty"`
}

func (x *UpsertInteractionEventGrpcRequest) Reset() {
	*x = UpsertInteractionEventGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertInteractionEventGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertInteractionEventGrpcRequest) ProtoMessage() {}

func (x *UpsertInteractionEventGrpcRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpsertInteractionEventGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertInteractionEventGrpcRequest) Descriptor() ([]byte, []int) {
	return file_interaction_event_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertInteractionEventGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetChannelData() string {
	if x != nil {
		return x.ChannelData
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetHide() string {
	if x != nil {
		return x.Hide
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertInteractionEventGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertInteractionEventGrpcRequest) GetPartOfIssueId() string {
	if x != nil && x.PartOfIssueId != nil {
		return *x.PartOfIssueId
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetPartOfSessionId() string {
	if x != nil && x.PartOfSessionId != nil {
		return *x.PartOfSessionId
	}
	return ""
}

func (x *UpsertInteractionEventGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *UpsertInteractionEventGrpcRequest) GetExternalSystemFields() *common.ExternalSystemFields {
	if x != nil {
		return x.ExternalSystemFields
	}
	return nil
}

type RequestGenerateSummaryGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant             string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	InteractionEventId string `protobuf:"bytes,2,opt,name=interactionEventId,proto3" json:"interactionEventId,omitempty"`
	LoggedInUserId     string `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *RequestGenerateSummaryGrpcRequest) Reset() {
	*x = RequestGenerateSummaryGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGenerateSummaryGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGenerateSummaryGrpcRequest) ProtoMessage() {}

func (x *RequestGenerateSummaryGrpcRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RequestGenerateSummaryGrpcRequest.ProtoReflect.Descriptor instead.
func (*RequestGenerateSummaryGrpcRequest) Descriptor() ([]byte, []int) {
	return file_interaction_event_proto_rawDescGZIP(), []int{1}
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

func (x *RequestGenerateSummaryGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type RequestGenerateActionItemsGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant             string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	InteractionEventId string `protobuf:"bytes,2,opt,name=interactionEventId,proto3" json:"interactionEventId,omitempty"`
	LoggedInUserId     string `protobuf:"bytes,3,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *RequestGenerateActionItemsGrpcRequest) Reset() {
	*x = RequestGenerateActionItemsGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGenerateActionItemsGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGenerateActionItemsGrpcRequest) ProtoMessage() {}

func (x *RequestGenerateActionItemsGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGenerateActionItemsGrpcRequest.ProtoReflect.Descriptor instead.
func (*RequestGenerateActionItemsGrpcRequest) Descriptor() ([]byte, []int) {
	return file_interaction_event_proto_rawDescGZIP(), []int{2}
}

func (x *RequestGenerateActionItemsGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *RequestGenerateActionItemsGrpcRequest) GetInteractionEventId() string {
	if x != nil {
		return x.InteractionEventId
	}
	return ""
}

func (x *RequestGenerateActionItemsGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
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
		mi := &file_interaction_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionEventIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionEventIdGrpcResponse) ProtoMessage() {}

func (x *InteractionEventIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_event_proto_msgTypes[3]
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
	return file_interaction_event_proto_rawDescGZIP(), []int{3}
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
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x05,
	0x0a, 0x21, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x64,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x64, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x29, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74,
	0x4f, 0x66, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x49,
	0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x4f, 0x66, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x93, 0x01, 0x0a, 0x21, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x25, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x1e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xc2, 0x02, 0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5d, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x65, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x16, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x42, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_interaction_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_interaction_event_proto_goTypes = []interface{}{
	(*UpsertInteractionEventGrpcRequest)(nil),     // 0: UpsertInteractionEventGrpcRequest
	(*RequestGenerateSummaryGrpcRequest)(nil),     // 1: RequestGenerateSummaryGrpcRequest
	(*RequestGenerateActionItemsGrpcRequest)(nil), // 2: RequestGenerateActionItemsGrpcRequest
	(*InteractionEventIdGrpcResponse)(nil),        // 3: InteractionEventIdGrpcResponse
	(*timestamppb.Timestamp)(nil),                 // 4: google.protobuf.Timestamp
	(*common.SourceFields)(nil),                   // 5: SourceFields
	(*common.ExternalSystemFields)(nil),           // 6: ExternalSystemFields
}
var file_interaction_event_proto_depIdxs = []int32{
	4, // 0: UpsertInteractionEventGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: UpsertInteractionEventGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	5, // 2: UpsertInteractionEventGrpcRequest.sourceFields:type_name -> SourceFields
	6, // 3: UpsertInteractionEventGrpcRequest.externalSystemFields:type_name -> ExternalSystemFields
	1, // 4: interactionEventGrpcService.RequestGenerateSummary:input_type -> RequestGenerateSummaryGrpcRequest
	2, // 5: interactionEventGrpcService.RequestGenerateActionItems:input_type -> RequestGenerateActionItemsGrpcRequest
	0, // 6: interactionEventGrpcService.UpsertInteractionEvent:input_type -> UpsertInteractionEventGrpcRequest
	3, // 7: interactionEventGrpcService.RequestGenerateSummary:output_type -> InteractionEventIdGrpcResponse
	3, // 8: interactionEventGrpcService.RequestGenerateActionItems:output_type -> InteractionEventIdGrpcResponse
	3, // 9: interactionEventGrpcService.UpsertInteractionEvent:output_type -> InteractionEventIdGrpcResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_interaction_event_proto_init() }
func file_interaction_event_proto_init() {
	if File_interaction_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interaction_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertInteractionEventGrpcRequest); i {
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
		file_interaction_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGenerateActionItemsGrpcRequest); i {
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
		file_interaction_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	file_interaction_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interaction_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
