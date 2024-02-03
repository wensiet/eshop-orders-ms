package client

import (
	"context"
	productv1 "eshop-orders-ms/gen/go/products"
	appError "eshop-orders-ms/internal/apperror"
)

func (c *GRPCClient) BeginTransaction(id string, quantity uint) (string, error) {
	transaction, err := c.productsClient.BeginOrder(context.Background(),
		&productv1.BeginOrderRequest{
			ProductId: id,
			Quantity:  int32(quantity),
		},
	)
	if err != nil {
		return "", err
	}
	return transaction.TransactionId, nil
}

func (c *GRPCClient) EndTransaction(transactionID string, status bool) error {
	response, err := c.productsClient.ApplyOrder(context.Background(),
		&productv1.ApplyOrderRequest{
			TransactionId: transactionID,
			Success:       status,
		})
	if err != nil {
		return err
	}
	if !response.Success {
		return appError.UnableToProceedOrder
	} else {
		return nil
	}
}
