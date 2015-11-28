package auth

import (
	"github.com/sctlee/hazel/daemon/service"
)

func (self *AuthAction) GetRouteList() service.RouteList {
	return service.RouteList{
		"setName":     self.SetUserName,
		"login":       self.Login,
		"logout":      self.Logout,
		"signup":      self.Signup,
		"getusername": self.GetUserName,
	}
}
