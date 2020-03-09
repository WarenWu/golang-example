package logger

import (
	"log"
	"os"
)


var FLogger = New("urlconv_log")

func New(filename string) *log.Logger {
	PFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	return log.New(PFile,"",log.LstdFlags)
}

