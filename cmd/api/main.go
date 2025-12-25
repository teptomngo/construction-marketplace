package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/teptomngo/construction-marketplace/internal/config"
	"github.com/teptomngo/construction-marketplace/internal/db"
	httpserver "github.com/teptomngo/construction-marketplace/internal/http"
	"github.com/teptomngo/construction-marketplace/internal/logging"
)

func main() {
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
	ctx := context.Background()
	database, err := db.Connect(ctx, cfg.DBURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect db")
	}

	// 4) Build HTTP server
	addr := ":" + cfg.AppPort
	srv := httpserver.New(addr, logger, database.Pool)

	// 5) Run server (non-blocking) so we can handle shutdown signals
	errCh := make(chan error, 1)

	go func() {
		logger.Info().Str("addr", addr).Msg("http server listening")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	// 6) Wait for SIGINT/SIGTERM or server error
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		logger.Info().Str("signal", sig.String()).Msg("shutdown signal received")
	case err := <-errCh:
		logger.Fatal().Err(err).Msg("http server failed")
		// server exited normally (unlikely in our flow, but handle it)
		//logger.Info().Msg("http server exited")
	}

	// 7) Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info().Msg("shutting down http server")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error().Err(err).Msg("http server shutdown error")
	}

	logger.Info().Msg("closing db pool")
	database.Close()

	logger.Info().Msg("shutdown complete")
}
