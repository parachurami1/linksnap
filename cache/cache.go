package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func NewCache() (*Cache, error) {
	var cache Cache
	cln := *redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	cache.Client = &cln
	stat := cache.Client.Ping(context.Background())
	_, err := stat.Result()
	if err != nil {
		return nil, err
	}
	return &cache, nil
}
