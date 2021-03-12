package handlers

import (
	"os"
	"log"
)

const connStr string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql"

const logFile = "logs/errors.log"
var ErrorLogger *log.Logger

func init() {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { log.Println(err); return }
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Category struct {
	Id int `json:id`
	Name string `json:name`
}

type Item struct {
	Name string `json:name`
	Description string `json:description`
	Images []string `json:images`
}
