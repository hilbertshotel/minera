package routes

import (
	"fmt"
	"strings"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"minera/data"
	"minera/logs"
	"minera/categories"
)

func Editor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	// load login if no 'session' cookie is present
	cookie, err := request.Cookie("session")
	if err != nil {
		err := data.EditorTemplates.ExecuteTemplate(writer, "login.html", nil)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
		}
		return
	}

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// check if session ID is in sessions else load login
	var id string
	err = db.QueryRow(`SELECT session_id FROM sessions WHERE session_id = $1`, cookie.Value).Scan(&id)
	if err != nil {
		err := data.EditorTemplates.ExecuteTemplate(writer, "login.html", nil)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
		}
		return
	}

	// local dispatcher
	url := request.URL.Path[1:]
	if url == "editor/" {
		categories.Dispatcher(writer, request)
		return
	}

	urlArray := strings.Split(url, "/")
	fmt.Println(urlArray)

	response := []byte("editor")
	writer.Write(response)
}
