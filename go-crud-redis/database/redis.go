package database

import (
	"fmt"

	"go-crud-redis/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "",
		DB:       0,
	})

	fmt.Println("Connected Successfully to the database (Redis)")

	return rdb
}
