package main

import (
	"fmt"
	"net/http"
	"os"

	apphttp "github.com/alepaezdev/tasksgo/internal/adapters/http"
	"github.com/alepaezdev/tasksgo/internal/infra/config"
	"github.com/alepaezdev/tasksgo/internal/infra/logger"
	"github.com/alepaezdev/tasksgo/internal/infra/shutdown"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	log := logger.New(os.Stdout, cfg.App.Env, logger.ParseLevel(cfg.LogLevel))
	router := apphttp.NewRouter(log, cfg.App.Env)

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Info("server starting", "addr", addr, "env", cfg.App.Env)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server listen failed", "error", err)
			os.Exit(1)
		}
	}()

	shutdown.Graceful(srv, log)
}
