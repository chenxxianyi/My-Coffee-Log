package database

import (
	"context"
	"log"

	"my-coffee-log/internal/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	cfg := config.AppConfig
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHost + ":" + cfg.REDISPort,
		Password: cfg.REDISPassword,
		DB:       0,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
		RDB = nil
		return
	}

	log.Println("Redis connected successfully")
}
