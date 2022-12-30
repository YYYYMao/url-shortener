package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var Client *redis.Client

type Config struct {
	Address string
	Pw      string
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	config := Config{
		Address: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Pw:      "",
	}
	Client = NewRedisClient(config)
}

func NewRedisClient(config Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Pw,
		DB:       0,
		PoolSize: 10,
	})
	return client
}
