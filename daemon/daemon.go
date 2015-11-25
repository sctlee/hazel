package daemon

import (
	"fmt"
	"log"

	"github.com/sctlee/tcpx/daemon/message"
	"github.com/sctlee/tcpx/daemon/service"
	"github.com/sctlee/tcpx/tcpx/client"
)

type ClientID string
type ClientMap map[ClientID]*client.Client

type ServiceList map[string]*service.Service

type Daemon struct {
	Clients ClientMap
	Pending chan *client.Client
	// quiting  chan net.Conn
	joinedNumber int

	config *DaemonConfig

	MsgManager message.Manager
	Services   ServiceList
}

func (self *Daemon) Listen() {
	for {
		select {
		case client := <-self.Pending:
			self.Join(client)
		}
	}
}

func (self *Daemon) Join(client *client.Client) {
	genClientID := func() ClientID {
		return ClientID(fmt.Sprintf("%s.%d", self.config.ServerName, self.joinedNumber))
	}

	cid := genClientID()
	fmt.Println("client id :" + cid)
	self.Clients[cid] = client
	self.joinedNumber++
	Logger.Println(fmt.Sprintf("one client joined, id:%s", string(cid)))
	fmt.Println("one client joined")

	go func(cid ClientID) {
		c := self.Clients[cid]
		defer func() {
			delete(self.Clients, cid)
			Logger.Println(fmt.Sprintf("one client quited, id:%s", string(cid)))
			fmt.Println("one client quited")
		}()

		for {
			rawData, ok := c.GetIncoming()
			if !ok {
				break
			}

			fmt.Println(rawData)

			err := self.MsgManager.PutMessage(
				message.NewMessage(
					self.config.Pt, string(cid), "", rawData, message.MESSAGE_TYPE_TOSERVICE))

			if err != nil {
				self.MsgManager.PutMessage(
					message.NewSimpleMessage(
						string(cid),
						fmt.Sprintf("server.error|msg:%s", err)))

			}
			// if !self.Routers.RouteMsg(cid, msg) {
			// 	c.PutOutgoing("command error, Usage:'chatroom join 1','chatroom send hello'")
			// self.incoming <- msg
			// }
		}
	}(cid)
}

func (self *Daemon) RegisterService(s *service.Service) {
	self.Services[s.Name] = s
	go s.Listen()
}

var Logger *log.Logger

func NewDaemon(daemonConfig *DaemonConfig) *Daemon {
	d := &Daemon{
		Clients:  make(ClientMap),
		Pending:  make(chan *client.Client),
		Services: make(ServiceList),
		config:   daemonConfig,
	}
	d.MsgManager = NewMessageManager(d)

	// set log
	Logger = daemonConfig.Logger

	go d.Listen()

	return d
}
