package backend

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

// EDITOR ROUTE
func Editor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	// load login if no session cookie is present
	cookie, err := request.Cookie("session")
	if err != nil {
		err := editor_templates.ExecuteTemplate(writer, "login.html", nil)
		if err != nil {
			ErrorLogger.Println(err)
			http.Error(writer, "Възникна грешка", 502)
		}
		return
	}

	// connect to database
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// check if session ID in sessions
	var id string
	err = db.QueryRow(`SELECT session_id FROM sessions WHERE session_id = $1`, cookie.Value).Scan(&id)
	if err != nil {
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	// if index - get categories template
	// else find path in category_paths and go there

	response := []byte("editor")
	writer.Write(response)
}


// AUTHENTICATION ROUTE
func Authentication(writer http.ResponseWriter, request *http.Request) {
	// parse request
	var user_data UserData
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(data, &user_data)

	// connect to database
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// validate username
	var username string
	err = db.QueryRow(`SELECT username FROM users
	WHERE username = $1`, user_data.Username).Scan(&username)
	if err != nil {
		writeResponse(writer, "НЕВАЛИДЕН ПОТРЕБИТЕЛ")
		return
	}

	// query database
	var hash string
	var attempts int
	err = db.QueryRow(`SELECT password, attempts FROM users
	WHERE username = $1`, username).Scan(&hash, &attempts)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	// validate attempts
	if attempts == max_attempts {
		writeResponse(writer, "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП")
		return 
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(user_data.Password))
	if err != nil {
		updateAttempts(writer, db, attempts+1, username)
		writeResponse(writer, "ГРЕШНА ПАРОЛА")
		return
	}

	// generate new session ID
	id := uuid.NewV4().String()
	cookie := http.Cookie{
		Name: "session",
		Value: id,
		// Secure: true,
		HttpOnly: true,
	}

	// write session ID to DB
	_, err = db.Exec(`INSERT INTO sessions (session_id) VALUES ($1)`, id)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	// handle response
	http.SetCookie(writer, &cookie)
	updateAttempts(writer, db, 0, username)
	writeResponse(writer, "ok")
}

func writeResponse(writer http.ResponseWriter, response string) {
	writer.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	writer.Write(output)
}

func updateAttempts(writer http.ResponseWriter, db *sql.DB, attempts int, username string) {
	_, err := db.Exec(`UPDATE users SET attempts = $1
	WHERE username = $2`, attempts, username)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
}