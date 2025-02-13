package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/felipedavid/saldop/handlers"
	"github.com/felipedavid/saldop/storage"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/felipedavid/saldop/internal/translations"
)

var Version = ""

// REMINDER: When going to push to production, remember to create a new google
// oauth thing not vinculated to my real name.

func runApp() error {
	cfg, err := newAppConfig()
	if err != nil {
		return err
	}

	slog.Info("Creating connection pool for the database", "user", cfg.DbUser, "dbname", cfg.DbName, "dbhost", cfg.DbHost, "port", cfg.DbPort)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %s", err)
	}
	defer connPool.Close()

	slog.Info("Pinging database")
	if err := connPool.Ping(context.Background()); err != nil {
		return err
	}

	storage.Init(connPool)

	slog.Info("Starting http server", "addr", cfg.Addr, "version", Version)

	s := http.Server{
		Handler: handlers.SetupMultiplexer(),
		Addr:    cfg.Addr,
	}

	err = s.ListenAndServe()
	return err
}

func main() {
	err := runApp()
	slog.Error("Application exited", "err", err)
}
