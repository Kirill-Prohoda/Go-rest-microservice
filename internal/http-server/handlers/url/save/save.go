package save

import (
	"errors"
	resp "go-rest-microservice/internal/lib/api/response"
	"go-rest-microservice/internal/lib/logger/slogAdapter"
	"go-rest-microservice/internal/lib/random"
	"go-rest-microservice/internal/storage"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
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
	Alias string `json:"alias,omitempty"`
}

// TODO настройка для уровня логирования лучше перенести в конфиг
const aliasLength = 6

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.saver.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slogAdapter.Err(err))
			render.JSON(w, r, resp.Error("failed to decode request"))
			// os.Exit(1)
			return
		}

		log.Info("request body decode", slog.Any("request", req))

    //TODO мешает в режиме дебага
		// if err := validator.New().Struct(req); err != nil {
		// 	validateErr := err.(validator.ValidationErrors)
		// 	log.Error("invalidate request", slogAdapter.Err(err))
		// 	render.JSON(w, r, resp.ValidationError(validateErr))
		// 	// os.Exit(1)
		// 	return
		// }

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

		id, err := urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info(storage.ErrURLExists.Error(), slog.String("url", req.URL))

			render.JSON(w, r, resp.Error(storage.ErrURLExists.Error()))
			return
		}

		log.Info("url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
