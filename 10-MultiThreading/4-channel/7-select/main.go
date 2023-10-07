package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	//Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{i, "Hello from Kafka "}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("rceived from rabbitMq: %d - %s\n", msg1.id, msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("received from kafka: %d - %s\n", msg2.id, msg2.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
