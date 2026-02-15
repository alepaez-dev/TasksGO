package http

import (
	"log/slog"

	"github.com/alepaezdev/tasksgo/internal/adapters/http/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(log *slog.Logger, appEnv string) *chi.Mux {
	r := chi.NewRouter()

	// TODO: priority (HIGH) - check api rate limiting.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logging(log))
	r.Use(middleware.SecurityHeaders)
	r.Use(middleware.CORS(appEnv))
	r.Use(middleware.Recoverer(log))

	r.Get("/healthz", HandleHealthz())

	return r
}
