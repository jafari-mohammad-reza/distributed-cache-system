package broker

import (
	"fmt"
	"testing"
	"time"
)

func TestBrokerService(t *testing.T) {

	go InitBroker(8091)
	msgBk := NewMsgBroker(8091)
	go func() {
		for msg := range msgBk.SubscribeToTopic("test", 3) {
			fmt.Println(msg.Payload)
		}
	}()
	time.Sleep(time.Second)
	msgBk.PublishMessage("test", []byte("test-msg"))
	msgBk.PublishMessage("test", []byte("test-msg-2"))
	msgBk.PublishMessage("test", []byte("test-msg-3"))
	// for msg := range receiver {
	// 	if len(receiver) == 2 {
	// 		close(receiver)
	// 		break
	// 	}
	// 	fmt.Printf("receiver msg %s\n", string(msg.Payload))
	// }
}
