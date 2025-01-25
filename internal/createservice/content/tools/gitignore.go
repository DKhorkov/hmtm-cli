package tools

const GitignoreContent = `# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
coverage

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env
.env.local
.env.example

# Local files^
*.log
logs/
*.db
bench.txt

# Environments
.venv
venv/
ENV/
env.bak/
venv.bak/

# IDE:
.idea/

# Backups:
postgres_backups
postgres_data
`
