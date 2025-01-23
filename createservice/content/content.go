package content

import (
	"strings"

	"hmtm-cli/createservice/content/api"
	"hmtm-cli/createservice/content/build"
	"hmtm-cli/createservice/content/cmd"
	"hmtm-cli/createservice/content/github"
	"hmtm-cli/createservice/content/internal"
	"hmtm-cli/createservice/content/internal/controllers"
	"hmtm-cli/createservice/content/scripts"
)

const (
	serviceNamePlaceholder      = "<service-name>"
	serviceNameUpperPlaceholder = "<service-name-upper>"
)

var pathsContent = map[string]string{
	".github/workflows/main.yml":                                  github.GithubContent,
	"api/protobuf/protofiles/<service-name>/<service-name>.proto": api.ProtoContent,
	"build/package/local/docker-compose.yml":                      build.LocalBuildContent,
	"build/package/prod/docker-compose.yml":                       build.ProdBuildContent,
	"build/package/Dockerfile":                                    build.DockerfileContent,
	"cmd/client/client.go":                                        cmd.ClientContent,
	"cmd/server/server.go":                                        cmd.ServerContent,
	"internal/app/app.go":                                         internal.AppContent,
	//"internal/config/config.go",
	//"internal/interfaces/controllers.go",
	//"internal/interfaces/repositories.go",
	//"internal/interfaces/services.go",
	//"internal/interfaces/usecases.go",
	//"internal/usecases/usecases.go",

	"internal/controllers/grpc/controller.go":            controllers.ControllerContent,
	"internal/controllers/grpc/<service-name>/server.go": controllers.ServerContent,
	"scripts/Taskfile.yml":                               scripts.TaskfileContent,
	"scripts/postgres/backup":                            scripts.BackupContent,
	"scripts/postgres/restore":                           scripts.RestoreContent,

	//".gitignore":
	//".golangci.yml":
	//"go.mod":
	//"LICENSE":
	//"README.md":
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

	return mapCopy
}
