package logger

import (
	"io"
	"log/slog"
)

// dev uses human-readable; all other environments use JSON.
func New(w io.Writer, env string, level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if env == "development" {
		handler = slog.NewTextHandler(w, opts)
	} else {
		handler = slog.NewJSONHandler(w, opts)
	}

	return slog.New(handler)
}

func ParseLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
