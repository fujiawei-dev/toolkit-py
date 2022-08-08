package query

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

var cache *ristretto.Cache

func init() {
	cache, _ = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
}

func SetCache(key, value any, ttl time.Duration) {
	log.Printf("query: cached %s", key)

	cache.SetWithTTL(key, value, 0, ttl)
}

func GetCache(key any) (any, bool) {
	value, exist := cache.Get(key)

	if exist {
		log.Printf("query: hit %s", key)
	} else {
		log.Printf("query: not hit %s", key)
	}

	return value, exist
}
