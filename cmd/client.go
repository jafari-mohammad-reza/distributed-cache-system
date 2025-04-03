package main

import (
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/client"
)

func main() {
	if err := client.InitClient(); err != nil {
		log.Fatal(err.Error())
	}
}
