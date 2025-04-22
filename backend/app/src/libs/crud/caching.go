package crud

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type IApplicationCache interface {
	Get(ctx *gin.Context, key string) (string, error)
	Set(ctx *gin.Context, key string, value string, expiration time.Duration) error
}

type CachingConfig struct {
	CacheDB    IApplicationCache
	Key        string
	Expiration time.Duration
	UseCache   bool
}

func IsHeaderUseCache(c *gin.Context, defaultValue bool) bool {
	header := c.Request.Header.Get("X-Use-Cache")
	if header == "" {
		return defaultValue
	}

	if strings.ToLower(header) == "true" {
		return true
	}

	return false
}
