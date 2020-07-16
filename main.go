package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

const (
	url  = "nats://127.0.0.1:4222"
	subj = "foo"
)

var (
	nc  *nats.Conn
	err error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err) {
		//
	}
}

//3 mode
//Publish-Subscribe
//Request-Reply
//Queue group
//Acknowledgements

func main() {

	//test pub sub
	//go TestPub()
	time.Sleep(time.Second)
	//go TestSub()

	//test request reply
	//nc, _ := nats.Connect(url)
	//defer nc.Close()

	nc.Subscribe("foo", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I will help you"))
	})
	reply, _ := nc.Request("foo", []byte("help"), 50*time.Millisecond)
	fmt.Println(string(reply.Data))

	//Queue
	received := 0
	nc.QueueSubscribe("foo", "worker_group", func(_ *nats.Msg) {
		received++
	})

	//acknowlage
	nc.Subscribe("foo", func(m *nats.Msg) {
		m.Respond([]byte(""))
	})

	reply, _ = nc.Request("foo", []byte("help"), 50*time.Millisecond)
	fmt.Println("ack:", string(reply.Data))
	fmt.Println("Main test ended!!")

}

func checkErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
