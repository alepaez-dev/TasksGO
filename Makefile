.PHONY: build run test lint vet docker-up docker-down migrate-up migrate-down

build:
	go build -o bin/api ./cmd/api

run: build
	./bin/api

test:
	go test ./... -v -race

lint:
	golangci-lint run ./...

vet:
	go vet ./...

docker-up:
	docker compose up --build

docker-down:
	docker compose down

migrate-up:
	@echo "migrations not yet configured"

migrate-down:
	@echo "migrations not yet configured"
