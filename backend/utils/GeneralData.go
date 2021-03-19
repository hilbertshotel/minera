package utils

// POSTGRESQL CONNECTION STRING
// const ConnStr string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql" // unix
const ConnStr string = "user=postgres dbname=minera_catalog sslmode=disable" // windows

// IMAGE DIRECTORY
const ImageDir = "frontend/catalog/images/"

// MAX LOGIN ATTEMPTS
const MaxAttempts = 10 // max login attempts