package client

import (
	productv1 "eshop-orders-ms/gen/go/products"
	"eshop-orders-ms/internal/config"
	"fmt"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	productsClient productv1.ProductServClient
}

func New() (*GRPCClient, error) {
	conf := config.Get()
	dsn := fmt.Sprintf("%s:%s", conf.Services.Products.Host, conf.Services.Products.Port)
	conn, err := grpc.Dial(dsn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := productv1.NewProductServClient(conn)
	return &GRPCClient{productsClient: client}, nil
}
