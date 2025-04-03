package client

import (
	"context"

	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"
)

type CommandService struct {
	pb.UnimplementedCommandServer
}

func NewCommandService() *CommandService {
	return &CommandService{}
}

func (c *CommandService) Set(ctx context.Context, req *pb.SetCmdRequest) (*pb.SetCmdResponse, error) {
	return nil, nil
}
func (c *CommandService) Get(ctx context.Context, req *pb.GetCmdRequest) (*pb.GetCmdResponse, error) {
	return nil, nil
}
func (c *CommandService) Del(ctx context.Context, req *pb.DeleteCmdRequest) (*pb.DeleteCmdResponse, error) {
	return nil, nil
}
