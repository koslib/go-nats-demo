package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strings"
	"time"
)

var nc *nats.Conn

func main() {
	natsServerAddr := getEnv("NATS_SERVER_ADDR", "127.0.0.1")

	// Although I'm using a single connection str, leaving this here just in case, for future-proofing
	natsClusterAddresses := []string{natsServerAddr}

	var err error
	nc, err = nats.Connect(strings.Join(natsClusterAddresses, ","), nats.Timeout(15*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("established connection with nats cluster")

	defer nc.Close()

	// Play around with each one of the following by commenting in/out
	testSimpleNats()
	//testStreaming()

}

func getEnv(name, defaultValue string) string {
	envValue := os.Getenv(name)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}
