package backend

import (
	"os"
	"log"
	"html/template"
)

// POSTGRESQL CONNECTION STRING
// const connection_string string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql" // unix
const connection_string = "user=postgres dbname=minera_catalog sslmode=disable" // windows

const Address = "127.0.0.1:8000"
var templates = template.Must(template.ParseGlob("templates/*"))

// LOGGER
const log_file = "logs/errors.log"
var Logger *log.Logger

func init() {
	file, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil { log.Println(err); return }
	Logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// STRUCTS
type Category struct {
	Name string
	Path string
	Function string
}

type SubCategory struct {
	Name string
	Path string
	Function string
}

type Item struct {
	Name string
	Description string
	Images []string
}

// FUNCTION LIST
// GLOBAL:
// get_categories OK
// add_category

// CATEGORIES:
// get_sub_categories NOW
// edit_category
// delete_category
// add_sub_category

// SUB_CATEGORIES
// get_products
// edit_sub_category
// delete_sub_category
// add_product

// PRODUCTS
// delete_product
// edit_product