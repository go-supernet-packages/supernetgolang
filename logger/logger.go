package logger

import (
	"log"
	"os"
)

func GetDebugLogger(fileName, appName string) (l *log.Logger) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println(err)
	}
	l = log.New(f, " DEBUG "+appName+" ", log.LstdFlags)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return
}

func GetInfoLogger(fileName, appName string) (l *log.Logger) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println(err)
	}
	l = log.New(f, " INFO "+appName+" ", log.LstdFlags)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return
}
