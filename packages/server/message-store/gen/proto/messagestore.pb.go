// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: messagestore.proto

package proto

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
	return file_messagestore_proto_enumTypes[0].Descriptor()
}

func (FeedItemState) Type() protoreflect.EnumType {
	return &file_messagestore_proto_enumTypes[0]
}

func (x FeedItemState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemState.Descriptor instead.
func (FeedItemState) EnumDescriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{0}
}

type SenderType int32

const (
	SenderType_CONTACT SenderType = 0
	SenderType_USER    SenderType = 1
)

// Enum value maps for SenderType.
var (
	SenderType_name = map[int32]string{
		0: "CONTACT",
		1: "USER",
	}
	SenderType_value = map[string]int32{
		"CONTACT": 0,
		"USER":    1,
	}
)

func (x SenderType) Enum() *SenderType {
	p := new(SenderType)
	*p = x
	return p
}

func (x SenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_messagestore_proto_enumTypes[1].Descriptor()
}

func (SenderType) Type() protoreflect.EnumType {
	return &file_messagestore_proto_enumTypes[1]
}

func (x SenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SenderType.Descriptor instead.
func (SenderType) EnumDescriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{1}
}

type MessageType int32

const (
	MessageType_MESSAGE MessageType = 0
	MessageType_FILE    MessageType = 1
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MESSAGE",
		1: "FILE",
	}
	MessageType_value = map[string]int32{
		"MESSAGE": 0,
		"FILE":    1,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messagestore_proto_enumTypes[2].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_messagestore_proto_enumTypes[2]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{2}
}

type MessageDirection int32

const (
	MessageDirection_INBOUND  MessageDirection = 0
	MessageDirection_OUTBOUND MessageDirection = 1
)

// Enum value maps for MessageDirection.
var (
	MessageDirection_name = map[int32]string{
		0: "INBOUND",
		1: "OUTBOUND",
	}
	MessageDirection_value = map[string]int32{
		"INBOUND":  0,
		"OUTBOUND": 1,
	}
)

func (x MessageDirection) Enum() *MessageDirection {
	p := new(MessageDirection)
	*p = x
	return p
}

func (x MessageDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_messagestore_proto_enumTypes[3].Descriptor()
}

func (MessageDirection) Type() protoreflect.EnumType {
	return &file_messagestore_proto_enumTypes[3]
}

func (x MessageDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageDirection.Descriptor instead.
func (MessageDirection) EnumDescriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{3}
}

type MessageChannel int32

const (
	MessageChannel_WIDGET   MessageChannel = 0
	MessageChannel_MAIL     MessageChannel = 1
	MessageChannel_WHATSAPP MessageChannel = 2
	MessageChannel_FACEBOOK MessageChannel = 3
	MessageChannel_TWITTER  MessageChannel = 4
	MessageChannel_VOICE    MessageChannel = 5
)

// Enum value maps for MessageChannel.
var (
	MessageChannel_name = map[int32]string{
		0: "WIDGET",
		1: "MAIL",
		2: "WHATSAPP",
		3: "FACEBOOK",
		4: "TWITTER",
		5: "VOICE",
	}
	MessageChannel_value = map[string]int32{
		"WIDGET":   0,
		"MAIL":     1,
		"WHATSAPP": 2,
		"FACEBOOK": 3,
		"TWITTER":  4,
		"VOICE":    5,
	}
)

func (x MessageChannel) Enum() *MessageChannel {
	p := new(MessageChannel)
	*p = x
	return p
}

func (x MessageChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_messagestore_proto_enumTypes[4].Descriptor()
}

func (MessageChannel) Type() protoreflect.EnumType {
	return &file_messagestore_proto_enumTypes[4]
}

