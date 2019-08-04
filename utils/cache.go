package utils

import (
	"time"

	"github.com/allegro/bigcache"
)

var Cache *bigcache.BigCache

func init() {
	Cache, _ = bigcache.NewBigCache(
		bigcache.DefaultConfig(CACHE_EXPIRED * time.Minute),
	)
}
