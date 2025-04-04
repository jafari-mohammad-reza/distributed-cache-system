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
	Queue       map[string]*Message // TODO: implement queue
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
	Port   int
	conn   *grpc.ClientConn
	client BrokerClient
}

func NewMsgBroker(port int) *MsgBroker {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	client := NewBrokerClient(conn)
	return &MsgBroker{
		Port:   port,
		conn:   conn,
		client: client,
	}
}

func (mb *MsgBroker) PublishMessage(topic string, data []byte) {
	_, err := mb.client.Pub(context.Background(), &PubRequest{
		Topic:   topic,
		Payload: data,
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Println("Message Published to topic:", topic)
}
func (mb *MsgBroker) SubscribeToTopic(topic string, limit int) <-chan *Message {
	receiver := make(chan *Message, limit)

	go func() {
		stream, err := mb.client.Sub(context.Background(), &SubRequest{Topic: topic})
		if err != nil {
			log.Printf("Failed to subscribe: %v", err)
			close(receiver)
			return
		}

		fmt.Println("Listening for messages on topic:", topic)

		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Printf("Stream closed or error: %v", err)
				close(receiver)
				return
			}

			fmt.Println("Received Message:", string(msg.Payload))

			select {
			case receiver <- msg:
			default:
				log.Println("Receiver channel full, dropping message")
			}
		}
	}()

	return receiver
}