func (x MessageChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageChannel.Descriptor instead.
func (MessageChannel) EnumDescriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{4}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      MessageType            `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type"`
	Message   string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Channel   MessageChannel         `protobuf:"varint,3,opt,name=channel,proto3,enum=MessageChannel" json:"channel"`
	Direction MessageDirection       `protobuf:"varint,4,opt,name=direction,proto3,enum=MessageDirection" json:"direction"`
	Time      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3,oneof" json:"time"`
	Id        *int64                 `protobuf:"varint,6,opt,name=id,proto3,oneof" json:"id"`
	FeedId    *int64                 `protobuf:"varint,7,opt,name=feedId,proto3,oneof" json:"feedId"`
	Username  *string                `protobuf:"bytes,8,opt,name=username,proto3,oneof" json:"username"`
	UserId    *string                `protobuf:"bytes,9,opt,name=userId,proto3,oneof" json:"userId"`
	ContactId *string                `protobuf:"bytes,10,opt,name=contactId,proto3,oneof" json:"contactId"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_MESSAGE
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetChannel() MessageChannel {
	if x != nil {
		return x.Channel
	}
	return MessageChannel_WIDGET
}

func (x *Message) GetDirection() MessageDirection {
	if x != nil {
		return x.Direction
	}
	return MessageDirection_INBOUND
}

func (x *Message) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Message) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Message) GetFeedId() int64 {
	if x != nil && x.FeedId != nil {
		return *x.FeedId
	}
	return 0
}

func (x *Message) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *Message) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *Message) GetContactId() string {
	if x != nil && x.ContactId != nil {
		return *x.ContactId
	}
	return ""
}

type MessagePagedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  []*Message             `protobuf:"bytes,1,rep,name=message,proto3" json:"message"`
	Before   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before,proto3,oneof" json:"before"`
	PageSize int32                  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *MessagePagedResponse) Reset() {
	*x = MessagePagedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePagedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePagedResponse) ProtoMessage() {}

func (x *MessagePagedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePagedResponse.ProtoReflect.Descriptor instead.
func (*MessagePagedResponse) Descriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{1}
}

func (x *MessagePagedResponse) GetMessage() []*Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MessagePagedResponse) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *MessagePagedResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FeedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ContactId        string                 `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId"`
	ContactFirstName string                 `protobuf:"bytes,3,opt,name=contactFirstName,proto3" json:"contactFirstName"`
	ContactLastName  string                 `protobuf:"bytes,4,opt,name=contactLastName,proto3" json:"contactLastName"`
	ContactEmail     string                 `protobuf:"bytes,5,opt,name=contactEmail,proto3" json:"contactEmail"`
	State            FeedItemState          `protobuf:"varint,6,opt,name=state,proto3,enum=FeedItemState" json:"state"`
	LastSenderId     string                 `protobuf:"bytes,7,opt,name=lastSenderId,proto3" json:"lastSenderId"`
	LastSenderType   SenderType             `protobuf:"varint,8,opt,name=lastSenderType,proto3,enum=SenderType" json:"lastSenderType"`
	Message          string                 `protobuf:"bytes,9,opt,name=message,proto3" json:"message"`
	UpdatedOn        *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updatedOn,proto3" json:"updatedOn"`
}

func (x *FeedItem) Reset() {
	*x = FeedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[2]
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
	return file_messagestore_proto_rawDescGZIP(), []int{2}
}

func (x *FeedItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedItem) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *FeedItem) GetContactFirstName() string {
	if x != nil {
		return x.ContactFirstName
	}
	return ""
}

func (x *FeedItem) GetContactLastName() string {
	if x != nil {
		return x.ContactLastName
	}
	return ""
}

func (x *FeedItem) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *FeedItem) GetState() FeedItemState {
	if x != nil {
		return x.State
	}
	return FeedItemState_NEW
}

func (x *FeedItem) GetLastSenderId() string {
	if x != nil {
		return x.LastSenderId
	}
	return ""
}

func (x *FeedItem) GetLastSenderType() SenderType {
	if x != nil {
		return x.LastSenderType
	}
	return SenderType_CONTACT
}

func (x *FeedItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FeedItem) GetUpdatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedOn
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
		mi := &file_messagestore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemPagedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemPagedResponse) ProtoMessage() {}

func (x *FeedItemPagedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[3]
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
	return file_messagestore_proto_rawDescGZIP(), []int{3}
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId int64                  `protobuf:"varint,1,opt,name=conversationId,proto3" json:"conversationId"`
	Before         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before,proto3,oneof" json:"before"`
	PageSize       int32                  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_messagestore_proto_rawDescGZIP(), []int{5}
}

func (x *GetMessagesRequest) GetConversationId() int64 {
	if x != nil {
		return x.ConversationId
	}
	return 0
}

func (x *GetMessagesRequest) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *GetMessagesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetFeedsPagedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateIn  []FeedItemState `protobuf:"varint,1,rep,packed,name=stateIn,proto3,enum=FeedItemState" json:"stateIn"`
	Page     int32           `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	PageSize int32           `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *GetFeedsPagedRequest) Reset() {
	*x = GetFeedsPagedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagestore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsPagedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsPagedRequest) ProtoMessage() {}

