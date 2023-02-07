package cacheRepository

import "context"

func (c *cache) AddToCache(ctx context.Context, key, value string) error

func (c *cache) GetUserFromCache(ctx context.Context, key string) (string, error)

func (c *cache) CleanCache(ctx context.Context) error
