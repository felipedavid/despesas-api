package storage

import (
	"context"
	"log/slog"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMain(t *testing.T) {
	connStr := "user=test password=test dbname=test_db host=localhost port=5432 sslmode=disable"
	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer connPool.Close()

	slog.Info("Pinging database")
	if err := connPool.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	Init(connPool)
}
