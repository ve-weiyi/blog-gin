package tokenx

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStore struct {
	client *redis.Client
}

// NewRedisStore 基于 Redis 的 Store 实现。
func NewRedisStore(client *redis.Client) Store {
	return &redisStore{client: client}
}

func (s *redisStore) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *redisStore) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (s *redisStore) Delete(ctx context.Context, keys ...string) error {
	return s.client.Del(ctx, keys...).Err()
}
