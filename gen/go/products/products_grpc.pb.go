// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: products/products.proto

package productv1

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

// ProductServClient is the client API for ProductServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	BeginOrder(ctx context.Context, in *BeginOrderRequest, opts ...grpc.CallOption) (*BeginOrderResponse, error)
	ApplyOrder(ctx context.Context, in *ApplyOrderRequest, opts ...grpc.CallOption) (*ApplyOrderResponse, error)
}

type productServClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServClient(cc grpc.ClientConnInterface) ProductServClient {
	return &productServClient{cc}
}

func (c *productServClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.ProductServ/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductServ/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductServ/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServClient) BeginOrder(ctx context.Context, in *BeginOrderRequest, opts ...grpc.CallOption) (*BeginOrderResponse, error) {
	out := new(BeginOrderResponse)
	err := c.cc.Invoke(ctx, "/product.ProductServ/BeginOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServClient) ApplyOrder(ctx context.Context, in *ApplyOrderRequest, opts ...grpc.CallOption) (*ApplyOrderResponse, error) {
	out := new(ApplyOrderResponse)
	err := c.cc.Invoke(ctx, "/product.ProductServ/ApplyOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServServer is the server API for ProductServ service.
// All implementations must embed UnimplementedProductServServer
// for forward compatibility
type ProductServServer interface {
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	BeginOrder(context.Context, *BeginOrderRequest) (*BeginOrderResponse, error)
	ApplyOrder(context.Context, *ApplyOrderRequest) (*ApplyOrderResponse, error)
	mustEmbedUnimplementedProductServServer()
}

// UnimplementedProductServServer must be embedded to have forward compatible implementations.
type UnimplementedProductServServer struct {
}

func (UnimplementedProductServServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductServServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServServer) BeginOrder(context.Context, *BeginOrderRequest) (*BeginOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginOrder not implemented")
}
func (UnimplementedProductServServer) ApplyOrder(context.Context, *ApplyOrderRequest) (*ApplyOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyOrder not implemented")
}
func (UnimplementedProductServServer) mustEmbedUnimplementedProductServServer() {}

// UnsafeProductServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServServer will
// result in compilation errors.
type UnsafeProductServServer interface {
	mustEmbedUnimplementedProductServServer()
}

func RegisterProductServServer(s grpc.ServiceRegistrar, srv ProductServServer) {
	s.RegisterService(&ProductServ_ServiceDesc, srv)
}

func _ProductServ_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductServ/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductServ_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductServ/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductServ_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductServ/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductServ_BeginOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServServer).BeginOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductServ/BeginOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServServer).BeginOrder(ctx, req.(*BeginOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductServ_ApplyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServServer).ApplyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductServ/ApplyOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServServer).ApplyOrder(ctx, req.(*ApplyOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductServ_ServiceDesc is the grpc.ServiceDesc for ProductServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductServ",
	HandlerType: (*ProductServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductServ_GetProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductServ_GetProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductServ_CreateProduct_Handler,
		},
		{
			MethodName: "BeginOrder",
			Handler:    _ProductServ_BeginOrder_Handler,
		},
		{
			MethodName: "ApplyOrder",
			Handler:    _ProductServ_ApplyOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products/products.proto",
}
