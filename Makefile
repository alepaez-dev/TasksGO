.PHONY: build run test lint vet docker-up docker-down migrate-up migrate-down setup-hooks check fmt verify verify-fmt verify-tidy

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

setup-hooks:
	git config core.hooksPath .githooks
	@echo "Git hooks configured to use .githooks/"

check: vet lint test build

fmt:
	@command -v goimports >/dev/null 2>&1 || { echo "ERROR: goimports not installed."; echo "  go install golang.org/x/tools/cmd/goimports@latest"; exit 1; }
	goimports -w .
	go mod tidy

verify: verify-fmt verify-tidy vet lint test build
	@echo ""
	@echo "All checks passed."

verify-fmt:
	@echo "==> goimports"
	@command -v goimports >/dev/null 2>&1 || { echo "ERROR: goimports not installed (go install golang.org/x/tools/cmd/goimports@latest)"; exit 1; }
	@bad=$$(git ls-files '*.go' | xargs goimports -l); \
	if [ -n "$$bad" ]; then echo "ERROR: files not formatted (run 'make fmt'):"; echo "$$bad" | sed 's/^/  /'; exit 1; fi

verify-tidy:
	@echo "==> go mod tidy"
	@tidy_exit=0; \
	tidy_out=$$(go mod tidy -diff 2>&1) || tidy_exit=$$?; \
	if [ "$$tidy_exit" -eq 0 ]; then \
	  exit 0; \
	fi; \
	if echo "$$tidy_out" | head -1 | grep -qE '^(diff |---|\+\+\+)'; then \
	  echo "ERROR: go.mod / go.sum not tidy (run 'make fmt'):"; \
	  echo ""; \
	  echo "$$tidy_out"; \
	  exit 1; \
	else \
	  echo "ERROR: 'go mod tidy -diff' failed (network or toolchain issue):"; \
	  echo ""; \
	  echo "$$tidy_out"; \
	  exit 1; \
	fi
