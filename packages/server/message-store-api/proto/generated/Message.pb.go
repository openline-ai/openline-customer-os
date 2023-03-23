// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: MessageDeprecate.proto

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
	return file_Message_proto_enumTypes[0].Descriptor()
}

func (SenderType) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[0]
}

func (x SenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SenderType.Descriptor instead.
func (SenderType) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

type MessageSubtype int32

const (
	MessageSubtype_MESSAGE MessageSubtype = 0
	MessageSubtype_FILE    MessageSubtype = 1
)

// Enum value maps for MessageSubtype.
var (
	MessageSubtype_name = map[int32]string{
		0: "MESSAGE",
		1: "FILE",
	}
	MessageSubtype_value = map[string]int32{
		"MESSAGE": 0,
		"FILE":    1,
	}
)

func (x MessageSubtype) Enum() *MessageSubtype {
	p := new(MessageSubtype)
	*p = x
	return p
}

func (x MessageSubtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageSubtype) Descriptor() protoreflect.EnumDescriptor {
	return file_Message_proto_enumTypes[1].Descriptor()
}

func (MessageSubtype) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[1]
}

func (x MessageSubtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageSubtype.Descriptor instead.
func (MessageSubtype) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
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
	return file_Message_proto_enumTypes[2].Descriptor()
}

func (MessageDirection) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[2]
}

func (x MessageDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageDirection.Descriptor instead.
func (MessageDirection) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{2}
}

type MessageType int32

const (
	MessageType_WEB_CHAT MessageType = 0
	MessageType_EMAIL    MessageType = 1
	MessageType_VOICE    MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "WEB_CHAT",
		1: "EMAIL",
		2: "VOICE",
	}
	MessageType_value = map[string]int32{
		"WEB_CHAT": 0,
		"EMAIL":    1,
		"VOICE":    2,
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
	return file_Message_proto_enumTypes[3].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[3]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{3}
}

type ParticipantIdType int32

const (
	ParticipantIdType_MAILTO ParticipantIdType = 0
	ParticipantIdType_TEL    ParticipantIdType = 1
)

// Enum value maps for ParticipantIdType.
var (
	ParticipantIdType_name = map[int32]string{
		0: "MAILTO",
		1: "TEL",
	}
	ParticipantIdType_value = map[string]int32{
		"MAILTO": 0,
		"TEL":    1,
	}
)

func (x ParticipantIdType) Enum() *ParticipantIdType {
	p := new(ParticipantIdType)
	*p = x
	return p
}

func (x ParticipantIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParticipantIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_Message_proto_enumTypes[4].Descriptor()
}

func (ParticipantIdType) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[4]
}

func (x ParticipantIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParticipantIdType.Descriptor instead.
func (ParticipantIdType) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{4}
}

type ParticipantId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       ParticipantIdType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ParticipantIdType" json:"type"`
	Identifier string            `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier"`
}

func (x *ParticipantId) Reset() {
	*x = ParticipantId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantId) ProtoMessage() {}

func (x *ParticipantId) ProtoReflect() protoreflect.MessageDeprecate {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantId.ProtoReflect.Descriptor instead.
func (*ParticipantId) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *ParticipantId) GetType() ParticipantIdType {
	if x != nil {
		return x.Type
	}
	return ParticipantIdType_MAILTO
}

