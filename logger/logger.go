package logger

import (
	"log"
	"os"
)


var RunLogger = New("run_log")
var DebugLogger = New("Debug_log")

func New(filename string) *log.Logger {
	PFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	return log.New(PFile,"",log.LstdFlags)
}

