package client

import (
	"log"
	"net"

	"github.com/sctlee/tcpx/sharedpreferences"
	"github.com/sctlee/tcpx/tcpx/base"
)

const (
	CLIENT_STATE_OPEN  = 1
	CLIENT_STATE_CLOSE = 2
)

type Xtime struct {
	isExist  bool
	question string
}

type OnCloseListener interface {
	OnClose(client *Client)
}
type IClient interface {
	TRead(incoming chan string) error
	TWrite(outgoing chan string) error
	Close()
}

type Client struct {
	c IClient

	Cid string

	incoming chan string
	outgoing chan string

	State int

	// extend features
	onCloseFuncs      []OnCloseListener
	sharedPreferences map[string]sharedpreferences.SharedPreferences
}

func CreateClient(conn net.Conn, cid string) (client *Client) {
	client = &Client{
		c:                 base.NewTCPClient(conn),
		incoming:          make(chan string),
		outgoing:          make(chan string),
		Cid:               cid,
		State:             CLIENT_STATE_OPEN,
		onCloseFuncs:      make([]OnCloseListener, 0),
		sharedPreferences: make(map[string]sharedpreferences.SharedPreferences),
	}

	go client.Read()
	go client.Write()

	return
}

func (self *Client) GetIncoming() (msg string, ok bool) {
	msg, ok = <-self.incoming
	return
}

func (self *Client) PutOutgoing(str string) {
	if self.State == CLIENT_STATE_OPEN {
		self.outgoing <- str
	}
}

func (self *Client) Read() {
	err := self.c.TRead(self.incoming)
	if err != nil {
		// logger.Printf("Read error %s\n", err)
		log.Printf("Read error %s\n", err)
		self.Close()
	}
}

func (self *Client) Write() {
	err := self.c.TWrite(self.outgoing)
	if err != nil {
		// logger.Printf("Write error %s\n", err)
		log.Printf("Write error %s\n", err)
	} else {
		log.Println("client writer closed")
		// logger.Println("client writer closed")
	}
}

func (self *Client) Close() {
	//trigger delegation event
	for _, f := range self.onCloseFuncs {
		f.OnClose(self)
	}

	self.c.Close()
	self.State = CLIENT_STATE_CLOSE
	// close mean to notify a receiver not to expect any more values to be sent.
	// but in a feature, it doesn't know the conn's stat, so it doesn't know if
	// the channel is useless, so it can't close the channel, so don't close it
	// here(it's not a producer)
	close(self.incoming)
	close(self.outgoing)

	// logger.Println("Client close")
	log.Println("Client close")
}

func (self *Client) SetOnCloseListener(onCloseListener OnCloseListener) {
	self.onCloseFuncs = append(self.onCloseFuncs, onCloseListener)
}

func (self *Client) GetSharedPreferences(key string) (sp sharedpreferences.SharedPreferences) {
	if sp, ok := self.sharedPreferences[key]; ok {
		return sp
	}
	sp = sharedpreferences.NewSharePreferences("map")
	self.sharedPreferences[key] = sp
	return sp
}
