// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: tap/points/v1/transfers.proto

package pointsv1

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
	TransfersService_ListTransfers_FullMethodName        = "/tap.points.v1.TransfersService/ListTransfers"
	TransfersService_GetTransfer_FullMethodName          = "/tap.points.v1.TransfersService/GetTransfer"
	TransfersService_MakeTransfer_FullMethodName         = "/tap.points.v1.TransfersService/MakeTransfer"
	TransfersService_MakeTransferToHolder_FullMethodName = "/tap.points.v1.TransfersService/MakeTransferToHolder"
	TransfersService_GetReceipt_FullMethodName           = "/tap.points.v1.TransfersService/GetReceipt"
)

// TransfersServiceClient is the client API for TransfersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Transfers management service definition.
type TransfersServiceClient interface {
	// ListTransfers lists transfer transactions.
	ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error)
	// GetTransfer returns a transfer transaction.
	GetTransfer(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*GetTransferResponse, error)
	// MakeTransfer creates a new transfer transaction.
	MakeTransfer(ctx context.Context, in *MakeTransferRequest, opts ...grpc.CallOption) (*MakeTransferResponse, error)
	// MakeTransferToHolder creates a new transfer transaction to a holder.
	MakeTransferToHolder(ctx context.Context, in *MakeTransferToHolderRequest, opts ...grpc.CallOption) (*MakeTransferToHolderResponse, error)
	// GetReceipt returns a transfer receipt.
	GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error)
}

type transfersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransfersServiceClient(cc grpc.ClientConnInterface) TransfersServiceClient {
	return &transfersServiceClient{cc}
}

func (c *transfersServiceClient) ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransfersResponse)
	err := c.cc.Invoke(ctx, TransfersService_ListTransfers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersServiceClient) GetTransfer(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*GetTransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransferResponse)
	err := c.cc.Invoke(ctx, TransfersService_GetTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersServiceClient) MakeTransfer(ctx context.Context, in *MakeTransferRequest, opts ...grpc.CallOption) (*MakeTransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeTransferResponse)
	err := c.cc.Invoke(ctx, TransfersService_MakeTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersServiceClient) MakeTransferToHolder(ctx context.Context, in *MakeTransferToHolderRequest, opts ...grpc.CallOption) (*MakeTransferToHolderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeTransferToHolderResponse)
	err := c.cc.Invoke(ctx, TransfersService_MakeTransferToHolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersServiceClient) GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReceiptResponse)
	err := c.cc.Invoke(ctx, TransfersService_GetReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServiceServer is the server API for TransfersService service.
// All implementations must embed UnimplementedTransfersServiceServer
// for forward compatibility
//
// Transfers management service definition.
type TransfersServiceServer interface {
	// ListTransfers lists transfer transactions.
	ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error)
	// GetTransfer returns a transfer transaction.
	GetTransfer(context.Context, *GetTransferRequest) (*GetTransferResponse, error)
	// MakeTransfer creates a new transfer transaction.
	MakeTransfer(context.Context, *MakeTransferRequest) (*MakeTransferResponse, error)
	// MakeTransferToHolder creates a new transfer transaction to a holder.
	MakeTransferToHolder(context.Context, *MakeTransferToHolderRequest) (*MakeTransferToHolderResponse, error)
	// GetReceipt returns a transfer receipt.
	GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error)
	mustEmbedUnimplementedTransfersServiceServer()
}

// UnimplementedTransfersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransfersServiceServer struct {
}

func (UnimplementedTransfersServiceServer) ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransfers not implemented")
}
func (UnimplementedTransfersServiceServer) GetTransfer(context.Context, *GetTransferRequest) (*GetTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer not implemented")
}
func (UnimplementedTransfersServiceServer) MakeTransfer(context.Context, *MakeTransferRequest) (*MakeTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransfer not implemented")
}
func (UnimplementedTransfersServiceServer) MakeTransferToHolder(context.Context, *MakeTransferToHolderRequest) (*MakeTransferToHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransferToHolder not implemented")
}
func (UnimplementedTransfersServiceServer) GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTransfersServiceServer) mustEmbedUnimplementedTransfersServiceServer() {}

// UnsafeTransfersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransfersServiceServer will
// result in compilation errors.
type UnsafeTransfersServiceServer interface {
	mustEmbedUnimplementedTransfersServiceServer()
}

func RegisterTransfersServiceServer(s grpc.ServiceRegistrar, srv TransfersServiceServer) {
	s.RegisterService(&TransfersService_ServiceDesc, srv)
}

func _TransfersService_ListTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).ListTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfersService_ListTransfers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).ListTransfers(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransfersService_GetTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).GetTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfersService_GetTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).GetTransfer(ctx, req.(*GetTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransfersService_MakeTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).MakeTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfersService_MakeTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).MakeTransfer(ctx, req.(*MakeTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransfersService_MakeTransferToHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeTransferToHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).MakeTransferToHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfersService_MakeTransferToHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).MakeTransferToHolder(ctx, req.(*MakeTransferToHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransfersService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfersService_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).GetReceipt(ctx, req.(*GetReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransfersService_ServiceDesc is the grpc.ServiceDesc for TransfersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransfersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.points.v1.TransfersService",
	HandlerType: (*TransfersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransfers",
			Handler:    _TransfersService_ListTransfers_Handler,
		},
		{
			MethodName: "GetTransfer",
			Handler:    _TransfersService_GetTransfer_Handler,
		},
		{
			MethodName: "MakeTransfer",
			Handler:    _TransfersService_MakeTransfer_Handler,
		},
		{
			MethodName: "MakeTransferToHolder",
			Handler:    _TransfersService_MakeTransferToHolder_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TransfersService_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/points/v1/transfers.proto",
}
