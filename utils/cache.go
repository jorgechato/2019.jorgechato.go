package utils

import (
	"time"

	"github.com/allegro/bigcache"
)

var Cache *bigcache.BigCache

func init() {
	Cache, _ = bigcache.NewBigCache(
		bigcache.Config{
			Shards:             1024,
			LifeWindow:         CACHE_EXPIRED * time.Minute,
			CleanWindow:        CACHE_CLEAN * time.Second,
			MaxEntriesInWindow: 1000 * 10 * 60,
			MaxEntrySize:       500,
			Verbose:            true,
			HardMaxCacheSize:   8192,
			OnRemove:           nil,
			OnRemoveWithReason: nil,
		},
	)
}
