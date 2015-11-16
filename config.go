package tcpx

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sctlee/tcpx/mlog"
	"github.com/sctlee/tcpx/protocol"

	"github.com/jackc/pgx"
	"gopkg.in/yaml.v2"
)

const (
	LOG_DEFAULT_FILE = "gen.log"
)

var pt protocol.Protocol
var logger *log.Logger

type Config struct {
	Host    string
	Port    string
	LogFile string
	Db      pgx.ConnConfig
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

	pt = new(protocol.SimpleProtocol)
	if len(config.LogFile) != 0 {
		logger = mlog.InitLogger(config.LogFile)
	} else {
		logger = mlog.InitLogger(LOG_DEFAULT_FILE)
	}

	return
}
