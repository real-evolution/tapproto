// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: buf/validate/expression.proto

package validate

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

// `Constraint` represents a validation rule written in the Common Expression
// Language (CEL) syntax. Each Constraint includes a unique identifier, an
// optional error message, and the CEL expression to evaluate. For more
// information on CEL, [see our documentation](https://github.com/bufbuild/protovalidate/blob/main/docs/cel.md).
//
// ```proto
//
//	message Foo {
//	  option (buf.validate.message).cel = {
//	    id: "foo.bar"
//	    message: "bar must be greater than 0"
//	    expression: "this.bar > 0"
//	  };
//	  int32 bar = 1;
//	}
//
// ```
type Constraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `id` is a string that serves as a machine-readable name for this Constraint.
	// It should be unique within its scope, which could be either a message or a field.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// `message` is an optional field that provides a human-readable error message
	// for this Constraint when the CEL expression evaluates to false. If a
	// non-empty message is provided, any strings resulting from the CEL
	// expression evaluation are ignored.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// `expression` is the actual CEL expression that will be evaluated for
	// validation. This string must resolve to either a boolean or a string
	// value. If the expression evaluates to false or a non-empty string, the
	// validation is considered failed, and the message is rejected.
	Expression string `protobuf:"bytes,3,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Constraint) Reset() {
	*x = Constraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_expression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint) ProtoMessage() {}

func (x *Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_expression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint.ProtoReflect.Descriptor instead.
func (*Constraint) Descriptor() ([]byte, []int) {
	return file_buf_validate_expression_proto_rawDescGZIP(), []int{0}
}

func (x *Constraint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Constraint) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Constraint) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

// `Violations` is a collection of `Violation` messages. This message type is returned by
// protovalidate when a proto message fails to meet the requirements set by the `Constraint` validation rules.
// Each individual violation is represented by a `Violation` message.
type Violations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `violations` is a repeated field that contains all the `Violation` messages corresponding to the violations detected.
	Violations []*Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *Violations) Reset() {
	*x = Violations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_expression_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Violations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Violations) ProtoMessage() {}

func (x *Violations) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_expression_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Violations.ProtoReflect.Descriptor instead.
func (*Violations) Descriptor() ([]byte, []int) {
	return file_buf_validate_expression_proto_rawDescGZIP(), []int{1}
}

func (x *Violations) GetViolations() []*Violation {
	if x != nil {
		return x.Violations
	}
	return nil
}

// `Violation` represents a single instance where a validation rule, expressed
// as a `Constraint`, was not met. It provides information about the field that
// caused the violation, the specific constraint that wasn't fulfilled, and a
// human-readable error message.
//
// ```json
//
//	{
//	  "fieldPath": "bar",
//	  "constraintId": "foo.bar",
//	  "message": "bar must be greater than 0"
//	}
//
// ```
type Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `field_path` is a machine-readable identifier that points to the specific field that failed the validation.
	// This could be a nested field, in which case the path will include all the parent fields leading to the actual field that caused the violation.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// `constraint_id` is the unique identifier of the `Constraint` that was not fulfilled.
	// This is the same `id` that was specified in the `Constraint` message, allowing easy tracing of which rule was violated.
	ConstraintId string `protobuf:"bytes,2,opt,name=constraint_id,json=constraintId,proto3" json:"constraint_id,omitempty"`
	// `message` is a human-readable error message that describes the nature of the violation.
	// This can be the default error message from the violated `Constraint`, or it can be a custom message that gives more context about the violation.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// `for_key` indicates whether the violation was caused by a map key, rather than a value.
	ForKey bool `protobuf:"varint,4,opt,name=for_key,json=forKey,proto3" json:"for_key,omitempty"`
}

func (x *Violation) Reset() {
	*x = Violation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_expression_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Violation) ProtoMessage() {}

func (x *Violation) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_expression_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Violation.ProtoReflect.Descriptor instead.
func (*Violation) Descriptor() ([]byte, []int) {
	return file_buf_validate_expression_proto_rawDescGZIP(), []int{2}
}

func (x *Violation) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (x *Violation) GetConstraintId() string {
	if x != nil {
		return x.ConstraintId
	}
	return ""
}

func (x *Violation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Violation) GetForKey() bool {
	if x != nil {
		return x.ForKey
	}
	return false
}

var File_buf_validate_expression_proto protoreflect.FileDescriptor

var file_buf_validate_expression_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0x56, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x0a, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x82, 0x01, 0x0a,
	0x09, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x4b, 0x65,
	0x79, 0x42, 0xbd, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0f, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x62, 0x75, 0x66, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0xa2, 0x02, 0x03, 0x42, 0x56, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0xca, 0x02, 0x0c, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x18, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_expression_proto_rawDescOnce sync.Once
	file_buf_validate_expression_proto_rawDescData = file_buf_validate_expression_proto_rawDesc
)

func file_buf_validate_expression_proto_rawDescGZIP() []byte {
	file_buf_validate_expression_proto_rawDescOnce.Do(func() {
		file_buf_validate_expression_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_expression_proto_rawDescData)
	})
	return file_buf_validate_expression_proto_rawDescData
}

var file_buf_validate_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_validate_expression_proto_goTypes = []interface{}{
	(*Constraint)(nil), // 0: buf.validate.Constraint
	(*Violations)(nil), // 1: buf.validate.Violations
	(*Violation)(nil),  // 2: buf.validate.Violation
}
var file_buf_validate_expression_proto_depIdxs = []int32{
	2, // 0: buf.validate.Violations.violations:type_name -> buf.validate.Violation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_validate_expression_proto_init() }
func file_buf_validate_expression_proto_init() {
	if File_buf_validate_expression_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_expression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint); i {
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
		file_buf_validate_expression_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Violations); i {
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
		file_buf_validate_expression_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Violation); i {
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
			RawDescriptor: file_buf_validate_expression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_expression_proto_goTypes,
		DependencyIndexes: file_buf_validate_expression_proto_depIdxs,
		MessageInfos:      file_buf_validate_expression_proto_msgTypes,
	}.Build()
	File_buf_validate_expression_proto = out.File
	file_buf_validate_expression_proto_rawDesc = nil
	file_buf_validate_expression_proto_goTypes = nil
	file_buf_validate_expression_proto_depIdxs = nil
}
