package main

import (
	"fmt"

	"github.com/EmreKb/fiber-boilerplate/internal/config"
)

func main() {
	config.Load()
	fmt.Println("Working...")
}
