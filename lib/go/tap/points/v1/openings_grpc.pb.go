// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/points/v1/openings.proto

package pointsv1

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

// OpeningsServiceClient is the client API for OpeningsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpeningsServiceClient interface {
	// Lsts openings.
	ListOpenings(ctx context.Context, in *ListOpeningsRequest, opts ...grpc.CallOption) (*ListOpeningsResponse, error)
	// Gets an opening.
	GetOpening(ctx context.Context, in *GetOpeningRequest, opts ...grpc.CallOption) (*GetOpeningResponse, error)
	// Checks if there is an active opening.
	HasActiveOpening(ctx context.Context, in *HasActiveOpeningRequest, opts ...grpc.CallOption) (*HasActiveOpeningResponse, error)
	// Creates an opening.
	CreateOpening(ctx context.Context, in *CreateOpeningRequest, opts ...grpc.CallOption) (*CreateOpeningResponse, error)
	// Cancels an opening.
	CancelOpening(ctx context.Context, in *CancelOpeningRequest, opts ...grpc.CallOption) (*CancelOpeningResponse, error)
}

type openingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpeningsServiceClient(cc grpc.ClientConnInterface) OpeningsServiceClient {
	return &openingsServiceClient{cc}
}

func (c *openingsServiceClient) ListOpenings(ctx context.Context, in *ListOpeningsRequest, opts ...grpc.CallOption) (*ListOpeningsResponse, error) {
	out := new(ListOpeningsResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.OpeningsService/ListOpenings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openingsServiceClient) GetOpening(ctx context.Context, in *GetOpeningRequest, opts ...grpc.CallOption) (*GetOpeningResponse, error) {
	out := new(GetOpeningResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.OpeningsService/GetOpening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openingsServiceClient) HasActiveOpening(ctx context.Context, in *HasActiveOpeningRequest, opts ...grpc.CallOption) (*HasActiveOpeningResponse, error) {
	out := new(HasActiveOpeningResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.OpeningsService/HasActiveOpening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openingsServiceClient) CreateOpening(ctx context.Context, in *CreateOpeningRequest, opts ...grpc.CallOption) (*CreateOpeningResponse, error) {
	out := new(CreateOpeningResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.OpeningsService/CreateOpening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openingsServiceClient) CancelOpening(ctx context.Context, in *CancelOpeningRequest, opts ...grpc.CallOption) (*CancelOpeningResponse, error) {
	out := new(CancelOpeningResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.OpeningsService/CancelOpening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpeningsServiceServer is the server API for OpeningsService service.
// All implementations must embed UnimplementedOpeningsServiceServer
// for forward compatibility
type OpeningsServiceServer interface {
	// Lsts openings.
	ListOpenings(context.Context, *ListOpeningsRequest) (*ListOpeningsResponse, error)
	// Gets an opening.
	GetOpening(context.Context, *GetOpeningRequest) (*GetOpeningResponse, error)
	// Checks if there is an active opening.
	HasActiveOpening(context.Context, *HasActiveOpeningRequest) (*HasActiveOpeningResponse, error)
	// Creates an opening.
	CreateOpening(context.Context, *CreateOpeningRequest) (*CreateOpeningResponse, error)
	// Cancels an opening.
	CancelOpening(context.Context, *CancelOpeningRequest) (*CancelOpeningResponse, error)
	mustEmbedUnimplementedOpeningsServiceServer()
}

// UnimplementedOpeningsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpeningsServiceServer struct {
}

func (UnimplementedOpeningsServiceServer) ListOpenings(context.Context, *ListOpeningsRequest) (*ListOpeningsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOpenings not implemented")
}
func (UnimplementedOpeningsServiceServer) GetOpening(context.Context, *GetOpeningRequest) (*GetOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpening not implemented")
}
func (UnimplementedOpeningsServiceServer) HasActiveOpening(context.Context, *HasActiveOpeningRequest) (*HasActiveOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasActiveOpening not implemented")
}
func (UnimplementedOpeningsServiceServer) CreateOpening(context.Context, *CreateOpeningRequest) (*CreateOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOpening not implemented")
}
func (UnimplementedOpeningsServiceServer) CancelOpening(context.Context, *CancelOpeningRequest) (*CancelOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOpening not implemented")
}
func (UnimplementedOpeningsServiceServer) mustEmbedUnimplementedOpeningsServiceServer() {}

// UnsafeOpeningsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpeningsServiceServer will
// result in compilation errors.
type UnsafeOpeningsServiceServer interface {
	mustEmbedUnimplementedOpeningsServiceServer()
}

func RegisterOpeningsServiceServer(s grpc.ServiceRegistrar, srv OpeningsServiceServer) {
	s.RegisterService(&OpeningsService_ServiceDesc, srv)
}

func _OpeningsService_ListOpenings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOpeningsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpeningsServiceServer).ListOpenings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.OpeningsService/ListOpenings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpeningsServiceServer).ListOpenings(ctx, req.(*ListOpeningsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpeningsService_GetOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpeningsServiceServer).GetOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.OpeningsService/GetOpening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpeningsServiceServer).GetOpening(ctx, req.(*GetOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpeningsService_HasActiveOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasActiveOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpeningsServiceServer).HasActiveOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.OpeningsService/HasActiveOpening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpeningsServiceServer).HasActiveOpening(ctx, req.(*HasActiveOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpeningsService_CreateOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpeningsServiceServer).CreateOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.OpeningsService/CreateOpening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpeningsServiceServer).CreateOpening(ctx, req.(*CreateOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpeningsService_CancelOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpeningsServiceServer).CancelOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.OpeningsService/CancelOpening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpeningsServiceServer).CancelOpening(ctx, req.(*CancelOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OpeningsService_ServiceDesc is the grpc.ServiceDesc for OpeningsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpeningsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.points.v1.OpeningsService",
	HandlerType: (*OpeningsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOpenings",
			Handler:    _OpeningsService_ListOpenings_Handler,
		},
		{
			MethodName: "GetOpening",
			Handler:    _OpeningsService_GetOpening_Handler,
		},
		{
			MethodName: "HasActiveOpening",
			Handler:    _OpeningsService_HasActiveOpening_Handler,
		},
		{
			MethodName: "CreateOpening",
			Handler:    _OpeningsService_CreateOpening_Handler,
		},
		{
			MethodName: "CancelOpening",
			Handler:    _OpeningsService_CancelOpening_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/points/v1/openings.proto",
}
