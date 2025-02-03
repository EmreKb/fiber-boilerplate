package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/EmreKb/fiber-boilerplate/api"
	"github.com/EmreKb/fiber-boilerplate/internal/config"
	"github.com/EmreKb/fiber-boilerplate/pkg/database"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	config.Load()
	database.NewPostgres(ctx, config.DatabaseUrl())

	go api.Bootstrap()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
	cancel()
}
