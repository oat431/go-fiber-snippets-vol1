# Go Fiber v3 — Production Reference Project

A personal, copy-paste–ready collection of Go Fiber v3 patterns I use across projects.
Every file is self-contained and demonstrates exactly one concern.
No database or external services required — everything runs with `go run cmd/api/main.go`.

## Features

| Feature | File |
|---|---|
| App bootstrap + graceful shutdown | `cmd/api/main.go`, `internal/server/server.go` |
| Structured logging (`log/slog` JSON) | `internal/logger/log.go` |
| Request logging middleware | `internal/middleware/logger.go` |
| Panic recovery | Built into `internal/server/server.go` |
| Health checks (livez, readyz, startupz) | Built into `internal/server/server.go` |
| 301 redirects | `internal/handler/redirect.go` |
| Webhook endpoint | `internal/handler/webhook.go` |
| WebSocket echo (with host guard) | `internal/handler/websocket.go`, `internal/middleware/ws_guard.go` |
| GraphQL (queries, schema, types) | `internal/graphql/` |
| REST endpoints (list + create books) | `internal/handler/book.go` |
| Request validation (struct tags) | `pkg/validator/` |
| Pagination (query params + meta envelope) | `pkg/pagination/` |
| Domain / repository / service layers | `internal/domain/`, `internal/repository/`, `internal/service/` |
| Background cron jobs (gocron v2) | `internal/cron/jobs.go` |
| Standardized JSON response envelope | `pkg/common/response.go` |
| Unit tests with testify mocks | `internal/service/book_service_test.go` |
| `.env` loading | `cmd/api/main.go` |

## Quick Start

```bash
# Copy .env
cp .env.example .env

# Run
go run cmd/api/main.go

# Test
go test ./...

# Build
make build
```

## Endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/api/v1/hello` | Smoke test |
| GET | `/api/v1/books?page=1&limit=20` | Paginated book list |
| POST | `/api/v1/books` | Create a book (with validation) |
| GET | `/api/v1/redirect/linkedin` | 301 → LinkedIn |
| GET | `/api/v1/redirect/github` | 301 → GitHub |
| GET | `/api/v1/redirect/facebook` | 301 → Facebook |
| POST | `/api/v1/hook` | Inbound webhook |
| GET | `/chat` | WebSocket echo |
| POST | `/graphql` | GraphQL endpoint |
| GET | `/livez` | Liveness probe |
| GET | `/readyz` | Readiness probe |
| GET | `/startupz` | Startup probe |

### GraphQL Example

```bash
curl -X POST http://localhost:8000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ books { id title } }"}'
```

### Paginated Books

```bash
curl "http://localhost:8000/api/v1/books?page=1&limit=2"
# → { "data": { "data": […], "meta": { "page":1,"limit":2,"total":3,"totalPages":2 } }, "status":"SUCCESS" }
```

### Validation Example

```bash
curl -X POST http://localhost:8000/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":""}'
# → 422: { "status":"FAIL", "error": {"message":"validation failed", "details":[{"field":"title","message":"this field is required"}]} }
```

## Project Layout

```
cmd/api/main.go           # Entry point — env, logging, wiring, graceful shutdown
internal/
  server/server.go        # Fiber app factory + middleware + routes
  handler/                 # HTTP + WebSocket handlers (hello, redirect, webhook, ws, book)
  middleware/              # Request logger, WS host guard
  graphql/                 # Schema, queries, types, HTTP handler
  domain/                  # Core domain models
  repository/              # Data access (mock implementation)
  service/                 # Business logic (+ tests)
  cron/                    # Background jobs (gocron v2)
  bootstrap/               # Dependency injection container
  logger/                  # Structured logging (slog JSON)
docs/                      # Pattern reference docs
pkg/
  common/response.go       # OK() / Fail() / Err() response constructors
  validator/               # Struct validation + error formatting
  pagination/              # Query params → offset + paginated response
```

## Docs

See [docs/](docs/) for per-pattern reference notes.

## Why This Exists

I copy patterns from this project into new services instead of regenerating boilerplate.
It saves tokens when working with coding assistants — paste a known-good file, tweak it, ship it.
If you find it useful, steal it. 😄
