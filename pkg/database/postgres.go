package database

import (
	"context"

	"github.com/EmreKb/fiber-boilerplate/pkg/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewPostgres(ctx context.Context, url string) *Postgres {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}

	queries := sqlc.New(pool)

	return &Postgres{
		Pool:    pool,
		Queries: queries,
	}
}
