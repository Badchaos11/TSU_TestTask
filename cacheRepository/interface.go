package cacheRepository

import "context"

type ICacheRepository interface {
	AddToCache(ctx context.Context, key, value string) error
	GetUserFromCache(ctx context.Context, key string) (string, error)
	CleanCache(ctx context.Context) error
}
