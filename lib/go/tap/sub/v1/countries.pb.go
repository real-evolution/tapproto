// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: tap/sub/v1/countries.proto

package subv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// The request message for the ListCountries method.
type ListCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The page to retrieve. Defaults to 0.
	Page *int32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// The maximum number of countries to return per page.
	PageSize *int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
}

func (x *ListCountriesRequest) Reset() {
	*x = ListCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesRequest) ProtoMessage() {}

func (x *ListCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesRequest.ProtoReflect.Descriptor instead.
func (*ListCountriesRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{0}
}

func (x *ListCountriesRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListCountriesRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

// The response message for the ListCountries method.
type ListCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of countries that match the request.
	Countries []*Country `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *ListCountriesResponse) Reset() {
	*x = ListCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesResponse) ProtoMessage() {}

func (x *ListCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesResponse.ProtoReflect.Descriptor instead.
func (*ListCountriesResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{1}
}

func (x *ListCountriesResponse) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

// The request message for the GetCountry method.
type GetCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ISO 3166-1 alpha 2 code of the country to retrieve.
	Iso2 string `protobuf:"bytes,1,opt,name=iso2,proto3" json:"iso2,omitempty"`
}

func (x *GetCountryRequest) Reset() {
	*x = GetCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryRequest) ProtoMessage() {}

func (x *GetCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryRequest.ProtoReflect.Descriptor instead.
func (*GetCountryRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{2}
}

func (x *GetCountryRequest) GetIso2() string {
	if x != nil {
		return x.Iso2
	}
	return ""
}

// The response message for the GetCountry method.
type GetCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The country.
	Country *Country `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *GetCountryResponse) Reset() {
	*x = GetCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryResponse) ProtoMessage() {}

func (x *GetCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryResponse.ProtoReflect.Descriptor instead.
func (*GetCountryResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{3}
}

func (x *GetCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

// The request message for the UpdateCountry method.
type UpdateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ISO 3166-1 alpha 2 code of the country to update.
	Iso2 string `protobuf:"bytes,1,opt,name=iso2,proto3" json:"iso2,omitempty"`
	// If set, the new display name of the country.
	DisplayName *string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
}

func (x *UpdateCountryRequest) Reset() {
	*x = UpdateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryRequest) ProtoMessage() {}

func (x *UpdateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCountryRequest) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCountryRequest) GetIso2() string {
	if x != nil {
		return x.Iso2
	}
	return ""
}

func (x *UpdateCountryRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

// The response message for the UpdateCountry method.
type UpdateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated country.
	Country *Country `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *UpdateCountryResponse) Reset() {
	*x = UpdateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryResponse) ProtoMessage() {}

func (x *UpdateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryResponse.ProtoReflect.Descriptor instead.
func (*UpdateCountryResponse) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

// A message to represent a country
type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ISO 3166-1 alpha 2 code
	Iso2 string `protobuf:"bytes,1,opt,name=iso2,proto3" json:"iso2,omitempty"`
	// The name of the country
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The country phone code
	PhoneCode int32 `protobuf:"varint,3,opt,name=phone_code,json=phoneCode,proto3" json:"phone_code,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tap_sub_v1_countries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_tap_sub_v1_countries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_tap_sub_v1_countries_proto_rawDescGZIP(), []int{6}
}

func (x *Country) GetIso2() string {
	if x != nil {
		return x.Iso2
	}
	return ""
}

func (x *Country) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Country) GetPhoneCode() int32 {
	if x != nil {
		return x.PhoneCode
	}
	return 0
}

var File_tap_sub_v1_countries_proto protoreflect.FileDescriptor

