package mlog

import (
	"log"
	"os"
)

var file *os.File

func InitLogger(filename string) *log.Logger {
	var logger *log.Logger

	_, err := os.Stat(filename)
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("log file set err")
	}

	logger = log.New(file, "logger: ", log.Lshortfile)

	return logger
}

func LogClose() {
	file.Close()
}
