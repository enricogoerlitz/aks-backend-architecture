package cache

import (
	"strings"
)

type CacheKeyAttribute struct {
	Attr  string
	Value string
}

func GenerateCacheKey(parts ...CacheKeyAttribute) string {
	cacheKeyParts := []string{}

	for _, part := range parts {
		cacheKeyParts = append(cacheKeyParts, part.Attr+":"+part.Value)
	}

	return strings.Join(cacheKeyParts, ",")
}
