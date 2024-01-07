// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/sub/v1/providers.proto

package subv1

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

// ProvidersServiceClient is the client API for ProvidersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvidersServiceClient interface {
	// Returns the list of providers matching the given filter.
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error)
	// Gets a provider by id.
	GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error)
	// Creates a provider.
	CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error)
	// Updates a provider.
	UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderResponse, error)
	// Deletes a provider.
	DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderResponse, error)
}

type providersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvidersServiceClient(cc grpc.ClientConnInterface) ProvidersServiceClient {
	return &providersServiceClient{cc}
}

func (c *providersServiceClient) ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error) {
	out := new(ListProvidersResponse)
	err := c.cc.Invoke(ctx, "/tap.sub.v1.ProvidersService/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersServiceClient) GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error) {
	out := new(GetProviderResponse)
	err := c.cc.Invoke(ctx, "/tap.sub.v1.ProvidersService/GetProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersServiceClient) CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderResponse, error) {
	out := new(CreateProviderResponse)
	err := c.cc.Invoke(ctx, "/tap.sub.v1.ProvidersService/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersServiceClient) UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderResponse, error) {
	out := new(UpdateProviderResponse)
	err := c.cc.Invoke(ctx, "/tap.sub.v1.ProvidersService/UpdateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providersServiceClient) DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderResponse, error) {
	out := new(DeleteProviderResponse)
	err := c.cc.Invoke(ctx, "/tap.sub.v1.ProvidersService/DeleteProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvidersServiceServer is the server API for ProvidersService service.
// All implementations must embed UnimplementedProvidersServiceServer
// for forward compatibility
type ProvidersServiceServer interface {
	// Returns the list of providers matching the given filter.
	ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error)
	// Gets a provider by id.
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error)
	// Creates a provider.
	CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error)
	// Updates a provider.
	UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderResponse, error)
	// Deletes a provider.
	DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderResponse, error)
	mustEmbedUnimplementedProvidersServiceServer()
}

// UnimplementedProvidersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProvidersServiceServer struct {
}

func (UnimplementedProvidersServiceServer) ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedProvidersServiceServer) GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvider not implemented")
}
func (UnimplementedProvidersServiceServer) CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProvidersServiceServer) UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (UnimplementedProvidersServiceServer) DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
func (UnimplementedProvidersServiceServer) mustEmbedUnimplementedProvidersServiceServer() {}

// UnsafeProvidersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvidersServiceServer will
// result in compilation errors.
type UnsafeProvidersServiceServer interface {
	mustEmbedUnimplementedProvidersServiceServer()
}

func RegisterProvidersServiceServer(s grpc.ServiceRegistrar, srv ProvidersServiceServer) {
	s.RegisterService(&ProvidersService_ServiceDesc, srv)
}

func _ProvidersService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.sub.v1.ProvidersService/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersServiceServer).ListProviders(ctx, req.(*ListProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersService_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersServiceServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.sub.v1.ProvidersService/GetProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersServiceServer).GetProvider(ctx, req.(*GetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersService_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersServiceServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.sub.v1.ProvidersService/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersServiceServer).CreateProvider(ctx, req.(*CreateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersService_UpdateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersServiceServer).UpdateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.sub.v1.ProvidersService/UpdateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersServiceServer).UpdateProvider(ctx, req.(*UpdateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvidersService_DeleteProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvidersServiceServer).DeleteProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.sub.v1.ProvidersService/DeleteProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvidersServiceServer).DeleteProvider(ctx, req.(*DeleteProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvidersService_ServiceDesc is the grpc.ServiceDesc for ProvidersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvidersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.sub.v1.ProvidersService",
	HandlerType: (*ProvidersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProviders",
			Handler:    _ProvidersService_ListProviders_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _ProvidersService_GetProvider_Handler,
		},
		{
			MethodName: "CreateProvider",
			Handler:    _ProvidersService_CreateProvider_Handler,
		},
		{
			MethodName: "UpdateProvider",
			Handler:    _ProvidersService_UpdateProvider_Handler,
		},
		{
			MethodName: "DeleteProvider",
			Handler:    _ProvidersService_DeleteProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/sub/v1/providers.proto",
}