package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Key   string
	Value interface{}
	TTL   int64
}

type Cache struct {
	Items map[string]CacheItem
	Mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		Items: make(map[string]CacheItem),
	}
}

func (c *Cache) Add(key string, value interface{}, duration time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	expiration := time.Now().Add(duration).UnixNano()
	c.Items[key] = CacheItem{
		Key:   key,
		Value: value,
		TTL:   expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	item, isPresent := c.Items[key]
	if !isPresent || item.TTL < time.Now().UnixNano() {
		return nil, false
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	delete(c.Items, key)
}
