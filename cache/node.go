package cache

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net"
	"strconv"
	"sync"

	"github.com/jafari-mohammad-reza/distributed-cache-system/broker"
	pb "github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
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
	return rpcServer
}
func InitCacheNode() error {
	node := NewNode()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", node.Port))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", node.Port, err.Error())
	}
	go func() {
		node.msgBroker.PublishMessage("service-discovery-register", fmt.Append(nil, node.Port))
		for msg := range node.msgBroker.SubscribeToTopic("service-discovery", 1) {
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
			fmt.Println("received port", port, node.Port, payload["Rule"])

			node.mu.Lock()
			if port == node.Port {
				node.Rule = rule
			} else {
				node.discoveredNodes[port] = rule
			}
			fmt.Println("NODE", node.Port, node.Rule)
			node.mu.Unlock()
		}
	}()
	fmt.Printf("cache running on port: %d\n", node.Port)
	return node.RpcServer.Serve(ln)
}
