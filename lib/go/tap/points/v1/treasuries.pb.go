// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: tap/points/v1/treasuries.proto

package pointsv1

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

// The request message for the ListTreasuries method.
type ListTreasuriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pagination parameters.
	Page *v1.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListTreasuriesRequest) Reset() {
	*x = ListTreasuriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTreasuriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTreasuriesRequest) ProtoMessage() {}

func (x *ListTreasuriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTreasuriesRequest.ProtoReflect.Descriptor instead.
func (*ListTreasuriesRequest) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{0}
}

func (x *ListTreasuriesRequest) GetPage() *v1.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

// The response message for the ListTreasuries method.
type ListTreasuriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The treasuries.
	Items []*Treasury `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// The next page cursor.
	NextCursor []byte `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3,oneof" json:"next_cursor,omitempty"`
}

func (x *ListTreasuriesResponse) Reset() {
	*x = ListTreasuriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTreasuriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTreasuriesResponse) ProtoMessage() {}

func (x *ListTreasuriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTreasuriesResponse.ProtoReflect.Descriptor instead.
func (*ListTreasuriesResponse) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{1}
}

func (x *ListTreasuriesResponse) GetItems() []*Treasury {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTreasuriesResponse) GetNextCursor() []byte {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

// The request message for the GetTreasury method.
type GetTreasuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the treasury to retrieve.
	Id *TreasuryID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTreasuryRequest) Reset() {
	*x = GetTreasuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreasuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreasuryRequest) ProtoMessage() {}

func (x *GetTreasuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreasuryRequest.ProtoReflect.Descriptor instead.
func (*GetTreasuryRequest) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{2}
}

func (x *GetTreasuryRequest) GetId() *TreasuryID {
	if x != nil {
		return x.Id
	}
	return nil
}

// The response message for the GetTreasury method.
type GetTreasuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The treasury.
	Treasury *Treasury `protobuf:"bytes,1,opt,name=treasury,proto3" json:"treasury,omitempty"`
}

func (x *GetTreasuryResponse) Reset() {
	*x = GetTreasuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreasuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreasuryResponse) ProtoMessage() {}

func (x *GetTreasuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreasuryResponse.ProtoReflect.Descriptor instead.
func (*GetTreasuryResponse) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{3}
}

func (x *GetTreasuryResponse) GetTreasury() *Treasury {
	if x != nil {
		return x.Treasury
	}
	return nil
}

// The request message for the CreateTreasury method.
type CreateTreasuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The display name of the treasury.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *CreateTreasuryRequest) Reset() {
	*x = CreateTreasuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTreasuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTreasuryRequest) ProtoMessage() {}

func (x *CreateTreasuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTreasuryRequest.ProtoReflect.Descriptor instead.
func (*CreateTreasuryRequest) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTreasuryRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// The response message for the CreateTreasury method.
type CreateTreasuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The treasury.
	Treasury *Treasury `protobuf:"bytes,1,opt,name=treasury,proto3" json:"treasury,omitempty"`
}

func (x *CreateTreasuryResponse) Reset() {
	*x = CreateTreasuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTreasuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTreasuryResponse) ProtoMessage() {}

func (x *CreateTreasuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTreasuryResponse.ProtoReflect.Descriptor instead.
func (*CreateTreasuryResponse) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTreasuryResponse) GetTreasury() *Treasury {
	if x != nil {
		return x.Treasury
	}
	return nil
}

// The request message for the UpdateTreasury method.
type UpdateTreasuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the treasury to update.
	Id *TreasuryID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the treasury.
	DisplayName *string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
}

func (x *UpdateTreasuryRequest) Reset() {
	*x = UpdateTreasuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTreasuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTreasuryRequest) ProtoMessage() {}

func (x *UpdateTreasuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTreasuryRequest.ProtoReflect.Descriptor instead.
func (*UpdateTreasuryRequest) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTreasuryRequest) GetId() *TreasuryID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateTreasuryRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

// The response message for the UpdateTreasury method.
type UpdateTreasuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated treasury.
	Treasury *Treasury `protobuf:"bytes,1,opt,name=treasury,proto3" json:"treasury,omitempty"`
}

func (x *UpdateTreasuryResponse) Reset() {
	*x = UpdateTreasuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTreasuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTreasuryResponse) ProtoMessage() {}

