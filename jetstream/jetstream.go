package jetstream

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

const (
	streamName     = "ORDERS"
	streamSubjects = "ORDERS.*"
	subjectName    = "ORDERS.created"
)

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

func CreateStream(js nats.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateOrder publishes stream of events with subject
// "ORDERS.created"
func CreateOrder(js nats.JetStreamContext) error {
	var order Order
	for i := 1; i <= 10; i++ {
		order = Order{
			OrderID:    i,
			CustomerID: "Cust-" + strconv.Itoa(i),
			Status: "created",
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(subjectName, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("Order with OrderID:%d has been published\n", i)
	}
	return nil
}
