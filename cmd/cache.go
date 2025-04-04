package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/cache"
)

func main() {
	go func() {
		if err := cache.InitCacheGateway(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	go func() {
		if err := cache.InitCache(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	select {}
}
