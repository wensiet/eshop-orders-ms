package ordersAPI

import (
	"context"
	"errors"
	ordersv1 "eshop-orders-ms/gen/go/orders"
	appError "eshop-orders-ms/internal/apperror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (o ordersAPI) CreateOrder(ctx context.Context, in *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	order, err := o.orderService.CreateNewOrder(in.ProductId, 1)
	if err != nil {
		if errors.Is(err, appError.UnableToProceedOrder) {
			return nil, status.Error(codes.Unknown, "failed to proceed order")
		}
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return &ordersv1.CreateOrderResponse{
		Id: strconv.Itoa(int(order)),
	}, nil
}

func (o ordersAPI) GetOrder(ctx context.Context, in *ordersv1.GetOrderRequest) (*ordersv1.Order, error) {
	order, err := o.orderService.GetOrder(in.Id)
	if err != nil {
		if errors.Is(err, appError.OrderNotFound) {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return &ordersv1.Order{
		Id:        strconv.Itoa(int(order.ID)),
		ProductId: strconv.Itoa(int(order.ProductID)),
		Status:    order.Status,
	}, err
}

func (o ordersAPI) ListOrders(ctx context.Context, in *ordersv1.ListOrdersRequest) (*ordersv1.ListOrdersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o ordersAPI) UpdateOrder(ctx context.Context, in *ordersv1.UpdateOrderStatusRequest) (*ordersv1.Empty, error) {
	err := o.orderService.UpdateOrderStatus(in.Id, in.Status)
	if err != nil {
		if errors.Is(err, appError.OrderNotFound) {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Error(codes.Internal, "internal error occurred")
	}
	return &ordersv1.Empty{}, nil
}
