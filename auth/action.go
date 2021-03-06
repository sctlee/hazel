package auth

import (
	"fmt"
	"strings"

	. "github.com/sctlee/hazel/auth/model"

	"github.com/sctlee/hazel"
	"github.com/sctlee/hazel/daemon"
	"github.com/sctlee/hazel/daemon/message"
	"github.com/sctlee/utils"
)

type AuthAction struct {
	AuthList map[string]string // [clientID]username
}

func NewAuthAction() *AuthAction {
	return &AuthAction{make(map[string]string)}
}

func (self *AuthAction) SetUserName(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "name") {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Please input name"))
		return
	}
	name := msg.Params["name"]

	//TODO: set user name

	hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, fmt.Sprintf("Hello, %s", name)))
}

func (self *AuthAction) Login(msg *message.Message) {
	/*
		use postgresql
	*/
	if !utils.IsExistInMap(msg.Params, "username", "password") {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "msg.Params error"))
		return
	}
	username := msg.Params["username"]
	password := msg.Params["password"]

	user, err := Exists(username, password)
	if err != nil {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Username or password error!"))
		return
	} else {
		// save login status in msg.Src.sharedPreferences
		// sp := msg.Src.GetSharedPreferences("Auth")
		// sp.Set("username", user.Name)
		self.AuthList[msg.Src] = user.Name
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Login Success!"))
	}
}

func (self *AuthAction) Logout(msg *message.Message) {
	// sp := msg.Src.GetSharedPreferences("Auth")
	// if _, ok := sp.Get("name"); ok {
	// sp.Del("name")
	if _, ok := self.AuthList[msg.Src]; ok {
		delete(self.AuthList, msg.Src)
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Logout success!"))
	} else {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Please login first!"))
	}
}

func (self *AuthAction) Signup(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "username", "password", "confitm") {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "msg.Params error"))
		return
	}
	username := msg.Params["username"]
	password := msg.Params["password"]
	confirm := msg.Params["confirm"]

	if strings.EqualFold(password, confirm) {
		user := &UserModel{
			Name:     username,
			Password: password,
		}
		if err := user.Save(); err != nil {
			hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Signup error!"))
		} else {
			hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "Signup success! Now you can login with your account!"))
		}
	} else {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src, "confirm is not equal to password"))
	}
}

func (self *AuthAction) GetUserName(msg *message.Message) {
	if username, ok := self.AuthList[msg.Params["cid"]]; ok {
		msg.Params["username"] = username
	} else {
		msg.Params["username"] = "匿名(you have not logined)"
	}
	msg.Des = msg.Src
	hazel.SendMessage(msg)
}
