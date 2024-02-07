// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tap/ads/v1/ads.proto

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
	AdsService_ListAds_FullMethodName  = "/tap.ads.v1.AdsService/ListAds"
	AdsService_GetAd_FullMethodName    = "/tap.ads.v1.AdsService/GetAd"
	AdsService_CreateAd_FullMethodName = "/tap.ads.v1.AdsService/CreateAd"
)

// AdsServiceClient is the client API for AdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdsServiceClient interface {
	ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (*ListAdsResponse, error)
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*GetAdResponse, error)
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*CreateAdResponse, error)
}

type adsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdsServiceClient(cc grpc.ClientConnInterface) AdsServiceClient {
	return &adsServiceClient{cc}
}

func (c *adsServiceClient) ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (*ListAdsResponse, error) {
	out := new(ListAdsResponse)
	err := c.cc.Invoke(ctx, AdsService_ListAds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adsServiceClient) GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*GetAdResponse, error) {
	out := new(GetAdResponse)
	err := c.cc.Invoke(ctx, AdsService_GetAd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adsServiceClient) CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*CreateAdResponse, error) {
	out := new(CreateAdResponse)
	err := c.cc.Invoke(ctx, AdsService_CreateAd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdsServiceServer is the server API for AdsService service.
// All implementations must embed UnimplementedAdsServiceServer
// for forward compatibility
type AdsServiceServer interface {
	ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error)
	GetAd(context.Context, *GetAdRequest) (*GetAdResponse, error)
	CreateAd(context.Context, *CreateAdRequest) (*CreateAdResponse, error)
	mustEmbedUnimplementedAdsServiceServer()
}

// UnimplementedAdsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdsServiceServer struct {
}

func (UnimplementedAdsServiceServer) ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAds not implemented")
}
func (UnimplementedAdsServiceServer) GetAd(context.Context, *GetAdRequest) (*GetAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (UnimplementedAdsServiceServer) CreateAd(context.Context, *CreateAdRequest) (*CreateAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedAdsServiceServer) mustEmbedUnimplementedAdsServiceServer() {}

// UnsafeAdsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdsServiceServer will
// result in compilation errors.
type UnsafeAdsServiceServer interface {
	mustEmbedUnimplementedAdsServiceServer()
}

func RegisterAdsServiceServer(s grpc.ServiceRegistrar, srv AdsServiceServer) {
	s.RegisterService(&AdsService_ServiceDesc, srv)
}

func _AdsService_ListAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).ListAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_ListAds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).ListAds(ctx, req.(*ListAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdsService_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_GetAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).GetAd(ctx, req.(*GetAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdsService_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdsService_CreateAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).CreateAd(ctx, req.(*CreateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdsService_ServiceDesc is the grpc.ServiceDesc for AdsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.AdsService",
	HandlerType: (*AdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAds",
			Handler:    _AdsService_ListAds_Handler,
		},
		{
			MethodName: "GetAd",
			Handler:    _AdsService_GetAd_Handler,
		},
		{
			MethodName: "CreateAd",
			Handler:    _AdsService_CreateAd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/ads.proto",
}
