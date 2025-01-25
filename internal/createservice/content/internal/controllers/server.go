package controllers

const ServerContent = `package <service-name>

import (
	"log/slog"

	"github.com/DKhorkov/hmtm-<service-name>/api/protobuf/generated/go/<service-name>"
	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
	"google.golang.org/grpc"
)

// RegisterServer handler (serverAPI) connects <service-name-title>Server to gRPC server:
func RegisterServer(gRPCServer *grpc.Server, useCases interfaces.UseCases, logger *slog.Logger) {
	<service-name>.Register<service-name-title>ServiceServer(gRPCServer, &ServerAPI{useCases: useCases, logger: logger})
}

type ServerAPI struct {
	// TODO add <service-name>.UnimplementedServer here
	useCases interfaces.UseCases
	logger   *slog.Logger
}
`
