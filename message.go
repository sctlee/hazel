package tcpx

import (
	"fmt"
)

type IMessage interface {
	Get() map[string]string // return map[string]string
	GetClient() *Client
	SetBoardClients(mc []*Client)
	exec()
	send()
	bcast()
}

type Message struct {
	rawData     string
	client      *Client
	multiClient []*Client
}

func NewMessage(c *Client, data interface{}) IMessage {
	var d string
	switch data.(type) {
	case string:
		d = data.(string)
	case map[string]string:
		d = pt.UnMarshal(data.(map[string]string))
	default:
		d = "Error: Can't parse the type of message!"
	}
	return &Message{
		rawData: d,
		client:  c,
	}
}

func NewBoardMessage(c *Client, data interface{}, mc []*Client) IMessage {
	msg := NewMessage(c, data)
	msg.SetBoardClients(mc)
	fmt.Println(msg)
	return msg
}

func (self *Message) Get() map[string]string {
	return pt.Marshal(self.rawData)
}

func (self *Message) GetClient() *Client {
	return self.client
}

func (self *Message) SetBoardClients(mc []*Client) {
	self.multiClient = make([]*Client, 0)
	self.multiClient = append(self.multiClient, mc[:]...)
}

func (self *Message) exec() {
	if self.multiClient != nil || len(self.multiClient) != 0 {
		self.bcast()
	} else {
		self.send()
	}
}
func (self *Message) send() {
	self.client.PutOutgoing(self.rawData)
}
func (self *Message) bcast() {
	for _, client := range self.multiClient {
		go client.PutOutgoing(self.rawData)
	}
}
