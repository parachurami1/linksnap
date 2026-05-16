package handlers

import (
	"encoding/json"
	"fmt"
	"linksnap/auth"
	"linksnap/cache"
	"linksnap/service"
	"log/slog"
	"net/http"
	"strings"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error("Error decoding body", "Error", err)
		return
	}
	err = auth.Register(user.Email, user.Password, "user")
	if err != nil {
		fmt.Println(err)
		if strings.Contains("Invalid email address", err.Error()) {
			http.Error(w, "Invalid email address", http.StatusBadRequest)
			slog.Error("Error Creating user", "Error", err)
			return
		}
		http.Error(w, "Problem registering user", http.StatusInternalServerError)
		slog.Error("Error Creating user", "Error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "User created successfully",
	})
	slog.Info("User created successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error("Error decoding body", "Error", err)
		return
	}

	cac, ok := r.Context().Value("cache").(*cache.Cache)
	if !ok {
		slog.Error("No cache found in context")
		return
	}
	person, err := service.GetUserByEmail(cac, user.Email)
	if err != nil {
		slog.Error("Error getting user", "Error", err)
		return
	}
	tok, err := auth.GenerateToken(person.Id, person.Role)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tok,
	})
}
