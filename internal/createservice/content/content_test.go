package content

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContent(t *testing.T) {
	t.Run("should create new content", func(t *testing.T) {
		var (
			serviceName = "orders"
			expected    = []string{
				".github/workflows/main.yml",
				fmt.Sprintf("api/protobuf/protofiles/%s/%s.proto", serviceName, serviceName),
				"build/package/local/docker-compose.yml",
				"build/package/prod/docker-compose.yml",
				"build/package/Dockerfile",
				"cmd/client/client.go",
				"cmd/server/server.go",
				"internal/interfaces/clients.go",
				"internal/interfaces/controllers.go",
				"internal/interfaces/repositories.go",
				"internal/interfaces/services.go",
				"internal/interfaces/usecases.go",
				"internal/config/config.go",
				"internal/app/app.go",
				"internal/usecases/usecases.go",
				fmt.Sprintf("internal/services/%s_service.go", serviceName),
				fmt.Sprintf("internal/repositories/%s_repository.go", serviceName),
				"internal/controllers/grpc/controller.go",
				fmt.Sprintf("internal/controllers/grpc/%s/server.go", serviceName),
				"scripts/Taskfile.yml",
				"scripts/postgres/backup",
				"scripts/postgres/restore",
				".gitignore",
				".golangci.yml",
				"go.mod",
				"LICENSE",
				"README.md",
			}
		)

		actual := New(serviceName)
		for _, path := range expected {
			_, ok := actual[path]
			assert.True(t, ok)
		}
	})
}
