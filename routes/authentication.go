package routes

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"minera/data"
)

type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}


func Authentication(writer http.ResponseWriter, request *http.Request) {
	// parse request
	var userData Credentials
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.LogErr(err, writer); return }
	json.Unmarshal(requestData, &userData)

	// editor access log
	data.LogAccess(request, userData.Username)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil { data.LogErr(err, writer); return }
	defer db.Close()

	// validate username
	var username string
	err = db.QueryRow(`SELECT username FROM users
	WHERE username = $1`, userData.Username).Scan(&username)
	if err != nil {
		writeResponse(writer, "НЕВАЛИДЕН ПОТРЕБИТЕЛ")
		return
	}

	// query database
	var hash string
	var attempts int
	err = db.QueryRow(`SELECT password, attempts FROM users
	WHERE username = $1`, username).Scan(&hash, &attempts)
	if err != nil { data.LogErr(err, writer); return }

	// validate attempts
	if attempts == data.MaxAttempts {
		writeResponse(writer, "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП")
		return 
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(userData.Password))
	if err != nil {
		updateAttempts(writer, db, attempts+1, username)
		writeResponse(writer, "ГРЕШНА ПАРОЛА")
		return
	}

	// generate new session ID
	sessionId := uuid.NewV4().String()
	cookie := http.Cookie{
		Name: data.CookieName,
		Value: sessionId,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	// write session ID to DB
	_, err = db.Exec(`INSERT INTO sessions (session_id) VALUES ($1)`, sessionId)
	if err != nil { data.LogErr(err, writer); return }

	// handle response
	http.SetCookie(writer, &cookie)
	updateAttempts(writer, db, 0, username)
	writeResponse(writer, "ok")
}


func writeResponse(writer http.ResponseWriter, response string) {
	writer.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { data.LogErr(err, writer); return }
	writer.Write(output)
}


func updateAttempts(writer http.ResponseWriter, db *sql.DB, attempts int, username string) {
	_, err := db.Exec(`UPDATE users SET attempts = $1
	WHERE username = $2`, attempts, username)
	if err != nil { data.LogErr(err, writer) }
}
