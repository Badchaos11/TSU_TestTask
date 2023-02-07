package cacheRepository

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (c *cache) AddToCache(ctx context.Context, key, value string) error {
	_, err := c.client.SetEx(ctx, key, value, c.ttl).Result()
	if err != nil {
		logrus.Errorf("error setting cache key %s: error %v", key, err)
		return err
	}
	return nil
}

func (c *cache) GetUsersFromCache(ctx context.Context, key string) (string, error) {
	rows, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logrus.Errorf("error getting users from cache key %s: error %v", key, err)
		return "", err
	}
	return rows, nil
}

func (c *cache) DeleteValue(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, key).Result()
	if err != nil {
		logrus.Errorf("error deleting cache %v", err)
		return err
	}
	return nil
}

func (c *cache) ClearCache(ctx context.Context) error {
	_, err := c.client.FlushAll(ctx).Result()
	if err != nil {
		logrus.Errorf("error clearing cache %v", err)
		return err
	}
	return nil
}
