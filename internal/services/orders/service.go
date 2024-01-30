package ordersService

import (
	"eshop-orders-ms/internal/grpc/client"
	"log/slog"
	"time"
)

type Orders struct {
	log           *slog.Logger
	gRPCClient    client.GRPCClient
	ordersStorage OrdersStorage
}

func New(log *slog.Logger, ordersStorage OrdersStorage) *Orders {
	gRPCClient, err := client.New()
	if err != nil {
		for {
			log.Warn("failed to init gRPC connection with eshop-products-ms. retrying...")
			log.Error("%s", err)
			time.Sleep(5 * time.Second)
			gRPCClient, err = client.New()
			if err == nil {
				break
			}
		}
	}
	return &Orders{log: log, ordersStorage: ordersStorage, gRPCClient: *gRPCClient}
}
