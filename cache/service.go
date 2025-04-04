package cache

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net"

	"github.com/jafari-mohammad-reza/distributed-cache-system/broker"
	"github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
)

var msgBroker *broker.MsgBroker

func init() {
	msgBroker = broker.NewMsgBroker(6091)
}
func InitCache() error {
	port := rand.IntN(7999-7000) + 7000

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", port, err.Error())
	}
	msg, _ := json.Marshal(map[string]string{"Port": fmt.Sprint(port)})
	msgBroker.PublishMessage("leader", msg)
	rpcServer := grpc.NewServer()
	pb.RegisterCommandServer(rpcServer, NewCommandService())

	fmt.Printf("cache running on port: %d\n", port)
	return rpcServer.Serve(ln)
}
func InitCacheGateway() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 6090))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", 6090, err.Error())
	}
	rpcServer := grpc.NewServer()
	pb.RegisterCommandServer(rpcServer, NewCommandService())
	go func() {
		for msg := range msgBroker.SubscribeToTopic("leader", 1) {
			fmt.Println("MSG", msg.Payload)
		}
	}()

	fmt.Println("gateway running on port 6090")
	return rpcServer.Serve(ln)
}
