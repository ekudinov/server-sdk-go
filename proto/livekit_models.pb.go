// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: livekit_models.proto

package livekit

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

type TrackType int32

const (
	TrackType_AUDIO TrackType = 0
	TrackType_VIDEO TrackType = 1
	TrackType_DATA  TrackType = 2
)

// Enum value maps for TrackType.
var (
	TrackType_name = map[int32]string{
		0: "AUDIO",
		1: "VIDEO",
		2: "DATA",
	}
	TrackType_value = map[string]int32{
		"AUDIO": 0,
		"VIDEO": 1,
		"DATA":  2,
	}
)

func (x TrackType) Enum() *TrackType {
	p := new(TrackType)
	*p = x
	return p
}

func (x TrackType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackType) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_models_proto_enumTypes[0].Descriptor()
}

func (TrackType) Type() protoreflect.EnumType {
	return &file_livekit_models_proto_enumTypes[0]
}

func (x TrackType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackType.Descriptor instead.
func (TrackType) EnumDescriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{0}
}

type ParticipantInfo_State int32

const (
	// websocket' connected, but not offered yet
	ParticipantInfo_JOINING ParticipantInfo_State = 0
	// server received client offer
	ParticipantInfo_JOINED ParticipantInfo_State = 1
	// ICE connectivity established
	ParticipantInfo_ACTIVE ParticipantInfo_State = 2
	// WS disconnected
	ParticipantInfo_DISCONNECTED ParticipantInfo_State = 3
)

// Enum value maps for ParticipantInfo_State.
var (
	ParticipantInfo_State_name = map[int32]string{
		0: "JOINING",
		1: "JOINED",
		2: "ACTIVE",
		3: "DISCONNECTED",
	}
	ParticipantInfo_State_value = map[string]int32{
		"JOINING":      0,
		"JOINED":       1,
		"ACTIVE":       2,
		"DISCONNECTED": 3,
	}
)

func (x ParticipantInfo_State) Enum() *ParticipantInfo_State {
	p := new(ParticipantInfo_State)
	*p = x
	return p
}

func (x ParticipantInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParticipantInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_models_proto_enumTypes[1].Descriptor()
}

func (ParticipantInfo_State) Type() protoreflect.EnumType {
	return &file_livekit_models_proto_enumTypes[1]
}

func (x ParticipantInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParticipantInfo_State.Descriptor instead.
func (ParticipantInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{2, 0}
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid             string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EmptyTimeout    uint32   `protobuf:"varint,3,opt,name=empty_timeout,json=emptyTimeout,proto3" json:"empty_timeout,omitempty"`
	MaxParticipants uint32   `protobuf:"varint,4,opt,name=max_participants,json=maxParticipants,proto3" json:"max_participants,omitempty"`
	CreationTime    int64    `protobuf:"varint,5,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	TurnPassword    string   `protobuf:"bytes,6,opt,name=turn_password,json=turnPassword,proto3" json:"turn_password,omitempty"`
	EnabledCodecs   []*Codec `protobuf:"bytes,7,rep,name=enabled_codecs,json=enabledCodecs,proto3" json:"enabled_codecs,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetEmptyTimeout() uint32 {
	if x != nil {
		return x.EmptyTimeout
	}
	return 0
}

func (x *Room) GetMaxParticipants() uint32 {
	if x != nil {
		return x.MaxParticipants
	}
	return 0
}

func (x *Room) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *Room) GetTurnPassword() string {
	if x != nil {
		return x.TurnPassword
	}
	return ""
}

func (x *Room) GetEnabledCodecs() []*Codec {
	if x != nil {
		return x.EnabledCodecs
	}
	return nil
}

type Codec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime     string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	FmtpLine string `protobuf:"bytes,2,opt,name=fmtp_line,json=fmtpLine,proto3" json:"fmtp_line,omitempty"`
}

func (x *Codec) Reset() {
	*x = Codec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Codec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Codec) ProtoMessage() {}

func (x *Codec) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Codec.ProtoReflect.Descriptor instead.
func (*Codec) Descriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{1}
}

func (x *Codec) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *Codec) GetFmtpLine() string {
	if x != nil {
		return x.FmtpLine
	}
	return ""
}

type ParticipantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid      string                `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Identity string                `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	State    ParticipantInfo_State `protobuf:"varint,3,opt,name=state,proto3,enum=livekit.ParticipantInfo_State" json:"state,omitempty"`
	Tracks   []*TrackInfo          `protobuf:"bytes,4,rep,name=tracks,proto3" json:"tracks,omitempty"`
	Metadata string                `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// timestamp when participant joined room
	JoinedAt int64 `protobuf:"varint,6,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
}

func (x *ParticipantInfo) Reset() {
	*x = ParticipantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantInfo) ProtoMessage() {}

func (x *ParticipantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantInfo.ProtoReflect.Descriptor instead.
func (*ParticipantInfo) Descriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{2}
}

func (x *ParticipantInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *ParticipantInfo) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *ParticipantInfo) GetState() ParticipantInfo_State {
	if x != nil {
		return x.State
	}
	return ParticipantInfo_JOINING
}

func (x *ParticipantInfo) GetTracks() []*TrackInfo {
	if x != nil {
		return x.Tracks
	}
	return nil
}

