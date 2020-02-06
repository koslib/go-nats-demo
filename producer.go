package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func MessagesProducer(nc *nats.Conn) {
	for i := 1; i <= 20000; i++ {
		msg := []byte(time.Now().String())
		if err := nc.Publish("updates", msg); err != nil {
			log.Fatal(err)
		}
	}

}
