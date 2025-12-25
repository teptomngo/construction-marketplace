package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/teptomngo/construction-marketplace/internal/health"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, logger zerolog.Logger, dbPool *pgxpool.Pool) *Server {
	mux := http.NewServeMux()

	// Health endpoints
	h := health.Handlers{DB: dbPool}
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/ready", h.Ready)

	// Basic HTTP server with sane timeouts
	srv := &http.Server{
		Addr:              addr,
		Handler:           requestLogMiddleware(logger, mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &Server{httpServer: srv}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func requestLogMiddleware(logger zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Dur("duration", time.Since(start)).
			Msg("http_request")
	})
}
