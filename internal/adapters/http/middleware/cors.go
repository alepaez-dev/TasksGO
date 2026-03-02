package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORS(appEnv string) func(http.Handler) http.Handler {
	var allowedOrigins []string

	// TODO: priority (LOW) - production, select allowed origins from configuration, as an env.
	if appEnv == "development" {
		allowedOrigins = []string{"http://localhost:3000"}
	}

	// block all cross-origin requests.
	if len(allowedOrigins) == 0 {
		return func(next http.Handler) http.Handler { return next }
	}

	return cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
