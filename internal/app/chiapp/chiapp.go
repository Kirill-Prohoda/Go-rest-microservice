package chiapp

import (
	"go-rest-microservice/internal/lib/logger/slogAdapter"
	"go-rest-microservice/internal/storage/sqlite"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiApp(storage *sqlite.Storage, log *slog.Logger) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger) //TODO нужно написать собственный логер
	r.Use(middleware.Recoverer)

	r.Get("/tube", func(w http.ResponseWriter, r *http.Request) {
		resURL, err := storage.GetURL("tube")
		if err != nil {
			log.Error("failed to init storage", slogAdapter.Err(err))
		}
		log.Info("get resURL", slog.String("resURL", resURL))

		w.Write([]byte(resURL))
	})

	println("start Chiap")

	http.ListenAndServe(":3333", r)

}
