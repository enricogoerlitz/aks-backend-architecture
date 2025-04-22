package cache

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type IApplicationRedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Ping(ctx context.Context) *redis.StatusCmd
}

type ApplicationRedisCache struct {
	client IApplicationRedisClient
}

func (c *ApplicationRedisCache) Get(ctx *gin.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *ApplicationRedisCache) Set(ctx *gin.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, 0).Err()
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

	return &ApplicationRedisCache{
		client: client,
	}
}

func NewApplicationRedisCache(config CacheConfig) IApplicationCache {
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
