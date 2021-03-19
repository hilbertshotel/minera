package utils

import (
	"log"
	"os"
)

const logFile = "backend/logs/errors.log"
var Logger *log.Logger

func init() {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { log.Println(err); return }
	Logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
