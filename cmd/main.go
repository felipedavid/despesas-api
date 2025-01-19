package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipedavid/saldop/handlers"
	"github.com/felipedavid/saldop/storage"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AppConfig struct {
	Addr                    string
	DbUser                  string
	DbPassword              string
	DbName                  string
	DbHost                  string
	DbPort                  string
	GoogleOauthClientID     string
	GoogleOauthClientSecret string
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
	cfg.GoogleOauthClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	cfg.GoogleOauthClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")

	return cfg, err
}

// REMINDER: When going to push to production, remember to create a new google oauth thing not vinculated to my real name.

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

	store := sessions.NewCookieStore([]byte("alsdfjasd8fajsdf8jasdfjkj"))
	gothic.Store = store

	gProvider := google.New(cfg.GoogleOauthClientID, cfg.GoogleOauthClientSecret, "http://localhost:8080/auth/google/callback")
	goth.UseProviders(gProvider)

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
