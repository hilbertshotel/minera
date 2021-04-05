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