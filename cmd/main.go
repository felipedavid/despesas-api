package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipedavid/saldop/handlers"
	"github.com/felipedavid/saldop/storage"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5"
)

type AppConfig struct {
	Addr       string
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

func newAppConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg.Addr = os.Getenv("ADDR")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = os.Getenv("DB_PORT")

	return cfg, err
}

func runApp() error {
	cfg, err := newAppConfig()
	if err != nil {
		return err
	}

	slog.Info("Connecting to the database", "user", cfg.DbUser, "dbname", cfg.DbName, "dbhost", cfg.DbHost, "port", cfg.DbPort)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}

	storage.Init(conn)

	slog.Info("Starting http server", "addr", cfg.Addr)

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
