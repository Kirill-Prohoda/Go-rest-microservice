package chiapp

import (
	"go-rest-microservice/internal/config"
	"go-rest-microservice/internal/storage/sqlite"
	"log/slog"
	"net/http"

	"go-rest-microservice/internal/http-server/handlers/url/fetchurl"
	"go-rest-microservice/internal/http-server/handlers/url/save"
	mw "go-rest-microservice/internal/http-server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApp(storage *sqlite.Storage, log *slog.Logger, cfg *config.Config) {

	log.Info("Start Chi app")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	// r.Use(middleware.Logger) //TODO нужно написать собственный логер
	r.Use(mw.NewLogEntry(log))  //TODO нужно написать собственный логер
	r.Use(middleware.Recoverer) // восстановление после паники
	r.Use(middleware.RealIP)
	r.Use(middleware.URLFormat) // позволяет использовать -> /post/{id}

	r.Post("/save-url", save.New(log, storage))
	r.Get("/get-url/{url}", fetchurl.FetchHandler(log, storage))

	log.Info("Listen: " + cfg.Address)
	http.ListenAndServe(cfg.Address, r)

}
