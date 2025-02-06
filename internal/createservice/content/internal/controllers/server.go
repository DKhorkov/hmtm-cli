package controllers

const ServerContent = `package <service-name>

import (
	"log/slog"

	"google.golang.org/grpc"

	"github.com/DKhorkov/hmtm-<service-name>/api/protobuf/generated/go/<service-name>"
	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
)

// RegisterServer handler (serverAPI) connects <service-name-title>Server to gRPC server:
func RegisterServer(gRPCServer *grpc.Server, useCases interfaces.UseCases, logger *slog.Logger) {
	<service-name>.Register<service-name-title>ServiceServer(gRPCServer, &ServerAPI{useCases: useCases, logger: logger})
}

type ServerAPI struct {
	// Helps to test single endpoints, if others is not implemented yet
	<service-name>.Unimplemented<service-name-title>ServiceServer
	useCases interfaces.UseCases
	logger   *slog.Logger
}
`
