## Usage

1) Install [Go](https://go.dev/) language.
2) Install hmtm-cli, using next command:
```shell
go install github.com/DKhorkov/hmtm-cli@latest
```
3) Choose command, which matches your purpose.

If CLI will not be found globally, 
check [official Go documentation](https://go.dev/doc/tutorial/compile-install) 
about `go install` command or 
read [next article](https://stackoverflow.com/questions/36083542/error-command-not-found-after-installing-go-eval).

## Linters

To run linters, use next command:
```shell
task linters -v
```

## Tests

To run test, use next command. Coverage info will be
recorded to ```coverage``` folder:
```shell
task tests -v
```

## Benchmarks

To run benchmarks, use next command:
```shell
task bench -v
```
