# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o dist/arkanum-cli ./cmd/cli/
