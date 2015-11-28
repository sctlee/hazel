package daemon

import (
	"errors"

	"github.com/sctlee/hazel/daemon/message"
	"github.com/sctlee/hazel/protocol"
)

const (
	MESSAGE_TYPE_TOCLIENT      = "message2client"
	MESSAGE_TYPE_TOMULTICLIENT = "message2clients"
	MESSAGE_TYPE_TOSERVICE     = "message2service"
)

type MessageManager struct {
	daemon *Daemon
}

func (self *MessageManager) PutMessage(msg *message.Message) error {
	switch msg.Type {
	case MESSAGE_TYPE_TOCLIENT:
		if client, ok := self.daemon.Clients[msg.Des]; ok {
			go client.PutOutgoing(msg.Params["msg"])
			return nil
		}
		return errors.New("no this client")
	case MESSAGE_TYPE_TOMULTICLIENT:
		for _, cid := range msg.MultiDes {
			if client, ok := self.daemon.Clients[cid]; ok {
				go client.PutOutgoing(msg.Params["msg"])
			}
		}
		return nil
	case MESSAGE_TYPE_TOSERVICE:
		if service, ok := self.daemon.SrvManager.GetService(msg.Des); ok {
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

func NewSimpleMessage(des string, msg string) *message.Message {
	params := map[string]string{"msg": msg}
	return message.NewMessage(&protocol.SimpleProtocol{},
		"", des, params, MESSAGE_TYPE_TOCLIENT)
}

func NewSimpleBoardMessage(mdes []string, msg string) *message.Message {
	params := map[string]string{"msg": msg}
	return message.NewBoardMessage(&protocol.SimpleProtocol{},
		"", mdes, params, MESSAGE_TYPE_TOMULTICLIENT)
}
