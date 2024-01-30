package storage

import (
	"errors"
	appError "eshop-orders-ms/internal/apperror"
	models "eshop-orders-ms/internal/models/order"
	"gorm.io/gorm"
	"strconv"
)

func (s Storage) CreateOrder(productID string, quantity uint) (uint, error) {
	mockedShippingDetails := models.ShippingDetails{
		FirstName: "John",
		LastName:  "Smith",
		Address:   "NY, 5th Avenue",
		ZipCode:   "12345",
	}
	parsedID, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return 0, err
	}
	order := models.Order{
		ProductID:       uint(parsedID),
		UserID:          "mocked",
		Quantity:        quantity,
		Status:          models.CreatedStatus,
		ShippingDetails: mockedShippingDetails,
	}
	err = s.Postgres.DB.Save(&order).Error

	return order.ID, err
}

func (s Storage) GetOrder(orderID string) (models.Order, error) {
	var order models.Order
	err := s.Postgres.DB.First(&order, orderID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Order{}, appError.OrderNotFound
		}
		return models.Order{}, err
	}
	return order, nil
}

func (s Storage) UpdateOrderStatus(orderID string, newStatus string) error {
	order, err := s.GetOrder(orderID)
	if err != nil {
		return err
	}
	order.Status = newStatus
	return s.Postgres.DB.Save(&order).Error
}
