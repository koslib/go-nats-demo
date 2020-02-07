package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

// generates a ton of dummy async messages to be consumed by a queue
func QueueMessageProducer(nc *nats.Conn) {
	for i := 1; i <= 20000; i++ {
		msg := []byte(time.Now().String())
		if err := nc.Publish("updates", msg); err != nil {
			log.Fatal(err)
		}
	}
}

// generates a ton of dummy messages that follow a sync model (request/reply)
func RequestReplyMessageProducer(nc *nats.Conn) {
	for i := 1; i <= 500; i++ {
		msg := []byte(time.Now().String())
		if err := nc.Publish("sync_updates", msg); err != nil {
			log.Fatal(err)
		}
	}
}
