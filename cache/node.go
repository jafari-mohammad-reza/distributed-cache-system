package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"strconv"
	"sync"

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
	RpcServer       *grpc.Server
	msgBroker       *broker.MsgBroker
	discoveredNodes map[int]NodeRule
}

func NewNode() *Node {
	port := rand.IntN(7999-7000) + 7000
	return &Node{
		Port:            port,
		RpcServer:       newRpcServer(port),
		storage:         NewStorage(),
		msgBroker:       broker.NewMsgBroker(6091),
		discoveredNodes: make(map[int]NodeRule),
	}
}
func newRpcServer(port int) *grpc.Server {
	rpcServer := grpc.NewServer()
	pb.RegisterCommandServer(rpcServer, NewCommandService())
	pb.RegisterNodeServer(rpcServer, NewNodeService())
	return rpcServer
}
func (n *Node) InitCacheNode() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", n.Port))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", n.Port, err.Error())
	}
	go n.registerNode()
	fmt.Printf("cache running on port: %d\n", n.Port)
	return n.RpcServer.Serve(ln)
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
	if len(n.discoveredNodes) == 0 || n.Rule == Master {
		return
	}
	var leader int
	for port, rule := range n.discoveredNodes {
		if rule == Master {
			leader = port
			break
		}
	}
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", leader), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to another node: %v", err)
	}

	client := pb.NewNodeClient(conn)
	resp, err := client.GetLog(context.Background(), &pb.GetLogRequest{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("respData", string(resp.Data), resp.Error)
}
