package backend

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"minera/backend/utils"
)

var cookies = make(map[string]string)

type UserData struct {
	Username string `json:username`
	Password string `json:password`
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	// request data
	var userData UserData
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { utils.Logger.Println(err); return }
	json.Unmarshal(request, &userData)

	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()

	// validate username
	var username string
	err = db.QueryRow(`SELECT username FROM login
	WHERE username = $1`, userData.Username).Scan(&username)
	if err != nil {
		writeResponse(w, "НЕВАЛИДЕН ПОТРЕБИТЕЛ")
		return
	}

	// query database
	var hash string
	var attempts int
	err = db.QueryRow(`SELECT password, attempts FROM login
	WHERE username = $1`, username).Scan(&hash, &attempts)
	if err != nil { utils.Logger.Println(err); return }

	// validate attempts
	if attempts == utils.MaxAttempts {
		writeResponse(w, "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП")
		return 
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(userData.Password))
	if err != nil {
		updateAttempts(db, attempts+1, username)
		writeResponse(w, "ГРЕШНА ПАРОЛА")
		return
	}

	// set cookie
	name, value := utils.GenerateCookie()
	cookie := utils.MakeCookie(name, value)
	cookies[name] = value
	http.SetCookie(w, &cookie)

	updateAttempts(db, 0, username)
	writeResponse(w, "ok")
}


func writeResponse(w http.ResponseWriter, response string) {
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { utils.Logger.Println(err); return }
	w.Write(output)
}

func updateAttempts(db *sql.DB, attempts int, username string) {
	_, err := db.Exec(`UPDATE login SET attempts = $1
	WHERE username = $2`, attempts, username)
	if err != nil { utils.Logger.Println(err); return }
}