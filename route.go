package tcpx

import (
	"fmt"

	"github.com/sctlee/tcpx/protocol"
	// "secret/chatroom"
)

var pt protocol.Protocol

type RouteFun func(params map[string]string, c *Client)

type Router struct {
	RouteList map[string]RouteFun
}

func (self *Router) Route(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)

	pt = new(protocol.SimpleProtocol)
	m := pt.Marshal(msg)
	fmt.Println(m)

	if f, ok := self.RouteList[m["feature"]]; ok {
		f(m, client)
		return true
	}
	return false
}
