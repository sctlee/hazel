package tcpx

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sctlee/tcpx/db"
	"github.com/sctlee/tcpx/mlog"
	"github.com/sctlee/tcpx/protocol"

	"github.com/jackc/pgx"
	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_SERVER_NAME = "default"
	LOG_DEFAULT_FILE    = "gen.log"
)

type Config struct {
	ServerName string
	Host       string
	Port       string
	LogFile    string
	Db         pgx.ConnConfig
	Redis      db.RedisConfig
}

var logger *log.Logger
var pt protocol.Protocol

func LoadConfig() (config *Config) {
	args := os.Args
	filePath := ""

	if len(args) == 2 {
		filePath = "config.yml"
	} else if len(args) == 3 {
		filePath = args[2]
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}

	// set server name
	if len(config.ServerName) != 0 {
		serverName = config.ServerName
	} else {
		serverName = DEFAULT_SERVER_NAME
	}

	// set protocol
	pt = new(protocol.SimpleProtocol)

	// set log file
	if len(config.LogFile) != 0 {
		logger = mlog.InitLogger(config.LogFile)
	} else {
		logger = mlog.InitLogger(LOG_DEFAULT_FILE)
	}

	// set database
	db.StartPool(config.Db)
	if config.Redis.Name != "" {
		db.StartRedisPool(config.Redis)
	}

	return
}
