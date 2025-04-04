package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/cache"
)

func main() {
	node := cache.NewNode()
	if err := node.InitCacheNode(); err != nil {
		log.Fatal(err.Error())
	}
}
