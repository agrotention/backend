package utils

import (
	"log"
	"os"
)

var (
	LogWarn *log.Logger
	LogInfo *log.Logger
	LogErr  *log.Logger
)

func init() {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	LogWarn = log.New(logFile, "INFO  : ", log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo = log.New(logFile, "WARN  : ", log.Ldate|log.Ltime|log.Lshortfile)
	LogErr = log.New(logFile, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}
