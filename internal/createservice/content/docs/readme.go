package docs

const ReadmeContent = "## Usage\n\n" +
	"Before usage need to create network for correct dependencies work:\n" +
	"```shell\ntask -d scripts network -v\n```\n\n" +
	"### Run via docker:\n\n" +
	"To run app and it's dependencies in docker, use next command:\n" +
	"```shell\ntask -d scripts prod -v\n```\n\n" +
	"### Run via source files:\n\n" +
	"To run application via source files, use next commands:\n" +
	"1) Run all application dependencies:\n" +
	"```shell\ntask -d scripts local -v\n```\n" +
	"2) Run application:\n" +
	"```shell\ngo run ./cmd/server/server.go\n```\n\n" +
	"## gRPC:\n\n" +
	"To setup protobuf, use next command:\n" +
	"```shell\ntask -d scripts setup_proto -v\n```\n\n" +
	"To generate files from.proto, use next command:\n" +
	"```shell\ntask -d scripts grpc_generate -v\n```\n\n" +
	"## Linters\n\n" +
	"To run linters, use next command:\n" +
	"```shell\ntask -d scripts linters -v\n```\n\n" +
	"## Tests\n\n" +
	"To run test, use next commands.Coverage docs will be\n" +
	"recorded to ```coverage``` folder:\n" +
	"```shell\ntask -d scripts tests -v\n```\n\n" +
	"## Benchmarks\n\n" +
	"To run benchmarks, use next command:\n" +
	"```shell\ntask -d scripts bench -v\n```\n\n" +
	"## Migrations\n\n" +
	"To create migration file, use next command:\n" +
	"```shell\ntask -d scripts makemigrations NAME={{migration name}}\n```\n\n" +
	"To apply all available migrations, use next command:\n" +
	"```shell\ntask -d scripts migrate\n```\n\n" +
	"To migrate up to a specific version, use next command:\n" +
	"```shell\ntask -d scripts migrate_to VERSION={{migration version}}\n```\n\n" +
	"To rollback migrations to a specific version, use next command:\n" +
	"```shell\ntask -d scripts downgrade_to VERSION={{migration version}}\n```\n\n" +
	"To rollback all migrations (careful!), use next command:\n" +
	"```shell\ntask -d scripts downgrade_to_base\n```\n\n" +
	"To print status of all migrations, use next command:\n" +
	"```shell\ntask -d scripts migrations_status\n```\n\n" +
	"## Tracing\n\n" +
	"To see tracing open\nnext [link](http://localhost:16686) in browser.\n"
