package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"sync"
)

var nc *nats.Conn

func main() {
	natsServerAddr := getEnv("NATS_SERVER_ADDR", "127.0.0.1")

	var err error
	nc, err = nats.Connect(natsServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("established connection with nats server")

	defer nc.Close()

	if _, err = nc.QueueSubscribe("updates", "updates", MessagesConsumer); err != nil {
		log.Fatal("could not subscribe to queue")
	}
	MessagesProducer(nc)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func getEnv(name, defaultValue string) string{
	envValue := os.Getenv(name)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}
