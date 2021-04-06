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

var mutex = &sync.Mutex{}

func Log(err error, writer http.ResponseWriter) {
	if err != nil {
  
		// lock mutex
		mutex.Lock()
		defer mutex.Unlock()
	
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
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil { log.Println(err); return }
		defer file.Close()
	
		// write to file
		_, err = file.WriteString(msg)
		if err != nil { log.Println(err); return }

		http.Error(writer, "Възникна грешка", 502)
	}
}
