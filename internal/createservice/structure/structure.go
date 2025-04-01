package structure

import "strings"

const serviceNamePlaceholder = "<service-name>"

var structure = []string{
	".github/workflows",
	"api/protobuf/generated/go/<service-name>",
	"api/protobuf/protofiles/<service-name>",
	"build/package/local",
	"build/package/prod",
	"cmd/client",
	"cmd/server",
	"internal/app",
	"internal/config",
	"internal/interfaces",
	"internal/clients",
	"internal/controllers/grpc/<service-name>",
	"internal/usecases",
	"internal/services",
	"internal/repositories",
	"internal/entities",
	"internal/errors",
	"migrations",
	"mocks",
	"scripts",
	"scripts/postgres",
}

func New(serviceName string) []string {
	for i := range structure {
		structure[i] = strings.ReplaceAll(structure[i], serviceNamePlaceholder, serviceName)
	}

	return structure
}
