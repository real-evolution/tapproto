// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: tap/verify/v1/verifiers.proto

package verifyv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VerifiersService_ListVerifiers_FullMethodName      = "/tap.verify.v1.VerifiersService/ListVerifiers"
	VerifiersService_GetVerifier_FullMethodName        = "/tap.verify.v1.VerifiersService/GetVerifier"
	VerifiersService_CreateVerifier_FullMethodName     = "/tap.verify.v1.VerifiersService/CreateVerifier"
	VerifiersService_UpdateVerifier_FullMethodName     = "/tap.verify.v1.VerifiersService/UpdateVerifier"
	VerifiersService_DeleteVerifier_FullMethodName     = "/tap.verify.v1.VerifiersService/DeleteVerifier"
	VerifiersService_AddVerifierTags_FullMethodName    = "/tap.verify.v1.VerifiersService/AddVerifierTags"
	VerifiersService_ListVerifierTags_FullMethodName   = "/tap.verify.v1.VerifiersService/ListVerifierTags"
	VerifiersService_DeleteVerifierTags_FullMethodName = "/tap.verify.v1.VerifiersService/DeleteVerifierTags"
)

// VerifiersServiceClient is the client API for VerifiersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Verifier's management service
type VerifiersServiceClient interface {
	// Lists verifiers
	ListVerifiers(ctx context.Context, in *ListVerifiersRequest, opts ...grpc.CallOption) (*ListVerifiersResponse, error)
	// Gets a verifier
	GetVerifier(ctx context.Context, in *GetVerifierRequest, opts ...grpc.CallOption) (*GetVerifierResponse, error)
	// Create a verifier
	CreateVerifier(ctx context.Context, in *CreateVerifierRequest, opts ...grpc.CallOption) (*CreateVerifierResponse, error)
	// Updates a verifier
	UpdateVerifier(ctx context.Context, in *UpdateVerifierRequest, opts ...grpc.CallOption) (*UpdateVerifierResponse, error)
	// Deletes a verifier
	DeleteVerifier(ctx context.Context, in *DeleteVerifierRequest, opts ...grpc.CallOption) (*DeleteVerifierResponse, error)
	// Add a tag
	AddVerifierTags(ctx context.Context, in *AddVerifierTagsRequest, opts ...grpc.CallOption) (*AddVerifierTagsResponse, error)
	// List tags
	ListVerifierTags(ctx context.Context, in *ListVerifierTagsRequest, opts ...grpc.CallOption) (*ListVerifierTagsResponse, error)
	// Remove tags
	DeleteVerifierTags(ctx context.Context, in *DeleteVerifierTagsRequest, opts ...grpc.CallOption) (*DeleteVerifierTagsResponse, error)
}

type verifiersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifiersServiceClient(cc grpc.ClientConnInterface) VerifiersServiceClient {
	return &verifiersServiceClient{cc}
}

func (c *verifiersServiceClient) ListVerifiers(ctx context.Context, in *ListVerifiersRequest, opts ...grpc.CallOption) (*ListVerifiersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVerifiersResponse)
	err := c.cc.Invoke(ctx, VerifiersService_ListVerifiers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) GetVerifier(ctx context.Context, in *GetVerifierRequest, opts ...grpc.CallOption) (*GetVerifierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVerifierResponse)
	err := c.cc.Invoke(ctx, VerifiersService_GetVerifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) CreateVerifier(ctx context.Context, in *CreateVerifierRequest, opts ...grpc.CallOption) (*CreateVerifierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVerifierResponse)
	err := c.cc.Invoke(ctx, VerifiersService_CreateVerifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) UpdateVerifier(ctx context.Context, in *UpdateVerifierRequest, opts ...grpc.CallOption) (*UpdateVerifierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVerifierResponse)
	err := c.cc.Invoke(ctx, VerifiersService_UpdateVerifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) DeleteVerifier(ctx context.Context, in *DeleteVerifierRequest, opts ...grpc.CallOption) (*DeleteVerifierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVerifierResponse)
	err := c.cc.Invoke(ctx, VerifiersService_DeleteVerifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) AddVerifierTags(ctx context.Context, in *AddVerifierTagsRequest, opts ...grpc.CallOption) (*AddVerifierTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVerifierTagsResponse)
	err := c.cc.Invoke(ctx, VerifiersService_AddVerifierTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) ListVerifierTags(ctx context.Context, in *ListVerifierTagsRequest, opts ...grpc.CallOption) (*ListVerifierTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVerifierTagsResponse)
	err := c.cc.Invoke(ctx, VerifiersService_ListVerifierTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiersServiceClient) DeleteVerifierTags(ctx context.Context, in *DeleteVerifierTagsRequest, opts ...grpc.CallOption) (*DeleteVerifierTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVerifierTagsResponse)
	err := c.cc.Invoke(ctx, VerifiersService_DeleteVerifierTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifiersServiceServer is the server API for VerifiersService service.
// All implementations must embed UnimplementedVerifiersServiceServer
// for forward compatibility
//
// Verifier's management service
type VerifiersServiceServer interface {
	// Lists verifiers
	ListVerifiers(context.Context, *ListVerifiersRequest) (*ListVerifiersResponse, error)
	// Gets a verifier
	GetVerifier(context.Context, *GetVerifierRequest) (*GetVerifierResponse, error)
	// Create a verifier
	CreateVerifier(context.Context, *CreateVerifierRequest) (*CreateVerifierResponse, error)
	// Updates a verifier
	UpdateVerifier(context.Context, *UpdateVerifierRequest) (*UpdateVerifierResponse, error)
	// Deletes a verifier
	DeleteVerifier(context.Context, *DeleteVerifierRequest) (*DeleteVerifierResponse, error)
	// Add a tag
	AddVerifierTags(context.Context, *AddVerifierTagsRequest) (*AddVerifierTagsResponse, error)
	// List tags
	ListVerifierTags(context.Context, *ListVerifierTagsRequest) (*ListVerifierTagsResponse, error)
	// Remove tags
	DeleteVerifierTags(context.Context, *DeleteVerifierTagsRequest) (*DeleteVerifierTagsResponse, error)
	mustEmbedUnimplementedVerifiersServiceServer()
}

// UnimplementedVerifiersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVerifiersServiceServer struct {
}

