package logger

import (
	"log/slog"
	"os"
	"strings"
)

// Config controls the logger setup.
type Config struct {
	Level  string // "debug", "info", "warn", "error" (default: "info")
	Format string // "json" or "text" (default: "json")
}

// New returns a structured *slog.Logger.
//
// Usage:
//
//	log := logger.New(logger.Config{Level: "debug"})
//	slog.SetDefault(log)
func New(cfg Config) *slog.Logger {
	level := parseLevel(cfg.Level)
	w := os.Stdout

	var handler slog.Handler
	if strings.ToLower(cfg.Format) == "text" {
		handler = slog.NewTextHandler(w, &slog.HandlerOptions{Level: level})
	} else {
		handler = slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level})
	}

	return slog.New(handler)
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
