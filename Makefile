ROOT := $(shell git rev-parse --show-toplevel)
OS := $(shell uname -s | awk '{print tolower($$0)}')
ARCH := amd64
PROJECT := go-polls
VERSION := $(shell git describe --abbrev=0 --tags)
LD_FLAGS := -X main.version=$(VERSION) -s -w

.PHONY: run
run: clean build
	./polls

clean: ### Clean build files
	@rm -rf ./bin
	@go clean -cache -testcache -modcache


deps: ### Optimize dependencies
	@go mod tidy
	@go mod vendor

.PHONY: build
build: clean ### Build binary
	@go build -v -race -ldflags "${LD_FLAGS}" -o ./bin/polls ./cmd/polls/*.go
	@chmod +x ./bin/*

.PHONY: test
test: ### Run tests
	@go test -v -coverprofile=cover.out -timeout 10s ./...

.PHONY: cover
cover: test ### Run tests and generate coverage
	@go tool cover -html=cover.out -o=cover.html

.PHONY: vendor
vendor: ### Vendor dependencies
	@go mod vendor

.PHONY: lint
lint:
	@staticcheck ./...
	@go vet ./...
	@golangci-lint run --enable=errcheck,gocritic,gofmt,goimports,gosec,gosimple,govet,ineffassign,revive,staticcheck,typecheck,unused,bodyclose ./...

.PHONY: fiximports
fiximports:
	@goimports -local "github.com/bharat-rajani/go-polls" -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
