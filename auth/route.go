package auth

import (
	"github.com/sctlee/tcpx/daemon/service"
)

var AuthService = NewAuthAction()

func (self *AuthAction) GetRouteList() service.RouteList {
	return service.RouteList{
		"setName":     self.SetUserName,
		"login":       self.Login,
		"logout":      self.Logout,
		"signup":      self.Signup,
		"getusername": self.GetUserName,
	}
}