func (x *UpdateTreasuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTreasuryResponse.ProtoReflect.Descriptor instead.
func (*UpdateTreasuryResponse) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTreasuryResponse) GetTreasury() *Treasury {
	if x != nil {
		return x.Treasury
	}
	return nil
}

// The request message for the DeleteTreasury method.
type DeleteTreasuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the treasury to delete.
	Id *TreasuryID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTreasuryRequest) Reset() {
	*x = DeleteTreasuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTreasuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTreasuryRequest) ProtoMessage() {}

func (x *DeleteTreasuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTreasuryRequest.ProtoReflect.Descriptor instead.
func (*DeleteTreasuryRequest) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTreasuryRequest) GetId() *TreasuryID {
	if x != nil {
		return x.Id
	}
	return nil
}

// The response message for the DeleteTreasury method.
type DeleteTreasuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTreasuryResponse) Reset() {
	*x = DeleteTreasuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTreasuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTreasuryResponse) ProtoMessage() {}

func (x *DeleteTreasuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTreasuryResponse.ProtoReflect.Descriptor instead.
func (*DeleteTreasuryResponse) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{9}
}

// A message to represent a treasury ID.
type TreasuryID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TreasuryID) Reset() {
	*x = TreasuryID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasuryID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasuryID) ProtoMessage() {}

func (x *TreasuryID) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasuryID.ProtoReflect.Descriptor instead.
func (*TreasuryID) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{10}
}

func (x *TreasuryID) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// A message to represent a treasury
type Treasury struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the treasury.
	Id *TreasuryID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the treasury.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The creation time of the treasury.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the treasury.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Treasury) Reset() {
	*x = Treasury{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_points_v1_treasuries_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Treasury) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Treasury) ProtoMessage() {}

func (x *Treasury) ProtoReflect() protoreflect.Message {
	mi := &file_tap_points_v1_treasuries_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Treasury.ProtoReflect.Descriptor instead.
func (*Treasury) Descriptor() ([]byte, []int) {
	return file_tap_points_v1_treasuries_proto_rawDescGZIP(), []int{11}
}

func (x *Treasury) GetId() *TreasuryID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Treasury) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Treasury) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Treasury) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_tap_points_v1_treasuries_proto protoreflect.FileDescriptor

var file_tap_points_v1_treasuries_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x61, 0x70, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3f, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x7d, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22,
	0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x52, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x22, 0x3a, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x08, 0x74,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x22, 0x7b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x79, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x0a, 0x0a, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xe5, 0x03, 0x0a, 0x11, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x21,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xbc, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x54, 0x50, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x61, 0x70, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x61, 0x70, 0x5c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x61, 0x70, 0x5c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x54, 0x61, 0x70, 0x3a, 0x3a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_points_v1_treasuries_proto_rawDescOnce sync.Once
	file_tap_points_v1_treasuries_proto_rawDescData = file_tap_points_v1_treasuries_proto_rawDesc
)

func file_tap_points_v1_treasuries_proto_rawDescGZIP() []byte {
	file_tap_points_v1_treasuries_proto_rawDescOnce.Do(func() {
		file_tap_points_v1_treasuries_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_points_v1_treasuries_proto_rawDescData)
	})
	return file_tap_points_v1_treasuries_proto_rawDescData
}

