package tcpx

import (
	"fmt"
)

type RouteFun func(params map[string]string, c *Client) IMessage

type RouterList map[string]RouteFun

func (self RouterList) RouteMsg(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)

	requestMsg := NewMessage(client, msg)
	params := requestMsg.Get()
	if route, ok := self[params["feature"]]; ok {
		// is it necessary of setting this in a goruntine?
		go func() {
			responseMsg := route(params, client)
			responseMsg.exec()
		}()
		return true
	}
	return false
}
