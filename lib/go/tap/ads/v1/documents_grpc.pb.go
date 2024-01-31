// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tap/ads/v1/documents.proto

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
	DocumentsService_ListDocuments_FullMethodName  = "/tap.ads.v1.DocumentsService/ListDocuments"
	DocumentsService_GetDocument_FullMethodName    = "/tap.ads.v1.DocumentsService/GetDocument"
	DocumentsService_CreateDocument_FullMethodName = "/tap.ads.v1.DocumentsService/CreateDocument"
)

// DocumentsServiceClient is the client API for DocumentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentsServiceClient interface {
	// Lists documents.
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error)
	// Gets a document.
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	// Creates a document.
	CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error)
}

type documentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentsServiceClient(cc grpc.ClientConnInterface) DocumentsServiceClient {
	return &documentsServiceClient{cc}
}

func (c *documentsServiceClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error) {
	out := new(ListDocumentsResponse)
	err := c.cc.Invoke(ctx, DocumentsService_ListDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsServiceClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, DocumentsService_GetDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsServiceClient) CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error) {
	out := new(CreateDocumentResponse)
	err := c.cc.Invoke(ctx, DocumentsService_CreateDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsServiceServer is the server API for DocumentsService service.
// All implementations must embed UnimplementedDocumentsServiceServer
// for forward compatibility
type DocumentsServiceServer interface {
	// Lists documents.
	ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error)
	// Gets a document.
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	// Creates a document.
	CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error)
	mustEmbedUnimplementedDocumentsServiceServer()
}

// UnimplementedDocumentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentsServiceServer struct {
}

func (UnimplementedDocumentsServiceServer) ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedDocumentsServiceServer) GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedDocumentsServiceServer) CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedDocumentsServiceServer) mustEmbedUnimplementedDocumentsServiceServer() {}

// UnsafeDocumentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentsServiceServer will
// result in compilation errors.
type UnsafeDocumentsServiceServer interface {
	mustEmbedUnimplementedDocumentsServiceServer()
}

func RegisterDocumentsServiceServer(s grpc.ServiceRegistrar, srv DocumentsServiceServer) {
	s.RegisterService(&DocumentsService_ServiceDesc, srv)
}

func _DocumentsService_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_ListDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).ListDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentsService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_GetDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentsService_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_CreateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).CreateDocument(ctx, req.(*CreateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentsService_ServiceDesc is the grpc.ServiceDesc for DocumentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tap.ads.v1.DocumentsService",
	HandlerType: (*DocumentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDocuments",
			Handler:    _DocumentsService_ListDocuments_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _DocumentsService_GetDocument_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _DocumentsService_CreateDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tap/ads/v1/documents.proto",
}
