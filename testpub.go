package main

import "github.com/nats-io/go-nats"

//TestPub : test func for pub msg
func TestPub() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Publish("foo", []byte("Hello World!"))
}