func (x *ParticipantId) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	Type SenderType `protobuf:"varint,2,opt,name=Type,proto3,enum=proto.SenderType" json:"Type"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.MessageDeprecate {
	mi := &file_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
}

func (x *Participant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Participant) GetType() SenderType {
	if x != nil {
		return x.Type
	}
	return SenderType_CONTACT
}

type MessageId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// postgres id
	ConversationEventId string `protobuf:"bytes,1,opt,name=conversationEventId,proto3" json:"conversationEventId"`
	// neo4j id
	ConversationId string `protobuf:"bytes,2,opt,name=conversationId,proto3" json:"conversationId"`
}

func (x *MessageId) Reset() {
	*x = MessageId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageId) ProtoMessage() {}

func (x *MessageId) ProtoReflect() protoreflect.MessageDeprecate {
	mi := &file_Message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageId.ProtoReflect.Descriptor instead.
func (*MessageId) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{2}
}

func (x *MessageId) GetConversationEventId() string {
	if x != nil {
		return x.ConversationEventId
	}
	return ""
}

func (x *MessageId) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

type MessageDeprecate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId         *MessageId             `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId"`
	InitiatorUsername *ParticipantId         `protobuf:"bytes,2,opt,name=initiatorUsername,proto3" json:"initiatorUsername"`
	Type              MessageType            `protobuf:"varint,3,opt,name=type,proto3,enum=proto.MessageType" json:"type"`
	Subtype           MessageSubtype         `protobuf:"varint,4,opt,name=subtype,proto3,enum=proto.MessageSubtype" json:"subtype"`
	Content           string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`
	Direction         MessageDirection       `protobuf:"varint,6,opt,name=direction,proto3,enum=proto.MessageDirection" json:"direction"`
	Time              *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=time,proto3" json:"time"`
	SenderType        SenderType             `protobuf:"varint,8,opt,name=senderType,proto3,enum=proto.SenderType" json:"senderType"`
	SenderId          string                 `protobuf:"bytes,9,opt,name=senderId,proto3" json:"senderId"`
	SenderUsername    *ParticipantId         `protobuf:"bytes,10,opt,name=senderUsername,proto3" json:"senderUsername"`
}

func (x *MessageDeprecate) Reset() {
	*x = MessageDeprecate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDeprecate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDeprecate) ProtoMessage() {}

func (x *MessageDeprecate) ProtoReflect() protoreflect.MessageDeprecate {
	mi := &file_Message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDeprecate.ProtoReflect.Descriptor instead.
func (*MessageDeprecate) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{3}
}

func (x *MessageDeprecate) GetMessageId() *MessageId {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *MessageDeprecate) GetInitiatorUsername() *ParticipantId {
	if x != nil {
		return x.InitiatorUsername
	}
	return nil
}

func (x *MessageDeprecate) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_WEB_CHAT
}

func (x *MessageDeprecate) GetSubtype() MessageSubtype {
	if x != nil {
		return x.Subtype
	}
	return MessageSubtype_MESSAGE
}

func (x *MessageDeprecate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageDeprecate) GetDirection() MessageDirection {
	if x != nil {
		return x.Direction
	}
	return MessageDirection_INBOUND
}

func (x *MessageDeprecate) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MessageDeprecate) GetSenderType() SenderType {
	if x != nil {
		return x.SenderType
	}
	return SenderType_CONTACT
}

func (x *MessageDeprecate) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *MessageDeprecate) GetSenderUsername() *ParticipantId {
	if x != nil {
		return x.SenderUsername
	}
	return nil
}

type InputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId          *string          `protobuf:"bytes,1,opt,name=conversationId,proto3,oneof" json:"conversationId"`
	InitiatorIdentifier     *ParticipantId   `protobuf:"bytes,2,opt,name=initiatorIdentifier,proto3,oneof" json:"initiatorIdentifier"`
	Type                    MessageType      `protobuf:"varint,3,opt,name=type,proto3,enum=proto.MessageType" json:"type"`
	Subtype                 MessageSubtype   `protobuf:"varint,4,opt,name=subtype,proto3,enum=proto.MessageSubtype" json:"subtype"`
	Content                 *string          `protobuf:"bytes,5,opt,name=content,proto3,oneof" json:"content"`
	Direction               MessageDirection `protobuf:"varint,6,opt,name=direction,proto3,enum=proto.MessageDirection" json:"direction"`
	ParticipantsIdentifiers []*ParticipantId `protobuf:"bytes,7,rep,name=participantsIdentifiers,proto3" json:"participantsIdentifiers"`
	ThreadId                *string          `protobuf:"bytes,8,opt,name=threadId,proto3,oneof" json:"threadId"`
}

func (x *InputMessage) Reset() {
	*x = InputMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputMessage) ProtoMessage() {}

func (x *InputMessage) ProtoReflect() protoreflect.MessageDeprecate {
	mi := &file_Message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputMessage.ProtoReflect.Descriptor instead.
func (*InputMessage) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{4}
}

