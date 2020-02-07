package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

// consumes a message - just displays the message data but that's fine for demonstration purposes
func MessagesConsumer(m *nats.Msg) {
	strMessage := string(m.Data)
	log.Printf("Async message data: %s", strMessage)
}

// consumes sync messages
func SyncMessageConsumer(sub *nats.Subscription) {
	// Run an endless loop as we need the consumer always open for processing incoming messages
	for {
		msg, err := sub.NextMsg(1 * time.Second)
		if err != nil {
			log.Printf("Error while consuming sync message: %s", err)
			if err.Error() == "timeout" {
				continue
			}
		}

		// if the message is nil, continue. Most probably caused by "no messages in the subscription" timeout
		if msg == nil {
			continue
		}

		// print the message data
		log.Printf("Reply data: %s", msg.Data)
		msg.Reply = "Got it and I'm replying back!"
		if err = msg.Respond([]byte("That's your byte data!")); err != nil {
			log.Printf("Error while responding a sync message: %s", err)
		}
	}
}
