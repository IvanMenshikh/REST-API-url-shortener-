package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/storage/postgres"
	//"url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	//fmt.Printf("%+v\n", cfg)

	// TODO: init Logger: slog
	log := setupLogger(cfg.Env)
	//log.Info("Starting URL Shortener service", slog.String("env", cfg.Env))
	//log.Debug("Debug logging enabled")

	// TODO: init storage: sqlite
		storage, err := postgres.NewPostgresStorage(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)
		if err != nil {
			log.Error("failed to init storage", slog.Any("error", err))
			os.Exit(1)
		}
		log.Info("Successfully connected to database")
		_ = storage
	// TODO: init router: chi, "chi render"

	// TODO: init server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
