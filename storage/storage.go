package storage

import "github.com/jackc/pgx/v5"

var conn *pgx.Conn

func Init(c *pgx.Conn) {
	conn = c
}
