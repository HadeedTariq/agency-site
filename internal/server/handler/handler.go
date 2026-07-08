package handler

import (
	"agency-site/internal/db"
	"agency-site/internal/types"
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// Handler handles requests.
type Handler struct {
	logger   *slog.Logger
	database db.Database
	Header   types.HeaderConfig
	todos    []types.TodoItem
	nextId   int
}

// New creates a new Handler.
func New(logger *slog.Logger, database db.Database, header types.HeaderConfig) *Handler {
	return &Handler{logger: logger, database: database, Header: header, todos: []types.TodoItem{}, nextId: 1}
}

func (h *Handler) html(ctx context.Context, w http.ResponseWriter, status int, t templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	// Use WithoutCancel so a client disconnect doesn't truncate a partially-written response.
	if err := t.Render(context.WithoutCancel(ctx), w); err != nil {
		h.logger.Error("Failed to render component", "error", err)
	}
}
