package data

import "html/template"

// POSTGRESQL CONNECTION STRING
const ConnectionString string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql" // unix
// const ConnectionString = "user=postgres dbname=minera_catalog sslmode=disable" // windows

// TEMPLATES
var EditorTemplates = template.Must(template.ParseGlob("templates/editor/*"))
var CatalogTemplates = template.Must(template.ParseGlob("templates/catalog/*"))

// IP ADDRESS
const Address = "127.0.0.1:8000"

// MAX LOGIN ATTEMPTS
const MaxAttempts = 10

// IMAGE DIRECTORY
const ImageDir = "images/"

// LOGGING FILES
const logFile = "logs/errors.log"