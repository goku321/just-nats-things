package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	userCredsPath := "/Users/deepak.sah/.nkeys/creds/server1-op/server1-admin/admin-user.creds"
	nc, err := nats.Connect("tls://localhost:443", nats.RootCAs("./cert/ca.pem"), nats.UserCredentials(userCredsPath))
	if err != nil {
		log.Fatal(err)
	}

	err = nc.Publish("msg.test", []byte("hello there"))
	if err != nil {
		log.Fatalf("failed to publish: %s", err)
	}

	nc.Close()
}
