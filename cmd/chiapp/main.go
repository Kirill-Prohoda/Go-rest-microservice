package main

import (
	"go-rest-microservice/internal/app/chiapp"
	"go-rest-microservice/internal/config"
	"go-rest-microservice/internal/lib/logger/slogAdapter"
	"go-rest-microservice/internal/logger"
	"go-rest-microservice/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	// chiapp.ChiApp()
	cfg := config.InitConfig()
	var log = logger.InitLogger(cfg.Env)

	// log = log.With("added for each log")
	storage, err := sqlite.New(cfg.StoragePath)
	// storage, err := sqlite.New(os.Getenv("STORAGE_PATH"))

	if err != nil {
		log.Error("failed to init storage", slogAdapter.Err(err))
		os.Exit(1)
	}

	chiapp.ChiApp(storage, log, cfg)

	// resURL, err := storage.GetURL("tube")
	// if err != nil {
	// 	log.Error("failed to init storage", slogAdapter.Err(err))
	// }
	// log.Info("get resURL", slog.String("resURL", resURL))

	// id, err := storage.SaveURL("http://youtube.com", "tube")
	// if err != nil {
	// 	log.Error("save url", slogAdapter.Err(err))
	// 	os.Exit(1)
	// }
	// log.Info("save url", slog.Int64("id", id))

	log.Info("start chi api", slog.String("env", cfg.Env))
	log.Debug("start chi api")

}
