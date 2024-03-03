package redisx

import "github.com/redis/go-redis/v9"

type RedisX struct {
	client *redis.Client
}

func NewRedisX(options *redis.Options) *RedisX {
	client := redis.NewClient(options)
	return &RedisX{
		client: client,
	}
}
