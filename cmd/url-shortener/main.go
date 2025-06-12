package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/handlers/slogpretty"

	//"url-shortener/internal/storage/postgres"
	"url-shortener/internal/storage/sqlite"

	"url-shortener/internal/http-server/middleware/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// log.Info("Starting URL Shortener service", slog.String("env", cfg.Env), slog.String("server", cfg.HTTPServer.Address))
	// log.Debug("Debug logging enabled")
	// log.Warn("This is a warning message")
	// log.Error("This is an error message")

	// TODO: init storage: sqlite
	storage, err := sqlite.NewSQLiteStorage(cfg.Sqlite.Path)
	if err != nil {
		log.Error("failed to init storage", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("Successfully connected to database")
	_ = storage

	// TODO: init router: chi, "chi render"
	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Генерирует уникальный ID для каждого запроса (будем использовать в graphql или kabana)
	router.Use(middleware.RealIP)    // Получаем реальный IP клиента
	//router.Use(middleware.Logger) на выбор, но не нужен, т.к. есть свой логгер
	router.Use(logger.New(log))      // Логгируем запросы
	router.Use(middleware.Recoverer) // Обрабатываем паники и ошибки
	router.Use(middleware.URLFormat) // Форматируем URL

	// TODO: init server

}

// Инициализация логгера в зависимости от окружения
func setupLogger(env string) *slog.Logger {
	var level slog.Level
	switch env {
	case envLocal:
		level = slog.LevelDebug
	case envDev:
		level = slog.LevelDebug
	case envProd:
		level = slog.LevelInfo
	}

	// Выбор типа логгера (pretty для локальной разработки, JSON для продакшена и девелопмента)
	if env == envLocal {
		return setupPrettySlog(level)
	}
	return newJSONLogger(level)
}

// Инициализация логгера с красивым выводом (с выбором уровня логирования)
func setupPrettySlog(level slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}

// Инициализация JSON логгера (для продакшена и девелопмента)
func newJSONLogger(level slog.Level) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(handler)
}
