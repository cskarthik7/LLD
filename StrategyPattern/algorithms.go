package main
import "fmt"

type LRU struct {}

func (l *LRU) evict (c *Cache) {
	fmt.Println("Removing element from Cache using LRU")
}

type FIFO struct {}

func (f *FIFO) evict (c *Cache) {
	fmt.Println("Removing element from Cache using FIFO")
}

type LFU struct {}

func (l *LFU) evict (c *Cache) {
	fmt.Println("Removing element from Cache using LFU")
}