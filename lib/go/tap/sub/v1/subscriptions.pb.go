// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: tap/sub/v1/subscriptions.proto

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

// Request message for `ListSubscriptions`.
type ListSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the provider to list subscriptions for.
	ProviderId *string `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3,oneof" json:"provider_id,omitempty"`
	// Filter by subscriber.
	SubscriberId *string `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3,oneof" json:"subscriber_id,omitempty"`
	// Filter by site.
	SiteId *string `protobuf:"bytes,3,opt,name=site_id,json=siteId,proto3,oneof" json:"site_id,omitempty"`
	// Creation timestamp after which to query.
	BeforeTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=before_timestamp,json=beforeTimestamp,proto3,oneof" json:"before_timestamp,omitempty"`
	// The maximum number of subscriptions to return.
	PageSize *int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
}

func (x *ListSubscriptionsRequest) Reset() {
	*x = ListSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsRequest) ProtoMessage() {}

func (x *ListSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{0}
}

func (x *ListSubscriptionsRequest) GetProviderId() string {
	if x != nil && x.ProviderId != nil {
		return *x.ProviderId
	}
	return ""
}

func (x *ListSubscriptionsRequest) GetSubscriberId() string {
	if x != nil && x.SubscriberId != nil {
		return *x.SubscriberId
	}
	return ""
}

func (x *ListSubscriptionsRequest) GetSiteId() string {
	if x != nil && x.SiteId != nil {
		return *x.SiteId
	}
	return ""
}

func (x *ListSubscriptionsRequest) GetBeforeTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BeforeTimestamp
	}
	return nil
}

func (x *ListSubscriptionsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

// Response message for `ListSubscriptions`.
type ListSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscriptions.
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *ListSubscriptionsResponse) Reset() {
	*x = ListSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsResponse) ProtoMessage() {}

func (x *ListSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

// Request message for `GetSubscription`.
type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscriber's phone number.
	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubscriptionRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

// Response message for `GetSubscription`.
type GetSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscription.
	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *GetSubscriptionResponse) Reset() {
	*x = GetSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionResponse) ProtoMessage() {}

