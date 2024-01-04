// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/ads/v1/requirements.proto

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

// RequirementsServiceClient is the client API for RequirementsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequirementsServiceClient interface {
	// Lists requirement.
	ListRequirements(ctx context.Context, in *ListRequirementsRequest, opts ...grpc.CallOption) (*ListRequirementsResponse, error)
	// Creates a requirement.
	CreateRequirement(ctx context.Context, in *CreateRequirementRequest, opts ...grpc.CallOption) (*CreateRequirementResponse, error)
}

type requirementsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequirementsServiceClient(cc grpc.ClientConnInterface) RequirementsServiceClient {
	return &requirementsServiceClient{cc}
}

func (c *requirementsServiceClient) ListRequirements(ctx context.Context, in *ListRequirementsRequest, opts ...grpc.CallOption) (*ListRequirementsResponse, error) {
	out := new(ListRequirementsResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.RequirementsService/ListRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requirementsServiceClient) CreateRequirement(ctx context.Context, in *CreateRequirementRequest, opts ...grpc.CallOption) (*CreateRequirementResponse, error) {
	out := new(CreateRequirementResponse)
	err := c.cc.Invoke(ctx, "/tap.ads.v1.RequirementsService/CreateRequirement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequirementsServiceServer is the server API for RequirementsService service.
// All implementations must embed UnimplementedRequirementsServiceServer
// for forward compatibility
type RequirementsServiceServer interface {
	// Lists requirement.
	ListRequirements(context.Context, *ListRequirementsRequest) (*ListRequirementsResponse, error)
	// Creates a requirement.
	CreateRequirement(context.Context, *CreateRequirementRequest) (*CreateRequirementResponse, error)
	mustEmbedUnimplementedRequirementsServiceServer()
}

// UnimplementedRequirementsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequirementsServiceServer struct {
}

func (UnimplementedRequirementsServiceServer) ListRequirements(context.Context, *ListRequirementsRequest) (*ListRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequirements not implemented")
}
func (UnimplementedRequirementsServiceServer) CreateRequirement(context.Context, *CreateRequirementRequest) (*CreateRequirementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequirement not implemented")
}
func (UnimplementedRequirementsServiceServer) mustEmbedUnimplementedRequirementsServiceServer() {}

// UnsafeRequirementsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequirementsServiceServer will
// result in compilation errors.
type UnsafeRequirementsServiceServer interface {
	mustEmbedUnimplementedRequirementsServiceServer()
}

func RegisterRequirementsServiceServer(s grpc.ServiceRegistrar, srv RequirementsServiceServer) {
	s.RegisterService(&RequirementsService_ServiceDesc, srv)
}

func _RequirementsService_ListRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequirementsServiceServer).ListRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.RequirementsService/ListRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequirementsServiceServer).ListRequirements(ctx, req.(*ListRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequirementsService_CreateRequirement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequirementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequirementsServiceServer).CreateRequirement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.ads.v1.RequirementsService/CreateRequirement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequirementsServiceServer).CreateRequirement(ctx, req.(*CreateRequirementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequirementsService_ServiceDesc is the grpc.ServiceDesc for RequirementsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequirementsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.RequirementsService",
	HandlerType: (*RequirementsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRequirements",
			Handler:    _RequirementsService_ListRequirements_Handler,
		},
		{
			MethodName: "CreateRequirement",
			Handler:    _RequirementsService_CreateRequirement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/requirements.proto",
}
