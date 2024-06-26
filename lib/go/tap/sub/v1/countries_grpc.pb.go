// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: tap/sub/v1/countries.proto

package subv1

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
	CountriesService_ListCountries_FullMethodName = "/tap.sub.v1.CountriesService/ListCountries"
	CountriesService_GetCountry_FullMethodName    = "/tap.sub.v1.CountriesService/GetCountry"
	CountriesService_UpdateCountry_FullMethodName = "/tap.sub.v1.CountriesService/UpdateCountry"
)

// CountriesServiceClient is the client API for CountriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The countries service definition.
type CountriesServiceClient interface {
	// Returns a list of countries.
	ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error)
	// Gets a country by its iso 3166-1 code.
	GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error)
	// Updates a country.
	UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error)
}

type countriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountriesServiceClient(cc grpc.ClientConnInterface) CountriesServiceClient {
	return &countriesServiceClient{cc}
}

func (c *countriesServiceClient) ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCountriesResponse)
	err := c.cc.Invoke(ctx, CountriesService_ListCountries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesServiceClient) GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCountryResponse)
	err := c.cc.Invoke(ctx, CountriesService_GetCountry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesServiceClient) UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCountryResponse)
	err := c.cc.Invoke(ctx, CountriesService_UpdateCountry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountriesServiceServer is the server API for CountriesService service.
// All implementations must embed UnimplementedCountriesServiceServer
// for forward compatibility
//
// The countries service definition.
type CountriesServiceServer interface {
	// Returns a list of countries.
	ListCountries(context.Context, *ListCountriesRequest) (*ListCountriesResponse, error)
	// Gets a country by its iso 3166-1 code.
	GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error)
	// Updates a country.
	UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error)
	mustEmbedUnimplementedCountriesServiceServer()
}

// UnimplementedCountriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountriesServiceServer struct {
}

func (UnimplementedCountriesServiceServer) ListCountries(context.Context, *ListCountriesRequest) (*ListCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCountries not implemented")
}
func (UnimplementedCountriesServiceServer) GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountry not implemented")
}
func (UnimplementedCountriesServiceServer) UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCountry not implemented")
}
func (UnimplementedCountriesServiceServer) mustEmbedUnimplementedCountriesServiceServer() {}

// UnsafeCountriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountriesServiceServer will
// result in compilation errors.
type UnsafeCountriesServiceServer interface {
	mustEmbedUnimplementedCountriesServiceServer()
}

func RegisterCountriesServiceServer(s grpc.ServiceRegistrar, srv CountriesServiceServer) {
	s.RegisterService(&CountriesService_ServiceDesc, srv)
}

func _CountriesService_ListCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServiceServer).ListCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountriesService_ListCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServiceServer).ListCountries(ctx, req.(*ListCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountriesService_GetCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServiceServer).GetCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountriesService_GetCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServiceServer).GetCountry(ctx, req.(*GetCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountriesService_UpdateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServiceServer).UpdateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountriesService_UpdateCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServiceServer).UpdateCountry(ctx, req.(*UpdateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountriesService_ServiceDesc is the grpc.ServiceDesc for CountriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.sub.v1.CountriesService",
	HandlerType: (*CountriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCountries",
			Handler:    _CountriesService_ListCountries_Handler,
		},
		{
			MethodName: "GetCountry",
			Handler:    _CountriesService_GetCountry_Handler,
		},
		{
			MethodName: "UpdateCountry",
			Handler:    _CountriesService_UpdateCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/sub/v1/countries.proto",
}