func (x *GetSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubscriptionResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

// Request message for `Subscribe`.
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscriber's phone number.
	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
	// The ID of the site to subscribe to.
	SiteId string `protobuf:"bytes,2,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *SubscribeRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

// Response message for `Subscribe`.
type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{5}
}

// Request message for `Unsubscribe`.
type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscriber's phone number.
	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

// Response message for `Unsubscribe`.
type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{7}
}

// A message to represent a subscription.
type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscriber's phone number associated with .
	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
	// The identifier of the subscriber who made the subscription.
	SubscriberId string `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	// The identifier of the site where the subscription was on.
	SiteId string `protobuf:"bytes,3,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	// The creation timestamp
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_subscriptions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_subscriptions_proto_rawDescGZIP(), []int{8}
}

func (x *Subscription) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *Subscription) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *Subscription) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *Subscription) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Subscription) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_tap_sub_v1_subscriptions_proto protoreflect.FileDescriptor

var file_tap_sub_v1_subscriptions_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x01, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x07, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x02, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x10, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0f, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x5b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x44, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xba, 0x48, 0x11, 0xc8, 0x01, 0x01,
	0x22, 0x0c, 0x18, 0xff, 0xff, 0x83, 0xfe, 0xa6, 0xde, 0xe1, 0x11, 0x28, 0x80, 0x20, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x57, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x14, 0xba, 0x48, 0x11, 0xc8, 0x01, 0x01, 0x22, 0x0c, 0x18, 0xff, 0xff, 0x83, 0xfe,
	0xa6, 0xde, 0xe1, 0x11, 0x28, 0x80, 0x20, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21,
	0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49,
	0x64, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x14, 0xba, 0x48, 0x11,
	0xc8, 0x01, 0x01, 0x22, 0x0c, 0x18, 0xff, 0xff, 0x83, 0xfe, 0xa6, 0xde, 0xe1, 0x11, 0x28, 0x80,
	0x20, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xdc, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xee,
	0x02, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1e,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xaa, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x75, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x61, 0x70,
	0x2e, 0x53, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x5c, 0x53, 0x75,
	0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x54, 0x61, 0x70, 0x5c, 0x53, 0x75, 0x62, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x54, 0x61, 0x70, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_sub_v1_subscriptions_proto_rawDescOnce sync.Once
	file_tap_sub_v1_subscriptions_proto_rawDescData = file_tap_sub_v1_subscriptions_proto_rawDesc
)

func file_tap_sub_v1_subscriptions_proto_rawDescGZIP() []byte {
	file_tap_sub_v1_subscriptions_proto_rawDescOnce.Do(func() {
		file_tap_sub_v1_subscriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_sub_v1_subscriptions_proto_rawDescData)
	})
	return file_tap_sub_v1_subscriptions_proto_rawDescData
}

var file_tap_sub_v1_subscriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tap_sub_v1_subscriptions_proto_goTypes = []interface{}{
	(*ListSubscriptionsRequest)(nil),  // 0: tap.sub.v1.ListSubscriptionsRequest
	(*ListSubscriptionsResponse)(nil), // 1: tap.sub.v1.ListSubscriptionsResponse
	(*GetSubscriptionRequest)(nil),    // 2: tap.sub.v1.GetSubscriptionRequest
	(*GetSubscriptionResponse)(nil),   // 3: tap.sub.v1.GetSubscriptionResponse
	(*SubscribeRequest)(nil),          // 4: tap.sub.v1.SubscribeRequest
	(*SubscribeResponse)(nil),         // 5: tap.sub.v1.SubscribeResponse
	(*UnsubscribeRequest)(nil),        // 6: tap.sub.v1.UnsubscribeRequest
	(*UnsubscribeResponse)(nil),       // 7: tap.sub.v1.UnsubscribeResponse
	(*Subscription)(nil),              // 8: tap.sub.v1.Subscription
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_tap_sub_v1_subscriptions_proto_depIdxs = []int32{
	9, // 0: tap.sub.v1.ListSubscriptionsRequest.before_timestamp:type_name -> google.protobuf.Timestamp
	8, // 1: tap.sub.v1.ListSubscriptionsResponse.subscriptions:type_name -> tap.sub.v1.Subscription
	8, // 2: tap.sub.v1.GetSubscriptionResponse.subscription:type_name -> tap.sub.v1.Subscription
	9, // 3: tap.sub.v1.Subscription.create_time:type_name -> google.protobuf.Timestamp
	9, // 4: tap.sub.v1.Subscription.update_time:type_name -> google.protobuf.Timestamp
	0, // 5: tap.sub.v1.SubscriptionsService.ListSubscriptions:input_type -> tap.sub.v1.ListSubscriptionsRequest
	2, // 6: tap.sub.v1.SubscriptionsService.GetSubscription:input_type -> tap.sub.v1.GetSubscriptionRequest
	4, // 7: tap.sub.v1.SubscriptionsService.Subscribe:input_type -> tap.sub.v1.SubscribeRequest
	6, // 8: tap.sub.v1.SubscriptionsService.Unsubscribe:input_type -> tap.sub.v1.UnsubscribeRequest
	1, // 9: tap.sub.v1.SubscriptionsService.ListSubscriptions:output_type -> tap.sub.v1.ListSubscriptionsResponse
	3, // 10: tap.sub.v1.SubscriptionsService.GetSubscription:output_type -> tap.sub.v1.GetSubscriptionResponse
	5, // 11: tap.sub.v1.SubscriptionsService.Subscribe:output_type -> tap.sub.v1.SubscribeResponse
	7, // 12: tap.sub.v1.SubscriptionsService.Unsubscribe:output_type -> tap.sub.v1.UnsubscribeResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tap_sub_v1_subscriptions_proto_init() }
func file_tap_sub_v1_subscriptions_proto_init() {
	if File_tap_sub_v1_subscriptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_sub_v1_subscriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsRequest); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsResponse); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionResponse); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeResponse); i {
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
		file_tap_sub_v1_subscriptions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
	file_tap_sub_v1_subscriptions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_sub_v1_subscriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_sub_v1_subscriptions_proto_goTypes,
		DependencyIndexes: file_tap_sub_v1_subscriptions_proto_depIdxs,
		MessageInfos:      file_tap_sub_v1_subscriptions_proto_msgTypes,
	}.Build()
	File_tap_sub_v1_subscriptions_proto = out.File
	file_tap_sub_v1_subscriptions_proto_rawDesc = nil
	file_tap_sub_v1_subscriptions_proto_goTypes = nil
	file_tap_sub_v1_subscriptions_proto_depIdxs = nil
}
