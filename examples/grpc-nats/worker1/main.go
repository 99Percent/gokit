package main

import (
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"

	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
)

const (
	queue   = "Orders.OrdersCreatedQueue"
	subject = "Order.OrderCreated"
)

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	eventStore := pb.EventStore{}
	// Subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
