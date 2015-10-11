package tcpx

import (
	"fmt"
	"strings"
)

var UserList map[*Client]string

func init() {
	UserList = make(map[*Client]string)
}

func GetUserName(client *Client) string {
	s := UserList[client]
	if s != "" {
		return s
	} else {
		return "匿名"
	}
}

func Route(url string, client *Client) {
	url = strings.TrimSpace(url)
	fmt.Println(url)
	i := strings.Index(url, " ")
	action := url[:i]
	fmt.Println(url)
	switch action {
	case "setName":
		name := strings.TrimSpace(url[i:])
		UserList[client] = name
		client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
	}
}
