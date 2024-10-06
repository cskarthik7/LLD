package main

type eviction interface {
	evict(c *Cache)
}