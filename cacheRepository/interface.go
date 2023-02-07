package cacheRepository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type ICacheRepository interface {
	AddToCache(ctx context.Context, key, value string) error
	GetUsersFromCache(ctx context.Context, key string) (string, error)
	DeleteValue(ctx context.Context, key string) error
	ClearCache(ctx context.Context) error
}

type cache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewCacheClient(addr, password string, ctx context.Context) (ICacheRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("Error connecting to redis: %v", err)
		return nil, err
	}

	return &cache{client: rdb, ttl: time.Hour * 1}, nil
}