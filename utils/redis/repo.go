package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Repository interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type repository struct {
	Client redis.Cmdable
}

func NewRedisRepository(Client redis.Cmdable) Repository {
	return &repository{Client}
}

func (r *repository) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(ctx, key, value, exp).Err()
}

func (r *repository) Get(ctx context.Context, key string) (string, error) {
	get := r.Client.Get(ctx, key)
	return get.Result()
}
