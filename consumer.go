package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func MessagesConsumer(m *nats.Msg) {
	strMessage := string(m.Data)
	log.Print(strMessage)

}
