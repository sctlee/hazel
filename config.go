package tcpx

import (
	"fmt"
	"io/ioutil"

	"github.com/jackc/pgx"
	"gopkg.in/yaml.v2"
)

const (
	LOG_FILE_NAME = "gen.log"
)

type Config struct {
	Host string
	Port string
	Db   pgx.ConnConfig
}

func LoadConfig(filePath string) (config Config) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}

	return
}
