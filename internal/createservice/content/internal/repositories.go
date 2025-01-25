package internal

const RepositoriesContent = `package repositories

import (
	"log/slog"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/tracing"
)

func NewCommon<service-name-title>Repository(
	dbConnector db.Connector,
	logger *slog.Logger,
	traceProvider tracing.TraceProvider,
	spanConfig tracing.SpanConfig,
) *Common<service-name-title>Repository {
	return &Common<service-name-title>Repository{
		dbConnector:   dbConnector,
		logger:        logger,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
	}
}

type Common<service-name-title>Repository struct {
	dbConnector   db.Connector
	logger        *slog.Logger
	traceProvider tracing.TraceProvider
	spanConfig    tracing.SpanConfig
}
`
