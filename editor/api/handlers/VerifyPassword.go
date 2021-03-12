package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	// request data
	var password string
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return }
	json.Unmarshal(request, &password)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()

	// query database
	var hash string
	row := db.QueryRow("SELECT password FROM login")
	err = row.Scan(&hash)
	if err != nil { ErrorLogger.Println(err); return }

	// compare passwords
	var response string
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		response = "error"
	} else {
		response = "ok"
	}

	// return response
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { ErrorLogger.Println(err); return }
	w.Write(output)
}
