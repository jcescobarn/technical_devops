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

func NewRedisConfig(address, password string, db int) *RedisConfig {
	return &RedisConfig{
		Address:  address,
		Password: password,
		DB:       db,
	}
}

func (rc *RedisConfig) Connect() (*redis.Client, error) {

	address := fmt.Sprintf("%s:%d", rc.Address, 6379)
	fmt.Printf("Connecting to Redis with:\n")
	fmt.Printf("Addr: %s\n", address)
	fmt.Printf("DB: %d\n", rc.DB)

	client := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   rc.DB,
	})

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println(err)
		log.Fatalf("Error de conexión")
		return nil, err
	}

	fmt.Println("Redis Connection Successful")
	return client, nil
}
