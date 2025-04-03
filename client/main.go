package client

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func InitGrpcServer() error {
	ln, err := net.Listen("tcp", ":6090")
	if err != nil {
		return fmt.Errorf("error in listening to port 6090: %s", err.Error())
	}
	rpcServer := grpc.NewServer()
	RegisterCommandServer(rpcServer, NewCommandService())

	return rpcServer.Serve(ln)
}
