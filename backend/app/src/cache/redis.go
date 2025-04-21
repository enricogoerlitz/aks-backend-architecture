package cache

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ApplicationRedisClusterCache struct {
	client *redis.ClusterClient
}

type ApplicationRedisCache struct {
	client *redis.Client
}

func NewApplicationRedisClusterCache(config CacheConfig) IApplicationCache {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    config.Hosts,
		Username: config.Username,
		Password: config.Password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("failed to connect to Redis cluster: " + err.Error())
	}

	return &ApplicationRedisClusterCache{
		client: client,
	}
}

func (c *ApplicationRedisClusterCache) Get(ctx *gin.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *ApplicationRedisClusterCache) Set(ctx *gin.Context, key string, value string) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func NewApplicationRedisCache(config CacheConfig) *ApplicationRedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Hosts[0],
		Username: config.Username,
		Password: config.Password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("failed to connect to Redis sigle-node: " + err.Error())
	}

	return &ApplicationRedisCache{
		client: client,
	}
}

func (c *ApplicationRedisCache) Get(ctx *gin.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *ApplicationRedisCache) Set(ctx *gin.Context, key string, value string) error {
	return c.client.Set(ctx, key, value, 0).Err()
}
