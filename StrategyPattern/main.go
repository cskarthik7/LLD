package main
import "fmt"

func main(){
	lru_object := &LRU{}
	cache_object := initCache(lru_object)
	cache_object.add("a", "1")
	cache_object.add("b", "2")
	cache_object.add("c", "3")
	cache_object.add("d", "6")
	cache_object.add("e", "7")
	cache_object.add("e", "7")
	fmt.Println(cache_object)
}