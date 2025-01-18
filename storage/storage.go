package storage

import "github.com/jackc/pgx/v5/pgxpool"

var conn *pgxpool.Pool

func Init(c *pgxpool.Pool) {
	conn = c
}
