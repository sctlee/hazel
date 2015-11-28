package model

import (
	"log"

	"github.com/sctlee/hazel/db"
)

type UserModel struct {
	id       int32
	Name     string
	Password string
}

func (self *UserModel) Save() error {
	_, err := db.Pool.Exec("insert into account(name, password) values($1, $2)",
		self.Name, self.Password)
	return err
}

func Exists(name string, password string) (user *UserModel, err error) {
	// user = nil
	// err = nil
	user = &UserModel{}

	row := db.Pool.QueryRow("select id from account where name = $1 and password = $2", name, password)
	err = row.Scan(&user.id)
	if err == nil {
		user.Name = name
	} else {
		log.Println(err)
	}
	return
}
