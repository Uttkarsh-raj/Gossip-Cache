package network

import (
	"errors"
	"sync"
	"time"
)

type CacheItem struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	TTL   int64       `json:"ttl"`
}

func NewCacheItem(key string, value interface{}, duration int64) *CacheItem {
	return &CacheItem{
		Key:   key,
		Value: value,
		TTL:   duration,
	}
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

func (c *Cache) Add(key string, value interface{}, duration int64) (string, CacheItem) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	expiration := time.Now().Add(time.Duration(duration) * time.Second).UnixNano()
	item, present := c.Items[key]
	if present {
		item.TTL = expiration
		item.Value = value
		return "Updated the value!!", item
	}
	c.Items[key] = *NewCacheItem(key, value, expiration)
	return "New value successfully added", c.Items[key]
}

func (c *Cache) Get(key string) (CacheItem, bool, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	item, isPresent := c.Items[key]
	currentTime := time.Now().UnixNano()
	if !isPresent {
		return CacheItem{}, false, errors.New("error: Key does not exist")
	}
	if item.TTL < currentTime {
		return CacheItem{}, false, errors.New("error: Key expired. TTL crossed")
	}
	return item, true, nil
}

func (c *Cache) Update(key string, updatedItem CacheItem) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Items[key] = updatedItem
}

func (c *Cache) Delete(key string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	delete(c.Items, key)
}
