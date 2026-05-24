# TasksGO

A modern task and ticket management system, designed to explore clean architecture and backend systems in Go.

## High-Level Concept

TasksGO will support:
- **Projects**, which group work
- **Tasks**, small actionable items, unrelated or related to tickets
- **Tickets**, issues, features or requests
- **Notes**, attached to tickets (role-dependent views)
- **Roles**: PM, DEV, QA

See [docs/tickets.md](docs/tickets.md) for detailed ticket design specs.

## Tech Stack (Initial)
- Backend: Go
- Frontend: React or Next.js (TBD)
- AI: Separate service (design TBD), python
- Database: PostgreSQL

## Getting Started

### Prerequisites
- [Docker](https://docs.docker.com/get-docker/) and Docker Compose
- [Go 1.24+](https://go.dev/dl/) (for local development without Docker)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) and [golangci-lint](https://golangci-lint.run/), used by the pre-commit hook:
  ```sh
  go install golang.org/x/tools/cmd/goimports@latest
  brew install golangci-lint   # or see https://golangci-lint.run/welcome/install/
  ```

### Setup

1. Clone the repository and copy the environment file:
   ```sh
   cp .env.example .env
   ```

2. Set up git hooks (required):
   ```sh
   make setup-hooks
   ```
   This registers `.githooks/` so every commit is checked for:
   - Conventional Commits format (`commit-msg`)
   - Secret/credential files (`.env`, `*.pem`, SSH keys), `goimports` formatting, `go mod tidy` freshness, `go vet`, and `golangci-lint` (`pre-commit`)
   - Branch name convention `<type>/<description>` (`pre-push`)

   The same checks run in CI, so anything skipped locally will fail the PR.

3. Start the API and database:
   ```sh
   make docker-up
   ```

4. Verify the server is running:
   ```sh
   curl localhost:8080/healthz
   ```
   You should see `{"status":"ok"}`.

5. To stop everything:
   ```sh
   make docker-down
   ```

### Available Make Commands

| Command | Description |
|---|---|
| `make docker-up` | Start API and PostgreSQL via Docker Compose |
| `make docker-down` | Stop all containers |
| `make build` | Compile the API binary to `bin/api` |
| `make run` | Build and run the API locally (requires env vars) |
| `make test` | Run all tests with race detection |
| `make lint` | Run golangci-lint |
| `make vet` | Run go vet |
| `make setup-hooks` | Configure git to use project hooks |
| `make check` | Run vet, lint, test, and build in sequence |
| `make fmt` | Auto-apply `goimports -w` and `go mod tidy` |
| `make verify` | Run all CI-equivalent gates locally (goimports, tidy, vet, lint, test, build). Read-only; run `make fmt` to fix format/tidy failures. |

### Hot Reload

When running wth `make docker-up`, [Air](https://github.com/air-verse/air) watches for `.go` file changes and automatically rebuilds the API inside the container

## License

Copyright (c) 2026 Alejandra Hernandez Paez

This project is licensed under the [PolyForm Noncommercial License 1.0.0](https://polyformproject.org/licenses/noncommercial/1.0.0/).

You may use this software for noncommercial purposes only.
Commercial use requires a separate license. Contact the author for inquiries.
