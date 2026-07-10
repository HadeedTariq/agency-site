package handler

import (
	"agency-site/internal/components/core"
	errorPage "agency-site/internal/components/error"
	"agency-site/internal/components/insights/newsroom"
	"agency-site/internal/db/queries"
	"net/http"
)

func (h *Handler) NewsRoomDetails(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	newsRoomDetails, err := queries.New(h.database.DB()).GetNewsRoomDetails(r.Context(), slug)

	if err != nil {
		h.html(
			r.Context(),
			w,
			http.StatusBadRequest,
			core.HTML(
				"Error Occured",
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
			h.Header,
			true,
			newsroom.DetailsPage(newsRoomDetails),
		),
	)
}
