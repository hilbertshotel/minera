package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// const connection_string = "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql" // linux
const connection_string = "user=postgres dbname=minera_catalog sslmode=disable" // windows

func main() {
	// data
	username := "asd"
	password := []byte("asd")

	// generate hash
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil { fmt.Println(err); return }

	// connect to database
	db, err := sql.Open("postgres", connection_string)
    if err != nil { fmt.Println(err); return }
	defer db.Close()

	// insert password into database
	_, err = db.Exec(`INSERT INTO users (username, password)
	VALUES ($1, $2)`, username, string(hash)) 
    if err != nil { fmt.Println(err); return }

	fmt.Println("ok")
}
