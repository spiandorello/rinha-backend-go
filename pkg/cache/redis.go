package cache

import (
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Cache *cache.Cache
}

func NewRedis() *Redis {
	client := redis.NewClient(&redis.Options{
		//Addr: "localhost:6379",
		Addr: "rb-redis:6379",
	})

	mycache := cache.New(&cache.Options{
		Redis: client,
	})

	return &Redis{
		Cache: mycache,
	}
}
