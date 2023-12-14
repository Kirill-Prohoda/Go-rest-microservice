package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func NewLogEntry(log *slog.Logger) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		log = log.With("component", "middleware/logger")

		log.Info("logger middleware enable")

		fn := func(w http.ResponseWriter, r *http.Request) {

			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()

			defer func() {
				entry.Info(
					"request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)

		}

		return http.HandlerFunc(fn)
	}
}
