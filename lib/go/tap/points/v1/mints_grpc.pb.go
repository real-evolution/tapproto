// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tap/points/v1/mints.proto

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

// MintsServiceClient is the client API for MintsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MintsServiceClient interface {
	// Lists mint transactions.
	ListMints(ctx context.Context, in *ListMintsRequest, opts ...grpc.CallOption) (*ListMintsResponse, error)
	// Gets a mint transaction by id.
	GetMint(ctx context.Context, in *GetMintRequest, opts ...grpc.CallOption) (*GetMintResponse, error)
	// Mints tokens with the given parameters.
	Mint(ctx context.Context, in *MintRequest, opts ...grpc.CallOption) (*MintResponse, error)
}

type mintsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMintsServiceClient(cc grpc.ClientConnInterface) MintsServiceClient {
	return &mintsServiceClient{cc}
}

func (c *mintsServiceClient) ListMints(ctx context.Context, in *ListMintsRequest, opts ...grpc.CallOption) (*ListMintsResponse, error) {
	out := new(ListMintsResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.MintsService/ListMints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mintsServiceClient) GetMint(ctx context.Context, in *GetMintRequest, opts ...grpc.CallOption) (*GetMintResponse, error) {
	out := new(GetMintResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.MintsService/GetMint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mintsServiceClient) Mint(ctx context.Context, in *MintRequest, opts ...grpc.CallOption) (*MintResponse, error) {
	out := new(MintResponse)
	err := c.cc.Invoke(ctx, "/tap.points.v1.MintsService/Mint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MintsServiceServer is the server API for MintsService service.
// All implementations must embed UnimplementedMintsServiceServer
// for forward compatibility
type MintsServiceServer interface {
	// Lists mint transactions.
	ListMints(context.Context, *ListMintsRequest) (*ListMintsResponse, error)
	// Gets a mint transaction by id.
	GetMint(context.Context, *GetMintRequest) (*GetMintResponse, error)
	// Mints tokens with the given parameters.
	Mint(context.Context, *MintRequest) (*MintResponse, error)
	mustEmbedUnimplementedMintsServiceServer()
}

// UnimplementedMintsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMintsServiceServer struct {
}

func (UnimplementedMintsServiceServer) ListMints(context.Context, *ListMintsRequest) (*ListMintsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMints not implemented")
}
func (UnimplementedMintsServiceServer) GetMint(context.Context, *GetMintRequest) (*GetMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMint not implemented")
}
func (UnimplementedMintsServiceServer) Mint(context.Context, *MintRequest) (*MintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (UnimplementedMintsServiceServer) mustEmbedUnimplementedMintsServiceServer() {}

// UnsafeMintsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MintsServiceServer will
// result in compilation errors.
type UnsafeMintsServiceServer interface {
	mustEmbedUnimplementedMintsServiceServer()
}

func RegisterMintsServiceServer(s grpc.ServiceRegistrar, srv MintsServiceServer) {
	s.RegisterService(&MintsService_ServiceDesc, srv)
}

func _MintsService_ListMints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMintsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MintsServiceServer).ListMints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.MintsService/ListMints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MintsServiceServer).ListMints(ctx, req.(*ListMintsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MintsService_GetMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MintsServiceServer).GetMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.MintsService/GetMint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MintsServiceServer).GetMint(ctx, req.(*GetMintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MintsService_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MintsServiceServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tap.points.v1.MintsService/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MintsServiceServer).Mint(ctx, req.(*MintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MintsService_ServiceDesc is the grpc.ServiceDesc for MintsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MintsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.points.v1.MintsService",
	HandlerType: (*MintsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMints",
			Handler:    _MintsService_ListMints_Handler,
		},
		{
			MethodName: "GetMint",
			Handler:    _MintsService_GetMint_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _MintsService_Mint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/points/v1/mints.proto",
}
