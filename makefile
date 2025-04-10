GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
GO_FLAGS=-ldflags="-X main.Version=$(VERSION)"
GO=$(GO_ENV) "$(shell which go)"

build:
	mkdir -p bin/ && $(GO) build $(GO_FLAGS) -o ./bin/ ./...