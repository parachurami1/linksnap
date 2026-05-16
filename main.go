package main

import (
	"fmt"
	"linksnap/cache"
	"linksnap/db"
	"linksnap/handlers"
	"linksnap/middleware"
	"log/slog"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	db.ConnectDB()
	db.RunMigrations(os.Getenv("POSTGRES_URL"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	cac, err := cache.NewCache()
	if err != nil {
		slog.Error("Error", "error", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "The server is running") })
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/links", middleware.Protect(handlers.LinksHandler))
	mux.HandleFunc("/links/{slug}", handlers.RedirectHandler)

	c := cors.AllowAll()

	handler := c.Handler(middleware.Logging(middleware.PassCache(cac, mux)))
	slog.Info("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
