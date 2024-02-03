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

// New Creates new instance of Orders service layer
func New(log *slog.Logger, ordersStorage OrdersStorage) *Orders {
	gRPCClient, err := client.New()
	if err != nil {
		for {
			log.With("error", err.Error()).Warn("failed to init gRPC connection with eshop-products-ms. retrying...")
			time.Sleep(5 * time.Second)
			gRPCClient, err = client.New()
			if err == nil {
				break
			}
		}
	}
	return &Orders{log: log, ordersStorage: ordersStorage, gRPCClient: *gRPCClient}
}
