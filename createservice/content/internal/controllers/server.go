package controllers

const ServerContent = `package <service-name>

import (
	"log/slog"

	"github.com/DKhorkov/hmtm-<service-name>/api/protobuf/generated/go/<service-name>"
	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
	"google.golang.org/grpc"
)

// RegisterServer handler (serverAPI) for <TODO add gRPC server name here> to gRPC server:
// TODO add <service-name>.RegisterServiceServer with ServerAPI inside
func RegisterServer(gRPCServer *grpc.Server, useCases interfaces.UseCases, logger *slog.Logger) {
}

type ServerAPI struct {
	// TODO add <service-name>.UnimplementedServer here
	useCases interfaces.UseCases
	logger   *slog.Logger
}
`
