package cache

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jafari-mohammad-reza/distributed-cache-system/broker"
	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeRule string

const (
	_      NodeRule = "NotRuled"
	Master          = "Master"
	Slave           = "Slave"
)

type Node struct {
	mu              sync.Mutex
	Rule            NodeRule
	Port            int
	storage         *Storage
	msgBroker       *broker.MsgBroker
	discoveredNodes map[int]NodeRule
}

func NewNode() *Node {
	port := rand.IntN(7999-7000) + 7000
	return &Node{
		Port:            port,
		storage:         NewStorage(),
		msgBroker:       broker.NewMsgBroker(6091),
		discoveredNodes: make(map[int]NodeRule),
	}
}
func (n *Node) InitCacheNode() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", n.Port))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", n.Port, err.Error())
	}
	rpcServer := grpc.NewServer()
	pb.RegisterCommandServer(rpcServer, NewCommandService(n))
	pb.RegisterNodeServer(rpcServer, NewNodeService(n))
	go n.registerNode()
	fmt.Printf("cache running on port: %d\n", n.Port)
	return rpcServer.Serve(ln)
}
func (n *Node) registerNode() {
	n.msgBroker.PublishMessage("service-discovery-register", fmt.Append(nil, n.Port))
	for msg := range n.msgBroker.SubscribeToTopic("service-discovery", 1) {
		var payload map[string]string
		err := json.Unmarshal(msg.Payload, &payload)
		if err != nil {
			continue
		}
		port, err := strconv.Atoi(payload["Port"])
		if err != nil {
			continue
		}
		rule := NodeRule(payload["Rule"])
		fmt.Println("received port", port, n.Port, payload["Rule"])

		n.mu.Lock()
		if port == n.Port {
			n.Rule = rule
		} else {
			n.discoveredNodes[port] = rule
		}
		fmt.Println("NODE", n.Port, n.Rule, n.discoveredNodes)
		go n.recoverLog()
		n.mu.Unlock()
	}
}
func (n *Node) recoverLog() {
	entries := n.loadLocalLog()
	for _, entry := range entries {
		n.execCommand(entry.Command, entry.Args)
	}

	if n.Rule == Master || len(n.discoveredNodes) == 0 {
		return
	}

	leaderPort := n.findLeader()
	if leaderPort == 0 {
		log.Println("No leader found for recovery")
		return
	}

	n.syncFromLeader(leaderPort, entries)
}
func (n *Node) loadLocalLog() []CommandLog {
	file, err := os.Open("aof.log")
	if err != nil {
		if os.IsNotExist(err) {
			f, _ := os.Create("aof.log")
			defer f.Close()
		}
		return nil
	}
	defer file.Close()

	var entries []CommandLog
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry CommandLog
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			continue
		}
		entries = append(entries, entry)
	}

	return entries
}
func (n *Node) findLeader() int {
	for port, rule := range n.discoveredNodes {
		if rule == Master {
			return port
		}
	}
	return 0
}
func (n *Node) syncFromLeader(leader int, localEntries []CommandLog) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", leader), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to leader node: %v", err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)

	var lastTimestamp string
	if len(localEntries) > 0 {
		lastTimestamp = localEntries[len(localEntries)-1].TimeStamp.Format(time.RFC3339Nano)
	}

	resp, err := client.GetLog(context.Background(), &pb.GetLogRequest{Start: lastTimestamp})
	if err != nil {
		log.Fatalf("Failed to sync logs from leader: %v", err)
	}

	if resp.Error != nil {
		log.Printf("Leader error: %s", resp.Error.Message)
		return
	}

	var remoteEntries []CommandLog
	if err := json.Unmarshal(resp.Data, &remoteEntries); err != nil {
		log.Printf("Failed to parse leader logs: %v", err)
		return
	}

	for _, entry := range remoteEntries {
		n.execCommand(entry.Command, entry.Args)
	}
}
func (n *Node) execCommand(command string, args []string) {
	switch command {
	case "SET":
		if len(args) < 3 {
			log.Println("Invalid SET args:", args)
			return
		}
		key := args[0]
		val := args[1]

		exp, err := strconv.Atoi(args[2])
		if err != nil {
			log.Println("Invalid TTL in SET:", args[2], err.Error())
			return
		}
		n.storage.Set(key, []byte(val), time.Duration(exp))
	case "DEL":
		if len(args) < 1 {
			log.Println("Invalid DEL args:", args)
			return
		}
		key := args[0]
		n.storage.Del(key)
	}
}
