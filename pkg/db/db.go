package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	Name     string
	QueryRow string
}

type DB struct {
	pool *pgxpool.Pool
}

func (db *DB) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, q.QueryRow, args...)
}
