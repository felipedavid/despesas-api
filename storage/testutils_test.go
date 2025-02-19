package storage

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	pgxmigrate "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func TestMain(m *testing.M) {
	connStr := "user=test password=test dbname=test_db host=localhost port=5433 sslmode=disable"
	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}
	defer connPool.Close()

	if err := connPool.Ping(context.Background()); err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	db := stdlib.OpenDBFromPool(connPool)
	driver, err := pgxmigrate.WithInstance(db, &pgxmigrate.Config{})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	mi, err := migrate.NewWithDatabaseInstance("file://../migrations", "postgres", driver)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	err = mi.Up()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	Init(connPool)

	exitCode := m.Run()
	os.Exit(exitCode)
}
