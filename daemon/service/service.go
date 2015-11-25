package service

import (
	"errors"

	"github.com/sctlee/tcpx/daemon/message"
)

const (
	MSG_QUEUE_NUM = 10
)

type RouteFun func(msg *message.Message)
type RouteList map[string]RouteFun

type Service struct {
	Name   string
	Routes RouteList

	MsgReceiver chan *message.Message
}

func (self *Service) Listen() {
	for m := range self.MsgReceiver {
		self.ExecMsg(m)
	}
}

func (self *Service) PutMessage(msg *message.Message) error {
	if _, ok := self.Routes[msg.Command]; ok {
		self.MsgReceiver <- msg
		return nil
	}
	return errors.New("no this command")
}

func (self *Service) ExecMsg(msg *message.Message) {
	if f, ok := self.Routes[msg.Command]; ok {
		f(msg)
	}
}

func NewService(name string, routes RouteList) *Service {
	return &Service{
		Name:        name,
		Routes:      routes,
		MsgReceiver: make(chan *message.Message, MSG_QUEUE_NUM),
	}
}
