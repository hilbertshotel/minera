package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	// request data
	var userData UserData
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return }
	json.Unmarshal(request, &userData)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()

	// validate username
	var username, response string
	err = db.QueryRow(`SELECT username FROM login
	WHERE username = $1`, userData.Username).Scan(&username)
	if err != nil {
		response = "НЕВАЛИДЕН ПОТРЕБИТЕЛ"
		writeResponse(w, response)
		return
	}

	// query database
	var hash string
	var attempts int
	err = db.QueryRow(`SELECT password, attempts FROM login
	WHERE username = $1`, username).Scan(&hash, &attempts)
	if err != nil { ErrorLogger.Println(err); return }

	// validate attempts
	if attempts == maxAttempts {
		response = "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП"
		writeResponse(w, response)
		return 
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(userData.Password))
	if err != nil {
		response = "ГРЕШНА ПАРОЛА"
		attempts++
	} else {
		response = "ok"
		attempts = 0
	}

	// update attempts
	_, err = db.Exec(`UPDATE login SET attempts = $1
	WHERE username = $2`, attempts, username)
	if err != nil { ErrorLogger.Println(err); return }
	
	writeResponse(w, response)
}

func writeResponse(w http.ResponseWriter, response string) {
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { ErrorLogger.Println(err); return }
	w.Write(output)
}
