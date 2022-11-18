package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func publish() {
	userCredsPath := "user.creds"
	natsURL := ""
	nc, err := nats.Connect(natsURL, nats.RootCAs("ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Simple publisher
	if err = nc.Publish("test", []byte("This is the way")); err != nil {
		log.Fatalf("publish error: %s\n", err)
	}

	// // Make sure it makes a round trip to the server before exiting
	nc.Flush()
}

func subscribe() {
	userCredsPath := "user.creds"
	natsURL := ""
	nc, err := nats.Connect(natsURL, nats.RootCAs("ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Sync subscriber
	sub, err := nc.SubscribeSync("test")
	if err != nil {
		log.Fatal(err)
	}

	m, err := sub.NextMsg(time.Second * 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received message: %s\n", string(m.Data))
}
