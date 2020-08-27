package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"strconv"
)

func testStreaming() {
	ncClientId, _ := nc.GetClientID()
	ncClientIdStr := strconv.FormatUint(ncClientId, 10)
	sc, err := stan.Connect("example-stan", ncClientIdStr, stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("established connection with nats streaming")

	defer sc.Close()

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	sub.Unsubscribe()
	sub.Close()

	// Simple Sync Publish
	sc.Publish("hello", []byte("one"))
	sc.Publish("hello", []byte("two"))
	sc.Publish("hello", []byte("three"))

	sc.Subscribe("hello", func(m *stan.Msg) {
		log.Printf("[Received] %+v", m)
	}, stan.DeliverAllAvailable())

	select {}
}
