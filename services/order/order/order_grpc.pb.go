// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: order.proto

package order

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
	OrderService_CreateOrder_FullMethodName                        = "/order.OrderService/CreateOrder"
	OrderService_CreateOrderRollback_FullMethodName                = "/order.OrderService/CreateOrderRollback"
	OrderService_CancelOrder_FullMethodName                        = "/order.OrderService/CancelOrder"
	OrderService_GetOrder_FullMethodName                           = "/order.OrderService/GetOrder"
	OrderService_ListOrders_FullMethodName                         = "/order.OrderService/ListOrders"
	OrderService_UpdateOrder2PaymentSuccess_FullMethodName         = "/order.OrderService/UpdateOrder2PaymentSuccess"
	OrderService_UpdateOrder2PaymentSuccessRollback_FullMethodName = "/order.OrderService/UpdateOrder2PaymentSuccessRollback"
	OrderService_UpdateOrder2PaymentStatus_FullMethodName          = "/order.OrderService/UpdateOrder2PaymentStatus"
	OrderService_UpdateOrder2PaymentStatusRollback_FullMethodName  = "/order.OrderService/UpdateOrder2PaymentStatusRollback"
	OrderService_GetOrder2Payment_FullMethodName                   = "/order.OrderService/GetOrder2Payment"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// // --------------- 服务接口定义 ---------------
type OrderServiceClient interface {
	// CreateOrder 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error)
	// CreateOrderRollback 补偿操作
	CreateOrderRollback(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	// CancelOrder 取消订单 由用户发起
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	// GetOrder 获取订单详情
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error)
	// ListOrders 分页查询订单列表
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	// --------------- 支付服务内部接口 ---------------
	//
	//	修改订单状态为支付成功
	//	UpdateOrder2PaymentSuccess 支付成功时（进行修改订单状态）
	UpdateOrder2PaymentSuccess(ctx context.Context, in *UpdateOrder2PaymentSuccessRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	// UpdateOrder2PaymentSuccessRollback 支付失败的补充操作
	UpdateOrder2PaymentSuccessRollback(ctx context.Context, in *UpdateOrder2PaymentSuccessRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	// UpdateOrder2Payment 更新订单（支付服务回调使用） 更新为支付中
	UpdateOrder2PaymentStatus(ctx context.Context, in *UpdateOrder2PaymentRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	// UpdateOrder2PaymentStatusRollback 补偿操作 更新订单（支付服务回调使用） 创建状态
	UpdateOrder2PaymentStatusRollback(ctx context.Context, in *UpdateOrder2PaymentRequest, opts ...grpc.CallOption) (*EmptyRes, error)
	GetOrder2Payment(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderDetail2PaymentResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderDetailResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrderRollback(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_CreateOrderRollback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderDetailResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, OrderService_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder2PaymentSuccess(ctx context.Context, in *UpdateOrder2PaymentSuccessRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder2PaymentSuccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder2PaymentSuccessRollback(ctx context.Context, in *UpdateOrder2PaymentSuccessRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder2PaymentSuccessRollback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder2PaymentStatus(ctx context.Context, in *UpdateOrder2PaymentRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder2PaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder2PaymentStatusRollback(ctx context.Context, in *UpdateOrder2PaymentRequest, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder2PaymentStatusRollback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder2Payment(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderDetail2PaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderDetail2PaymentResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrder2Payment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
//
// // --------------- 服务接口定义 ---------------
type OrderServiceServer interface {
	// CreateOrder 创建订单
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderDetailResponse, error)
	// CreateOrderRollback 补偿操作
	CreateOrderRollback(context.Context, *CreateOrderRequest) (*EmptyRes, error)
	// CancelOrder 取消订单 由用户发起
	CancelOrder(context.Context, *CancelOrderRequest) (*EmptyRes, error)
	// GetOrder 获取订单详情
	GetOrder(context.Context, *GetOrderRequest) (*OrderDetailResponse, error)
	// ListOrders 分页查询订单列表
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	// --------------- 支付服务内部接口 ---------------
	//
	//	修改订单状态为支付成功
	//	UpdateOrder2PaymentSuccess 支付成功时（进行修改订单状态）
	UpdateOrder2PaymentSuccess(context.Context, *UpdateOrder2PaymentSuccessRequest) (*EmptyRes, error)
	// UpdateOrder2PaymentSuccessRollback 支付失败的补充操作
	UpdateOrder2PaymentSuccessRollback(context.Context, *UpdateOrder2PaymentSuccessRequest) (*EmptyRes, error)
	// UpdateOrder2Payment 更新订单（支付服务回调使用） 更新为支付中
	UpdateOrder2PaymentStatus(context.Context, *UpdateOrder2PaymentRequest) (*EmptyRes, error)
	// UpdateOrder2PaymentStatusRollback 补偿操作 更新订单（支付服务回调使用） 创建状态
	UpdateOrder2PaymentStatusRollback(context.Context, *UpdateOrder2PaymentRequest) (*EmptyRes, error)
	GetOrder2Payment(context.Context, *GetOrderRequest) (*OrderDetail2PaymentResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*OrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderRollback(context.Context, *CreateOrderRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderRollback not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetOrderRequest) (*OrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder2PaymentSuccess(context.Context, *UpdateOrder2PaymentSuccessRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder2PaymentSuccess not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder2PaymentSuccessRollback(context.Context, *UpdateOrder2PaymentSuccessRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder2PaymentSuccessRollback not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder2PaymentStatus(context.Context, *UpdateOrder2PaymentRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder2PaymentStatus not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder2PaymentStatusRollback(context.Context, *UpdateOrder2PaymentRequest) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder2PaymentStatusRollback not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder2Payment(context.Context, *GetOrderRequest) (*OrderDetail2PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder2Payment not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrderRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrderRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrderRollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrderRollback(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder2PaymentSuccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrder2PaymentSuccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder2PaymentSuccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder2PaymentSuccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder2PaymentSuccess(ctx, req.(*UpdateOrder2PaymentSuccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder2PaymentSuccessRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrder2PaymentSuccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder2PaymentSuccessRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder2PaymentSuccessRollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder2PaymentSuccessRollback(ctx, req.(*UpdateOrder2PaymentSuccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder2PaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrder2PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder2PaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder2PaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder2PaymentStatus(ctx, req.(*UpdateOrder2PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder2PaymentStatusRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrder2PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder2PaymentStatusRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder2PaymentStatusRollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder2PaymentStatusRollback(ctx, req.(*UpdateOrder2PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder2Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder2Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder2Payment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder2Payment(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "CreateOrderRollback",
			Handler:    _OrderService_CreateOrderRollback_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "UpdateOrder2PaymentSuccess",
			Handler:    _OrderService_UpdateOrder2PaymentSuccess_Handler,
		},
		{
			MethodName: "UpdateOrder2PaymentSuccessRollback",
			Handler:    _OrderService_UpdateOrder2PaymentSuccessRollback_Handler,
		},
		{
			MethodName: "UpdateOrder2PaymentStatus",
			Handler:    _OrderService_UpdateOrder2PaymentStatus_Handler,
		},
		{
			MethodName: "UpdateOrder2PaymentStatusRollback",
			Handler:    _OrderService_UpdateOrder2PaymentStatusRollback_Handler,
		},
		{
			MethodName: "GetOrder2Payment",
			Handler:    _OrderService_GetOrder2Payment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
