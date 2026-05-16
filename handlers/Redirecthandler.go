package handlers

import (
	"linksnap/cache"
	"linksnap/service"
	"linksnap/storage"
	"log/slog"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	cac, ok := r.Context().Value("cache").(*cache.Cache)
	if !ok {
		slog.Error("No cache found in context")
		return
	}
	link, err := service.GetLinkBySlug(cac, slug)
	if err != nil {
		http.Error(w, "Link not found", http.StatusBadRequest)
		slog.Error("Link not found", "Error", err)
		return
	}
	http.Redirect(w, r, link.Url, http.StatusFound)
	storage.UpdateCount(cac, slug)

}
