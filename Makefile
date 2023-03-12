export

run: build
	./polls

build: fmt lint sec
	go build cmd/polls.go

fmt:
	gofmt -s -w .

lint:
	staticcheck ./...
	go vet ./...
	golangci-lint run ./...

sec:
	gosec ./...

govendor:
	go mod vendor && go mod tidy