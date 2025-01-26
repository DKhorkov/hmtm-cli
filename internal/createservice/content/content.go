package content

import (
	"strings"

	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/api"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/build"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/build/dockercompose"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/cmd"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/docs"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/github"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/internal"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/internal/controllers"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/internal/interfaces"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/scripts"
	"github.com/DKhorkov/hmtm-cli/internal/createservice/content/tools"
)

const (
	serviceNamePlaceholder      = "<service-name>"
	serviceNameUpperPlaceholder = "<service-name-upper>"
	serviceNameTitlePlaceholder = "<service-name-title>"
)

var pathsContent = map[string]string{
	".github/workflows/main.yml":                                  github.GithubContent,
	"api/protobuf/protofiles/<service-name>/<service-name>.proto": api.ProtoContent,
	"build/package/local/docker-compose.yml":                      dockercompose.LocalContent,
	"build/package/prod/docker-compose.yml":                       dockercompose.ProdContent,
	"build/package/Dockerfile":                                    build.DockerfileContent,
	"cmd/client/client.go":                                        cmd.ClientContent,
	"cmd/server/server.go":                                        cmd.ServerContent,
	"internal/interfaces/clients.go":                              interfaces.ClientsContent,
	"internal/interfaces/controllers.go":                          interfaces.ControllersContent,
	"internal/interfaces/repositories.go":                         interfaces.RepositoriesContent,
	"internal/interfaces/services.go":                             interfaces.ServicesContent,
	"internal/interfaces/usecases.go":                             interfaces.UsecasesContent,
	"internal/config/config.go":                                   internal.ConfigContent,
	"internal/app/app.go":                                         internal.AppContent,
	"internal/usecases/usecases.go":                               internal.UsecasesContent,
	"internal/services/<service-name>_service.go":                 internal.ServicesContent,
	"internal/repositories/<service-name>_repository.go":          internal.RepositoriesContent,
	"internal/controllers/grpc/controller.go":                     controllers.ControllerContent,
	"internal/controllers/grpc/<service-name>/server.go":          controllers.ServerContent,
	"scripts/Taskfile.yml":                                        scripts.TaskfileContent,
	"scripts/postgres/backup":                                     scripts.BackupContent,
	"scripts/postgres/restore":                                    scripts.RestoreContent,
	".gitignore":                                                  tools.GitignoreContent,
	".golangci.yml":                                               tools.LintersContent,
	"go.mod":                                                      tools.GoModContent,
	"LICENSE":                                                     docs.LicenseContent,
	"README.md":                                                   docs.ReadmeContent,
}

func New(serviceName string) map[string]string {
	// Coping map for correct paths replacement:
	mapCopy := make(map[string]string)

	// Replace service placeholder on real service name in every path:
	for path, content := range pathsContent {
		mapCopy[strings.Replace(path, serviceNamePlaceholder, serviceName, -1)] = content
	}

	for path, content := range mapCopy {
		mapCopy[path] = strings.Replace(
			content,
			serviceNamePlaceholder,
			serviceName,
			-1,
		)
	}

	// in each iteration for correct work of replacement:
	for path, content := range mapCopy {
		mapCopy[path] = strings.Replace(
			content,
			serviceNameUpperPlaceholder,
			strings.ToUpper(serviceName),
			-1,
		)
	}

	for path, content := range mapCopy {
		mapCopy[path] = strings.Replace(
			content,
			serviceNameTitlePlaceholder,
			strings.ToUpper(string(serviceName[0]))+serviceName[1:], // Only ASCII letters should be in service name
			-1,
		)
	}

	return mapCopy
}
