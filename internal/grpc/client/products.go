package client

import (
	"context"
	productv1 "eshop-orders-ms/gen/go/products"
	"log"
)

func (c *GRPCClient) ProceedOrder(id string, quantity uint) error {
	_, err := c.productsClient.ProceedOrder(context.Background(),
		&productv1.ProceedOrderRequest{
			Id:       id,
			Quantity: int32(quantity),
		},
	)
	return err
}

func (c *GRPCClient) Test() {
	product, err := c.productsClient.GetProduct(context.Background(), &productv1.GetProductRequest{Id: "1"})
	if err != nil {
		panic(err)
	}
	log.Fatal(product)
}
