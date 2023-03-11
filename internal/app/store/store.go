package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	rdbConfig *RedisConfig
	rdb       *redis.Client
	rctx      context.Context
}

func NewRedisStore(config *RedisConfig) *RedisStore {
	return &RedisStore{
		rdbConfig: config,
	}
}

func (s *RedisStore) ConfigureStore() error {
	s.rctx = context.Background()
	s.rdb = redis.NewClient(&redis.Options{
		Addr:     s.rdbConfig.RedisURL,
		Password: s.rdbConfig.Password,
		DB:       0,
	})

	return s.rdb.Ping(s.rctx).Err()
}
