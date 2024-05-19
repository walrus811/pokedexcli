package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]CacheEntry
	mu    *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newValue := Cache{
		store: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}
	go newValue.reapLoop(interval)
	return newValue
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.store[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		for k, v := range c.store {
			if time.Since(v.createdAt) > interval {
				c.mu.Lock()
				delete(c.store, k)
				c.mu.Unlock()
			}
		}
	}
}
