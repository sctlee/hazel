package daemon

import (
	"errors"

	"github.com/sctlee/tcpx/daemon/message"
)

type MessageManager struct {
	daemon *Daemon
}

func (self *MessageManager) PutMessage(msg *message.Message) error {
	switch msg.Type {
	case message.MESSAGE_TYPE_TOCLIENT:
		if client, ok := self.daemon.Clients[ClientID(msg.Des)]; ok {
			go client.PutOutgoing(msg.Params["msg"])
			return nil
		}
		return errors.New("no this client")
	case message.MESSAGE_TYPE_TOMULTICLIENT:
		for _, cid := range msg.MultiDes {
			if client, ok := self.daemon.Clients[ClientID(cid)]; ok {
				go client.PutOutgoing(msg.Params["msg"])
			}
		}
		return nil
	case message.MESSAGE_TYPE_TOSERVICE:
		if service, ok := self.daemon.Services[msg.Des]; ok {
			return service.PutMessage(msg)
		}
		return errors.New("no this service")
	default:
		return errors.New("no this type")
	}
}

func NewMessageManager(d *Daemon) message.Manager {
	return &MessageManager{
		daemon: d,
	}
}