var file_tap_points_v1_treasuries_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tap_points_v1_treasuries_proto_goTypes = []interface{}{
	(*ListTreasuriesRequest)(nil),  // 0: tap.points.v1.ListTreasuriesRequest
	(*ListTreasuriesResponse)(nil), // 1: tap.points.v1.ListTreasuriesResponse
	(*GetTreasuryRequest)(nil),     // 2: tap.points.v1.GetTreasuryRequest
	(*GetTreasuryResponse)(nil),    // 3: tap.points.v1.GetTreasuryResponse
	(*CreateTreasuryRequest)(nil),  // 4: tap.points.v1.CreateTreasuryRequest
	(*CreateTreasuryResponse)(nil), // 5: tap.points.v1.CreateTreasuryResponse
	(*UpdateTreasuryRequest)(nil),  // 6: tap.points.v1.UpdateTreasuryRequest
	(*UpdateTreasuryResponse)(nil), // 7: tap.points.v1.UpdateTreasuryResponse
	(*DeleteTreasuryRequest)(nil),  // 8: tap.points.v1.DeleteTreasuryRequest
	(*DeleteTreasuryResponse)(nil), // 9: tap.points.v1.DeleteTreasuryResponse
	(*TreasuryID)(nil),             // 10: tap.points.v1.TreasuryID
	(*Treasury)(nil),               // 11: tap.points.v1.Treasury
	(*v1.Page)(nil),                // 12: tap.types.v1.Page
	(*timestamppb.Timestamp)(nil),  // 13: google.protobuf.Timestamp
}
var file_tap_points_v1_treasuries_proto_depIdxs = []int32{
	12, // 0: tap.points.v1.ListTreasuriesRequest.page:type_name -> tap.types.v1.Page
	11, // 1: tap.points.v1.ListTreasuriesResponse.items:type_name -> tap.points.v1.Treasury
	10, // 2: tap.points.v1.GetTreasuryRequest.id:type_name -> tap.points.v1.TreasuryID
	11, // 3: tap.points.v1.GetTreasuryResponse.treasury:type_name -> tap.points.v1.Treasury
	11, // 4: tap.points.v1.CreateTreasuryResponse.treasury:type_name -> tap.points.v1.Treasury
	10, // 5: tap.points.v1.UpdateTreasuryRequest.id:type_name -> tap.points.v1.TreasuryID
	11, // 6: tap.points.v1.UpdateTreasuryResponse.treasury:type_name -> tap.points.v1.Treasury
	10, // 7: tap.points.v1.DeleteTreasuryRequest.id:type_name -> tap.points.v1.TreasuryID
	10, // 8: tap.points.v1.Treasury.id:type_name -> tap.points.v1.TreasuryID
	13, // 9: tap.points.v1.Treasury.create_time:type_name -> google.protobuf.Timestamp
	13, // 10: tap.points.v1.Treasury.update_time:type_name -> google.protobuf.Timestamp
	0,  // 11: tap.points.v1.TreasuriesService.ListTreasuries:input_type -> tap.points.v1.ListTreasuriesRequest
	2,  // 12: tap.points.v1.TreasuriesService.GetTreasury:input_type -> tap.points.v1.GetTreasuryRequest
	4,  // 13: tap.points.v1.TreasuriesService.CreateTreasury:input_type -> tap.points.v1.CreateTreasuryRequest
	6,  // 14: tap.points.v1.TreasuriesService.UpdateTreasury:input_type -> tap.points.v1.UpdateTreasuryRequest
	8,  // 15: tap.points.v1.TreasuriesService.DeleteTreasury:input_type -> tap.points.v1.DeleteTreasuryRequest
	1,  // 16: tap.points.v1.TreasuriesService.ListTreasuries:output_type -> tap.points.v1.ListTreasuriesResponse
	3,  // 17: tap.points.v1.TreasuriesService.GetTreasury:output_type -> tap.points.v1.GetTreasuryResponse
	5,  // 18: tap.points.v1.TreasuriesService.CreateTreasury:output_type -> tap.points.v1.CreateTreasuryResponse
	7,  // 19: tap.points.v1.TreasuriesService.UpdateTreasury:output_type -> tap.points.v1.UpdateTreasuryResponse
	9,  // 20: tap.points.v1.TreasuriesService.DeleteTreasury:output_type -> tap.points.v1.DeleteTreasuryResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tap_points_v1_treasuries_proto_init() }
func file_tap_points_v1_treasuries_proto_init() {
	if File_tap_points_v1_treasuries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_points_v1_treasuries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTreasuriesRequest); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTreasuriesResponse); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreasuryRequest); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreasuryResponse); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTreasuryRequest); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTreasuryResponse); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTreasuryRequest); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTreasuryResponse); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTreasuryRequest); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTreasuryResponse); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasuryID); i {
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
		file_tap_points_v1_treasuries_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Treasury); i {
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
	file_tap_points_v1_treasuries_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tap_points_v1_treasuries_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_points_v1_treasuries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_points_v1_treasuries_proto_goTypes,
		DependencyIndexes: file_tap_points_v1_treasuries_proto_depIdxs,
		MessageInfos:      file_tap_points_v1_treasuries_proto_msgTypes,
	}.Build()
	File_tap_points_v1_treasuries_proto = out.File
	file_tap_points_v1_treasuries_proto_rawDesc = nil
	file_tap_points_v1_treasuries_proto_goTypes = nil
	file_tap_points_v1_treasuries_proto_depIdxs = nil
}
