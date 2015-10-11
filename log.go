package tcpx

import (
	"log"
	"os"
)

var logger *log.Logger
var file *os.File

func init() {
	filename := "gen.log"

	_, err := os.Stat(filename)
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("log file set err")
	}

	logger = log.New(file, "logger: ", log.Lshortfile)
}

func LogClose() {
	file.Close()
}
