package main

import (
	"analytics-service/internal/app"
	"analytics-service/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	cfg := config.MustLoad()

	application, err := app.New(logger, cfg)
	if err != nil {
		slog.Error("Failed to create application", slog.Any("error", err))
		os.Exit(1)
	}

	go func() {
		application.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	application.Stop()
}
