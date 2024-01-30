package ordersAPI

import (
	ordersv1 "eshop-orders-ms/gen/go/orders"
	ordersService "eshop-orders-ms/internal/services/orders"
	"google.golang.org/grpc"
)

type ordersAPI struct {
	ordersv1.UnimplementedOrdersServer
	orderService ordersService.Orders
}

func Register(serv *grpc.Server, orders ordersService.Orders) {
	ordersAPI := &ordersAPI{
		orderService: orders,
	}

	ordersv1.RegisterOrdersServer(serv, ordersAPI)
}
