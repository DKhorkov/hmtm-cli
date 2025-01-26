package structure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DKhorkov/hmtm-cli/internal/createservice/structure"
)

func TestNewStructure(t *testing.T) {
	t.Run("should return a new structure", func(t *testing.T) {
		var (
			serviceName = "orders"
			expected    = []string{
				".github/workflows",
				"api/protobuf/generated/go/" + serviceName,
				"api/protobuf/protofiles/" + serviceName,
				"build/package/local",
				"build/package/prod",
				"cmd/client",
				"cmd/server",
				"internal/app",
				"internal/config",
				"internal/interfaces",
				"internal/clients",
				"internal/controllers/grpc/" + serviceName,
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
		)

		actual := structure.New(serviceName)
		assert.Equal(t, expected, actual)
	})
}
