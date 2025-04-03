package main

import (
	"fmt"
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/client"
)

func main() {
	go func() {
		err := client.InitGrpcServer()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	fmt.Println("Initializing grpc server")
	select {}
}
