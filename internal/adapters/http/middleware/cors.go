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

	return cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
