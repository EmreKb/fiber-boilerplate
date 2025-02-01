package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var required = []string{
	"PORT",
	"NODE_ENV",

	"DATABASE_USER",
	"DATABASE_PASSWORD",
	"DATABASE_DB",

	"JWT_ACCESS_SECRET",
	"JWT_ACCESS_EXPIRE",

	"JWT_REFRESH_SECRET",
	"JWT_REFRESH_EXPIRE",
}

func Load() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	for _, key := range required {
		if Get(key) == "" {
			panic(fmt.Sprintf("Missing environment variable: %s", key))
		}
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
