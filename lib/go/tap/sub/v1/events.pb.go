// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: tap/sub/v1/events.proto

package subv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Event_Type int32

const (
	// Unspecified event type.
	Event_TYPE_UNSPECIFIED Event_Type = 0
	// Event type A.
	Event_TYPE_ACTIVATE Event_Type = 1
	// Event type B.
	Event_TYPE_DEACTIVATE Event_Type = 2
)

// Enum value maps for Event_Type.
var (
	Event_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_ACTIVATE",
		2: "TYPE_DEACTIVATE",
	}
	Event_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_ACTIVATE":    1,
		"TYPE_DEACTIVATE":  2,
	}
)

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}

func (x Event_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tap_sub_v1_events_proto_enumTypes[0].Descriptor()
}

func (Event_Type) Type() protoreflect.EnumType {
	return &file_tap_sub_v1_events_proto_enumTypes[0]
}

func (x Event_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Type.Descriptor instead.
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{6, 0}
}

// A request to get a batch of events.
type GetEventBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the site to get events for.
	SiteId string `protobuf:"bytes,1,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	// The maximum number of events to return in the batch.
	MaxEvents int32 `protobuf:"varint,2,opt,name=max_events,json=maxEvents,proto3" json:"max_events,omitempty"`
}

func (x *GetEventBatchRequest) Reset() {
	*x = GetEventBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventBatchRequest) ProtoMessage() {}

func (x *GetEventBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventBatchRequest.ProtoReflect.Descriptor instead.
func (*GetEventBatchRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventBatchRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *GetEventBatchRequest) GetMaxEvents() int32 {
	if x != nil {
		return x.MaxEvents
	}
	return 0
}

// The response to a stream events request.
type GetEventBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The events in the batch.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventBatchResponse) Reset() {
	*x = GetEventBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventBatchResponse) ProtoMessage() {}

func (x *GetEventBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventBatchResponse.ProtoReflect.Descriptor instead.
func (*GetEventBatchResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventBatchResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type CommitEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the site to get events for.
	SiteId string `protobuf:"bytes,1,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	// The ID of the event.
	EventId int64 `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *CommitEventRequest) Reset() {
	*x = CommitEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitEventRequest) ProtoMessage() {}

func (x *CommitEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitEventRequest.ProtoReflect.Descriptor instead.
func (*CommitEventRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *CommitEventRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *CommitEventRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

// The response to a commit event request.
type CommitEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitEventResponse) Reset() {
	*x = CommitEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitEventResponse) ProtoMessage() {}

func (x *CommitEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitEventResponse.ProtoReflect.Descriptor instead.
func (*CommitEventResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{3}
}

type RevertEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the site to get events for.
	SiteId string `protobuf:"bytes,1,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	// The ID of the event.
	EventId int64 `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *RevertEventRequest) Reset() {
	*x = RevertEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertEventRequest) ProtoMessage() {}

func (x *RevertEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertEventRequest.ProtoReflect.Descriptor instead.
func (*RevertEventRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{4}
}

func (x *RevertEventRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *RevertEventRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

// The response to a commit event request.
type RevertEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevertEventResponse) Reset() {
	*x = RevertEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertEventResponse) ProtoMessage() {}

func (x *RevertEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertEventResponse.ProtoReflect.Descriptor instead.
func (*RevertEventResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{5}
}

// An event.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the event.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of site event occurred on.
	Phone int64 `protobuf:"varint,2,opt,name=phone,proto3" json:"phone,omitempty"`
	// The event type.
	Type Event_Type `protobuf:"varint,3,opt,name=type,proto3,enum=tap.sub.v1.Event_Type" json:"type,omitempty"`
	// The timestamp of the event.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_events_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *Event) GetType() Event_Type {
	if x != nil {
		return x.Type
	}
	return Event_TYPE_UNSPECIFIED
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_tap_sub_v1_events_proto protoreflect.FileDescriptor

var file_tap_sub_v1_events_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x1a, 0x05, 0x18, 0x80, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x5e, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba, 0x48, 0x07,
	0xc8, 0x01, 0x01, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x15, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd9,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x44, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0x85, 0x02, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x20, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xa3, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x75, 0x62, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x2e, 0x53, 0x75, 0x62,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x5c, 0x53, 0x75, 0x62, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x16, 0x54, 0x61, 0x70, 0x5c, 0x53, 0x75, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54, 0x61, 0x70, 0x3a,
	0x3a, 0x53, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_sub_v1_events_proto_rawDescOnce sync.Once
	file_tap_sub_v1_events_proto_rawDescData = file_tap_sub_v1_events_proto_rawDesc
)

func file_tap_sub_v1_events_proto_rawDescGZIP() []byte {
	file_tap_sub_v1_events_proto_rawDescOnce.Do(func() {
		file_tap_sub_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_sub_v1_events_proto_rawDescData)
	})
	return file_tap_sub_v1_events_proto_rawDescData
}

var file_tap_sub_v1_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tap_sub_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tap_sub_v1_events_proto_goTypes = []interface{}{
	(Event_Type)(0),               // 0: tap.sub.v1.Event.Type
	(*GetEventBatchRequest)(nil),  // 1: tap.sub.v1.GetEventBatchRequest
	(*GetEventBatchResponse)(nil), // 2: tap.sub.v1.GetEventBatchResponse
	(*CommitEventRequest)(nil),    // 3: tap.sub.v1.CommitEventRequest
	(*CommitEventResponse)(nil),   // 4: tap.sub.v1.CommitEventResponse
	(*RevertEventRequest)(nil),    // 5: tap.sub.v1.RevertEventRequest
	(*RevertEventResponse)(nil),   // 6: tap.sub.v1.RevertEventResponse
	(*Event)(nil),                 // 7: tap.sub.v1.Event
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_tap_sub_v1_events_proto_depIdxs = []int32{
	7, // 0: tap.sub.v1.GetEventBatchResponse.events:type_name -> tap.sub.v1.Event
	0, // 1: tap.sub.v1.Event.type:type_name -> tap.sub.v1.Event.Type
	8, // 2: tap.sub.v1.Event.timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: tap.sub.v1.EventsService.GetEventBatch:input_type -> tap.sub.v1.GetEventBatchRequest
	3, // 4: tap.sub.v1.EventsService.CommitEvent:input_type -> tap.sub.v1.CommitEventRequest
	5, // 5: tap.sub.v1.EventsService.RevertEvent:input_type -> tap.sub.v1.RevertEventRequest
	2, // 6: tap.sub.v1.EventsService.GetEventBatch:output_type -> tap.sub.v1.GetEventBatchResponse
	4, // 7: tap.sub.v1.EventsService.CommitEvent:output_type -> tap.sub.v1.CommitEventResponse
	6, // 8: tap.sub.v1.EventsService.RevertEvent:output_type -> tap.sub.v1.RevertEventResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tap_sub_v1_events_proto_init() }
func file_tap_sub_v1_events_proto_init() {
	if File_tap_sub_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_sub_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventBatchRequest); i {
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
		file_tap_sub_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventBatchResponse); i {
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
		file_tap_sub_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitEventRequest); i {
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
		file_tap_sub_v1_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitEventResponse); i {
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
		file_tap_sub_v1_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertEventRequest); i {
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
		file_tap_sub_v1_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertEventResponse); i {
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
		file_tap_sub_v1_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_tap_sub_v1_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_sub_v1_events_proto_goTypes,
		DependencyIndexes: file_tap_sub_v1_events_proto_depIdxs,
		EnumInfos:         file_tap_sub_v1_events_proto_enumTypes,
		MessageInfos:      file_tap_sub_v1_events_proto_msgTypes,
	}.Build()
	File_tap_sub_v1_events_proto = out.File
	file_tap_sub_v1_events_proto_rawDesc = nil
	file_tap_sub_v1_events_proto_goTypes = nil
	file_tap_sub_v1_events_proto_depIdxs = nil
}
