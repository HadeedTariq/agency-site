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

	if r.Header.Get("HX-Trigger") == "search-input" {
		page = 1
	}

	totalCount, err := queries.New(h.database.DB()).GetCaseStudiesCount(r.Context(), search)

	const limit = int64(4)

	totalPages := 1
	if totalCount > 0 {
		totalPages = int((totalCount + limit - 1) / limit)
	}

	if page > totalPages {
		page = 1
	}

	caseStudies, err := queries.New(h.database.DB()).GetPaginatedCaseStudies(
		r.Context(),
		queries.GetPaginatedCaseStudiesParams{
			Search:    search,
			LimitVal:  limit,
			OffsetVal: int64((page - 1) * int(limit)),
		},
	)

	pageComponent := case_studies.CaseStudyPage(caseStudies, page, totalPages)
	isHX := r.Header.Get("HX-Request") == "true"

	if isHX {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		case_studies.CaseStudiesCards(caseStudies, page, totalPages).Render(r.Context(), w)
		return
	}

	if err != nil {
		h.renderError(w, r, sentryDsn, "Failed to load case studies")
		return // ALWAYS return on error
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
