package handler

import (
	"agency-site/internal/components/core"
	"agency-site/internal/components/home"
	"agency-site/internal/db/queries"
	"agency-site/internal/server/config"
	"net/http"
)

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	carouselData, _ := config.LoadCarousel()

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
			h.Header,
			true,
			home.Page(carouselData, insights),
		),
	)
}
