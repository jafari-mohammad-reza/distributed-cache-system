package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/cache"
)

func main() {

	if err := cache.InitCacheNode(); err != nil {
		log.Fatal(err.Error())
	}
}
