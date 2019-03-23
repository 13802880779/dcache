package main

import (
	//	"fmt"
	"strconv"
	"time"

	cache "github.com/dcache/cache"
	rest "github.com/dcache/http/rest"
)

func main() {

	//println("hereeeee")
	ct := cache.NewCacheTable(100)

	for i := 0; i < 100; i++ {
		ci := cache.NewCacheItem("key"+strconv.Itoa(i),
			"value"+strconv.Itoa(i),
			time.Duration(i)*time.Second,
			cache.RESET_TTL_ON_ACCESS)
		ct.Add(ci)
	}

	restapiserver := rest.NewRestApiServer("8081", ct)
	restapiserver.Start()
	//	fmt.Scanln()

}
