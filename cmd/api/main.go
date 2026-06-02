package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-fiber-snippets/internal/bootstrap"
	"go-fiber-snippets/internal/cron"
	"go-fiber-snippets/internal/logger"
	"go-fiber-snippets/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// --- Structured logging ---
	logCfg := logger.Config{
		Level:  envOrDefault("LOG_LEVEL", "info"),
		Format: envOrDefault("LOG_FORMAT", "json"),
	}
	slog.SetDefault(logger.New(logCfg))
	slog.Info("starting server", "level", logCfg.Level, "format", logCfg.Format)

	// --- Environment ---
	if err := godotenv.Load(); err != nil {
		slog.Debug(".env not found — using system environment variables")
	}
	port := envOrDefault("PORT", "8000")

	// --- Scheduler ---
	sched, err := cron.RegisterAll()
	if err != nil {
		slog.Error("failed to start scheduler", "error", err)
		os.Exit(1)
	}
	sched.Start()
	defer func() { _ = sched.Shutdown() }()

	// --- Wire ---
	container := bootstrap.NewContainer()

	// --- Fiber app ---
	app := server.New(server.Config{
		Port:           port,
		GraphQLHandler: container.GraphQLHandler,
		BookHandler:    container.BookHandler,
	})

	// --- Start ---
	go func() {
		slog.Info("listening", "port", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()

	// --- Graceful shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down…")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("forced shutdown", "error", err)
	}
	slog.Info("stopped")
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
