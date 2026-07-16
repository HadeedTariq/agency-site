package handler

import (
	"agency-site/internal/components/core"
	"agency-site/internal/components/home"
	"agency-site/internal/db/queries"
	"agency-site/internal/server/config"
	"fmt"
	"net/http"
	"os"
)

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	carouselData, _ := config.LoadCarousel()
	advantageData, _ := config.LoadAdvantages()
	testimonialData, _ := config.LoadTestimonial()
	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")
	fmt.Println(sentryDsn)

	insights, err := queries.New(h.database.DB()).GetHomePageInsights(r.Context())
	if err != nil {
		insights = []queries.GetHomePageInsightsRow{}
	}

	h.html(
		r.Context(),
		w,
		http.StatusOK,
		core.HTML(
			"System Limited",
			sentryDsn,
			h.Header,
			true,
			home.Page(carouselData, insights, advantageData, testimonialData),
		),
	)
}
