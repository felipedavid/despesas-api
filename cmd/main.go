package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipedavid/saldop/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	slog.Info("Connecting to the database", "user", cfg.DbUser, "dbname", cfg.DbName, "dbhost", cfg.DbHost, "port", cfg.DbPort)

	if err := conn.Ping(); err != nil {
		return err
	}

	slog.Info("Starting http server", "addr", cfg.Addr)

	s := http.Server{
		Handler: handlers.SetupRoutes(),
		Addr:    cfg.Addr,
	}

	err = s.ListenAndServe()
	return err
}

func main() {
	err := runApp()
	slog.Error("Application exited", "err", err)
}
