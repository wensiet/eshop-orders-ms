package main

import (
	grpcApp "eshop-orders-ms/internal/app/grpc"
	"eshop-orders-ms/internal/config"
	"eshop-orders-ms/internal/migrations"
	ordersService "eshop-orders-ms/internal/services/orders"
	"eshop-orders-ms/internal/storage/storage"
	"github.com/getsentry/sentry-go"
	"log/slog"
)

var logger *slog.Logger
var store *storage.Storage
var orderService ordersService.Orders
var conf config.Config

func init() {
	conf = config.Get()

	logger = config.GetLogger()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://1b3afd7e713f420d4e4899c66bebc4f5@o4506655030378496.ingest.sentry.io/4506661215797248",
		Environment:      conf.Env,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		logger.Error("failed to init sentry %s", err)
	}

	store = storage.New()

	orderService = *ordersService.New(logger, store)

	migrations.MustMigrate(*store)
}

func main() {
	server := grpcApp.New(logger, orderService, conf.GRPC.Port)
	server.MustRun()
}
