package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting sso", 
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)
	log.Debug("debug messages are enabled")
	log.Error("error message")
	log.Warn("warn message")

	// init logger
	// init app
	// init storage
	//
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}