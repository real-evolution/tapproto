// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tap/ads/v1/requirements.proto

package adsv1

import (
	v1 "github.com/real-evolution/tapproto/lib/go/tap/types/v1"
	v11 "github.com/real-evolution/tapproto/lib/go/tap/verify/v1"
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

// The request message for the ListRequirements method.
type ListRequirementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pagination parameters.
	Page *v1.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRequirementsRequest) Reset() {
	*x = ListRequirementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequirementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequirementsRequest) ProtoMessage() {}

func (x *ListRequirementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequirementsRequest.ProtoReflect.Descriptor instead.
func (*ListRequirementsRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequirementsRequest) GetPage() *v1.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

// The response message for the ListRequirements method.
type ListRequirementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requirements.
	Items []*Requirement `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// The next page cursor.
	NextCursor []byte `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3,oneof" json:"next_cursor,omitempty"`
}

func (x *ListRequirementsResponse) Reset() {
	*x = ListRequirementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequirementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequirementsResponse) ProtoMessage() {}

func (x *ListRequirementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequirementsResponse.ProtoReflect.Descriptor instead.
func (*ListRequirementsResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequirementsResponse) GetItems() []*Requirement {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRequirementsResponse) GetNextCursor() []byte {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

// The request message for the CreateRequirement method.
type CreateRequirementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The display name of the requirement.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the requirement.
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Required verification tags of the requirement.
	RequiredTags []*v11.TagCode `protobuf:"bytes,3,rep,name=required_tags,json=requiredTags,proto3" json:"required_tags,omitempty"`
}

func (x *CreateRequirementRequest) Reset() {
	*x = CreateRequirementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequirementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequirementRequest) ProtoMessage() {}

func (x *CreateRequirementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequirementRequest.ProtoReflect.Descriptor instead.
func (*CreateRequirementRequest) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequirementRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CreateRequirementRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateRequirementRequest) GetRequiredTags() []*v11.TagCode {
	if x != nil {
		return x.RequiredTags
	}
	return nil
}

// The response message for the CreateRequirement method.
type CreateRequirementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requirement.
	Requirement *Requirement `protobuf:"bytes,1,opt,name=requirement,proto3" json:"requirement,omitempty"`
}

func (x *CreateRequirementResponse) Reset() {
	*x = CreateRequirementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequirementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequirementResponse) ProtoMessage() {}

func (x *CreateRequirementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequirementResponse.ProtoReflect.Descriptor instead.
func (*CreateRequirementResponse) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequirementResponse) GetRequirement() *Requirement {
	if x != nil {
		return x.Requirement
	}
	return nil
}

// A message to represent a requirement ID.
type RequirementID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the requirement.
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RequirementID) Reset() {
	*x = RequirementID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequirementID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequirementID) ProtoMessage() {}

func (x *RequirementID) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequirementID.ProtoReflect.Descriptor instead.
func (*RequirementID) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{4}
}

func (x *RequirementID) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// A message to represent a requirement
type Requirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the requirement.
	Id *RequirementID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the requirement.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the requirement.
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Required verification tags of the requirement.
	RequiredTags []*v11.TagCode `protobuf:"bytes,4,rep,name=required_tags,json=requiredTags,proto3" json:"required_tags,omitempty"`
	// The creation time of the requirement.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the requirement.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Requirement) Reset() {
	*x = Requirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_ads_v1_requirements_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Requirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Requirement) ProtoMessage() {}

func (x *Requirement) ProtoReflect() protoreflect.Message {
	mi := &file_tap_ads_v1_requirements_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Requirement.ProtoReflect.Descriptor instead.
func (*Requirement) Descriptor() ([]byte, []int) {
	return file_tap_ads_v1_requirements_proto_rawDescGZIP(), []int{5}
}

func (x *Requirement) GetId() *RequirementID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Requirement) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Requirement) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Requirement) GetRequiredTags() []*v11.TagCode {
	if x != nil {
		return x.RequiredTags
	}
	return nil
}

func (x *Requirement) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Requirement) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_tap_ads_v1_requirements_proto protoreflect.FileDescriptor

