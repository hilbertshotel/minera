package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var connStr string = "user=postgres dbname=minera_catalog host=/run/postgresql"

func main() {
	// generate hash
	password := []byte("password")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil { fmt.Println(err); return }

	username := "username"
	attempts := 0

	// connect to database
	db, err := sql.Open("postgres", connStr)
    if err != nil { fmt.Println(err); return }
	defer db.Close()

	// insert password into database
	_, err = db.Exec(`INSERT INTO login (username, password, attempts)
	VALUES ($1, $2, $3)`, username, string(hash), attempts) 
    if err != nil { fmt.Println(err); return }

	fmt.Println("ok")
}
