package main

import (
	"fmt"

	"github.com/nats-io/go-nats"
)

//TestSub : test function for sub
func TestSub() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}
