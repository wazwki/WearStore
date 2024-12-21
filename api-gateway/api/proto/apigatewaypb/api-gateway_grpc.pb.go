// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: github.com/wazwki/WearStore/api-gateway/api/proto/api-gateway.proto

package apigatewaypb

import (
	context "context"
	cartpb "github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	productpb "github.com/wazwki/WearStore/product-service/api/proto/productpb"
	userpb "github.com/wazwki/WearStore/user-service/api/proto/userpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ApiGatewayService_RegisterUser_FullMethodName   = "/apigatewaypb.ApiGatewayService/RegisterUser"
	ApiGatewayService_LoginUser_FullMethodName      = "/apigatewaypb.ApiGatewayService/LoginUser"
	ApiGatewayService_GetUser_FullMethodName        = "/apigatewaypb.ApiGatewayService/GetUser"
	ApiGatewayService_UpdateUser_FullMethodName     = "/apigatewaypb.ApiGatewayService/UpdateUser"
	ApiGatewayService_DeleteUser_FullMethodName     = "/apigatewaypb.ApiGatewayService/DeleteUser"
	ApiGatewayService_GetProduct_FullMethodName     = "/apigatewaypb.ApiGatewayService/GetProduct"
	ApiGatewayService_ListProducts_FullMethodName   = "/apigatewaypb.ApiGatewayService/ListProducts"
	ApiGatewayService_AddProduct_FullMethodName     = "/apigatewaypb.ApiGatewayService/AddProduct"
	ApiGatewayService_UpdateProduct_FullMethodName  = "/apigatewaypb.ApiGatewayService/UpdateProduct"
	ApiGatewayService_DeleteProduct_FullMethodName  = "/apigatewaypb.ApiGatewayService/DeleteProduct"
	ApiGatewayService_AddToCart_FullMethodName      = "/apigatewaypb.ApiGatewayService/AddToCart"
	ApiGatewayService_RemoveFromCart_FullMethodName = "/apigatewaypb.ApiGatewayService/RemoveFromCart"
	ApiGatewayService_GetCart_FullMethodName        = "/apigatewaypb.ApiGatewayService/GetCart"
	ApiGatewayService_ClearCart_FullMethodName      = "/apigatewaypb.ApiGatewayService/ClearCart"
	ApiGatewayService_Checkout_FullMethodName       = "/apigatewaypb.ApiGatewayService/Checkout"
)

// ApiGatewayServiceClient is the client API for ApiGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiGatewayServiceClient interface {
	// User Service
	RegisterUser(ctx context.Context, in *userpb.RegisterUserRequest, opts ...grpc.CallOption) (*userpb.RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *userpb.LoginUserRequest, opts ...grpc.CallOption) (*userpb.LoginUserResponse, error)
	GetUser(ctx context.Context, in *userpb.GetUserRequest, opts ...grpc.CallOption) (*userpb.GetUserResponse, error)
	UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest, opts ...grpc.CallOption) (*userpb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest, opts ...grpc.CallOption) (*userpb.DeleteUserResponse, error)
	// Product Service
	GetProduct(ctx context.Context, in *productpb.GetProductRequest, opts ...grpc.CallOption) (*productpb.GetProductResponse, error)
	ListProducts(ctx context.Context, in *productpb.ListProductsRequest, opts ...grpc.CallOption) (*productpb.ListProductsResponse, error)
	AddProduct(ctx context.Context, in *productpb.AddProductRequest, opts ...grpc.CallOption) (*productpb.AddProductResponse, error)
	UpdateProduct(ctx context.Context, in *productpb.UpdateProductRequest, opts ...grpc.CallOption) (*productpb.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *productpb.DeleteProductRequest, opts ...grpc.CallOption) (*productpb.DeleteProductResponse, error)
	// Cart Service
	AddToCart(ctx context.Context, in *cartpb.AddToCartRequest, opts ...grpc.CallOption) (*cartpb.AddToCartResponse, error)
	RemoveFromCart(ctx context.Context, in *cartpb.RemoveFromCartRequest, opts ...grpc.CallOption) (*cartpb.RemoveFromCartResponse, error)
	GetCart(ctx context.Context, in *cartpb.GetCartRequest, opts ...grpc.CallOption) (*cartpb.GetCartResponse, error)
	ClearCart(ctx context.Context, in *cartpb.ClearCartRequest, opts ...grpc.CallOption) (*cartpb.ClearCartResponse, error)
	Checkout(ctx context.Context, in *cartpb.CheckoutRequest, opts ...grpc.CallOption) (*cartpb.CheckoutResponse, error)
}

type apiGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiGatewayServiceClient(cc grpc.ClientConnInterface) ApiGatewayServiceClient {
	return &apiGatewayServiceClient{cc}
}

func (c *apiGatewayServiceClient) RegisterUser(ctx context.Context, in *userpb.RegisterUserRequest, opts ...grpc.CallOption) (*userpb.RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(userpb.RegisterUserResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) LoginUser(ctx context.Context, in *userpb.LoginUserRequest, opts ...grpc.CallOption) (*userpb.LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(userpb.LoginUserResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetUser(ctx context.Context, in *userpb.GetUserRequest, opts ...grpc.CallOption) (*userpb.GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(userpb.GetUserResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest, opts ...grpc.CallOption) (*userpb.UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(userpb.UpdateUserResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest, opts ...grpc.CallOption) (*userpb.DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(userpb.DeleteUserResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetProduct(ctx context.Context, in *productpb.GetProductRequest, opts ...grpc.CallOption) (*productpb.GetProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(productpb.GetProductResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) ListProducts(ctx context.Context, in *productpb.ListProductsRequest, opts ...grpc.CallOption) (*productpb.ListProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(productpb.ListProductsResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) AddProduct(ctx context.Context, in *productpb.AddProductRequest, opts ...grpc.CallOption) (*productpb.AddProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(productpb.AddProductResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) UpdateProduct(ctx context.Context, in *productpb.UpdateProductRequest, opts ...grpc.CallOption) (*productpb.UpdateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(productpb.UpdateProductResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) DeleteProduct(ctx context.Context, in *productpb.DeleteProductRequest, opts ...grpc.CallOption) (*productpb.DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(productpb.DeleteProductResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) AddToCart(ctx context.Context, in *cartpb.AddToCartRequest, opts ...grpc.CallOption) (*cartpb.AddToCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(cartpb.AddToCartResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_AddToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) RemoveFromCart(ctx context.Context, in *cartpb.RemoveFromCartRequest, opts ...grpc.CallOption) (*cartpb.RemoveFromCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(cartpb.RemoveFromCartResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_RemoveFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) GetCart(ctx context.Context, in *cartpb.GetCartRequest, opts ...grpc.CallOption) (*cartpb.GetCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(cartpb.GetCartResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) ClearCart(ctx context.Context, in *cartpb.ClearCartRequest, opts ...grpc.CallOption) (*cartpb.ClearCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(cartpb.ClearCartResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_ClearCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiGatewayServiceClient) Checkout(ctx context.Context, in *cartpb.CheckoutRequest, opts ...grpc.CallOption) (*cartpb.CheckoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(cartpb.CheckoutResponse)
	err := c.cc.Invoke(ctx, ApiGatewayService_Checkout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiGatewayServiceServer is the server API for ApiGatewayService service.
// All implementations must embed UnimplementedApiGatewayServiceServer
// for forward compatibility.
type ApiGatewayServiceServer interface {
	// User Service
	RegisterUser(context.Context, *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error)
	LoginUser(context.Context, *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error)
	GetUser(context.Context, *userpb.GetUserRequest) (*userpb.GetUserResponse, error)
	UpdateUser(context.Context, *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error)
	DeleteUser(context.Context, *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error)
	// Product Service
	GetProduct(context.Context, *productpb.GetProductRequest) (*productpb.GetProductResponse, error)
	ListProducts(context.Context, *productpb.ListProductsRequest) (*productpb.ListProductsResponse, error)
	AddProduct(context.Context, *productpb.AddProductRequest) (*productpb.AddProductResponse, error)
	UpdateProduct(context.Context, *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error)
	DeleteProduct(context.Context, *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error)
	// Cart Service
	AddToCart(context.Context, *cartpb.AddToCartRequest) (*cartpb.AddToCartResponse, error)
	RemoveFromCart(context.Context, *cartpb.RemoveFromCartRequest) (*cartpb.RemoveFromCartResponse, error)
	GetCart(context.Context, *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error)
	ClearCart(context.Context, *cartpb.ClearCartRequest) (*cartpb.ClearCartResponse, error)
	Checkout(context.Context, *cartpb.CheckoutRequest) (*cartpb.CheckoutResponse, error)
	mustEmbedUnimplementedApiGatewayServiceServer()
}

// UnimplementedApiGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiGatewayServiceServer struct{}

func (UnimplementedApiGatewayServiceServer) RegisterUser(context.Context, *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedApiGatewayServiceServer) LoginUser(context.Context, *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetUser(context.Context, *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedApiGatewayServiceServer) UpdateUser(context.Context, *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedApiGatewayServiceServer) DeleteUser(context.Context, *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetProduct(context.Context, *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedApiGatewayServiceServer) ListProducts(context.Context, *productpb.ListProductsRequest) (*productpb.ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedApiGatewayServiceServer) AddProduct(context.Context, *productpb.AddProductRequest) (*productpb.AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedApiGatewayServiceServer) UpdateProduct(context.Context, *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedApiGatewayServiceServer) DeleteProduct(context.Context, *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedApiGatewayServiceServer) AddToCart(context.Context, *cartpb.AddToCartRequest) (*cartpb.AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedApiGatewayServiceServer) RemoveFromCart(context.Context, *cartpb.RemoveFromCartRequest) (*cartpb.RemoveFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedApiGatewayServiceServer) GetCart(context.Context, *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedApiGatewayServiceServer) ClearCart(context.Context, *cartpb.ClearCartRequest) (*cartpb.ClearCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedApiGatewayServiceServer) Checkout(context.Context, *cartpb.CheckoutRequest) (*cartpb.CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedApiGatewayServiceServer) mustEmbedUnimplementedApiGatewayServiceServer() {}
func (UnimplementedApiGatewayServiceServer) testEmbeddedByValue()                           {}

// UnsafeApiGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiGatewayServiceServer will
// result in compilation errors.
type UnsafeApiGatewayServiceServer interface {
	mustEmbedUnimplementedApiGatewayServiceServer()
}

func RegisterApiGatewayServiceServer(s grpc.ServiceRegistrar, srv ApiGatewayServiceServer) {
	// If the following call pancis, it indicates UnimplementedApiGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApiGatewayService_ServiceDesc, srv)
}

func _ApiGatewayService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userpb.RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).RegisterUser(ctx, req.(*userpb.RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userpb.LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).LoginUser(ctx, req.(*userpb.LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userpb.GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetUser(ctx, req.(*userpb.GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userpb.UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).UpdateUser(ctx, req.(*userpb.UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(userpb.DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).DeleteUser(ctx, req.(*userpb.DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productpb.GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetProduct(ctx, req.(*productpb.GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productpb.ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).ListProducts(ctx, req.(*productpb.ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productpb.AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).AddProduct(ctx, req.(*productpb.AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productpb.UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).UpdateProduct(ctx, req.(*productpb.UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(productpb.DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).DeleteProduct(ctx, req.(*productpb.DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cartpb.AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).AddToCart(ctx, req.(*cartpb.AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cartpb.RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_RemoveFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).RemoveFromCart(ctx, req.(*cartpb.RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cartpb.GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).GetCart(ctx, req.(*cartpb.GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cartpb.ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).ClearCart(ctx, req.(*cartpb.ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiGatewayService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cartpb.CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGatewayServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiGatewayService_Checkout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGatewayServiceServer).Checkout(ctx, req.(*cartpb.CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiGatewayService_ServiceDesc is the grpc.ServiceDesc for ApiGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apigatewaypb.ApiGatewayService",
	HandlerType: (*ApiGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _ApiGatewayService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _ApiGatewayService_LoginUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _ApiGatewayService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ApiGatewayService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ApiGatewayService_DeleteUser_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ApiGatewayService_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _ApiGatewayService_ListProducts_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _ApiGatewayService_AddProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ApiGatewayService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ApiGatewayService_DeleteProduct_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _ApiGatewayService_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _ApiGatewayService_RemoveFromCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _ApiGatewayService_GetCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _ApiGatewayService_ClearCart_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _ApiGatewayService_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wazwki/WearStore/api-gateway/api/proto/api-gateway.proto",
}
