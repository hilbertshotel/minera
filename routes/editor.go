package routes

import (
	"strconv"
	"strings"
	"net/http"
	"database/sql"
	"minera/data"
	"minera/logs"
	"minera/methods"
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

	// handle index
	url := request.URL.Path[1:]
	if url == "editor/" {
		methods.CategoriesDispatcher(writer, request)
		return
	}

	// handle sub categories
	urlArray := strings.Split(url, "/")
	if len(urlArray) == 2 {
		id, err := strconv.Atoi(urlArray[1])
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		methods.SubCategoriesDispatcher(writer, request, id)
		return
	}

	// handle products
	if len(urlArray) == 3 {
		categoryId, err := strconv.Atoi(urlArray[1])
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}

		subCategoryId, err := strconv.Atoi(urlArray[2])
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
	
		methods.ProductsDispatcher(writer, request, categoryId, subCategoryId)
		return
	}

	http.Error(writer, "Възникна грешка", 502)
}
