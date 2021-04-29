package routes

import (
	"log"
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"minera/config"
)

type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}

func authentication(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	cfg *config.Config,
	db *sql.DB) {

	// parse request
	var userData Credentials
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}
	json.Unmarshal(request, &userData)

	// validate username
	var username string
	err = db.QueryRow(`SELECT username FROM users
	WHERE username = $1`, userData.Username).Scan(&username)
	if err != nil {
		writeResponse(w, "НЕВАЛИДЕН ПОТРЕБИТЕЛ", log)
		log.Println("WARNING: registered login attempt by unknown user", userData.Username)
		return
	}

	// query database
	var hash string
	var attempts int
	err = db.QueryRow(`SELECT password, attempts FROM users
	WHERE username = $1`, username).Scan(&hash, &attempts)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}

	// validate attempts
	if attempts == cfg.MaxAtt {
		writeResponse(w, "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП", log)
		return 
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(userData.Password))
	if err != nil {
		updateAttempts(w, db, attempts+1, username, log)
		writeResponse(w, "ГРЕШНА ПАРОЛА", log)
		return
	}

	// generate new session ID
	sessionId := uuid.NewV4().String()
	cookie := http.Cookie{
		Name: cfg.CookieName,
		Value: sessionId,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	// write session ID to DB
	_, err = db.Exec(`INSERT INTO sessions (session_id) VALUES ($1)`, sessionId)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}

	// handle response
	log.Println("successful login")
	http.SetCookie(w, &cookie)
	updateAttempts(w, db, 0, username, log)
	writeResponse(w, "ok", log)
}


func writeResponse(w http.ResponseWriter, response string, log *log.Logger) {
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}


func updateAttempts(w http.ResponseWriter, db *sql.DB, attempts int, name string, log *log.Logger) {
	_, err := db.Exec(`UPDATE users SET attempts = $1
	WHERE username = $2`, attempts, name)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}
}
