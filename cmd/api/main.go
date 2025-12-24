package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/teptomngo/construction-marketplace/internal/config"
	"github.com/teptomngo/construction-marketplace/internal/db"
	"github.com/teptomngo/construction-marketplace/internal/http"
	"github.com/teptomngo/construction-marketplace/internal/logging"
)

func main() {
	ctx := context.Background()

	// 1) Load config
	cfg, err := config.Load()
	if err != nil {
		// Fallback logger (stdout) in case logging isn't initialized yet
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// 2) Setup logger
	logger := logging.Setup(logging.Options{Env: cfg.AppEnv})
	logger.Info().
		Str("app_env", cfg.AppEnv).
		Str("app_port", cfg.AppPort).
		Msg("starting api")

	// 3) Connect DB
	database, err := db.Connect(ctx, cfg.DBURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect db")
	}
	defer database.Close()

	// 4) Build HTTP server
	addr := ":" + cfg.AppPort
	srv := httpserver.New(addr, logger, database.Pool)

	// 5) Run server (blocking)
	logger.Info().Str("addr", addr).Msg("http server listening")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal().Err(err).Msg("http server failed")
	}

	// NOTE: graceful shutdown is Step 6.9 (signals + Shutdown(ctx))
	_ = os.Stdout // keeps imports stable if you remove fallback logging later
}
