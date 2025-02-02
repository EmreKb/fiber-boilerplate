package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var required = []string{
	"PORT",
	"NODE_ENV",

	"DATABASE_HOST",
	"DATABASE_PORT",
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

func DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", Get("DATABASE_USER"), Get("DATABASE_PASSWORD"), Get("DATABASE_HOST"), Get("DATABASE_PORT"), Get("DATABASE_DB"))
}
