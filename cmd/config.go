package main

import (
	"os"

	"github.com/joho/godotenv"
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

	_ = godotenv.Load()

	cfg.Addr = os.Getenv("ADDR")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.GoogleOauthClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	cfg.GoogleOauthClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")

	return cfg, nil
}
