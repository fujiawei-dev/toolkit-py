{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

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

func SetCache(key, value interface{}, ttl time.Duration) {
	log.Printf("query: cached %s", key)

	cache.SetWithTTL(key, value, 1, ttl)
}

func GetCache(key interface{}) (interface{}, bool) {
	value, exist := cache.Get(key)

	if exist {
		log.Printf("query: hit %s", key)
	} else {
		log.Printf("query: not hit %s", key)
	}

	return value, exist
}
