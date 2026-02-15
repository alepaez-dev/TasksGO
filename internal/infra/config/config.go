package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App      App
	Postgres Postgres
	LogLevel string
}

type App struct {
	Env  string
	Port int
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func (p Postgres) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DB,
	)
}

func Load() (Config, error) {
	var errs []error

	appEnv := requireEnv("APP_ENV", &errs)
	appPort := requireIntEnv("APP_PORT", &errs)
	pgHost := requireEnv("POSTGRES_HOST", &errs)
	pgPort := requireIntEnv("POSTGRES_PORT", &errs)
	pgUser := requireEnv("POSTGRES_USER", &errs)
	pgPassword := requireEnv("POSTGRES_PASSWORD", &errs)
	pgDB := requireEnv("POSTGRES_DB", &errs)
	logLevel := envOrDefault("LOG_LEVEL", "info")

	if len(errs) > 0 {
		return Config{}, fmt.Errorf("config validation failed: %w", errors.Join(errs...))
	}

	return Config{
		App: App{
			Env:  appEnv,
			Port: appPort,
		},
		Postgres: Postgres{
			Host:     pgHost,
			Port:     pgPort,
			User:     pgUser,
			Password: pgPassword,
			DB:       pgDB,
		},
		LogLevel: logLevel,
	}, nil
}

func requireEnv(key string, errs *[]error) string {
	val := os.Getenv(key)
	if val == "" {
		*errs = append(*errs, fmt.Errorf("%s is required", key))
	}
	return val
}

func requireIntEnv(key string, errs *[]error) int {
	raw := requireEnv(key, errs)
	if raw == "" {
		return 0
	}
	val, err := strconv.Atoi(raw)
	if err != nil {
		*errs = append(*errs, fmt.Errorf("%s must be an integer, got %q", key, raw))
		return 0
	}
	return val
}

func envOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