func (x *ParticipantInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *ParticipantInfo) GetJoinedAt() int64 {
	if x != nil {
		return x.JoinedAt
	}
	return 0
}

type TrackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   string    `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Type  TrackType `protobuf:"varint,2,opt,name=type,proto3,enum=livekit.TrackType" json:"type,omitempty"`
	Name  string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Muted bool      `protobuf:"varint,4,opt,name=muted,proto3" json:"muted,omitempty"`
	// original width of video (unset for audio)
	// clients may receive a lower resolution version with simulcast
	Width uint32 `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	// original height of video (unset for audio)
	Height uint32 `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	// true if track is simulcasted
	Simulcast bool `protobuf:"varint,7,opt,name=simulcast,proto3" json:"simulcast,omitempty"`
}

func (x *TrackInfo) Reset() {
	*x = TrackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackInfo) ProtoMessage() {}

func (x *TrackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackInfo.ProtoReflect.Descriptor instead.
func (*TrackInfo) Descriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{3}
}

func (x *TrackInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *TrackInfo) GetType() TrackType {
	if x != nil {
		return x.Type
	}
	return TrackType_AUDIO
}

func (x *TrackInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrackInfo) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

func (x *TrackInfo) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *TrackInfo) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TrackInfo) GetSimulcast() bool {
	if x != nil {
		return x.Simulcast
	}
	return false
}

// old DataTrack message
type DataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*DataMessage_Text
	//	*DataMessage_Binary
	Value isDataMessage_Value `protobuf_oneof:"value"`
}

func (x *DataMessage) Reset() {
	*x = DataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMessage) ProtoMessage() {}

func (x *DataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMessage.ProtoReflect.Descriptor instead.
func (*DataMessage) Descriptor() ([]byte, []int) {
	return file_livekit_models_proto_rawDescGZIP(), []int{4}
}

func (m *DataMessage) GetValue() isDataMessage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *DataMessage) GetText() string {
	if x, ok := x.GetValue().(*DataMessage_Text); ok {
		return x.Text
	}
	return ""
}

func (x *DataMessage) GetBinary() []byte {
	if x, ok := x.GetValue().(*DataMessage_Binary); ok {
		return x.Binary
	}
	return nil
}

type isDataMessage_Value interface {
	isDataMessage_Value()
}

type DataMessage_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type DataMessage_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*DataMessage_Text) isDataMessage_Value() {}

func (*DataMessage_Binary) isDataMessage_Value() {}

var File_livekit_models_proto protoreflect.FileDescriptor

var file_livekit_models_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x22,
	0xfd, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d,
	0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x75, 0x72, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63,
	0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x22,
	0x38, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x6d, 0x74, 0x70, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x6d, 0x74, 0x70, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x9a, 0x02, 0x0a, 0x0f, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x4a, 0x4f, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x63,
	0x61, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x63, 0x61, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x2b, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x44,
	0x49, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_livekit_models_proto_rawDescOnce sync.Once
	file_livekit_models_proto_rawDescData = file_livekit_models_proto_rawDesc
)

func file_livekit_models_proto_rawDescGZIP() []byte {
	file_livekit_models_proto_rawDescOnce.Do(func() {
		file_livekit_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_models_proto_rawDescData)
	})
	return file_livekit_models_proto_rawDescData
}

var file_livekit_models_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_livekit_models_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_livekit_models_proto_goTypes = []interface{}{
	(TrackType)(0),             // 0: livekit.TrackType
	(ParticipantInfo_State)(0), // 1: livekit.ParticipantInfo.State
	(*Room)(nil),               // 2: livekit.Room
	(*Codec)(nil),              // 3: livekit.Codec
	(*ParticipantInfo)(nil),    // 4: livekit.ParticipantInfo
	(*TrackInfo)(nil),          // 5: livekit.TrackInfo
	(*DataMessage)(nil),        // 6: livekit.DataMessage
}
var file_livekit_models_proto_depIdxs = []int32{
	3, // 0: livekit.Room.enabled_codecs:type_name -> livekit.Codec
	1, // 1: livekit.ParticipantInfo.state:type_name -> livekit.ParticipantInfo.State
	5, // 2: livekit.ParticipantInfo.tracks:type_name -> livekit.TrackInfo
	0, // 3: livekit.TrackInfo.type:type_name -> livekit.TrackType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_livekit_models_proto_init() }
func file_livekit_models_proto_init() {
	if File_livekit_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_livekit_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_livekit_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Codec); i {
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
		file_livekit_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantInfo); i {
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
		file_livekit_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackInfo); i {
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
		file_livekit_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMessage); i {
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
	file_livekit_models_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*DataMessage_Text)(nil),
		(*DataMessage_Binary)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_livekit_models_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_models_proto_goTypes,
		DependencyIndexes: file_livekit_models_proto_depIdxs,
		EnumInfos:         file_livekit_models_proto_enumTypes,
		MessageInfos:      file_livekit_models_proto_msgTypes,
	}.Build()
	File_livekit_models_proto = out.File
	file_livekit_models_proto_rawDesc = nil
	file_livekit_models_proto_goTypes = nil
	file_livekit_models_proto_depIdxs = nil
}
