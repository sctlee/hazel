package service

import (
	"errors"
	"time"

	"github.com/sctlee/hazel/daemon/message"

	"github.com/nu7hatch/gouuid"
)

const (
	MSG_QUEUE_NUM = 10
)

type RouteFun func(msg *message.Message)

type RouteList map[string]RouteFun

type Servicer interface {
	GetRouteList() RouteList
}

type Manager interface {
	RegisterService(service *Service)
	GetService(name string) (*Service, bool)
	TriggerEvent(eventType string, params ...string)
}

type Service struct {
	Name   string
	Routes RouteList

	// original service
	s Servicer

	Sessions    map[*uuid.UUID]chan *message.Message
	MsgReceiver chan *message.Message
}

func (self *Service) Listen() {
	for m := range self.MsgReceiver {
		self.ExecMsg(m)
	}
}

func (self *Service) PutMessage(msg *message.Message) error {
	if _, ok := self.Sessions[msg.Session]; ok {
		// 防止写入超时
		select {
		case <-time.After(time.Second * 2):
			return errors.New("write channel timeout")
		case self.Sessions[msg.Session] <- msg:
			return nil
		}
	}
	self.Sessions[msg.Session] = msg.Response
	if _, ok := self.Routes[msg.Command]; ok {
		self.MsgReceiver <- msg
		return nil
	}
	return errors.New("no this command")
}

func (self *Service) ExecMsg(msg *message.Message) {
	if f, ok := self.Routes[msg.Command]; ok {
		go func() {
			f(msg)
			delete(self.Sessions, msg.Session)
		}()
	}
}

func (self *Service) GetOriginalService() Servicer {
	return self.s
}

func NewService(name string, s Servicer) *Service {
	return &Service{
		Name:        name,
		s:           s,
		Routes:      s.GetRouteList(),
		Sessions:    make(map[*uuid.UUID]chan *message.Message),
		MsgReceiver: make(chan *message.Message, MSG_QUEUE_NUM),
	}
}
