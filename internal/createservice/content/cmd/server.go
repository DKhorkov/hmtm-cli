package cmd

const ServerContent = `package main

import (
	"context"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"

	"github.com/DKhorkov/hmtm-<service-name>/internal/app"
	"github.com/DKhorkov/hmtm-<service-name>/internal/config"
	grpccontroller "github.com/DKhorkov/hmtm-<service-name>/internal/controllers/grpc"
	"github.com/DKhorkov/hmtm-<service-name>/internal/repositories"
	"github.com/DKhorkov/hmtm-<service-name>/internal/services"
	"github.com/DKhorkov/hmtm-<service-name>/internal/usecases"
)

func main() {
	settings := config.New()
	logger := logging.GetInstance(
		settings.Logging.Level,
		settings.Logging.LogFilePath,
	)

	dbConnector, err := db.New(
		db.BuildDsn(settings.Database),
		settings.Database.Driver,
		logger,
		db.WithMaxOpenConnections(settings.Database.Pool.MaxOpenConnections),
		db.WithMaxIdleConnections(settings.Database.Pool.MaxIdleConnections),
		db.WithMaxConnectionLifetime(settings.Database.Pool.MaxConnectionLifetime),
		db.WithMaxConnectionIdleTime(settings.Database.Pool.MaxConnectionIdleTime),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = dbConnector.Close(); err != nil {
			logging.LogError(logger, "Failed to close db connections pool", err)
		}
	}()

	traceProvider, err := tracing.New(settings.Tracing.Server)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = traceProvider.Shutdown(context.Background()); err != nil {
			logging.LogError(logger, "Error shutting down tracer", err)
		}
	}()

	<service-name>Repository := repositories.NewCommon<service-name-title>Repository(
		dbConnector,
		logger,
		traceProvider,
		settings.Tracing.Spans.Repositories.<service-name-title>,
	)

	<service-name>Service := services.NewCommon<service-name-title>Service(
		<service-name>Repository,
		logger,
	)

	useCases := usecases.NewCommonUseCases(<service-name>Service)
	controller := grpccontroller.New(
		settings.HTTP.Host,
		settings.HTTP.Port,
		useCases,
		logger,
		traceProvider,
		settings.Tracing.Spans.Root,
	)

	application := app.New(controller)
	application.Run()
}
`
