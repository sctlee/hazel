package tcpx

import (
	"fmt"
)

type RouteFun func(client *Client, params map[string]string) IMessage

type RouterList map[string]map[string]RouteFun

func (self RouterList) RouteMsg(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)

	requestMsg := NewMessage(client, msg)
	params := requestMsg.Get()
	if router, ok := self[params["feature"]]; ok {
		// is it necessary of setting this in a goruntine?
		go func() {
			var responseMsg IMessage
			if f, ok := router[params["command"]]; ok {
				responseMsg = f(client, params)
			} else {
				responseMsg = NewMessage(client, fmt.Sprintf("no '%s' command", params["command"]))
			}
			responseMsg.exec()
		}()
		return true
	}
	return false
}
