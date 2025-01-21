package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/felipedavid/saldop/handlers"
	"github.com/felipedavid/saldop/storage"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/felipedavid/saldop/translations"
)

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

	//store := sessions.NewCookieStore([]byte("alsdfjasd8fajsdf8jasdfjkj"))
	//gothic.Store = store

	//gProvider := google.New(cfg.GoogleOauthClientID, cfg.GoogleOauthClientSecret, "http://localhost:8080/auth/google/callback")
	//goth.UseProviders(gProvider)

	sessionManager := scs.New()
	//sessionManager.Lifetime = 24 * time.Hour
	//sessionManager.Store = connPool

	handlers.Init(sessionManager)

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
