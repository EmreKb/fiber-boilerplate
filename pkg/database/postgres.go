package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, url string) *Postgres {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")

	return &Postgres{
		Pool: pool,
	}
}
