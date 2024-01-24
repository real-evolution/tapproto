// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/ads/v1/quotas.proto

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

// QuotasServiceClient is the client API for QuotasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotasServiceClient interface {
	ListQuotas(ctx context.Context, in *ListQuotasRequest, opts ...grpc.CallOption) (*ListQuotasResponse, error)
	GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error)
}

type quotasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotasServiceClient(cc grpc.ClientConnInterface) QuotasServiceClient {
	return &quotasServiceClient{cc}
}

func (c *quotasServiceClient) ListQuotas(ctx context.Context, in *ListQuotasRequest, opts ...grpc.CallOption) (*ListQuotasResponse, error) {
	out := new(ListQuotasResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.QuotasService/ListQuotas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotasServiceClient) GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error) {
	out := new(GetQuotaResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.QuotasService/GetQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotasServiceServer is the server API for QuotasService service.
// All implementations must embed UnimplementedQuotasServiceServer
// for forward compatibility
type QuotasServiceServer interface {
	ListQuotas(context.Context, *ListQuotasRequest) (*ListQuotasResponse, error)
	GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error)
	mustEmbedUnimplementedQuotasServiceServer()
}

// UnimplementedQuotasServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuotasServiceServer struct {
}

func (UnimplementedQuotasServiceServer) ListQuotas(context.Context, *ListQuotasRequest) (*ListQuotasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuotas not implemented")
}
func (UnimplementedQuotasServiceServer) GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuota not implemented")
}
func (UnimplementedQuotasServiceServer) mustEmbedUnimplementedQuotasServiceServer() {}

// UnsafeQuotasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotasServiceServer will
// result in compilation errors.
type UnsafeQuotasServiceServer interface {
	mustEmbedUnimplementedQuotasServiceServer()
}

func RegisterQuotasServiceServer(s grpc.ServiceRegistrar, srv QuotasServiceServer) {
	s.RegisterService(&QuotasService_ServiceDesc, srv)
}

func _QuotasService_ListQuotas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuotasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotasServiceServer).ListQuotas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.QuotasService/ListQuotas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotasServiceServer).ListQuotas(ctx, req.(*ListQuotasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuotasService_GetQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotasServiceServer).GetQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.QuotasService/GetQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotasServiceServer).GetQuota(ctx, req.(*GetQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuotasService_ServiceDesc is the grpc.ServiceDesc for QuotasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuotasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.QuotasService",
	HandlerType: (*QuotasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQuotas",
			Handler:    _QuotasService_ListQuotas_Handler,
		},
		{
			MethodName: "GetQuota",
			Handler:    _QuotasService_GetQuota_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/quotas.proto",
}