var file_tap_ads_v1_requirements_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x61, 0x70, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x61,
	0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x61, 0x70,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x7f, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x3b, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc9, 0x02, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x3b, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd6, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xa9, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69,
	0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x64, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x61,
	0x70, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x54, 0x61, 0x70, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x54, 0x61, 0x70, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_ads_v1_requirements_proto_rawDescOnce sync.Once
	file_tap_ads_v1_requirements_proto_rawDescData = file_tap_ads_v1_requirements_proto_rawDesc
)

func file_tap_ads_v1_requirements_proto_rawDescGZIP() []byte {
	file_tap_ads_v1_requirements_proto_rawDescOnce.Do(func() {
		file_tap_ads_v1_requirements_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_ads_v1_requirements_proto_rawDescData)
	})
	return file_tap_ads_v1_requirements_proto_rawDescData
}

var file_tap_ads_v1_requirements_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tap_ads_v1_requirements_proto_goTypes = []interface{}{
	(*ListRequirementsRequest)(nil),   // 0: tap.ads.v1.ListRequirementsRequest
	(*ListRequirementsResponse)(nil),  // 1: tap.ads.v1.ListRequirementsResponse
	(*CreateRequirementRequest)(nil),  // 2: tap.ads.v1.CreateRequirementRequest
	(*CreateRequirementResponse)(nil), // 3: tap.ads.v1.CreateRequirementResponse
	(*RequirementID)(nil),             // 4: tap.ads.v1.RequirementID
	(*Requirement)(nil),               // 5: tap.ads.v1.Requirement
	(*v1.Page)(nil),                   // 6: tap.types.v1.Page
	(*v11.TagCode)(nil),               // 7: tap.verify.v1.TagCode
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_tap_ads_v1_requirements_proto_depIdxs = []int32{
	6,  // 0: tap.ads.v1.ListRequirementsRequest.page:type_name -> tap.types.v1.Page
	5,  // 1: tap.ads.v1.ListRequirementsResponse.items:type_name -> tap.ads.v1.Requirement
	7,  // 2: tap.ads.v1.CreateRequirementRequest.required_tags:type_name -> tap.verify.v1.TagCode
	5,  // 3: tap.ads.v1.CreateRequirementResponse.requirement:type_name -> tap.ads.v1.Requirement
	4,  // 4: tap.ads.v1.Requirement.id:type_name -> tap.ads.v1.RequirementID
	7,  // 5: tap.ads.v1.Requirement.required_tags:type_name -> tap.verify.v1.TagCode
	8,  // 6: tap.ads.v1.Requirement.create_time:type_name -> google.protobuf.Timestamp
	8,  // 7: tap.ads.v1.Requirement.update_time:type_name -> google.protobuf.Timestamp
	0,  // 8: tap.ads.v1.RequirementsService.ListRequirements:input_type -> tap.ads.v1.ListRequirementsRequest
	2,  // 9: tap.ads.v1.RequirementsService.CreateRequirement:input_type -> tap.ads.v1.CreateRequirementRequest
	1,  // 10: tap.ads.v1.RequirementsService.ListRequirements:output_type -> tap.ads.v1.ListRequirementsResponse
	3,  // 11: tap.ads.v1.RequirementsService.CreateRequirement:output_type -> tap.ads.v1.CreateRequirementResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_tap_ads_v1_requirements_proto_init() }
func file_tap_ads_v1_requirements_proto_init() {
	if File_tap_ads_v1_requirements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_ads_v1_requirements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequirementsRequest); i {
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
		file_tap_ads_v1_requirements_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequirementsResponse); i {
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
		file_tap_ads_v1_requirements_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequirementRequest); i {
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
		file_tap_ads_v1_requirements_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequirementResponse); i {
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
		file_tap_ads_v1_requirements_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequirementID); i {
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
		file_tap_ads_v1_requirements_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Requirement); i {
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
	file_tap_ads_v1_requirements_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tap_ads_v1_requirements_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_tap_ads_v1_requirements_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_ads_v1_requirements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_ads_v1_requirements_proto_goTypes,
		DependencyIndexes: file_tap_ads_v1_requirements_proto_depIdxs,
		MessageInfos:      file_tap_ads_v1_requirements_proto_msgTypes,
	}.Build()
	File_tap_ads_v1_requirements_proto = out.File
	file_tap_ads_v1_requirements_proto_rawDesc = nil
	file_tap_ads_v1_requirements_proto_goTypes = nil
	file_tap_ads_v1_requirements_proto_depIdxs = nil
}