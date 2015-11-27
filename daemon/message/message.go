package message

import (
	"fmt"

	"github.com/sctlee/tcpx/protocol"

	"github.com/nu7hatch/gouuid"
)

type Message struct {
	Type     string
	Src      string
	Des      string
	MultiDes []string
	Command  string
	Params   map[string]string

	Session  *uuid.UUID
	Response chan *Message
}

type Manager interface {
	PutMessage(msg *Message) error
}

func CopySession(src *Message, des *Message) {
	des.Session = src.Session
	des.Response = src.Response
}

func NewMessage(pt protocol.Protocol, src string, des string, rawData interface{}, t string) *Message {
	var params map[string]string
	switch rawData.(type) {
	case string:
		params = pt.Marshal(rawData.(string))
	case map[string]string:
		params = rawData.(map[string]string)
	default:
		params = nil
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	msg := &Message{
		Type:    t,
		Src:     src,
		Des:     des,
		Command: params["command"],
		Params:  params,

		Session:  uuid,
		Response: make(chan *Message),
	}

	if len(des) == 0 {
		msg.Des = params["feature"]
	}

	return msg
}

func NewBoardMessage(pt protocol.Protocol, src string, mdes []string, rawData interface{}, t string) *Message {
	msg := NewMessage(pt, src, "", rawData, t)
	msg.MultiDes = mdes
	return msg
}
