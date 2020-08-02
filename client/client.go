package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

//Client :nats client which enclose a connection
type Client struct {
	address string
	creds   string
	conn    *nats.Conn
}

//NewClient :nats client to send msg
func NewClient(address, creds string) (*Client, error) {
	client := &Client{
		address: address,
		creds:   creds,
	}
	if err := client.buildConnect(20); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) buildConnect(retryTimes int) (e error) {
	opts := []nats.Option{
		nats.Name("NATS Sample Publisher"),
		nats.Timeout(5 * time.Second),
		nats.FlusherTimeout(2 * time.Second),
		nats.DrainTimeout(2 * time.Second),
	}
	if c.creds != "" {
		opts = append(opts, nats.UserCredentials(c.creds))
	}
	for i := 0; i < retryTimes; i++ {
		nc, e := nats.Connect(c.address, opts...)
		if e != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		c.Close()
		c.conn = nc
		return nil
	}
	return fmt.Errorf(fmt.Sprintf("connect error,retry times:%d", retryTimes))
}

//Pub :pub msg.
//subj: subject with the format like a.b.c, a.*.c, a.b.>
//Publish publishes the data argument to the given subject. The data argument is left untouched and needs to be correctly interpreted on the receiver.
func (c *Client) Pub(subj string, payload []byte) (e error) {
	if c == nil {
		return err.Errno(err.EInvPointer)
	}
	if !c.conn.IsConnected() {
		return fmt.Errorf("nats not connected")
	}

	if e = c.conn.Publish(subj, payload); e != nil {
		return e
	}
	//TBD: Flush will block the main process, this need to be modified in the future.
	if e = c.conn.Flush(); e != nil {
		return e
	}
	return c.conn.LastError()
}

//Sub :sub msg.
//subj: subject with the format like a.b.c, a.*.c, a.b.>
func (c *Client) Sub(subj string, handler func(msg *nats.Msg)) (e error) {
	if c == nil {
		return err.Errno(err.EInvPointer)
	}
	if !c.conn.IsConnected() {
		return fmt.Errorf("nats not connected")
	}

	if _, e := c.conn.Subscribe(subj, handler); e != nil {
		return e
	}
	if e := c.conn.Flush(); e != nil {
		return e
	}

	return c.conn.LastError()
}

//Request :
//subj: subject with the format like a.b.c, a.*.c, a.b.>
//Request will send a request payload and deliver the response message, or an error, including a timeout if no message was received properly.
func (c *Client) Request(subj string, payload []byte) error {
	if c == nil {
		return err.Errno(err.EInvPointer)
	}
	if c.conn == nil {
		return err.Errno(err.EInvPointer)
	}

	_, e := c.conn.Request(subj, payload, 2*time.Second)
	if e != nil {
		return fmt.Errorf("%v for request", e)
	}
	return nil
}

//IsConnected : check the connection status
func (c *Client) IsConnected() bool {
	return c.conn.IsConnected()
}

//Close :close the connection before you quit
func (c *Client) Close() {
	if c != nil && c.conn != nil {
		c.conn.Close()
	}
}
