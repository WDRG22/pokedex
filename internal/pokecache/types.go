package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	mu sync.mutex 
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
