package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func (rc *RedisConfig) Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Address,
		Password: rc.Password,
		DB:       rc.DB,
	})

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis Connection Successful")
	return client
}
