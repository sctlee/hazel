package auth

//
// import (
// 	"reflect"
//
// 	"github.com/sctlee/tcpx/daemon/service"
// )
//
// const (
// 	IsLogin = "IsLogin"
// )
//
// type Permission struct {
// }
//
// func (self Permission) IsLogin(client *tcpx.Client) bool {
// 	auth := client.GetSharedPreferences("Auth")
// 	if _, ok := auth.Get("name"); ok {
// 		return true
// 	}
// 	return false
// }
//
// var permission = new(Permission)
//
// func PermissionDecorate(msg *message.Message, permissions ...string) (pf tcpx.RouteFun) {
// 	for _, methodName := range permissions {
// 		method := reflect.ValueOf(permission).MethodByName(methodName)
// 		if method.Interface().(func(client *tcpx.Client) bool)(client) {
// 			pf = f
// 		} else {
// 			pf = func(client *tcpx.Client, params map[string]string) tcpx.IMessage {
// 				return tcpx.NewMessage(client, "Permission refused")
// 			}
// 			break
// 		}
// 	}
// 	return
// }
