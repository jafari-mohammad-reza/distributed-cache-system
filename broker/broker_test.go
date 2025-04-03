package broker

import (
	"fmt"
	"testing"
	"time"
)

func TestBrokerService(t *testing.T) {

	go InitBroker(8091)
	msgBk := NewMsgBroker(8091)
	receiver := make(chan *Message, 3)
	go msgBk.SubscribeToTopic("test", receiver)
	time.Sleep(time.Second)
	msgBk.PublishMessage("test", "test-msg")
	msgBk.PublishMessage("test", "test-msg-2")
	msgBk.PublishMessage("test", "test-msg-3")
	for msg := range receiver {
		if len(receiver) == 2 {
			close(receiver)
			break
		}
		fmt.Printf("receiver msg %s\n", string(msg.Payload))
	}
}
