package tcpx

import (
	"log"
	"os"
)

var logger *log.Logger
var file *os.File

func init() {
	_, err := os.Stat(LOG_FILE_NAME)
	file, err = os.OpenFile(LOG_FILE_NAME, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("log file set err")
	}

	logger = log.New(file, "logger: ", log.Lshortfile)
}

func LogClose() {
	file.Close()
}
