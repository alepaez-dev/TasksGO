package logger

import (
	"io"
	"log/slog"
)

// New returns a slog.Logger configured for the given environment: human-readable
// text in development, JSON elsewhere, both filtered by the given level.
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

// ParseLevel converts a level name (debug, info, warn, error) to a slog.Level,
// defaulting to LevelInfo for unrecognized values.
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
