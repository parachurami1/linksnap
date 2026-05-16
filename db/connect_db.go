package db

import (
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func ConnectDB() {
	DB, err = sqlx.Open("pgx", os.Getenv("POSTGRES_URL"))
	if err != nil {
		slog.Error("Databese error", "Error", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		slog.Error("Database error", "Error", err)
		return
	}
	slog.Info("Database Connected")
}
