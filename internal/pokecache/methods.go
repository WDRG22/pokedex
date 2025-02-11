package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go newCache.reaploop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.lock)
	defer c.mu.unlock()

	c.entries[key] := cacheEntry{
		createdAt: time.Now(),
		val: val
	}
}

func (c *Cache) Get(key string) []byte bool {
	c.mu.lock()
	defer c.mu.unlock()

	entry, ok := c.entries[key]
	if ok != nil {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.lock()
	defer c.mu.unlock()

	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
			delete(c.entries, key)
	}
}
