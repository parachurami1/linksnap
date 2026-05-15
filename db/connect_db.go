package db

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func ConnectDB() {
	DB, err = sqlx.Open("pgx", "host=localhost port=5432 user=postgres password=p4ssw0rd database=Linksnap")
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