func (x *InputMessage) GetConversationId() string {
	if x != nil && x.ConversationId != nil {
		return *x.ConversationId
	}
	return ""
}

func (x *InputMessage) GetInitiatorIdentifier() *ParticipantId {
	if x != nil {
		return x.InitiatorIdentifier
	}
	return nil
}

func (x *InputMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_WEB_CHAT
}

func (x *InputMessage) GetSubtype() MessageSubtype {
	if x != nil {
		return x.Subtype
	}
	return MessageSubtype_MESSAGE
}

func (x *InputMessage) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *InputMessage) GetDirection() MessageDirection {
	if x != nil {
		return x.Direction
	}
	return MessageDirection_INBOUND
}

func (x *InputMessage) GetParticipantsIdentifiers() []*ParticipantId {
	if x != nil {
		return x.ParticipantsIdentifiers
	}
	return nil
}

func (x *InputMessage) GetThreadId() string {
	if x != nil && x.ThreadId != nil {
		return *x.ThreadId
	}
	return ""
}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x65, 0x0a, 0x09,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xe4, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x42, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xec, 0x03, 0x0a, 0x0c, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x13, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x48, 0x01, 0x52, 0x13, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x75,
	0x62, 0x74, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x17, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x2a, 0x23, 0x0a, 0x0a, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x43, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x27,
	0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x2a, 0x2d, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x45, 0x42, 0x5f, 0x43, 0x48, 0x41,
	0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x02, 0x2a, 0x28, 0x0a, 0x11, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x41, 0x49, 0x4c, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45,
	0x4c, 0x10, 0x01, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x61, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d,
	0x6f, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Message_proto_goTypes = []interface{}{
	(SenderType)(0),               // 0: proto.SenderType
	(MessageSubtype)(0),           // 1: proto.MessageSubtype
	(MessageDirection)(0),         // 2: proto.MessageDirection
	(MessageType)(0),              // 3: proto.MessageType
	(ParticipantIdType)(0),        // 4: proto.ParticipantIdType
	(*ParticipantId)(nil),         // 5: proto.ParticipantId
	(*Participant)(nil),           // 6: proto.Participant
	(*MessageId)(nil),             // 7: proto.MessageId
	(*MessageDeprecate)(nil),               // 8: proto.MessageDeprecate
	(*InputMessage)(nil),          // 9: proto.InputMessage
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_Message_proto_depIdxs = []int32{
	4,  // 0: proto.ParticipantId.type:type_name -> proto.ParticipantIdType
	0,  // 1: proto.Participant.Type:type_name -> proto.SenderType
	7,  // 2: proto.MessageDeprecate.messageId:type_name -> proto.MessageId
	5,  // 3: proto.MessageDeprecate.initiatorUsername:type_name -> proto.ParticipantId
	3,  // 4: proto.MessageDeprecate.type:type_name -> proto.MessageType
	1,  // 5: proto.MessageDeprecate.subtype:type_name -> proto.MessageSubtype
	2,  // 6: proto.MessageDeprecate.direction:type_name -> proto.MessageDirection
	10, // 7: proto.MessageDeprecate.time:type_name -> google.protobuf.Timestamp
	0,  // 8: proto.MessageDeprecate.senderType:type_name -> proto.SenderType
	5,  // 9: proto.MessageDeprecate.senderUsername:type_name -> proto.ParticipantId
	5,  // 10: proto.InputMessage.initiatorIdentifier:type_name -> proto.ParticipantId
	3,  // 11: proto.InputMessage.type:type_name -> proto.MessageType
	1,  // 12: proto.InputMessage.subtype:type_name -> proto.MessageSubtype
	2,  // 13: proto.InputMessage.direction:type_name -> proto.MessageDirection
	5,  // 14: proto.InputMessage.participantsIdentifiers:type_name -> proto.ParticipantId
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantId); i {
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
		file_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
		file_Message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageId); i {
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
		file_Message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDeprecate); i {
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
		file_Message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputMessage); i {
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
	file_Message_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		EnumInfos:         file_Message_proto_enumTypes,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}
