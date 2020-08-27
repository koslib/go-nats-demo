package main

import (
	"log"
	"time"
)

func testSimpleNats() {
	// async model example
	if _, err := nc.QueueSubscribe("updates", "updates", MessagesConsumer); err != nil {
		log.Fatal("could not subscribe to queue")
	}

	QueueMessageProducer(nc)

	// sync model example
	sub, err := nc.SubscribeSync("sync_updates")
	if err != nil {
		log.Fatal(err)
	}
	if err = sub.SetPendingLimits(20000, 50*1024*1024); err != nil {
		log.Fatalf("unable to set nats limits with error: %s", err)
	}

	// create a channel so that it only blocks exiting - an alternative to an endless loop
	channel := make(chan struct{})

	RequestReplyMessageProducer(nc)

	// run this in a goroutine so that we don't block the main thread, and keep the consumer running
	go SyncMessageConsumer(sub)

	log.Print("sleeping and then trying the second batch")
	// simulate a pause in sync messages, and then fire back again
	time.Sleep(5 * time.Second)

	RequestReplyMessageProducer(nc)

	// block until the consumer is done
	<-channel
}
