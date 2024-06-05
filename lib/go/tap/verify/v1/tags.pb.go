// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: tap/verify/v1/tags.proto

package verifyv1

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

// Request message for `ListTags`
type ListTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter by code
	Code *string `protobuf:"bytes,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	// Creation timestamp after which to query.
	BeforeTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before_timestamp,json=beforeTimestamp,proto3,oneof" json:"before_timestamp,omitempty"`
	// The maximum number of products to return
	PageSize *int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
}

func (x *ListTagsRequest) Reset() {
	*x = ListTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsRequest) ProtoMessage() {}

func (x *ListTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsRequest.ProtoReflect.Descriptor instead.
func (*ListTagsRequest) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{0}
}

func (x *ListTagsRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *ListTagsRequest) GetBeforeTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BeforeTimestamp
	}
	return nil
}

func (x *ListTagsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

// Response message for `ListTags`
type ListTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tags.
	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ListTagsResponse) Reset() {
	*x = ListTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsResponse) ProtoMessage() {}

func (x *ListTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsResponse.ProtoReflect.Descriptor instead.
func (*ListTagsResponse) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{1}
}

func (x *ListTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Request message for `GetTag`
type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The code of the tag.
	Code *TagCode `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagRequest) GetCode() *TagCode {
	if x != nil {
		return x.Code
	}
	return nil
}

// Response message for `GetTag`
type GetTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tag.
	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTagResponse) Reset() {
	*x = GetTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagResponse) ProtoMessage() {}

func (x *GetTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagResponse.ProtoReflect.Descriptor instead.
func (*GetTagResponse) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

// Request message for `CreateTag`
type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tag code.
	Code *TagCode `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTagRequest) GetCode() *TagCode {
	if x != nil {
		return x.Code
	}
	return nil
}

// Response message for `CreateTag`
type CreateTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The created tag.
	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CreateTagResponse) Reset() {
	*x = CreateTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagResponse) ProtoMessage() {}

func (x *CreateTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagResponse.ProtoReflect.Descriptor instead.
func (*CreateTagResponse) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

// Request message for `UpdateTag`
type UpdateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The old code.
	OldCode *TagCode `protobuf:"bytes,1,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	// The new code
	NewCode *TagCode `protobuf:"bytes,2,opt,name=new_code,json=newCode,proto3" json:"new_code,omitempty"`
}

func (x *UpdateTagRequest) Reset() {
	*x = UpdateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagRequest) ProtoMessage() {}

func (x *UpdateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateTagRequest) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTagRequest) GetOldCode() *TagCode {
	if x != nil {
		return x.OldCode
	}
	return nil
}

func (x *UpdateTagRequest) GetNewCode() *TagCode {
	if x != nil {
		return x.NewCode
	}
	return nil
}

// Response message for `UpdateTag`
type UpdateTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated tag.
	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *UpdateTagResponse) Reset() {
	*x = UpdateTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagResponse) ProtoMessage() {}

func (x *UpdateTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagResponse.ProtoReflect.Descriptor instead.
func (*UpdateTagResponse) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

// Request message for `DeleteTag`
type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tag code.
	Code *TagCode `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTagRequest) GetCode() *TagCode {
	if x != nil {
		return x.Code
	}
	return nil
}

// Response message for `DeleteTag`
type DeleteTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagResponse) Reset() {
	*x = DeleteTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagResponse) ProtoMessage() {}

func (x *DeleteTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagResponse.ProtoReflect.Descriptor instead.
func (*DeleteTagResponse) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{9}
}

// A message representing a tag code.
type TagCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The code of the tag.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TagCode) Reset() {
	*x = TagCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCode) ProtoMessage() {}

func (x *TagCode) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCode.ProtoReflect.Descriptor instead.
func (*TagCode) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{10}
}

