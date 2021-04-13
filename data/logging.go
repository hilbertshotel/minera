package data

import (
	"os"
	"log"
	"sync"
	"fmt"
	"time"
	"net/http"
	"runtime"
)

var errorLock = &sync.Mutex{}
var accessLock = &sync.Mutex{}
var requestLock = &sync.Mutex{}


func LogErr(err error, writer http.ResponseWriter) {
	// handle error mutex
	errorLock.Lock()
	defer errorLock.Unlock()
	
	_, filepath, line, _ := runtime.Caller(1)

	// prepare log
	t := time.Now().Format("2 Jan 2006 15:04:05")
	msg := fmt.Sprintf(`________________________________________
DATETIME: %s
ERROR MSG: %s
FILE NAME: %s
LINE NUMBER: %d
`, t, err, filepath, line)
	
	// open file
	file, err := os.OpenFile(errorLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { log.Println(err); return }
	defer file.Close()
	
	// write to file
	_, err = file.WriteString(msg)
	if err != nil { log.Println(err); return }

	http.Error(writer, "Възникна грешка", 502)
}


func LogAccess(request *http.Request, username string) {
	// handle access mutex
	accessLock.Lock()
	defer accessLock.Unlock()

	// prepare log
	t := time.Now().Format("2 Jan 2006 15:04:05")
	msg := fmt.Sprintf(`________________________________________
DATETIME: %s
USERNAME: %s
REMOTE ADDRESS: %s
USER AGENT: %s
`, t, username, request.RemoteAddr, request.UserAgent())

	// open file
	file, err := os.OpenFile(accessLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { log.Println(err); return }
	defer file.Close()

	// write to file
	_, err = file.WriteString(msg)
	if err != nil { log.Println(err); return }
}


func LogRequest(writer http.ResponseWriter, request *http.Request) {
	// handle request mutex
	requestLock.Lock()
	defer requestLock.Unlock()

	// prepare log
	t := time.Now().Format("2 Jan 2006 15:04:05")
	msg := fmt.Sprintf(`________________________________________
DATETIME: %s
URL: %s
REMOTE ADDRESS: %s
USER AGENT: %s
`, t, request.URL.Path, request.RemoteAddr, request.UserAgent())

	// open file
	file, err := os.OpenFile(requestLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { log.Println(err); return }
	defer file.Close()

	// write to file
	_, err = file.WriteString(msg)
	if err != nil { log.Println(err); return }

	http.Error(writer, "Възникна грешка", 502)
}	