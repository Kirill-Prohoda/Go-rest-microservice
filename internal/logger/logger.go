package logger

import (
	"go-rest-microservice/internal/config"
	slogpretty "go-rest-microservice/internal/lib/logger/handlers"
	"log/slog"
	"os"
)

func InitLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.ENV_LOCAL:
		log = setupPrettySlog()
		// log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.ENV_DEV:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.ENV_PROD:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