func (UnimplementedVerifiersServiceServer) ListVerifiers(context.Context, *ListVerifiersRequest) (*ListVerifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVerifiers not implemented")
}
func (UnimplementedVerifiersServiceServer) GetVerifier(context.Context, *GetVerifierRequest) (*GetVerifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerifier not implemented")
}
func (UnimplementedVerifiersServiceServer) CreateVerifier(context.Context, *CreateVerifierRequest) (*CreateVerifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVerifier not implemented")
}
func (UnimplementedVerifiersServiceServer) UpdateVerifier(context.Context, *UpdateVerifierRequest) (*UpdateVerifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVerifier not implemented")
}
func (UnimplementedVerifiersServiceServer) DeleteVerifier(context.Context, *DeleteVerifierRequest) (*DeleteVerifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVerifier not implemented")
}
func (UnimplementedVerifiersServiceServer) AddVerifierTags(context.Context, *AddVerifierTagsRequest) (*AddVerifierTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVerifierTags not implemented")
}
func (UnimplementedVerifiersServiceServer) ListVerifierTags(context.Context, *ListVerifierTagsRequest) (*ListVerifierTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVerifierTags not implemented")
}
func (UnimplementedVerifiersServiceServer) DeleteVerifierTags(context.Context, *DeleteVerifierTagsRequest) (*DeleteVerifierTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVerifierTags not implemented")
}
func (UnimplementedVerifiersServiceServer) mustEmbedUnimplementedVerifiersServiceServer() {}

// UnsafeVerifiersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifiersServiceServer will
// result in compilation errors.
type UnsafeVerifiersServiceServer interface {
	mustEmbedUnimplementedVerifiersServiceServer()
}

func RegisterVerifiersServiceServer(s grpc.ServiceRegistrar, srv VerifiersServiceServer) {
	s.RegisterService(&VerifiersService_ServiceDesc, srv)
}

func _VerifiersService_ListVerifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVerifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).ListVerifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_ListVerifiers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).ListVerifiers(ctx, req.(*ListVerifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_GetVerifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVerifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).GetVerifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_GetVerifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).GetVerifier(ctx, req.(*GetVerifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_CreateVerifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVerifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).CreateVerifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_CreateVerifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).CreateVerifier(ctx, req.(*CreateVerifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_UpdateVerifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVerifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).UpdateVerifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_UpdateVerifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).UpdateVerifier(ctx, req.(*UpdateVerifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_DeleteVerifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVerifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).DeleteVerifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_DeleteVerifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).DeleteVerifier(ctx, req.(*DeleteVerifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_AddVerifierTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVerifierTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).AddVerifierTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_AddVerifierTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).AddVerifierTags(ctx, req.(*AddVerifierTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_ListVerifierTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVerifierTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).ListVerifierTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_ListVerifierTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).ListVerifierTags(ctx, req.(*ListVerifierTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiersService_DeleteVerifierTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVerifierTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiersServiceServer).DeleteVerifierTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifiersService_DeleteVerifierTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiersServiceServer).DeleteVerifierTags(ctx, req.(*DeleteVerifierTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerifiersService_ServiceDesc is the grpc.ServiceDesc for VerifiersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerifiersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.verify.v1.VerifiersService",
	HandlerType: (*VerifiersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVerifiers",
			Handler:    _VerifiersService_ListVerifiers_Handler,
		},
		{
			MethodName: "GetVerifier",
			Handler:    _VerifiersService_GetVerifier_Handler,
		},
		{
			MethodName: "CreateVerifier",
			Handler:    _VerifiersService_CreateVerifier_Handler,
		},
		{
			MethodName: "UpdateVerifier",
			Handler:    _VerifiersService_UpdateVerifier_Handler,
		},
		{
			MethodName: "DeleteVerifier",
			Handler:    _VerifiersService_DeleteVerifier_Handler,
		},
		{
			MethodName: "AddVerifierTags",
			Handler:    _VerifiersService_AddVerifierTags_Handler,
		},
		{
			MethodName: "ListVerifierTags",
			Handler:    _VerifiersService_ListVerifierTags_Handler,
		},
		{
			MethodName: "DeleteVerifierTags",
			Handler:    _VerifiersService_DeleteVerifierTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/verify/v1/verifiers.proto",
}
