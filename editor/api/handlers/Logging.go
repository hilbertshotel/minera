package handlers

import (
	"os"
	"log"
)

var ErrorLogger *log.Logger

func init() {
	file, err := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { log.Println(err); return }
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
