package cache

import (
	"time"

	"github.com/gin-gonic/gin"
)

var CacheFactory = ApplicationCacheFactory{}

type ApplicationCacheFactory struct{}

func (c *ApplicationCacheFactory) Create(cacheType string, config CacheConfig) IApplicationCache {
	switch cacheType {
	case "REDIS_SINGLE_NODE":
		return NewApplicationRedisCache(config)
	case "REDIS_CLUSTER":
		return NewApplicationRedisClusterCache(config)
	default:
		return nil
	}
}

type IApplicationCache interface {
	Get(ctx *gin.Context, key string) (string, error)
	Set(ctx *gin.Context, key string, value string, expiration time.Duration) error
}

type CacheConfig struct {
	Hosts    []string
	Port     string
	Username string
	Password string
}
