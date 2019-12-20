package cache

import (
	"graze/pkg/cache"
)

var cacheHandler cache.Cache

func NewCacheHandler() {
	cacheHandler = cache.New()
}