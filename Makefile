.DEFAULT_GOAL := run

run:
	gofmt -w .
	goimports -w .
	go run ./cmd/budget-bot/

tests:
	gofmt -w .
	goimports -w .
	go test

build:
	go build .

fmt:
	gofmt -w .

imports:
	goimports -w .

lint:
	golangci-lint run ./



