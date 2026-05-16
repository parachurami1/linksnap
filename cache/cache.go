package cache

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func NewCache() (*Cache, error) {
	var cache Cache
	cln := *redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
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

func (c *Cache) DeleteByPattern(pattern string) error {
	var cursor uint64
	for {
		keys, nextCursor, err := c.Client.Scan(context.Background(), cursor, pattern, 10).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			c.Client.Del(context.Background(), keys...)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
