package main

import (
	"context"

	"github.com/EmreKb/fiber-boilerplate/api"
	"github.com/EmreKb/fiber-boilerplate/internal/config"
	"github.com/EmreKb/fiber-boilerplate/pkg/database"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config.Load()

	database.NewPostgres(ctx, config.DatabaseUrl())

	api.Bootstrap()
}
