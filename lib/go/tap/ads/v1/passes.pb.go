// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tap/ads/v1/passes.proto

package adsv1

import (
	v1 "github.com/real-evolution/tapproto/lib/go/tap/types/v1"
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

type ListPassesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *v1.Page `protobuf:"bytes,15,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListPassesRequest) Reset() {
	*x = ListPassesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPassesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassesRequest) ProtoMessage() {}

func (x *ListPassesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassesRequest.ProtoReflect.Descriptor instead.
func (*ListPassesRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{0}
}

func (x *ListPassesRequest) GetPage() *v1.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListPassesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*Pass `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextCursor []byte  `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3,oneof" json:"next_cursor,omitempty"`
}

func (x *ListPassesResponse) Reset() {
	*x = ListPassesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassesResponse) ProtoMessage() {}

func (x *ListPassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassesResponse.ProtoReflect.Descriptor instead.
func (*ListPassesResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{1}
}

func (x *ListPassesResponse) GetItems() []*Pass {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPassesResponse) GetNextCursor() []byte {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

type GetPassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *PassID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPassRequest) Reset() {
	*x = GetPassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPassRequest) ProtoMessage() {}

func (x *GetPassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPassRequest.ProtoReflect.Descriptor instead.
func (*GetPassRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{2}
}

func (x *GetPassRequest) GetId() *PassID {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetPassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass *Pass `protobuf:"bytes,1,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *GetPassResponse) Reset() {
	*x = GetPassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPassResponse) ProtoMessage() {}

func (x *GetPassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPassResponse.ProtoReflect.Descriptor instead.
func (*GetPassResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{3}
}

func (x *GetPassResponse) GetPass() *Pass {
	if x != nil {
		return x.Pass
	}
	return nil
}

type CreatePassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string       `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ValidFor    *v1.Interval `protobuf:"bytes,3,opt,name=valid_for,json=validFor,proto3" json:"valid_for,omitempty"`
}

func (x *CreatePassRequest) Reset() {
	*x = CreatePassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePassRequest) ProtoMessage() {}

func (x *CreatePassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePassRequest.ProtoReflect.Descriptor instead.
func (*CreatePassRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePassRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CreatePassRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePassRequest) GetValidFor() *v1.Interval {
	if x != nil {
		return x.ValidFor
	}
	return nil
}

type CreatePassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass *Pass `protobuf:"bytes,1,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *CreatePassResponse) Reset() {
	*x = CreatePassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePassResponse) ProtoMessage() {}

func (x *CreatePassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePassResponse.ProtoReflect.Descriptor instead.
func (*CreatePassResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePassResponse) GetPass() *Pass {
	if x != nil {
		return x.Pass
	}
	return nil
}

type UpdatePassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *PassID      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName *string      `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	Description *string      `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ValidFor    *v1.Interval `protobuf:"bytes,4,opt,name=valid_for,json=validFor,proto3,oneof" json:"valid_for,omitempty"`
}

func (x *UpdatePassRequest) Reset() {
	*x = UpdatePassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePassRequest) ProtoMessage() {}

func (x *UpdatePassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePassRequest.ProtoReflect.Descriptor instead.
func (*UpdatePassRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePassRequest) GetId() *PassID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdatePassRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *UpdatePassRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdatePassRequest) GetValidFor() *v1.Interval {
	if x != nil {
		return x.ValidFor
	}
	return nil
}

type UpdatePassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass *Pass `protobuf:"bytes,1,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *UpdatePassResponse) Reset() {
	*x = UpdatePassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePassResponse) ProtoMessage() {}

func (x *UpdatePassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePassResponse.ProtoReflect.Descriptor instead.
func (*UpdatePassResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePassResponse) GetPass() *Pass {
	if x != nil {
		return x.Pass
	}
	return nil
}

type DeletePassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *PassID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePassRequest) Reset() {
	*x = DeletePassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePassRequest) ProtoMessage() {}

func (x *DeletePassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePassRequest.ProtoReflect.Descriptor instead.
func (*DeletePassRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePassRequest) GetId() *PassID {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeletePassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePassResponse) Reset() {
	*x = DeletePassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePassResponse) ProtoMessage() {}

func (x *DeletePassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePassResponse.ProtoReflect.Descriptor instead.
func (*DeletePassResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{9}
}

type PassID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PassID) Reset() {
	*x = PassID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassID) ProtoMessage() {}

func (x *PassID) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassID.ProtoReflect.Descriptor instead.
func (*PassID) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{10}
}

func (x *PassID) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Pass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *PassID                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ValidFor    *v1.Interval           `protobuf:"bytes,4,opt,name=valid_for,json=validFor,proto3" json:"valid_for,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Pass) Reset() {
	*x = Pass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_passes_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pass) ProtoMessage() {}

func (x *Pass) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_passes_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pass.ProtoReflect.Descriptor instead.
func (*Pass) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_passes_proto_rawDescGZIP(), []int{11}
}

func (x *Pass) GetId() *PassID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Pass) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Pass) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Pass) GetValidFor() *v1.Interval {
	if x != nil {
		return x.ValidFor
	}
	return nil
}

func (x *Pass) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Pass) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_tap_ads_v1_passes_proto protoreflect.FileDescriptor

var file_tap_ads_v1_passes_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x61, 0x70, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61, 0x70, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x0f, 0x22, 0x72, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x22, 0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x04,
	0x70, 0x61, 0x73, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x46, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x70, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x22, 0x37, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x0a, 0x06, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x9e, 0x02, 0x0a, 0x04, 0x50, 0x61, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x32, 0x87, 0x03, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x12, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1d, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa3, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d,
	0x65, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x41, 0x58,
	0xaa, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a,
	0x54, 0x61, 0x70, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x54, 0x61, 0x70,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54, 0x61, 0x70, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_ads_v1_passes_proto_rawDescOnce sync.Once
	file_tap_ads_v1_passes_proto_rawDescData = file_tap_ads_v1_passes_proto_rawDesc
)

func file_tap_ads_v1_passes_proto_rawDescGZIP() []byte {
	file_tap_ads_v1_passes_proto_rawDescOnce.Do(func() {
		file_tap_ads_v1_passes_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_ads_v1_passes_proto_rawDescData)
	})
	return file_tap_ads_v1_passes_proto_rawDescData
}

var file_tap_ads_v1_passes_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tap_ads_v1_passes_proto_goTypes = []interface{}{
	(*ListPassesRequest)(nil),     // 0: tap.ads.v1.ListPassesRequest
	(*ListPassesResponse)(nil),    // 1: tap.ads.v1.ListPassesResponse
	(*GetPassRequest)(nil),        // 2: tap.ads.v1.GetPassRequest
	(*GetPassResponse)(nil),       // 3: tap.ads.v1.GetPassResponse
	(*CreatePassRequest)(nil),     // 4: tap.ads.v1.CreatePassRequest
	(*CreatePassResponse)(nil),    // 5: tap.ads.v1.CreatePassResponse
	(*UpdatePassRequest)(nil),     // 6: tap.ads.v1.UpdatePassRequest
	(*UpdatePassResponse)(nil),    // 7: tap.ads.v1.UpdatePassResponse
	(*DeletePassRequest)(nil),     // 8: tap.ads.v1.DeletePassRequest
	(*DeletePassResponse)(nil),    // 9: tap.ads.v1.DeletePassResponse
	(*PassID)(nil),                // 10: tap.ads.v1.PassID
	(*Pass)(nil),                  // 11: tap.ads.v1.Pass
	(*v1.Page)(nil),               // 12: tap.types.v1.Page
	(*v1.Interval)(nil),           // 13: tap.types.v1.Interval
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_tap_ads_v1_passes_proto_depIdxs = []int32{
	12, // 0: tap.ads.v1.ListPassesRequest.page:type_name -> tap.types.v1.Page
	11, // 1: tap.ads.v1.ListPassesResponse.items:type_name -> tap.ads.v1.Pass
	10, // 2: tap.ads.v1.GetPassRequest.id:type_name -> tap.ads.v1.PassID
	11, // 3: tap.ads.v1.GetPassResponse.pass:type_name -> tap.ads.v1.Pass
	13, // 4: tap.ads.v1.CreatePassRequest.valid_for:type_name -> tap.types.v1.Interval
	11, // 5: tap.ads.v1.CreatePassResponse.pass:type_name -> tap.ads.v1.Pass
	10, // 6: tap.ads.v1.UpdatePassRequest.id:type_name -> tap.ads.v1.PassID
	13, // 7: tap.ads.v1.UpdatePassRequest.valid_for:type_name -> tap.types.v1.Interval
	11, // 8: tap.ads.v1.UpdatePassResponse.pass:type_name -> tap.ads.v1.Pass
	10, // 9: tap.ads.v1.DeletePassRequest.id:type_name -> tap.ads.v1.PassID
	10, // 10: tap.ads.v1.Pass.id:type_name -> tap.ads.v1.PassID
	13, // 11: tap.ads.v1.Pass.valid_for:type_name -> tap.types.v1.Interval
	14, // 12: tap.ads.v1.Pass.create_time:type_name -> google.protobuf.Timestamp
	14, // 13: tap.ads.v1.Pass.update_time:type_name -> google.protobuf.Timestamp
	0,  // 14: tap.ads.v1.PassesService.ListPasses:input_type -> tap.ads.v1.ListPassesRequest
	2,  // 15: tap.ads.v1.PassesService.GetPass:input_type -> tap.ads.v1.GetPassRequest
	4,  // 16: tap.ads.v1.PassesService.CreatePass:input_type -> tap.ads.v1.CreatePassRequest
	6,  // 17: tap.ads.v1.PassesService.UpdatePass:input_type -> tap.ads.v1.UpdatePassRequest
	8,  // 18: tap.ads.v1.PassesService.DeletePass:input_type -> tap.ads.v1.DeletePassRequest
	1,  // 19: tap.ads.v1.PassesService.ListPasses:output_type -> tap.ads.v1.ListPassesResponse
	3,  // 20: tap.ads.v1.PassesService.GetPass:output_type -> tap.ads.v1.GetPassResponse
	5,  // 21: tap.ads.v1.PassesService.CreatePass:output_type -> tap.ads.v1.CreatePassResponse
	7,  // 22: tap.ads.v1.PassesService.UpdatePass:output_type -> tap.ads.v1.UpdatePassResponse
	9,  // 23: tap.ads.v1.PassesService.DeletePass:output_type -> tap.ads.v1.DeletePassResponse
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tap_ads_v1_passes_proto_init() }
func file_tap_ads_v1_passes_proto_init() {
	if File_tap_ads_v1_passes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_ads_v1_passes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPassesRequest); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPassesResponse); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPassRequest); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPassResponse); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePassRequest); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePassResponse); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePassRequest); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePassResponse); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePassRequest); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePassResponse); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassID); i {
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
		file_tap_ads_v1_passes_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pass); i {
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
	file_tap_ads_v1_passes_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tap_ads_v1_passes_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_ads_v1_passes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_ads_v1_passes_proto_goTypes,
		DependencyIndexes: file_tap_ads_v1_passes_proto_depIdxs,
		MessageInfos:      file_tap_ads_v1_passes_proto_msgTypes,
	}.Build()
	File_tap_ads_v1_passes_proto = out.File
	file_tap_ads_v1_passes_proto_rawDesc = nil
	file_tap_ads_v1_passes_proto_goTypes = nil
	file_tap_ads_v1_passes_proto_depIdxs = nil
}
