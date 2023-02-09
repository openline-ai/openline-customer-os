// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: MessageStore.proto

package generated

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

type FeedItemState int32

const (
	FeedItemState_NEW         FeedItemState = 0
	FeedItemState_IN_PROGRESS FeedItemState = 1
	FeedItemState_CLOSED      FeedItemState = 2
)

// Enum value maps for FeedItemState.
var (
	FeedItemState_name = map[int32]string{
		0: "NEW",
		1: "IN_PROGRESS",
		2: "CLOSED",
	}
	FeedItemState_value = map[string]int32{
		"NEW":         0,
		"IN_PROGRESS": 1,
		"CLOSED":      2,
	}
)

func (x FeedItemState) Enum() *FeedItemState {
	p := new(FeedItemState)
	*p = x
	return p
}

func (x FeedItemState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemState) Descriptor() protoreflect.EnumDescriptor {
	return file_MessageStore_proto_enumTypes[0].Descriptor()
}

func (FeedItemState) Type() protoreflect.EnumType {
	return &file_MessageStore_proto_enumTypes[0]
}

func (x FeedItemState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemState.Descriptor instead.
func (FeedItemState) EnumDescriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{0}
}

type MessageListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages"`
}

func (x *MessageListResponse) Reset() {
	*x = MessageListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageListResponse) ProtoMessage() {}

func (x *MessageListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageListResponse.ProtoReflect.Descriptor instead.
func (*MessageListResponse) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{0}
}

func (x *MessageListResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ParticipantsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participants []string `protobuf:"bytes,1,rep,name=participants,proto3" json:"participants"`
}

func (x *ParticipantsListResponse) Reset() {
	*x = ParticipantsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantsListResponse) ProtoMessage() {}

func (x *ParticipantsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantsListResponse.ProtoReflect.Descriptor instead.
func (*ParticipantsListResponse) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{1}
}

func (x *ParticipantsListResponse) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type FeedId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *FeedId) Reset() {
	*x = FeedId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedId) ProtoMessage() {}

func (x *FeedId) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedId.ProtoReflect.Descriptor instead.
func (*FeedId) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{2}
}

func (x *FeedId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FeedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// used to produce a record in the feed
	InitiatorFirstName string `protobuf:"bytes,2,opt,name=initiatorFirstName,proto3" json:"initiatorFirstName"`
	InitiatorLastName  string `protobuf:"bytes,3,opt,name=initiatorLastName,proto3" json:"initiatorLastName"`
	InitiatorUsername  string `protobuf:"bytes,4,opt,name=initiatorUsername,proto3" json:"initiatorUsername"`
	InitiatorType      string `protobuf:"bytes,5,opt,name=initiatorType,proto3" json:"initiatorType"`
	// used to produce the preview in the feed
	LastSenderFirstName string                 `protobuf:"bytes,6,opt,name=lastSenderFirstName,proto3" json:"lastSenderFirstName"`
	LastSenderLastName  string                 `protobuf:"bytes,7,opt,name=lastSenderLastName,proto3" json:"lastSenderLastName"`
	LastContentPreview  string                 `protobuf:"bytes,8,opt,name=lastContentPreview,proto3" json:"lastContentPreview"`
	LastTimestamp       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=lastTimestamp,proto3" json:"lastTimestamp"`
}

func (x *FeedItem) Reset() {
	*x = FeedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItem.ProtoReflect.Descriptor instead.
func (*FeedItem) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{3}
}

func (x *FeedItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FeedItem) GetInitiatorFirstName() string {
	if x != nil {
		return x.InitiatorFirstName
	}
	return ""
}

func (x *FeedItem) GetInitiatorLastName() string {
	if x != nil {
		return x.InitiatorLastName
	}
	return ""
}

func (x *FeedItem) GetInitiatorUsername() string {
	if x != nil {
		return x.InitiatorUsername
	}
	return ""
}

func (x *FeedItem) GetInitiatorType() string {
	if x != nil {
		return x.InitiatorType
	}
	return ""
}

func (x *FeedItem) GetLastSenderFirstName() string {
	if x != nil {
		return x.LastSenderFirstName
	}
	return ""
}

func (x *FeedItem) GetLastSenderLastName() string {
	if x != nil {
		return x.LastSenderLastName
	}
	return ""
}

func (x *FeedItem) GetLastContentPreview() string {
	if x != nil {
		return x.LastContentPreview
	}
	return ""
}

func (x *FeedItem) GetLastTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTimestamp
	}
	return nil
}

type FeedItemPagedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedItems     []*FeedItem `protobuf:"bytes,1,rep,name=feedItems,proto3" json:"feedItems"`
	TotalElements int32       `protobuf:"varint,2,opt,name=totalElements,proto3" json:"totalElements"`
}

func (x *FeedItemPagedResponse) Reset() {
	*x = FeedItemPagedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemPagedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemPagedResponse) ProtoMessage() {}

func (x *FeedItemPagedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemPagedResponse.ProtoReflect.Descriptor instead.
func (*FeedItemPagedResponse) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{4}
}

func (x *FeedItemPagedResponse) GetFeedItems() []*FeedItem {
	if x != nil {
		return x.FeedItems
	}
	return nil
}

func (x *FeedItemPagedResponse) GetTotalElements() int32 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

type GetFeedsPagedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateIn  []FeedItemState `protobuf:"varint,1,rep,packed,name=stateIn,proto3,enum=proto.FeedItemState" json:"stateIn"`
	Page     int32           `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	PageSize int32           `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *GetFeedsPagedRequest) Reset() {
	*x = GetFeedsPagedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageStore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsPagedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsPagedRequest) ProtoMessage() {}

func (x *GetFeedsPagedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MessageStore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsPagedRequest.ProtoReflect.Descriptor instead.
func (*GetFeedsPagedRequest) Descriptor() ([]byte, []int) {
	return file_MessageStore_proto_rawDescGZIP(), []int{5}
}

func (x *GetFeedsPagedRequest) GetStateIn() []FeedItemState {
	if x != nil {
		return x.StateIn
	}
	return nil
}

func (x *GetFeedsPagedRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFeedsPagedRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_MessageStore_proto protoreflect.FileDescriptor

var file_MessageStore_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x13, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3e,
	0x0a, 0x18, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x18,
	0x0a, 0x06, 0x46, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa0, 0x03, 0x0a, 0x08, 0x46, 0x65, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x40, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6c, 0x0a, 0x15, 0x46,
	0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x73, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x2a, 0x35, 0x0a, 0x0d, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02, 0x32, 0xfd, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x50, 0x61, 0x67,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x67, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x73,
	0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d,
	0x61, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_MessageStore_proto_rawDescOnce sync.Once
	file_MessageStore_proto_rawDescData = file_MessageStore_proto_rawDesc
)

func file_MessageStore_proto_rawDescGZIP() []byte {
	file_MessageStore_proto_rawDescOnce.Do(func() {
		file_MessageStore_proto_rawDescData = protoimpl.X.CompressGZIP(file_MessageStore_proto_rawDescData)
	})
	return file_MessageStore_proto_rawDescData
}

var file_MessageStore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MessageStore_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_MessageStore_proto_goTypes = []interface{}{
	(FeedItemState)(0),               // 0: proto.FeedItemState
	(*MessageListResponse)(nil),      // 1: proto.MessageListResponse
	(*ParticipantsListResponse)(nil), // 2: proto.ParticipantsListResponse
	(*FeedId)(nil),                   // 3: proto.FeedId
	(*FeedItem)(nil),                 // 4: proto.FeedItem
	(*FeedItemPagedResponse)(nil),    // 5: proto.FeedItemPagedResponse
	(*GetFeedsPagedRequest)(nil),     // 6: proto.GetFeedsPagedRequest
	(*Message)(nil),                  // 7: proto.Message
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*MessageId)(nil),                // 9: proto.MessageId
	(*InputMessage)(nil),             // 10: proto.InputMessage
}
var file_MessageStore_proto_depIdxs = []int32{
	7,  // 0: proto.MessageListResponse.messages:type_name -> proto.Message
	8,  // 1: proto.FeedItem.lastTimestamp:type_name -> google.protobuf.Timestamp
	4,  // 2: proto.FeedItemPagedResponse.feedItems:type_name -> proto.FeedItem
	0,  // 3: proto.GetFeedsPagedRequest.stateIn:type_name -> proto.FeedItemState
	6,  // 4: proto.MessageStoreService.getFeeds:input_type -> proto.GetFeedsPagedRequest
	3,  // 5: proto.MessageStoreService.getFeed:input_type -> proto.FeedId
	3,  // 6: proto.MessageStoreService.getMessagesForFeed:input_type -> proto.FeedId
	9,  // 7: proto.MessageStoreService.getMessage:input_type -> proto.MessageId
	10, // 8: proto.MessageStoreService.saveMessage:input_type -> proto.InputMessage
	3,  // 9: proto.MessageStoreService.getParticipants:input_type -> proto.FeedId
	5,  // 10: proto.MessageStoreService.getFeeds:output_type -> proto.FeedItemPagedResponse
	4,  // 11: proto.MessageStoreService.getFeed:output_type -> proto.FeedItem
	1,  // 12: proto.MessageStoreService.getMessagesForFeed:output_type -> proto.MessageListResponse
	7,  // 13: proto.MessageStoreService.getMessage:output_type -> proto.Message
	9,  // 14: proto.MessageStoreService.saveMessage:output_type -> proto.MessageId
	2,  // 15: proto.MessageStoreService.getParticipants:output_type -> proto.ParticipantsListResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_MessageStore_proto_init() }
func file_MessageStore_proto_init() {
	if File_MessageStore_proto != nil {
		return
	}
	file_Message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MessageStore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageListResponse); i {
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
		file_MessageStore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantsListResponse); i {
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
		file_MessageStore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedId); i {
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
		file_MessageStore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItem); i {
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
		file_MessageStore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemPagedResponse); i {
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
		file_MessageStore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsPagedRequest); i {
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
			RawDescriptor: file_MessageStore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MessageStore_proto_goTypes,
		DependencyIndexes: file_MessageStore_proto_depIdxs,
		EnumInfos:         file_MessageStore_proto_enumTypes,
		MessageInfos:      file_MessageStore_proto_msgTypes,
	}.Build()
	File_MessageStore_proto = out.File
	file_MessageStore_proto_rawDesc = nil
	file_MessageStore_proto_goTypes = nil
	file_MessageStore_proto_depIdxs = nil
}
