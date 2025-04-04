package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/jafari-mohammad-reza/distributed-cache-system/broker"
	"github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DiscoveryService struct {
	pb.UnimplementedDiscoveryServer
	mu              sync.Mutex
	Port            int
	LeaderNode      int // leader node port
	DiscoveredNodes []int
	msgBroker       *broker.MsgBroker
}

func NewDiscoveryService(port int) *DiscoveryService {
	return &DiscoveryService{
		Port:            port,
		LeaderNode:      0,
		DiscoveredNodes: []int{},
		msgBroker:       broker.NewMsgBroker(6091),
	}
}
func (ds *DiscoveryService) GetLeader(ctx context.Context, req *emptypb.Empty) (*pb.DiscoveryResponse, error) {
	return &pb.DiscoveryResponse{
		Port: int32(ds.LeaderNode),
	}, nil
}

func (ds *DiscoveryService) InitDiscoveryService() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", ds.Port))
	if err != nil {
		return fmt.Errorf("error in listening to port %d: %s", ds.Port, err.Error())
	}
	rpcServer := grpc.NewServer()
	pb.RegisterDiscoveryServer(rpcServer, ds)
	fmt.Printf("discovery running on port %d \n", ds.Port)
	go ds.healthCheck()
	go ds.discover()
	return rpcServer.Serve(ln)
}

func (ds *DiscoveryService) healthCheck() {}
func (ds *DiscoveryService) discover() {
	for msg := range ds.msgBroker.SubscribeToTopic("service-discovery-register", 2) {
		fmt.Println("discover register msg", string(msg.Payload))
		port, err := strconv.Atoi(string(msg.Payload))
		if err != nil {
			continue
		}
		ds.mu.Lock()
		if ds.LeaderNode == 0 {
			ds.LeaderNode = port
			payload, _ := json.Marshal(map[string]string{"Port": strconv.Itoa(port), "Rule": string(Master)})
			ds.msgBroker.PublishMessage("service-discovery", payload)
			fmt.Println("new leader assigned")
		} else {
			payload, _ := json.Marshal(map[string]string{"Port": strconv.Itoa(port), "Rule": string(Slave)})
			ds.msgBroker.PublishMessage("service-discovery", payload)
			ds.DiscoveredNodes = append(ds.DiscoveredNodes, port)
		}
		ds.mu.Unlock()

	}

}
