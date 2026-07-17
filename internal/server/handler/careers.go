package handler

import (
	"agency-site/internal/components/careers"
	"agency-site/internal/components/core"
	"agency-site/internal/server/config"
	"net/http"
	"os"
)

func (h *Handler) CareersPage(w http.ResponseWriter, r *http.Request) {
	benefitsData, _ := config.LoadBenefits()

	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")

	h.html(
		r.Context(),
		w,
		http.StatusOK,
		core.HTML(
			"System Limited",
			sentryDsn,
			h.Header,
			true,
			careers.Page(benefitsData),
		),
	)
}
