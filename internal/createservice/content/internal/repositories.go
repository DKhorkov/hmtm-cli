package internal

const RepositoriesContent = `package repositories

import (
	"log/slog"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/tracing"
)

func New<service-name-title>Repository(
	dbConnector db.Connector,
	logger *slog.Logger,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
) *<service-name-title>Repository {
	return &<service-name-title>Repository{
		dbConnector:   dbConnector,
		logger:        logger,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
	}
}

type <service-name-title>Repository struct {
	dbConnector   db.Connector
	logger        *slog.Logger
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
}
`
