package hazel

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net"
// 	"strings"
//
// 	"gopkg.in/yaml.v2"
// )
//
// var ConnectionMap map[string]*Client = make(map[string]*Client)
//
// const SERVER_LIST_FILE = "servers.yml"
//
// type HarborServer struct {
// 	Name string
// 	Addr string
// }
//
// type HarborServerList struct {
// 	Servers []HarborServer
// }
//
// func HarborStart() {
// 	data, err := ioutil.ReadFile(SERVER_LIST_FILE)
// 	if err != nil {
// 		fmt.Println("Can't find servers.conf. It may be not in a cluster.")
// 		return
// 	}
//
// 	var hsl HarborServerList
// 	err = yaml.Unmarshal(data, &hsl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	// create harbor server
// 	hb := CreateServer()
// 	hb.Routers["harbor"] = HarborRouter
// 	go hb.Start(hsl.Servers[0].Addr)
//
// 	for _, s := range hsl.Servers[1:] {
// 		for {
// 			conn, err := net.Dial("tcp", ":"+s.Addr)
// 			if err == nil {
// 				client := CreateClient(conn)
// 				ConnectionMap[s.Name] = client
// 				break
// 			}
// 		}
// 	}
// }
//
// func PutMessage2Harbor(des ClientID, msg string) {
// 	sname := strings.Split(string(des), ".")[0]
// 	fmt.Println("124:" + msg)
// 	paramsString := fmt.Sprintf("harbor.pop|cid:%s;msg:%s", des, msg)
// 	ConnectionMap[sname].PutOutgoing(paramsString)
// }
