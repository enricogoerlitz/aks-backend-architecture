package crud

import (
	"explore-aks-backend-app-air/src/cache"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CachingConfig struct {
	CacheDB    cache.IApplicationCache
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
