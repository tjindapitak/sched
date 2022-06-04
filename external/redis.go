package external

import (
	"context"
	"log"
	"sched/config"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:       config.Get().Redis.URI,
		PoolSize:   5,
		MaxRetries: 2,
		Password:   config.Get().Redis.Password,
		DB:         0,
	})

	if cmd := client.Ping(); cmd.Err() != nil {
		log.Fatalf("failed to ping redis: %s\n", cmd.Err())
	}

	return client
}
