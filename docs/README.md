# Pattern Reference

Self-contained code notes for common Go Fiber v3 patterns.

## Patterns

| Doc | What | Key File |
|---|---|---|
| [Validation](validation.md) | Struct validation + error formatting | `pkg/validator/` |
| [Pagination](pagination.md) | Query params → offset + meta envelope | `pkg/pagination/` |
| [Structured Logging](logging.md) | `log/slog` JSON logging | `internal/logger/` |

## Core Patterns (in code)

| Pattern | Key File |
|---|---|
| App bootstrap + graceful shutdown | `cmd/api/main.go` |
| Fiber app factory + middleware | `internal/server/server.go` |
| Route grouping (v1, redirects, webhooks) | `internal/server/server.go` |
| 301 redirects | `internal/handler/redirect.go` |
| Webhook receiver (plain REST endpoint) | `internal/handler/webhook.go` |
| WebSocket echo with host guard | `internal/handler/websocket.go` + `internal/middleware/ws_guard.go` |
| GraphQL (schema, queries, types, HTTP handler) | `internal/graphql/` |
| Domain → Repository → Service → Handler (clean layers) | `internal/domain/` → `repository/` → `service/` → `handler/book.go` |
| Cron jobs (gocron v2) | `internal/cron/jobs.go` |
| Dependency injection container | `internal/bootstrap/wire.go` |
| Standardized JSON response envelope | `pkg/common/response.go` |
| Unit tests with testify mocks | `internal/service/book_service_test.go` |
| `.env` loading + envOrDefault | `cmd/api/main.go` |
| Health checks (livez, readyz, startupz) | `internal/server/server.go` |
| Request logging middleware | `internal/middleware/logger.go` |
