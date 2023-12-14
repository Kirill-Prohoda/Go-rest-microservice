package fetchurl

import (
	resp "go-rest-microservice/internal/lib/api/response"
	"go-rest-microservice/internal/lib/logger/slogAdapter"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Request struct {
	URL   string `json:"url" validate:"requires,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	// Status string `json:"status"`
	// Error string `json:"error,omitempty"`
	resp.Response
	URL string `json:"url,omitempty"`
}
type URLFetcher interface {
	GetURL(alias string) (string, error)
}

func FetchHandler(log *slog.Logger, urlFetcher URLFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("get resURL", slog.String("reschi.URLParam", chi.URLParam(r, "id")))

		resURL, err := urlFetcher.GetURL(chi.URLParam(r, "url"))
		if err != nil {
			log.Error("failed to init storage", slogAdapter.Err(err))
		}
		log.Info("get resURL", slog.String("resURL", resURL))

		render.JSON(w, r, Response{
			Response: resp.OK(),
			URL:      resURL,
		})
	}
}
