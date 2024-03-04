package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func initDefaultClient(options *redis.Options) {
	redisClient = redis.NewClient(options)
}

func getDefaultClient() *redis.Client {
	return redisClient
}

func RedisX() *redis.Client {
	redisOptions := &redis.Options{}
	initDefaultClient(redisOptions)
	client := getDefaultClient()
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis: ", err)
	} else {
		fmt.Println("Connected to Redis")
	}

	return client
}
