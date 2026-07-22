package handler

import (
	"agency-site/internal/components/core"
	errorPage "agency-site/internal/components/error"
	case_studies "agency-site/internal/components/insights/case-studies"
	"agency-site/internal/components/insights/newsroom"
	"agency-site/internal/db/queries"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (h *Handler) NewsRoomDetails(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")
	newsRoomDetails, err := queries.New(h.database.DB()).GetNewsRoomDetails(r.Context(), slug)

	if err != nil {
		h.html(
			r.Context(),
			w,
			http.StatusBadRequest,
			core.HTML(
				"Error Occured",
				sentryDsn,
				h.Header,
				true,
				errorPage.ErrorPage("Something went wrong"),
			),
		)
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
			newsroom.DetailsPage(newsRoomDetails),
		),
	)
}

func (h *Handler) BlogDetails(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")
	blogDetails, err := queries.New(h.database.DB()).GetBlogDetails(r.Context(), slug)

	if err != nil {
		h.html(
			r.Context(),
			w,
			http.StatusBadRequest,
			core.HTML(
				"Error Occured",
				sentryDsn,
				h.Header,
				true,
				errorPage.ErrorPage("Something went wrong"),
			),
		)
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
			newsroom.DetailsPage(blogDetails),
		),
	)
}

func (h *Handler) CaseStudyDetails(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")
	caseStudyDetails, err := queries.New(h.database.DB()).GetCaseStudyDetails(r.Context(), slug)

	if err != nil {
		h.html(
			r.Context(),
			w,
			http.StatusBadRequest,
			core.HTML(
				"Error Occured",
				sentryDsn,
				h.Header,
				true,
				errorPage.ErrorPage("Something went wrong"),
			),
		)
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
			newsroom.DetailsPage(caseStudyDetails),
		),
	)
}

func (h *Handler) CaseStudies(w http.ResponseWriter, r *http.Request) {
	sentryDsn := os.Getenv("SENTRY_BROWSER_DSN")

	search := strings.TrimSpace(r.URL.Query().Get("search"))

	page := 1
	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	// Reset page to 1 whenever the search input triggers the request
	isSearchTrigger := r.Header.Get("HX-Trigger") == "search-input"
	if isSearchTrigger {
		page = 1
	}

	const limit = int64(6) // Set your preferred batch size per load
	db := queries.New(h.database.DB())

	// 1. Fetch Total Count for Infinite Scroll Boundary
	totalCount, err := db.GetCaseStudiesCount(r.Context(), search)
	if err != nil {
		h.renderError(w, r, sentryDsn, "Failed to load count")
		return
	}

	hasMore := int64(page*int(limit)) < totalCount

	// 2. Fetch Data
	caseStudies, err := db.GetPaginatedCaseStudies(
		r.Context(),
		queries.GetPaginatedCaseStudiesParams{
			Search:    search,
			LimitVal:  limit,
			OffsetVal: int64((page - 1) * int(limit)),
		},
	)
	if err != nil {
		h.renderError(w, r, sentryDsn, "Failed to load case studies")
		return
	}

	isHX := r.Header.Get("HX-Request") == "true"

	if isHX {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// SCENARIO A: Search Input Triggered (Page 1)
		// We re-render the wrapper inner HTML to clear previous results
		if isSearchTrigger || page == 1 {
			case_studies.SearchResults(caseStudies, page, hasMore, search).Render(r.Context(), w)
			return
		}

		// SCENARIO B: Scroll Triggered (Page > 1)
		// Append only the new batch of cards
		case_studies.CaseStudyItems(caseStudies, page, hasMore, search).Render(r.Context(), w)
		return
	}

	// Direct Page Load / Refresh
	pageComponent := case_studies.CaseStudyPage(caseStudies, page, hasMore, search)
	h.html(
		r.Context(),
		w,
		http.StatusOK,
		core.HTML(
			"System Limited",
			sentryDsn,
			h.Header,
			true,
			pageComponent,
		),
	)
}

func (h *Handler) renderError(w http.ResponseWriter, r *http.Request, sentryDsn, msg string) {
	h.html(
		r.Context(),
		w,
		http.StatusBadRequest,
		core.HTML(
			"Error Occurred",
			sentryDsn,
			h.Header,
			true,
			errorPage.ErrorPage(msg),
		),
	)
}
