// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tap/verify/v1/tags.proto

package verifyv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TagsService_ListTags_FullMethodName  = "/tap.verify.v1.TagsService/ListTags"
	TagsService_GetTag_FullMethodName    = "/tap.verify.v1.TagsService/GetTag"
	TagsService_CreateTag_FullMethodName = "/tap.verify.v1.TagsService/CreateTag"
	TagsService_UpdateTag_FullMethodName = "/tap.verify.v1.TagsService/UpdateTag"
	TagsService_DeleteTag_FullMethodName = "/tap.verify.v1.TagsService/DeleteTag"
)

// TagsServiceClient is the client API for TagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsServiceClient interface {
	// Lists tags
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	// Gets a tag
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error)
	// Create a tag
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error)
	// Updates a tag
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error)
	// Deletes a tag
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error)
}

type tagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsServiceClient(cc grpc.ClientConnInterface) TagsServiceClient {
	return &tagsServiceClient{cc}
}

func (c *tagsServiceClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, TagsService_ListTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error) {
	out := new(GetTagResponse)
	err := c.cc.Invoke(ctx, TagsService_GetTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error) {
	out := new(CreateTagResponse)
	err := c.cc.Invoke(ctx, TagsService_CreateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error) {
	out := new(UpdateTagResponse)
	err := c.cc.Invoke(ctx, TagsService_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error) {
	out := new(DeleteTagResponse)
	err := c.cc.Invoke(ctx, TagsService_DeleteTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServiceServer is the server API for TagsService service.
// All implementations must embed UnimplementedTagsServiceServer
// for forward compatibility
type TagsServiceServer interface {
	// Lists tags
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	// Gets a tag
	GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error)
	// Create a tag
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
	// Updates a tag
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error)
	// Deletes a tag
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error)
	mustEmbedUnimplementedTagsServiceServer()
}

// UnimplementedTagsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagsServiceServer struct {
}

func (UnimplementedTagsServiceServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedTagsServiceServer) GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagsServiceServer) CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagsServiceServer) UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagsServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagsServiceServer) mustEmbedUnimplementedTagsServiceServer() {}

// UnsafeTagsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsServiceServer will
// result in compilation errors.
type UnsafeTagsServiceServer interface {
	mustEmbedUnimplementedTagsServiceServer()
}

func RegisterTagsServiceServer(s grpc.ServiceRegistrar, srv TagsServiceServer) {
	s.RegisterService(&TagsService_ServiceDesc, srv)
}

func _TagsService_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_ListTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagsService_ServiceDesc is the grpc.ServiceDesc for TagsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.verify.v1.TagsService",
	HandlerType: (*TagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTags",
			Handler:    _TagsService_ListTags_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagsService_GetTag_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _TagsService_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagsService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagsService_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/verify/v1/tags.proto",
}