var file_tap_sub_v1_countries_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61,
	0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x1a, 0x02, 0x28, 0x00, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x20, 0x00, 0x48, 0x01, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x4a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x73, 0x6f, 0x32, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x14, 0xba, 0x48, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0x52, 0x04, 0x69, 0x73, 0x6f, 0x32, 0x22, 0x43,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x22, 0x79, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69,
	0x73, 0x6f, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba, 0x48, 0x11, 0x72, 0x0f,
	0x32, 0x0d, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0x52,
	0x04, 0x69, 0x73, 0x6f, 0x32, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5f, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x6f, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x73, 0x6f, 0x32, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x1d, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x20, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa6, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x75, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x0a, 0x54,
	0x61, 0x70, 0x2e, 0x53, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x61, 0x70, 0x5c,
	0x53, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x54, 0x61, 0x70, 0x5c, 0x53, 0x75, 0x62,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x54, 0x61, 0x70, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tap_sub_v1_countries_proto_rawDescOnce sync.Once
	file_tap_sub_v1_countries_proto_rawDescData = file_tap_sub_v1_countries_proto_rawDesc
)

func file_tap_sub_v1_countries_proto_rawDescGZIP() []byte {
	file_tap_sub_v1_countries_proto_rawDescOnce.Do(func() {
		file_tap_sub_v1_countries_proto_rawDescData = protoimpl.X.CompressGZIP(file_tap_sub_v1_countries_proto_rawDescData)
	})
	return file_tap_sub_v1_countries_proto_rawDescData
}

var file_tap_sub_v1_countries_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tap_sub_v1_countries_proto_goTypes = []interface{}{
	(*ListCountriesRequest)(nil),  // 0: tap.sub.v1.ListCountriesRequest
	(*ListCountriesResponse)(nil), // 1: tap.sub.v1.ListCountriesResponse
	(*GetCountryRequest)(nil),     // 2: tap.sub.v1.GetCountryRequest
	(*GetCountryResponse)(nil),    // 3: tap.sub.v1.GetCountryResponse
	(*UpdateCountryRequest)(nil),  // 4: tap.sub.v1.UpdateCountryRequest
	(*UpdateCountryResponse)(nil), // 5: tap.sub.v1.UpdateCountryResponse
	(*Country)(nil),               // 6: tap.sub.v1.Country
}
var file_tap_sub_v1_countries_proto_depIdxs = []int32{
	6, // 0: tap.sub.v1.ListCountriesResponse.countries:type_name -> tap.sub.v1.Country
	6, // 1: tap.sub.v1.GetCountryResponse.country:type_name -> tap.sub.v1.Country
	6, // 2: tap.sub.v1.UpdateCountryResponse.country:type_name -> tap.sub.v1.Country
	0, // 3: tap.sub.v1.CountriesService.ListCountries:input_type -> tap.sub.v1.ListCountriesRequest
	2, // 4: tap.sub.v1.CountriesService.GetCountry:input_type -> tap.sub.v1.GetCountryRequest
	4, // 5: tap.sub.v1.CountriesService.UpdateCountry:input_type -> tap.sub.v1.UpdateCountryRequest
	1, // 6: tap.sub.v1.CountriesService.ListCountries:output_type -> tap.sub.v1.ListCountriesResponse
	3, // 7: tap.sub.v1.CountriesService.GetCountry:output_type -> tap.sub.v1.GetCountryResponse
	5, // 8: tap.sub.v1.CountriesService.UpdateCountry:output_type -> tap.sub.v1.UpdateCountryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tap_sub_v1_countries_proto_init() }
func file_tap_sub_v1_countries_proto_init() {
	if File_tap_sub_v1_countries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tap_sub_v1_countries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountriesRequest); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountriesResponse); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryRequest); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryResponse); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryRequest); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryResponse); i {
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
		file_tap_sub_v1_countries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
	file_tap_sub_v1_countries_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_tap_sub_v1_countries_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tap_sub_v1_countries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tap_sub_v1_countries_proto_goTypes,
		DependencyIndexes: file_tap_sub_v1_countries_proto_depIdxs,
		MessageInfos:      file_tap_sub_v1_countries_proto_msgTypes,
	}.Build()
	File_tap_sub_v1_countries_proto = out.File
	file_tap_sub_v1_countries_proto_rawDesc = nil
	file_tap_sub_v1_countries_proto_goTypes = nil
	file_tap_sub_v1_countries_proto_depIdxs = nil
}
