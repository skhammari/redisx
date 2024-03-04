package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type StringOps struct {
	client *redis.Client
}

func NewStringOps() *StringOps {
	client := RedisX()
	return &StringOps{
		client: client,
	}
}

func (s *StringOps) Set(key string, value string) error {
	return s.client.Set(context.Background(), key, value, 0).Err()
}

func (s *StringOps) Get(key string) (string, error) {
	return s.client.Get(context.Background(), key).Result()
}

func (s *StringOps) Delete(key string) error {
	return s.client.Del(context.Background(), key).Err()
}

func (s *StringOps) Increment(key string) error {
	return s.client.Incr(context.Background(), key).Err()
}

func (s *StringOps) IncrementBy(key string, value int) error {
	return s.client.IncrBy(context.Background(), key, int64(value)).Err()
}

func (s *StringOps) GetSet(key string, value string) (string, error) {
	return s.client.GetSet(context.Background(), key, value).Result()
}

func (s *StringOps) MSet(key string, value string) error {
	return s.client.MSet(context.Background(), key, value).Err()
}

// todo: add MGet
/*func (s *StringOps) MGet(keys ...string) ([]string, error) {
	return s.client.MGet(context.Background(), keys...).Result()
}*/
