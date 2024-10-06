package main
import "fmt"

type Cache struct {
	storage map[string]string
	evictionStrategy eviction
	capacity int
	max_capacity int
}

func initCache(e eviction) *Cache {
	fmt.Println("Using eviction startegy")
	storage := make(map[string]string)
	return &Cache{
		storage: storage,
		evictionStrategy: e,
		capacity: 0,
		max_capacity: 5,
	}
}

func (c *Cache) add (key, value string) {
	if c.capacity == c.max_capacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get (key string) string{
	value, exists := c.storage[key]
	if exists {
		return value
	}
	return ""
}

func (c *Cache) setEvictionStrategy(e eviction) {
	c.evictionStrategy = e
}

func (c *Cache) evict() {
	c.evictionStrategy.evict(c)
	c.capacity--
}

