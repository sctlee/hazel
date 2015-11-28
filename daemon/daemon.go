package daemon

import (
	"fmt"
	"log"

	"github.com/sctlee/hazel/daemon/message"
	"github.com/sctlee/hazel/daemon/service"
	"github.com/sctlee/hazel/tcpx/client"
)

type ClientMap map[string]*client.Client // string:Cid

type ServiceList map[string]*service.Service

type Daemon struct {
	Clients      ClientMap
	Pending      chan *client.Client
	Quiting      chan *client.Client // string : cid
	joinedNumber int

	config *DaemonConfig

	MsgManager message.Manager
	SrvManager service.Manager
	// Services       ServiceManagereList
}

func (self *Daemon) Listen() {
	for {
		select {
		case client := <-self.Pending:
			self.Join(client)
		case client := <-self.Quiting:
			self.SrvManager.TriggerEvent(EVENT_CLIENT_QUIT, client.Cid)
			self.Quit(client)
		}
	}
}

func (self *Daemon) Join(client *client.Client) {
	self.Clients[client.Cid] = client
	Logger.Println(fmt.Sprintf("one client joined, id:%s", client.Cid))
	fmt.Println("one client joined")
}

func (self *Daemon) Quit(client *client.Client) {
	delete(self.Clients, client.Cid)
	Logger.Println(fmt.Sprintf("one client quited, id:%s", client.Cid))
	fmt.Println("one client quited")
}

var Logger *log.Logger

func NewDaemon(daemonConfig *DaemonConfig) *Daemon {
	d := &Daemon{
		Clients: make(ClientMap),
		Pending: make(chan *client.Client),
		Quiting: make(chan *client.Client),
		config:  daemonConfig,
	}
	d.MsgManager = NewMessageManager(d)
	d.SrvManager = NewServiceManager(d)

	// set log
	Logger = daemonConfig.Logger

	go d.Listen()

	return d
}
