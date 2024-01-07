// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/ads/v1/advertisers.proto

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

// AdvertisersServiceClient is the client API for AdvertisersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertisersServiceClient interface {
	// Lists advertisers.
	ListAdvertisers(ctx context.Context, in *ListAdvertisersRequest, opts ...grpc.CallOption) (*ListAdvertisersResponse, error)
	// Gets a advertiser.
	GetAdvertiser(ctx context.Context, in *GetAdvertiserRequest, opts ...grpc.CallOption) (*GetAdvertiserResponse, error)
	// Creates a advertiser.
	CreateAdvertiser(ctx context.Context, in *CreateAdvertiserRequest, opts ...grpc.CallOption) (*CreateAdvertiserResponse, error)
	// Updates a advertiser.
	UpdateAdvertiser(ctx context.Context, in *UpdateAdvertiserRequest, opts ...grpc.CallOption) (*UpdateAdvertiserResponse, error)
	// Deletes a advertiser.
	DeleteAdvertiser(ctx context.Context, in *DeleteAdvertiserRequest, opts ...grpc.CallOption) (*DeleteAdvertiserResponse, error)
}

type advertisersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertisersServiceClient(cc grpc.ClientConnInterface) AdvertisersServiceClient {
	return &advertisersServiceClient{cc}
}

func (c *advertisersServiceClient) ListAdvertisers(ctx context.Context, in *ListAdvertisersRequest, opts ...grpc.CallOption) (*ListAdvertisersResponse, error) {
	out := new(ListAdvertisersResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.AdvertisersService/ListAdvertisers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisersServiceClient) GetAdvertiser(ctx context.Context, in *GetAdvertiserRequest, opts ...grpc.CallOption) (*GetAdvertiserResponse, error) {
	out := new(GetAdvertiserResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.AdvertisersService/GetAdvertiser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisersServiceClient) CreateAdvertiser(ctx context.Context, in *CreateAdvertiserRequest, opts ...grpc.CallOption) (*CreateAdvertiserResponse, error) {
	out := new(CreateAdvertiserResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.AdvertisersService/CreateAdvertiser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisersServiceClient) UpdateAdvertiser(ctx context.Context, in *UpdateAdvertiserRequest, opts ...grpc.CallOption) (*UpdateAdvertiserResponse, error) {
	out := new(UpdateAdvertiserResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.AdvertisersService/UpdateAdvertiser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisersServiceClient) DeleteAdvertiser(ctx context.Context, in *DeleteAdvertiserRequest, opts ...grpc.CallOption) (*DeleteAdvertiserResponse, error) {
	out := new(DeleteAdvertiserResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.AdvertisersService/DeleteAdvertiser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertisersServiceServer is the server API for AdvertisersService service.
// All implementations must embed UnimplementedAdvertisersServiceServer
// for forward compatibility
type AdvertisersServiceServer interface {
	// Lists advertisers.
	ListAdvertisers(context.Context, *ListAdvertisersRequest) (*ListAdvertisersResponse, error)
	// Gets a advertiser.
	GetAdvertiser(context.Context, *GetAdvertiserRequest) (*GetAdvertiserResponse, error)
	// Creates a advertiser.
	CreateAdvertiser(context.Context, *CreateAdvertiserRequest) (*CreateAdvertiserResponse, error)
	// Updates a advertiser.
	UpdateAdvertiser(context.Context, *UpdateAdvertiserRequest) (*UpdateAdvertiserResponse, error)
	// Deletes a advertiser.
	DeleteAdvertiser(context.Context, *DeleteAdvertiserRequest) (*DeleteAdvertiserResponse, error)
	mustEmbedUnimplementedAdvertisersServiceServer()
}

// UnimplementedAdvertisersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdvertisersServiceServer struct {
}

func (UnimplementedAdvertisersServiceServer) ListAdvertisers(context.Context, *ListAdvertisersRequest) (*ListAdvertisersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdvertisers not implemented")
}
func (UnimplementedAdvertisersServiceServer) GetAdvertiser(context.Context, *GetAdvertiserRequest) (*GetAdvertiserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdvertiser not implemented")
}
func (UnimplementedAdvertisersServiceServer) CreateAdvertiser(context.Context, *CreateAdvertiserRequest) (*CreateAdvertiserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvertiser not implemented")
}
func (UnimplementedAdvertisersServiceServer) UpdateAdvertiser(context.Context, *UpdateAdvertiserRequest) (*UpdateAdvertiserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdvertiser not implemented")
}
func (UnimplementedAdvertisersServiceServer) DeleteAdvertiser(context.Context, *DeleteAdvertiserRequest) (*DeleteAdvertiserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdvertiser not implemented")
}
func (UnimplementedAdvertisersServiceServer) mustEmbedUnimplementedAdvertisersServiceServer() {}

// UnsafeAdvertisersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertisersServiceServer will
// result in compilation errors.
type UnsafeAdvertisersServiceServer interface {
	mustEmbedUnimplementedAdvertisersServiceServer()
}

func RegisterAdvertisersServiceServer(s grpc.ServiceRegistrar, srv AdvertisersServiceServer) {
	s.RegisterService(&AdvertisersService_ServiceDesc, srv)
}

func _AdvertisersService_ListAdvertisers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdvertisersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisersServiceServer).ListAdvertisers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.AdvertisersService/ListAdvertisers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisersServiceServer).ListAdvertisers(ctx, req.(*ListAdvertisersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisersService_GetAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisersServiceServer).GetAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.AdvertisersService/GetAdvertiser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisersServiceServer).GetAdvertiser(ctx, req.(*GetAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisersService_CreateAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisersServiceServer).CreateAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.AdvertisersService/CreateAdvertiser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisersServiceServer).CreateAdvertiser(ctx, req.(*CreateAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisersService_UpdateAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisersServiceServer).UpdateAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.AdvertisersService/UpdateAdvertiser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisersServiceServer).UpdateAdvertiser(ctx, req.(*UpdateAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertisersService_DeleteAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisersServiceServer).DeleteAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.AdvertisersService/DeleteAdvertiser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisersServiceServer).DeleteAdvertiser(ctx, req.(*DeleteAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdvertisersService_ServiceDesc is the grpc.ServiceDesc for AdvertisersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdvertisersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.AdvertisersService",
	HandlerType: (*AdvertisersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAdvertisers",
			Handler:    _AdvertisersService_ListAdvertisers_Handler,
		},
		{
			MethodName: "GetAdvertiser",
			Handler:    _AdvertisersService_GetAdvertiser_Handler,
		},
		{
			MethodName: "CreateAdvertiser",
			Handler:    _AdvertisersService_CreateAdvertiser_Handler,
		},
		{
			MethodName: "UpdateAdvertiser",
			Handler:    _AdvertisersService_UpdateAdvertiser_Handler,
		},
		{
			MethodName: "DeleteAdvertiser",
			Handler:    _AdvertisersService_DeleteAdvertiser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/advertisers.proto",
}