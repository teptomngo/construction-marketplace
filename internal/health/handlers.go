package health

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handlers struct {
	DB *pgxpool.Pool
}

func (h Handlers) Health(w http.ResponseWriter, r *http.Request) {
	// Liveness: process is up
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h Handlers) Ready(w http.ResponseWriter, r *http.Request) {
	// Readiness: dependencies are ready (DB reachable)
	if h.DB == nil {
		http.Error(w, "db not configured", http.StatusServiceUnavailable)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.DB.Ping(ctx); err != nil {
		http.Error(w, "db not ready", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
