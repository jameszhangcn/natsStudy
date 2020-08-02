package main

import (
	"fmt"
        "time"
	"github.com/nats-io/nats.go"
)
var NatsServerUrl = "mynats:4222"
//TestSub : test function for sub
func testSub() {
	nc, err := nats.Connect(NatsServerUrl)
	defer func() {
		fmt.Println("Sub nats connect err: ", err)
		nc.Close()
	}()
        for {
	    time.Sleep(time.Second)
	    nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	    })
         }
}

func testPub() {
	nc, err := nats.Connect(NatsServerUrl)
	defer func() {
		fmt.Println("Pub nats connect err: ", err)
		nc.Close()
	}()
	for {
            time.Sleep(time.Second)
	    nc.Publish("foo", []byte("Hello World!"))
         }

}

func main(){
	go testPub()
	testSub()
}
