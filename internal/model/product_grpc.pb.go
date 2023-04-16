// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: product.proto

package model

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	ListProducts(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductListResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/model.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/model.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/model.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/model.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProducts(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductListResponse, error) {
	out := new(ProductListResponse)
	err := c.cc.Invoke(ctx, "/model.ProductService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *ProductCreateRequest) (*ProductResponse, error)
	GetProduct(context.Context, *ProductRequest) (*ProductResponse, error)
	UpdateProduct(context.Context, *ProductCreateRequest) (*ProductResponse, error)
	DeleteProduct(context.Context, *ProductRequest) (*ProductResponse, error)
	ListProducts(context.Context, *Product) (*ProductListResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *ProductCreateRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *ProductCreateRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) ListProducts(context.Context, *Product) (*ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ProductService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProducts(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _ProductService_ListProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}

// ProductImageServiceClient is the client API for ProductImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductImageServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (ProductImageService_UploadClient, error)
	Download(ctx context.Context, in *ProductImageDownloadRequest, opts ...grpc.CallOption) (ProductImageService_DownloadClient, error)
}

type productImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductImageServiceClient(cc grpc.ClientConnInterface) ProductImageServiceClient {
	return &productImageServiceClient{cc}
}

func (c *productImageServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (ProductImageService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductImageService_ServiceDesc.Streams[0], "/model.ProductImageService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &productImageServiceUploadClient{stream}
	return x, nil
}

type ProductImageService_UploadClient interface {
	Send(*ProductImageUploadRequest) error
	CloseAndRecv() (*ProductImageUploadResponse, error)
	grpc.ClientStream
}

type productImageServiceUploadClient struct {
	grpc.ClientStream
}

func (x *productImageServiceUploadClient) Send(m *ProductImageUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productImageServiceUploadClient) CloseAndRecv() (*ProductImageUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductImageUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productImageServiceClient) Download(ctx context.Context, in *ProductImageDownloadRequest, opts ...grpc.CallOption) (ProductImageService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductImageService_ServiceDesc.Streams[1], "/model.ProductImageService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &productImageServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductImageService_DownloadClient interface {
	Recv() (*ProductImageDownloadResponse, error)
	grpc.ClientStream
}

type productImageServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *productImageServiceDownloadClient) Recv() (*ProductImageDownloadResponse, error) {
	m := new(ProductImageDownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductImageServiceServer is the server API for ProductImageService service.
// All implementations must embed UnimplementedProductImageServiceServer
// for forward compatibility
type ProductImageServiceServer interface {
	Upload(ProductImageService_UploadServer) error
	Download(*ProductImageDownloadRequest, ProductImageService_DownloadServer) error
	mustEmbedUnimplementedProductImageServiceServer()
}

// UnimplementedProductImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductImageServiceServer struct {
}

func (UnimplementedProductImageServiceServer) Upload(ProductImageService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedProductImageServiceServer) Download(*ProductImageDownloadRequest, ProductImageService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedProductImageServiceServer) mustEmbedUnimplementedProductImageServiceServer() {}

// UnsafeProductImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductImageServiceServer will
// result in compilation errors.
type UnsafeProductImageServiceServer interface {
	mustEmbedUnimplementedProductImageServiceServer()
}

func RegisterProductImageServiceServer(s grpc.ServiceRegistrar, srv ProductImageServiceServer) {
	s.RegisterService(&ProductImageService_ServiceDesc, srv)
}

func _ProductImageService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductImageServiceServer).Upload(&productImageServiceUploadServer{stream})
}

type ProductImageService_UploadServer interface {
	SendAndClose(*ProductImageUploadResponse) error
	Recv() (*ProductImageUploadRequest, error)
	grpc.ServerStream
}

type productImageServiceUploadServer struct {
	grpc.ServerStream
}

func (x *productImageServiceUploadServer) SendAndClose(m *ProductImageUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productImageServiceUploadServer) Recv() (*ProductImageUploadRequest, error) {
	m := new(ProductImageUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductImageService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductImageDownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductImageServiceServer).Download(m, &productImageServiceDownloadServer{stream})
}

type ProductImageService_DownloadServer interface {
	Send(*ProductImageDownloadResponse) error
	grpc.ServerStream
}

type productImageServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *productImageServiceDownloadServer) Send(m *ProductImageDownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProductImageService_ServiceDesc is the grpc.ServiceDesc for ProductImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.ProductImageService",
	HandlerType: (*ProductImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _ProductImageService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _ProductImageService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "product.proto",
}
