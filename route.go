package hazel

// import (
// 	"fmt"
// )
//
// type RouteFun func(cid ClientID, params map[string]string)
//
// type RouterList map[string]map[string]RouteFun
//
// func (self RouterList) RouteMsg(cid ClientID, msg string) bool {
// 	fmt.Printf("route %v msg:%s", cid, msg)
//
// 	requestMsg := NewMessage(cid, msg)
// 	params := requestMsg.Get()
// 	if router, ok := self[params["feature"]]; ok {
// 		// is it necessary of setting this in a goruntine?
// 		go func() {
// 			if f, ok := router[params["command"]]; ok {
// 				f(cid, params)
// 			} else {
// 				SendMessage(NewMessage(cid, fmt.Sprintf("no '%s' command", params["command"])))
// 			}
// 		}()
// 		return true
// 	}
// 	return false
// }
