package cache

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommandLog struct {
	Command   string
	Args      []string
	TimeStamp time.Time
}

type CommandService struct {
	pb.UnimplementedCommandServer
	node *Node
}

func NewCommandService(node *Node) *CommandService {
	return &CommandService{
		node: node,
	}
}

func (c *CommandService) Set(ctx context.Context, req *pb.SetCmdRequest) (*pb.SetCmdResponse, error) {
	err := c.node.storage.Set(req.Key, req.Value, time.Duration(req.Ttl)*time.Second)
	if err != nil {
		return &pb.SetCmdResponse{Error: &pb.Error{
			Message: err.Error(),
		}}, nil
	}
	AppendCommandToLog("SET", req.Key, string(req.Value), strconv.Itoa(int(req.Ttl)))
	args := []string{}
	args = append(args, req.Key)
	args = append(args, string(req.Value))
	args = append(args, strconv.Itoa(int(req.Ttl)))
	go c.sendLogToFollowers(CommandLog{Command: "SET", Args: args, TimeStamp: time.Now()})
	return &pb.SetCmdResponse{}, nil
}
func (c *CommandService) Get(ctx context.Context, req *pb.GetCmdRequest) (*pb.GetCmdResponse, error) {
	res, err := c.node.storage.Get(req.Key)
	if err != nil {
		return &pb.GetCmdResponse{Error: &pb.Error{
			Message: err.Error(),
		}}, nil
	}
	return &pb.GetCmdResponse{
		Value: res,
	}, nil
}
func (c *CommandService) Del(ctx context.Context, req *pb.DeleteCmdRequest) (*pb.DeleteCmdResponse, error) {
	err := c.node.storage.Del(req.Key)
	if err != nil {
		return &pb.DeleteCmdResponse{Error: &pb.Error{
			Message: err.Error(),
		}}, nil
	}
	AppendCommandToLog("DEL", req.Key)
	return &pb.DeleteCmdResponse{}, nil
}
func (c *CommandService) sendLogToFollowers(clog CommandLog) {
	for port, rule := range c.node.discoveredNodes {
		if rule != Master {
			conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatal(err.Error())
			}
			client := pb.NewNodeClient(conn)
			defer conn.Close()
			_, err = client.SendLog(context.Background(), &pb.SendLogRequest{Command: clog.Command, Args: clog.Args, Time: clog.TimeStamp.String()})
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

type NodeService struct {
	pb.UnimplementedNodeServer
	node *Node
}

func NewNodeService(node *Node) *NodeService {
	return &NodeService{
		node: node,
	}
}

func (n *NodeService) GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	var start, end time.Time
	if req.Start != "" {
		s, err := time.Parse(time.RFC3339, req.Start)
		if err != nil {
			return &pb.GetLogResponse{Error: &pb.Error{Message: err.Error()}}, nil
		}
		start = s
	}
	if req.End != "" {
		e, err := time.Parse(time.RFC3339, req.End)
		if err != nil {
			return &pb.GetLogResponse{Error: &pb.Error{Message: err.Error()}}, nil
		}
		end = e
	}
	file, err := os.Open("aof.log")
	if err != nil {
		return &pb.GetLogResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}
	defer file.Close()

	var matchedEntries []CommandLog

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		var entry CommandLog
		if err := json.Unmarshal(line, &entry); err != nil {
			continue
		}

		include := true

		if !start.IsZero() && entry.TimeStamp.Before(start) {
			include = false
		}
		if !end.IsZero() && entry.TimeStamp.After(end) {
			include = false
		}

		if include {
			matchedEntries = append(matchedEntries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return &pb.GetLogResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}

	data, err := json.Marshal(matchedEntries)
	if err != nil {
		return &pb.GetLogResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}

	return &pb.GetLogResponse{Data: data}, nil
}
func (n *NodeService) SendLog(ctx context.Context, req *pb.SendLogRequest) (*emptypb.Empty, error) {
	fmt.Println("log received")
	if err := AppendCommandToLog(req.Command, req.Args...); err != nil {
		log.Fatal("follower appending log", err.Error())
	}
	n.node.execCommand(req.Command, req.Args)
	return nil, nil
}
func AppendCommandToLog(cmd string, args ...string) error {
	entry := CommandLog{
		Command:   cmd,
		Args:      args,
		TimeStamp: time.Now().UTC(),
	}

	file, err := os.OpenFile("aof.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(entry); err != nil {
		return err
	}

	return nil
}
