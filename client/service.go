package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClientService struct {
	CommandClient pb.CommandClient
}

func NewClientService() *ClientService {
	leader := getLeader()
	commandClient := getCommandClient(leader)
	return &ClientService{
		CommandClient: commandClient,
	}
}
func getLeader() int {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", 6090), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	client := pb.NewDiscoveryClient(conn)
	resp, err := client.GetLeader(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	fmt.Println("LEADER PORT", resp)
	return int(resp.Port)
}
func getCommandClient(leader int) pb.CommandClient {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", leader), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	client := pb.NewCommandClient(conn)
	return client
}
func (c *ClientService) Set(args []string) error {
	ttl, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	resp, err := c.CommandClient.Set(context.Background(), &pb.SetCmdRequest{Key: args[0], Value: []byte(args[1]), Ttl: int32(ttl)})
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New(resp.Error.Message)
	}
	return nil
}
func (c *ClientService) Get(args []string) (string, error) {
	resp, err := c.CommandClient.Get(context.Background(), &pb.GetCmdRequest{Key: args[0]})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}
	return string(resp.Value), nil
}
func (c *ClientService) Del(args []string) error {
	resp, err := c.CommandClient.Del(context.Background(), &pb.DeleteCmdRequest{Key: args[0]})
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New(resp.Error.Message)
	}
	return nil
}
