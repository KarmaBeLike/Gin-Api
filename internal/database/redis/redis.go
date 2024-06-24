package redis

import (
	"fmt"

	"Gin-Api/config"

	"github.com/go-redis/redis"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(ping)

	return client, nil
}
