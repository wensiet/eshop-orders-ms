package ordersService

import (
	"context"
	appError "eshop-orders-ms/internal/apperror"
	models "eshop-orders-ms/internal/models/order"
	"fmt"
	"github.com/getsentry/sentry-go"
)

type OrdersStorage interface {
	CreateOrder(productID string, quantity uint) (uint, error)
	GetOrder(orderID string) (models.Order, error)
	UpdateOrderStatus(orderID string, newStatus string) error
}

func (o Orders) CreateNewOrder(productID string, quantity uint) (uint, error) {
	const op = "ordersService.Orders.CreateOrder"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := o.log.With(
		"op", op,
	)

	log.Info("creating new order")

	trID, err := o.gRPCClient.BeginTransaction(productID, quantity)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	orderID, err := o.ordersStorage.CreateOrder(productID, quantity)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = o.gRPCClient.EndTransaction(trID, true)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return orderID, nil
}

func (o Orders) GetOrder(orderID string) (models.Order, error) {
	const op = "ordersService.Orders.GetOrder"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := o.log.With(
		"op", op,
	)

	log.Info("getting order")

	order, err := o.ordersStorage.GetOrder(orderID)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (o Orders) UpdateOrderStatus(orderID string, newStatus string) error {
	const op = "ordersService.Orders.UpdateOrderStatus"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := o.log.With(
		"op", op,
	)

	log.Info("updating order status")

	err := models.IsValidStatus(newStatus)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return fmt.Errorf("%s: %w", op, err)
	}

	err = o.ordersStorage.UpdateOrderStatus(orderID, newStatus)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