func (x *TagCode) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A message representing a tag.
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The code of the tag.
	Code *TagCode `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// The creation time of the tag.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the tag.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_verify_v1_tags_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tap_verify_v1_tags_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tap_verify_v1_tags_proto_rawDescGZIP(), []int{11}
}

func (x *Tag) GetCode() *TagCode {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Tag) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tag) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_tap_verify_v1_tags_proto protoreflect.FileDescriptor

var file_tap_verify_v1_tags_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x10, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3a,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x43, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x39, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x39, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a,
	0x07, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x2a,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0x91, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x67, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1c, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb6, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x54, 0x56, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x61, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x61, 0x70, 0x5c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x61, 0x70, 0x5c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x54, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_verify_v1_tags_proto_rawDescOnce sync.Once
	file_tap_verify_v1_tags_proto_rawDescData = file_tap_verify_v1_tags_proto_rawDesc
)

func file_tap_verify_v1_tags_proto_rawDescGZIP() []byte {
	file_tap_verify_v1_tags_proto_rawDescOnce.Do(func() {
		file_tap_verify_v1_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_verify_v1_tags_proto_rawDescData)
	})
	return file_tap_verify_v1_tags_proto_rawDescData
}

var file_tap_verify_v1_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tap_verify_v1_tags_proto_goTypes = []interface{}{
	(*ListTagsRequest)(nil),       // 0: tap.verify.v1.ListTagsRequest
	(*ListTagsResponse)(nil),      // 1: tap.verify.v1.ListTagsResponse
	(*GetTagRequest)(nil),         // 2: tap.verify.v1.GetTagRequest
	(*GetTagResponse)(nil),        // 3: tap.verify.v1.GetTagResponse
	(*CreateTagRequest)(nil),      // 4: tap.verify.v1.CreateTagRequest
	(*CreateTagResponse)(nil),     // 5: tap.verify.v1.CreateTagResponse
	(*UpdateTagRequest)(nil),      // 6: tap.verify.v1.UpdateTagRequest
	(*UpdateTagResponse)(nil),     // 7: tap.verify.v1.UpdateTagResponse
	(*DeleteTagRequest)(nil),      // 8: tap.verify.v1.DeleteTagRequest
	(*DeleteTagResponse)(nil),     // 9: tap.verify.v1.DeleteTagResponse
	(*TagCode)(nil),               // 10: tap.verify.v1.TagCode
	(*Tag)(nil),                   // 11: tap.verify.v1.Tag
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_tap_verify_v1_tags_proto_depIdxs = []int32{
	12, // 0: tap.verify.v1.ListTagsRequest.before_timestamp:type_name -> google.protobuf.Timestamp
	11, // 1: tap.verify.v1.ListTagsResponse.tags:type_name -> tap.verify.v1.Tag
	10, // 2: tap.verify.v1.GetTagRequest.code:type_name -> tap.verify.v1.TagCode
	11, // 3: tap.verify.v1.GetTagResponse.tag:type_name -> tap.verify.v1.Tag
	10, // 4: tap.verify.v1.CreateTagRequest.code:type_name -> tap.verify.v1.TagCode
	11, // 5: tap.verify.v1.CreateTagResponse.tag:type_name -> tap.verify.v1.Tag
	10, // 6: tap.verify.v1.UpdateTagRequest.old_code:type_name -> tap.verify.v1.TagCode
	10, // 7: tap.verify.v1.UpdateTagRequest.new_code:type_name -> tap.verify.v1.TagCode
	11, // 8: tap.verify.v1.UpdateTagResponse.tag:type_name -> tap.verify.v1.Tag
	10, // 9: tap.verify.v1.DeleteTagRequest.code:type_name -> tap.verify.v1.TagCode
	10, // 10: tap.verify.v1.Tag.code:type_name -> tap.verify.v1.TagCode
	12, // 11: tap.verify.v1.Tag.create_time:type_name -> google.protobuf.Timestamp
	12, // 12: tap.verify.v1.Tag.update_time:type_name -> google.protobuf.Timestamp
	0,  // 13: tap.verify.v1.TagsService.ListTags:input_type -> tap.verify.v1.ListTagsRequest
	2,  // 14: tap.verify.v1.TagsService.GetTag:input_type -> tap.verify.v1.GetTagRequest
	4,  // 15: tap.verify.v1.TagsService.CreateTag:input_type -> tap.verify.v1.CreateTagRequest
	6,  // 16: tap.verify.v1.TagsService.UpdateTag:input_type -> tap.verify.v1.UpdateTagRequest
	8,  // 17: tap.verify.v1.TagsService.DeleteTag:input_type -> tap.verify.v1.DeleteTagRequest
	1,  // 18: tap.verify.v1.TagsService.ListTags:output_type -> tap.verify.v1.ListTagsResponse
	3,  // 19: tap.verify.v1.TagsService.GetTag:output_type -> tap.verify.v1.GetTagResponse
	5,  // 20: tap.verify.v1.TagsService.CreateTag:output_type -> tap.verify.v1.CreateTagResponse
	7,  // 21: tap.verify.v1.TagsService.UpdateTag:output_type -> tap.verify.v1.UpdateTagResponse
	9,  // 22: tap.verify.v1.TagsService.DeleteTag:output_type -> tap.verify.v1.DeleteTagResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_tap_verify_v1_tags_proto_init() }
func file_tap_verify_v1_tags_proto_init() {
	if File_tap_verify_v1_tags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_verify_v1_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsRequest); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsResponse); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagResponse); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagResponse); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagRequest); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagResponse); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagResponse); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCode); i {
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
		file_tap_verify_v1_tags_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
	file_tap_verify_v1_tags_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_verify_v1_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_verify_v1_tags_proto_goTypes,
		DependencyIndexes: file_tap_verify_v1_tags_proto_depIdxs,
		MessageInfos:      file_tap_verify_v1_tags_proto_msgTypes,
	}.Build()
	File_tap_verify_v1_tags_proto = out.File
	file_tap_verify_v1_tags_proto_rawDesc = nil
	file_tap_verify_v1_tags_proto_goTypes = nil
	file_tap_verify_v1_tags_proto_depIdxs = nil
}
