package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]cacheEntry
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data:  make(map[string]cacheEntry),
		mutex: sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	fmt.Println("adding new value to cache")
	c.data[key] = cacheEntry{time.Now().UTC(), value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("looking in cache for key: %v\n", key)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.data[key]
	return val.value, ok
}

//todo find error
//func (c *Cache) reapLoop(interval time.Duration) {
//	for range time.Tick(interval) {
//		fmt.Println("starting reap loop")
//		c.mutex.Lock()
//		defer c.mutex.Unlock()
//		for key, value := range c.data {
//			if time.Since(value.createdAt) > interval {
//				delete(c.data, key)
//				fmt.Printf("Deleting entry with key %v\n", key)
//			}
//		}
//	}
//}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			fmt.Printf("Deleting entry with key: %v \n", k)
			delete(c.data, k)
		}
	}
}
