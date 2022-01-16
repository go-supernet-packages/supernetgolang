package logger

import (
	"log"
	"os"
)

var Exception *log.Logger
var Error *log.Logger
var Warning *log.Logger
var Info *log.Logger
var Debug *log.Logger

func InitLogger(fileName, appName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
		log.Println("Error opening file:", err)
		os.Exit(1)
	}
	Exception = log.New(f, appName+" EXCEPTION ", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
	Error = log.New(f, appName+" ERROR ", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
	Warning = log.New(f, appName+" WARRING ", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
	Info = log.New(f, appName+" INFO ", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
	Debug = log.New(f, appName+" DEBUG ", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
}
