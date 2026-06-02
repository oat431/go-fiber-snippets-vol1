# Structured Logging

Uses Go 1.21+ `log/slog` with JSON output by default.

## Setup

In `main.go`:

```go
import (
    "log/slog"
    "go-fiber-snippets/internal/logger"
)

func main() {
    slog.SetDefault(logger.New(logger.Config{
        Level:  os.Getenv("LOG_LEVEL"),  // "debug" | "info" | "warn" | "error"
        Format: os.Getenv("LOG_FORMAT"), // "json" | "text"
    }))
}
```

## In handlers & services

```go
import "log/slog"

slog.Info("book created", "id", book.ID, "title", book.Title)
slog.Debug("cache miss", "key", cacheKey)
slog.Warn("slow query", "duration_ms", elapsed, "query", sql)
slog.Error("db connection failed", "error", err, "retry", 3)
```

### When to use each level

| Level | Use for |
|---|---|
| `Debug` | Development-only detail (SQL, cache, request bodies) |
| `Info` | Key events (user login, payment, resource created) |
| `Warn` | Recoverable anomalies (retry, fallback, deprecation) |
| `Error` | Things that need attention (panic recovery, DB down) |

## Integrate with Fiber

Fiber v3 uses its own `log` package. To bridge, set the env vars above — structured logs go to stdout, Fiber internal logs stay separate on stderr. For full unification, configure Fiber's logger with a custom writer that feeds into `slog`.
