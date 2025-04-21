package cache

import (
	"explore-aks-backend-app-air/constants"
	"strings"
)

var Cache IApplicationCache

func ConnectCache() {
	cacheConfig := CacheConfig{
		Hosts:    strings.Split(constants.ENV_CACHE_HOSTS, ","),
		Port:     constants.ENV_CACHE_PORT,
		Username: constants.ENV_CACHE_USER,
		Password: constants.ENV_CACHE_PASSWORD,
	}

	Cache = CacheFactory.Create(constants.ENV_CACHE_TYPE, cacheConfig)

	if Cache == nil {
		panic("failed to initialize cache: unsupported cache type")
	}
}
