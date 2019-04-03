package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"time"

	cache "github.com/dcache/cache"
)

func main() {
	ct := cache.NewCacheTable(0)
	for i := 0; i < 100; i++ {
		go testAdd((ct))
	}
	//time.Sleep(3 * time.Second)
	//for i := 0; i < 1000; i++ {
	//	go testGet(ct)
	//}

	time.Sleep(5 * time.Second)
	for i := 0; i < 100; i++ {
		go testAdd((ct))
	}
	//fmt.Println(ct.Size())
	fmt.Scanln()
}

func testAdd(ct *cache.CacheTable) {

	for i := 0; i < 10000; i++ {
		d := rand.Int31n(1000000)
		//fmt.Println("adding item" + strconv.Itoa(int(d)))
		item := cache.NewCacheItem("key"+strconv.Itoa(int(d)), "value"+strconv.Itoa(int(d)))
		ct.Add(item)
	}
}

func testGet(ct *cache.CacheTable) {
	d := rand.Int31n(1000000)
	item := ct.Get("key" + strconv.Itoa(int(d)))
	if item != nil {
		fmt.Println("Get:", item.Key, item.Value, " found")

	} else {
		//fmt.Println("Get:", "key"+strconv.Itoa(int(d)), " not found")
	}

}
