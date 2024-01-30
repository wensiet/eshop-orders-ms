package models

import (
	appError "eshop-orders-ms/internal/apperror"
	"gorm.io/gorm"
)

const CreatedStatus = "created"
const ShippingStatus = "shipping"
const DeliveredStatus = "delivered"

func IsValidStatus(status string) error {
	switch status {
	case CreatedStatus, ShippingStatus, DeliveredStatus:
		return nil
	default:
		return appError.InvalidStatus
	}
}

type Order struct {
	gorm.Model
	ProductID         uint            `json:"product_id"`
	Quantity          uint            `json:"quantity"`
	UserID            string          `json:"user_id"`
	Status            string          `json:"status"`
	ShippingDetails   ShippingDetails `json:"shipping_details" gorm:"foreignKey:ShippingDetailsID"`
	ShippingDetailsID uint            `json:"shipping_details_id"`
}

type ShippingDetails struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	ZipCode   string `json:"zip_code"`
}
