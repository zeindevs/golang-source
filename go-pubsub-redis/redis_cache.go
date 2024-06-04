package main

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(c *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: c,
		ttl:    ttl,
	}
}

func (c *RedisCache) Get(key int) (string, bool) {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	val, err := c.client.Get(ctx, keyStr).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (c *RedisCache) Set(key int, val string) error {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	_, err := c.client.Set(ctx, keyStr, val, c.ttl).Result()
	return err
}

func (c *RedisCache) Remove(key int) error {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	_, err := c.client.Del(ctx, keyStr).Result()
	return err
}
