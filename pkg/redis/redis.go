package redis

import (
	"context"
	"log"

	"bingo-backend/internal/config"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() {

	opt, err := redis.ParseURL(config.App.RedisURL)

	if err != nil {
		log.Fatal(err)
	}

	Client = redis.NewClient(opt)

	_, err = Client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("Connected to Upstash Redis")
}