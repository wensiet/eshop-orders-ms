package migrations

import (
	models "eshop-orders-ms/internal/models/order"
	"eshop-orders-ms/internal/storage/storage"
)

func MustMigrate(storage storage.Storage) {
	err := storage.Postgres.DB.AutoMigrate(&models.Order{})
	if err != nil {
		panic(err)
	}
	err = storage.Postgres.DB.AutoMigrate(&models.ShippingDetails{})
	if err != nil {
		panic(err)
	}
}
