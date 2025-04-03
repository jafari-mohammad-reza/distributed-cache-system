package client

import (
	"fmt"
	"net"

	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"

	"google.golang.org/grpc"
)

func InitGrpcServer() error {
	ln, err := net.Listen("tcp", ":6090")
	if err != nil {
		return fmt.Errorf("error in listening to port 6090: %s", err.Error())
	}
	rpcServer := grpc.NewServer()
	pb.RegisterCommandServer(rpcServer, NewCommandService())

	fmt.Println("client running on port 6090")
	return rpcServer.Serve(ln)
}
