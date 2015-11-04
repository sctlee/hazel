package tcpx

import (
	"fmt"
)

type RouteFun func(params map[string]string, c *Client)

type RouterList map[string]RouteFun

func (self RouterList) RouteMsg(client *Client, params map[string]string) bool {
	fmt.Printf("route %v msg:%t", client, params)

	if route, ok := self[params["feature"]]; ok {
		route(params, client)
		return true
	}
	return false
}
