package github

const GithubLintersContent = `name: linters-workflow
run-name: ${{ github.actor }} is testing changes
on: [push, pull_request]
jobs:
  linters:
    name: golangci-lint
    runs-on: ubuntu-latest
    steps:
      - name: Check out source repository
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: stable
      - name: Install dependencies
        run: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
      - name: Run Linters
        run: golangci-lint run -v --fix
`
