package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/cache"
)

func main() {
	ds := cache.NewDiscoveryService(6090)
	if err := ds.InitDiscoveryService(); err != nil {
		log.Fatal(err.Error())
	}
}