func (x *GetFeedsPagedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagestore_proto_msgTypes[6]
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
	return file_messagestore_proto_rawDescGZIP(), []int{6}
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

var File_messagestore_proto protoreflect.FileDescriptor

var file_messagestore_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x13,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22,
	0x85, 0x03, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x66, 0x0a, 0x15, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09,
	0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73,
	0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0x35, 0x0a, 0x0d, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x23, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x01, 0x2a, 0x24, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x2a, 0x2d, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x5a, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x49, 0x44,
	0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x57, 0x48, 0x41, 0x54, 0x53, 0x41, 0x50, 0x50, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x41, 0x43, 0x45, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x57, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x4f, 0x49, 0x43,
	0x45, 0x10, 0x05, 0x32, 0xf0, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x73,
	0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x13, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x64,
	0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08,
	0x67, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x73, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x61, 0x69,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2d, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagestore_proto_rawDescOnce sync.Once
	file_messagestore_proto_rawDescData = file_messagestore_proto_rawDesc
)

func file_messagestore_proto_rawDescGZIP() []byte {
	file_messagestore_proto_rawDescOnce.Do(func() {
		file_messagestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagestore_proto_rawDescData)
	})
	return file_messagestore_proto_rawDescData
}

var file_messagestore_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_messagestore_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_messagestore_proto_goTypes = []interface{}{
	(FeedItemState)(0),            // 0: FeedItemState
	(SenderType)(0),               // 1: SenderType
	(MessageType)(0),              // 2: MessageType
	(MessageDirection)(0),         // 3: MessageDirection
	(MessageChannel)(0),           // 4: MessageChannel
	(*Message)(nil),               // 5: Message
	(*MessagePagedResponse)(nil),  // 6: MessagePagedResponse
	(*FeedItem)(nil),              // 7: FeedItem
	(*FeedItemPagedResponse)(nil), // 8: FeedItemPagedResponse
	(*Id)(nil),                    // 9: Id
	(*GetMessagesRequest)(nil),    // 10: GetMessagesRequest
	(*GetFeedsPagedRequest)(nil),  // 11: GetFeedsPagedRequest
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_messagestore_proto_depIdxs = []int32{
	2,  // 0: Message.type:type_name -> MessageType
	4,  // 1: Message.channel:type_name -> MessageChannel
	3,  // 2: Message.direction:type_name -> MessageDirection
	12, // 3: Message.time:type_name -> google.protobuf.Timestamp
	5,  // 4: MessagePagedResponse.message:type_name -> Message
	12, // 5: MessagePagedResponse.before:type_name -> google.protobuf.Timestamp
	0,  // 6: FeedItem.state:type_name -> FeedItemState
	1,  // 7: FeedItem.lastSenderType:type_name -> SenderType
	12, // 8: FeedItem.updatedOn:type_name -> google.protobuf.Timestamp
	7,  // 9: FeedItemPagedResponse.feedItems:type_name -> FeedItem
	12, // 10: GetMessagesRequest.before:type_name -> google.protobuf.Timestamp
	0,  // 11: GetFeedsPagedRequest.stateIn:type_name -> FeedItemState
	5,  // 12: MessageStoreService.saveMessage:input_type -> Message
	10, // 13: MessageStoreService.getMessages:input_type -> GetMessagesRequest
	9,  // 14: MessageStoreService.getMessage:input_type -> Id
	11, // 15: MessageStoreService.getFeeds:input_type -> GetFeedsPagedRequest
	9,  // 16: MessageStoreService.getFeed:input_type -> Id
	5,  // 17: MessageStoreService.saveMessage:output_type -> Message
	6,  // 18: MessageStoreService.getMessages:output_type -> MessagePagedResponse
	5,  // 19: MessageStoreService.getMessage:output_type -> Message
	8,  // 20: MessageStoreService.getFeeds:output_type -> FeedItemPagedResponse
	7,  // 21: MessageStoreService.getFeed:output_type -> FeedItem
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_messagestore_proto_init() }
func file_messagestore_proto_init() {
	if File_messagestore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagestore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_messagestore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePagedResponse); i {
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
		file_messagestore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_messagestore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_messagestore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_messagestore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesRequest); i {
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
		file_messagestore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	file_messagestore_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_messagestore_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_messagestore_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messagestore_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messagestore_proto_goTypes,
		DependencyIndexes: file_messagestore_proto_depIdxs,
		EnumInfos:         file_messagestore_proto_enumTypes,
		MessageInfos:      file_messagestore_proto_msgTypes,
	}.Build()
	File_messagestore_proto = out.File
	file_messagestore_proto_rawDesc = nil
	file_messagestore_proto_goTypes = nil
	file_messagestore_proto_depIdxs = nil
}
