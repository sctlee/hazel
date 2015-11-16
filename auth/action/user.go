package action

import (
	"fmt"
	"strings"

	. "github.com/sctlee/tcpx/auth/model"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/utils"
)

type UserAction struct {
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func (self *UserAction) SetUserName(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "name") {
		return tcpx.NewMessage(client, "Please input name")
	}
	name := params["name"]

	//TODO: set user name

	return tcpx.NewMessage(client, fmt.Sprintf("Hello, %s", name))
}

func (self *UserAction) Login(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	/*
		use postgresql
	*/
	if !utils.IsExistInMap(params, "username", "password") {
		return tcpx.NewMessage(client, "params error")
	}
	username := params["username"]
	password := params["password"]

	user, err := Exists(username, password)
	if err != nil {
		return tcpx.NewMessage(client, "Username or password error!")
	} else {
		// save login status in client.sharedPreferences
		sp := client.GetSharedPreferences("Auth")
		sp.Set("username", user.Name)

		return tcpx.NewMessage(client, "Login Success!")
	}
}

func (self *UserAction) Logout(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	sp := client.GetSharedPreferences("Auth")
	if _, ok := sp.Get("name"); ok {
		sp.Del("name")
		return tcpx.NewMessage(client, "Logout success!")
	} else {
		return tcpx.NewMessage(client, "Please login first!")
	}
}

func (self *UserAction) Signup(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "username", "password", "confitm") {
		return tcpx.NewMessage(client, "params error")
	}
	username := params["username"]
	password := params["password"]
	confirm := params["confirm"]

	if strings.EqualFold(password, confirm) {
		user := &UserModel{
			Name:     username,
			Password: password,
		}
		if err := user.Save(); err != nil {
			return tcpx.NewMessage(client, "Signup error!")
		} else {
			return tcpx.NewMessage(client, "Signup success! Now you can login with your account!")
		}
	} else {
		return tcpx.NewMessage(client, "confirm is not equal to password")
	}
}
