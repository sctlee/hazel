package message

import (
	"github.com/sctlee/tcpx/protocol"
)

const (
	MESSAGE_TYPE_TOCLIENT      = "message2client"
	MESSAGE_TYPE_TOMULTICLIENT = "message2clients"
	MESSAGE_TYPE_TOSERVICE     = "message2service"
)

type Message struct {
	Type     string
	Src      string
	Des      string
	MultiDes []string
	Command  string
	Params   map[string]string
}

type Manager interface {
	PutMessage(msg *Message) error
}

func NewMessage(pt protocol.Protocol, src string, des string, rawData string, t string) *Message {
	params := pt.Marshal(rawData)

	msg := &Message{
		Type:    t,
		Src:     src,
		Des:     des,
		Command: params["command"],
		Params:  params,
	}

	if len(des) == 0 {
		msg.Des = params["feature"]
	}

	return msg
}

func NewBoardMessage(pt protocol.Protocol, src string, mdes []string, rawData string, t string) *Message {
	params := pt.Marshal(rawData)
	return &Message{
		Type:     t,
		Src:      src,
		MultiDes: mdes,
		Command:  params["command"],
		Params:   params,
	}
}

func NewSimpleMessage(des string, msg string) *Message {
	return &Message{
		Type:    MESSAGE_TYPE_TOCLIENT,
		Src:     "",
		Des:     des,
		Command: "",
		Params:  map[string]string{"msg": msg},
	}
}

func NewSimpleBoardMessage(mdes []string, msg string) *Message {
	return &Message{
		Type:     MESSAGE_TYPE_TOMULTICLIENT,
		Src:      "",
		MultiDes: mdes,
		Command:  "",
		Params:   map[string]string{"msg": msg},
	}
}
