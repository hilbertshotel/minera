package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var connStr string = "user=postgres dbname=minera_catalog host=/run/postgresql"

func main() {
	password := []byte("asd")

	// generate hash
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil { fmt.Println(err); return }

	// connect to database
	db, err := sql.Open("postgres", connStr)
    if err != nil { fmt.Println(err); return }
	defer db.Close()

	// insert password into database
	_, err = db.Exec(`INSERT INTO login (password) VALUES ($1)`, string(hash)) 
    if err != nil { fmt.Println(err); return }

	fmt.Println("ok")
}
