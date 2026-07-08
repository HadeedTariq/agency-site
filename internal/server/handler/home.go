package handler

import (
	"agency-site/internal/components/core"
	"agency-site/internal/components/home"
	"net/http"
)

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.html(r.Context(), w, http.StatusOK, core.HTML("Example Site", h.Header, true, home.Page()))
}
