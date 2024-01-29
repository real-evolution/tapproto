// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tap/ads/v1/passes.proto

package adsv1

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
	PassesService_ListPasses_FullMethodName = "/tap.ads.v1.PassesService/ListPasses"
	PassesService_GetPass_FullMethodName    = "/tap.ads.v1.PassesService/GetPass"
	PassesService_CreatePass_FullMethodName = "/tap.ads.v1.PassesService/CreatePass"
	PassesService_UpdatePass_FullMethodName = "/tap.ads.v1.PassesService/UpdatePass"
	PassesService_DeletePass_FullMethodName = "/tap.ads.v1.PassesService/DeletePass"
)

// PassesServiceClient is the client API for PassesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PassesServiceClient interface {
	ListPasses(ctx context.Context, in *ListPassesRequest, opts ...grpc.CallOption) (*ListPassesResponse, error)
	GetPass(ctx context.Context, in *GetPassRequest, opts ...grpc.CallOption) (*GetPassResponse, error)
	CreatePass(ctx context.Context, in *CreatePassRequest, opts ...grpc.CallOption) (*CreatePassResponse, error)
	UpdatePass(ctx context.Context, in *UpdatePassRequest, opts ...grpc.CallOption) (*UpdatePassResponse, error)
	DeletePass(ctx context.Context, in *DeletePassRequest, opts ...grpc.CallOption) (*DeletePassResponse, error)
}

type passesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPassesServiceClient(cc grpc.ClientConnInterface) PassesServiceClient {
	return &passesServiceClient{cc}
}

func (c *passesServiceClient) ListPasses(ctx context.Context, in *ListPassesRequest, opts ...grpc.CallOption) (*ListPassesResponse, error) {
	out := new(ListPassesResponse)
	err := c.cc.Invoke(ctx, PassesService_ListPasses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passesServiceClient) GetPass(ctx context.Context, in *GetPassRequest, opts ...grpc.CallOption) (*GetPassResponse, error) {
	out := new(GetPassResponse)
	err := c.cc.Invoke(ctx, PassesService_GetPass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passesServiceClient) CreatePass(ctx context.Context, in *CreatePassRequest, opts ...grpc.CallOption) (*CreatePassResponse, error) {
	out := new(CreatePassResponse)
	err := c.cc.Invoke(ctx, PassesService_CreatePass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passesServiceClient) UpdatePass(ctx context.Context, in *UpdatePassRequest, opts ...grpc.CallOption) (*UpdatePassResponse, error) {
	out := new(UpdatePassResponse)
	err := c.cc.Invoke(ctx, PassesService_UpdatePass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passesServiceClient) DeletePass(ctx context.Context, in *DeletePassRequest, opts ...grpc.CallOption) (*DeletePassResponse, error) {
	out := new(DeletePassResponse)
	err := c.cc.Invoke(ctx, PassesService_DeletePass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassesServiceServer is the server API for PassesService service.
// All implementations must embed UnimplementedPassesServiceServer
// for forward compatibility
type PassesServiceServer interface {
	ListPasses(context.Context, *ListPassesRequest) (*ListPassesResponse, error)
	GetPass(context.Context, *GetPassRequest) (*GetPassResponse, error)
	CreatePass(context.Context, *CreatePassRequest) (*CreatePassResponse, error)
	UpdatePass(context.Context, *UpdatePassRequest) (*UpdatePassResponse, error)
	DeletePass(context.Context, *DeletePassRequest) (*DeletePassResponse, error)
	mustEmbedUnimplementedPassesServiceServer()
}

// UnimplementedPassesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPassesServiceServer struct {
}

func (UnimplementedPassesServiceServer) ListPasses(context.Context, *ListPassesRequest) (*ListPassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPasses not implemented")
}
func (UnimplementedPassesServiceServer) GetPass(context.Context, *GetPassRequest) (*GetPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPass not implemented")
}
func (UnimplementedPassesServiceServer) CreatePass(context.Context, *CreatePassRequest) (*CreatePassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePass not implemented")
}
func (UnimplementedPassesServiceServer) UpdatePass(context.Context, *UpdatePassRequest) (*UpdatePassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePass not implemented")
}
func (UnimplementedPassesServiceServer) DeletePass(context.Context, *DeletePassRequest) (*DeletePassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePass not implemented")
}
func (UnimplementedPassesServiceServer) mustEmbedUnimplementedPassesServiceServer() {}

// UnsafePassesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassesServiceServer will
// result in compilation errors.
type UnsafePassesServiceServer interface {
	mustEmbedUnimplementedPassesServiceServer()
}

func RegisterPassesServiceServer(s grpc.ServiceRegistrar, srv PassesServiceServer) {
	s.RegisterService(&PassesService_ServiceDesc, srv)
}

func _PassesService_ListPasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassesServiceServer).ListPasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassesService_ListPasses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassesServiceServer).ListPasses(ctx, req.(*ListPassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassesService_GetPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassesServiceServer).GetPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassesService_GetPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassesServiceServer).GetPass(ctx, req.(*GetPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassesService_CreatePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassesServiceServer).CreatePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassesService_CreatePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassesServiceServer).CreatePass(ctx, req.(*CreatePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassesService_UpdatePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassesServiceServer).UpdatePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassesService_UpdatePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassesServiceServer).UpdatePass(ctx, req.(*UpdatePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassesService_DeletePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassesServiceServer).DeletePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassesService_DeletePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassesServiceServer).DeletePass(ctx, req.(*DeletePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PassesService_ServiceDesc is the grpc.ServiceDesc for PassesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PassesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.PassesService",
	HandlerType: (*PassesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPasses",
			Handler:    _PassesService_ListPasses_Handler,
		},
		{
			MethodName: "GetPass",
			Handler:    _PassesService_GetPass_Handler,
		},
		{
			MethodName: "CreatePass",
			Handler:    _PassesService_CreatePass_Handler,
		},
		{
			MethodName: "UpdatePass",
			Handler:    _PassesService_UpdatePass_Handler,
		},
		{
			MethodName: "DeletePass",
			Handler:    _PassesService_DeletePass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/passes.proto",
}
