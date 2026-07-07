package router

import (
	"context"
	"log/slog"
	"net/http"

	"agency-site/internal/db"
	"agency-site/internal/dist"
	"agency-site/internal/server/config"
	"agency-site/internal/server/handler"
	"agency-site/internal/server/middleware"
	"agency-site/internal/version"
)

// New creates a new router with the given context, logger, database, and rate limit.
func New(ctx context.Context, logger *slog.Logger, database db.Database, rateLimit int) http.Handler {
	headerData, err := config.LoadHeader()

	if err != nil {
		logger.Error("error loading header data: %s", err.Error())
	}
	h := handler.New(logger, database, headerData)

	ipCfg := middleware.IPConfig{
		TrustProxyHeaders: version.Value != "dev",
	}

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc(newPath(http.MethodGet, "/health"), h.Health)
	mux.Handle(newPath(http.MethodGet, "/assets/"), middleware.CacheMiddleware(http.FileServer(http.FS(dist.AssetsDir))))
	mux.HandleFunc(newPath(http.MethodGet, "/{$}"), h.Home)
	mux.HandleFunc(newPath(http.MethodPost, "/count"), h.Count)

	// Middleware chain
	hdlr := http.Handler(mux)
	hdlr = middleware.Chain(
		middleware.Recovery(logger, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		})),
		middleware.Logging(logger, ipCfg),
		middleware.Security(logger, ipCfg),
		middleware.RateLimit(ctx, logger, rateLimit, middleware.DefaultMaxEntries, ipCfg),
		middleware.CSRF(logger, ipCfg),
	)(hdlr)

	return hdlr
}

func newPath(method string, path string) string {
	return method + " " + path
}
