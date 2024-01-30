package storage

import (
	"eshop-orders-ms/internal/storage/database/postgres"
)

type Storage struct {
	Postgres *postgres.Postgres
}

func New() *Storage {
	postgresInstance, err := postgres.New()
	if err != nil {
		panic(err)
	}
	return &Storage{
		Postgres: postgresInstance,
	}
}
