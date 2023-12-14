package chiapp

import (
	"go-rest-microservice/internal/config"
	"go-rest-microservice/internal/storage/sqlite"
	"log/slog"
	"net/http"

	"go-rest-microservice/internal/http-server/handlers/redirect"
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

	r.Route("/url", func(r chi.Router) {

		//TODO скрин прилагается в docs
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.HTTPServer.User: cfg.HTTPServer.Password,
		}))

		r.Post("/", save.New(log, storage))
		r.Get("/{url}", fetchurl.FetchHandler(log, storage))
	})

	r.Get("/{url}", redirect.MakeRedirect(log, storage))

	log.Info("Listen: " + cfg.Address)
	// http.ListenAndServe(cfg.Address, r)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stoped")
}
