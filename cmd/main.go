package main

import (
	"github.com/EmreKb/fiber-boilerplate/api"
	"github.com/EmreKb/fiber-boilerplate/internal/config"
)

func main() {
	config.Load()
	api.Bootstrap()
}
