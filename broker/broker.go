package broker

import (
	"context"
	"fmt"
	"log"
	"net"
	sync "sync"

	"slices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BrokerService struct {
	UnimplementedBrokerServer
	broker *Broker
}

type Subscriber struct {
	Channel     chan *Message
	Unsubscribe chan bool
}

type Broker struct {
	subscribers map[string][]*Subscriber
	mutex       sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*Subscriber),
	}
}

func (b *BrokerService) Pub(ctx context.Context, req *PubRequest) (*Error, error) {
	b.broker.mutex.Lock()
	defer b.broker.mutex.Unlock()

	if subs, exists := b.broker.subscribers[req.Topic]; exists {
		for _, sub := range subs {
			select {
			case sub.Channel <- &Message{Payload: req.Payload}:
			default:
				fmt.Println("Dropped message (slow subscriber)")
			}
		}
		return &Error{Message: ""}, nil
	}
	return &Error{Message: "No subscribers found"}, nil
}

func (b *BrokerService) Sub(req *SubRequest, stream Broker_SubServer) error {
	sub := &Subscriber{
		Channel:     make(chan *Message, 10), // TODO make size dynamic
		Unsubscribe: make(chan bool),
	}

	b.broker.mutex.Lock()
	b.broker.subscribers[req.Topic] = append(b.broker.subscribers[req.Topic], sub)
	b.broker.mutex.Unlock()

	for {
		select {
		case msg := <-sub.Channel:
			if err := stream.Send(msg); err != nil {
				b.UnSub(req.Topic, sub)
				return err
			}
		case <-sub.Unsubscribe:
			b.UnSub(req.Topic, sub)
			return nil
		}
	}
}

func (b *BrokerService) UnSub(topic string, sub *Subscriber) {
	b.broker.mutex.Lock()
	defer b.broker.mutex.Unlock()

	subs := b.broker.subscribers[topic]
	for i, s := range subs {
		if s == sub {
			b.broker.subscribers[topic] = slices.Delete(subs, i, i+1)
			close(s.Channel)
			break
		}
	}
}

func InitBroker(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	rpcServer := grpc.NewServer()
	broker := &BrokerService{broker: NewBroker()}

	RegisterBrokerServer(rpcServer, broker)
	fmt.Println("Broker running on port:", port)
	return rpcServer.Serve(ln)
}

type MsgBroker struct {
	Port int
}

func NewMsgBroker(port int) *MsgBroker {
	return &MsgBroker{
		Port: port,
	}
}

func (mb *MsgBroker) PublishMessage(topic, message string) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", mb.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}
	defer conn.Close()

	client := NewBrokerClient(conn)

	_, err = client.Pub(context.Background(), &PubRequest{
		Topic:   topic,
		Payload: []byte(message),
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Println("Message Published to topic:", topic)
}
func (mb *MsgBroker) SubscribeToTopic(topic string, receiver chan *Message) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", mb.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewBrokerClient(conn)

	stream, err := client.Sub(context.Background(), &SubRequest{Topic: topic})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	fmt.Println("Listening for messages on topic:", topic)
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Println("Received Message:", string(msg.Payload))
		receiver <- msg
	}
}
