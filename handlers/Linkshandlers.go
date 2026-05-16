package handlers

import (
	"encoding/json"
	"fmt"
	"linksnap/cache"
	"linksnap/models"
	"linksnap/service"
	"log/slog"
	"net/http"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var url struct {
			Link string `json:"link"`
		}
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			http.Error(w, "Could not parse request", http.StatusBadRequest)
			return
		}
		clms, ok := r.Context().Value("claims").(*models.Claim)
		if !ok {
			slog.Error("Could not find claims")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cac, ok := r.Context().Value("cache").(*cache.Cache)
		if !ok {
			slog.Error("No cache found in context")
			return
		}
		err = service.CreateLink(cac, url.Link, service.Shorten(url.Link), clms.User_ID)
		if err != nil {
			http.Error(w, "Could not create link", http.StatusInternalServerError)
			slog.Error("Could not create link", "Error", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Successfully created link")
	case http.MethodGet:
		cac := r.Context().Value("cache").(*cache.Cache)
		clms, ok := r.Context().Value("claims").(*models.Claim)
		if !ok {
			slog.Error("Could not find claims")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		urls, err := service.GetLinksByID(cac, clms.User_ID)
		if err != nil {
			slog.Error("Failed to get links", "Error", err)
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&urls)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
