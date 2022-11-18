package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func jetStreamPublish() {
	userCredsPath := "user.creds"
	natsURL := ""
	nc, err := nats.Connect(natsURL, nats.RootCAs("ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Create JetStream Context and set the maximum number
	// of inflight (at one time) async publishes.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	// Stream Publisher
	_, err = js.Publish("FORCE.hope", []byte("Never tell me the odds"))
	if err != nil {
		log.Fatal(err)
	}
}

func jetStreamCreate() {
	userCredsPath := "user.creds"
	natsURL := ""
	nc, err := nats.Connect(natsURL, nats.RootCAs("ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Create JetStream Context and set the maximum number
	// of inflight (at one time) async publishes.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	// Create a Stream
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "FORCE",
		Subjects: []string{"FORCE.*"},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func jetStreamSubscribe() {
	userCredsPath := "user.creds"
	natsURL := "nats://localhost:4222"
	nc, err := nats.Connect(natsURL, nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Create JetStream Context and set the maximum number
	// of inflight (at one time) async publishes.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	_, err = js.AddConsumer("stream1", &nats.ConsumerConfig{
		Durable:       "go-consumer",
		AckPolicy:     nats.AckExplicitPolicy,
		Replicas:      1,
	})
	if err != nil {
		log.Fatalf("failed to create consumer: %s\n", err)
	}

	// Sync ephemeral
	// startSeq := 1011
	// sub1, err := js.SubscribeSync("stream1.*", nats.StartSequence(uint64(startSeq)), nats.MaxDeliver(3))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 0; i < 5; i++ {
	// 	m, err := sub1.NextMsg(time.Second * 3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if err = m.Ack(); err == nil {
	// 		startSeq++
	// 	}
	// 	log.Println("Received message: ", string(m.Data))
	// }

	// sub1.Drain()

	// log.Println(startSeq)
	// sub2, err := js.SubscribeSync("stream1.*", nats.StartSequence(uint64(startSeq)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 0; i < 5; i++ {
	// 	m, err := sub2.NextMsg(time.Second * 3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if err = m.Ack(); err == nil {
	// 		startSeq++
	// 	}
	// 	log.Println("Received message: ", string(m.Data))
	// }
}

func jetStreamAsyncSubscribe() {
	userCredsPath := "user.creds"
	natsURL := ""
	nc, err := nats.Connect(natsURL, nats.RootCAs("ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Create JetStream Context and set the maximum number
	// of inflight (at one time) async publishes.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	// Async ephemeral consumer
	_, err = js.Subscribe("FORCE.*", func(m *nats.Msg) {
		log.Printf("Received Message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}
