package cache

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"time"

	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"
)

type CommandLog struct {
	Command   string
	Args      []interface{}
	TimeStamp time.Time
}

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

type NodeService struct {
	pb.UnimplementedNodeServer
}

func NewNodeService() *NodeService {
	return &NodeService{}
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
func AppendCommandToLog(cmd string, args ...interface{}) error {
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
