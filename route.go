package tcpx

import (
	"fmt"

	"github.com/sctlee/tcpx/protocol"
	// "secret/chatroom"
)

var pt protocol.Protocol

type RouteFun func(params map[string]string, c *Client)

type RouterList map[string]RouteFun

func (self RouterList) RouteMsg(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)

	pt = new(protocol.SimpleProtocol)
	params := pt.Marshal(msg)
	fmt.Println(params)

	if route, ok := self[params["feature"]]; ok {
		route(params, client)
		return true
	}
	return false
}
