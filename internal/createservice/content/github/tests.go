package github

const GithubTestsContent = `name: tests-workflow
run-name: ${{ github.actor }} is testing changes
on: [push, pull_request]
jobs:
  tests:
    runs-on: ubuntu-latest
    name: Tests
    steps:
      - name: Check out source repository
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: stable
      - name: Install Taskfile
        run: go install github.com/go-task/task/v3/cmd/task@latest
      - name: Install dependencies
        run: go get ./...
      - name: Run tests
        run: task -d scripts tests -v
`
