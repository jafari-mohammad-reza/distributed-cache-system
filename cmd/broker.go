package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/broker"
)

func main() {
	go func() {
		err := broker.InitBroker(6091)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	select {}
}
