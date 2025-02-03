package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(addr string, password string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return &Redis{
		Client: client,
	}
}

func (c *Redis) Set(ctx context.Context, key string, value interface{}) error {
	return c.Client.Set(ctx, key, value, 0).Err()
}

func (c *Redis) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}
