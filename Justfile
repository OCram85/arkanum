# Variablen
GO := "go"
GOFLAGS := "-ldflags='-s -w'"

default: build-cli

# Default Build (Linux AMD64)
build-cli:
    just build-with linux amd64

# build with parameters
build-with os arch:
    CGO_ENABLED=0 GOOS={{os}} GOARCH={{arch}} {{GO}} build {{GOFLAGS}} -o dist/arkanum ./cmd/cli/
