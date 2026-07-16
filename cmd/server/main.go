package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"agency-site/internal/db"
	"agency-site/internal/log"
	"agency-site/internal/server"
	"agency-site/internal/server/router"

	sentry "github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"

	"github.com/joho/godotenv"
)

const defaultRateLimit = 50

var errInvalidRateLimit = errors.New("invalid RATE_LIMIT value")

func main() {
	logger := log.New(
		log.GetLevel(),
		log.GetOutput(),
	)
	err := godotenv.Load()
	if err != nil {
		logger.Error("error  env variables: %s", err.Error())
	}

	if err := run(logger); err != nil {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	sentryDsn := envOrDefault("SENTRY_DSN", "")

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         sentryDsn,
		Environment: "development",
		Release:     "v1.0.0",
	})

	if err != nil {
		return err
	}

	defer sentry.Flush(2 * time.Second)

	database, err := openDatabase()
	if err != nil {
		return err
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			logger.Error("failed to close the database", "error", cerr)
		}
	}()

	port := envOrDefault("PORT", "8080")
	rateLimit, err := parseRateLimit()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sentryHandler := sentryhttp.New(
		sentryhttp.Options{},
	)

	svr := server.New(
		logger,
		":"+port,
		server.WithRouter(sentryHandler.Handle(router.New(ctx, logger, database, rateLimit))),
	)

	return svr.StartAndWait()
}

func openDatabase() (db.Database, error) {
	url := envOrDefault("DB_URL", "./db.sqlite3")
	return db.New(url)
}

func parseRateLimit() (int, error) {
	rateLimitStr := os.Getenv("RATE_LIMIT")
	if rateLimitStr == "" {
		return defaultRateLimit, nil
	}
	parsed, err := strconv.Atoi(rateLimitStr)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("%w: %s", errInvalidRateLimit, rateLimitStr)
	}
	return parsed, nil
}

func envOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
