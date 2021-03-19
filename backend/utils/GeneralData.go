package utils

// POSTGRESQL CONNECTION STRING
// unix
// const ConnStr string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql"
// windows
const ConnStr string = "user=postgres dbname=minera_catalog sslmode=disable"

// IMAGE DIRECTORY
const ImageDir = "./frontend/catalog/images"

// MAX LOGIN ATTEMPTS
const MaxAttempts = 10 // max login attempts