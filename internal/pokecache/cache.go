// Concerns api caching

package pokecache

import (
	"time"
	"sync"
)


type Cache struct {
        Entries map[string]cacheEntry
        sync.Mutex
}

type cacheEntry struct {
        createdAt time.Time
        val []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		Entries: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.Lock()
	defer c.Unlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.Unlock()
	}
}
