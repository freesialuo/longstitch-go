APP_NAME=longstitch

.PHONY: run build test fmt tidy

run:
	go run ./cmd/longstitch

build:
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd/longstitch

test:
	go test ./internal/...

fmt:
	gofmt -w ./cmd ./internal

tidy:
	go mod tidy
