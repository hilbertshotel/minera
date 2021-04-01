package logs

import (
	"os"
	"log"
)

// ERROR LOGGER
const logFile = "logs/errors.log"
var Errors *log.Logger

func init() {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { log.Println(err); return }
	Errors = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
