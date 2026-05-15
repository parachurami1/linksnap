package db

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dburl string) {
	pth, _ := os.Getwd()
	fil := filepath.Join(pth, "migrations")
	migrationsURL := "file://" + filepath.ToSlash(fil) // ← three slashes

	slog.Info("Running migrations from", "path", migrationsURL)

	m, err := migrate.New(migrationsURL, dburl)
	if err != nil {
		slog.Error("Error running migrations", "Error", err.Error())
		return
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("Error running migrations", "Error", err.Error())
		return
	}

	slog.Info("Migrations run successfully")
}
