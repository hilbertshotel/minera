package routes

import (
	"strconv"
	"strings"
	"net/http"
	"database/sql"
	"minera/data"
	"minera/methods"
)

func Editor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	// load login if no cookie is present
	cookie, err := request.Cookie(data.CookieName)
	if err != nil {
		err := data.EditorTemplates.ExecuteTemplate(writer, "login.html", nil)
		if err != nil { data.Log(err, writer) }
		return
	}

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil { data.Log(err, writer); return }
	defer db.Close()

	// check if session ID is in sessions else load login
	var id string
	err = db.QueryRow(`SELECT session_id FROM sessions WHERE session_id = $1`, cookie.Value).Scan(&id)
	if err != nil {
		err := data.EditorTemplates.ExecuteTemplate(writer, "login.html", nil)
		if err != nil { data.Log(err, writer) }
		return
	}

	// handle categories
	url := request.URL.Path[1:]
	if url == "editor/" {
		methods.CategoriesDispatcher(db, writer, request)
		return
	}

	// handle sub categories
	urlArray := strings.Split(url, "/")
	if len(urlArray) == 2 {
		categoryId, err := strconv.Atoi(urlArray[1])
		if err != nil { data.Log(err, writer); return }
		methods.SubCategoriesDispatcher(db, writer, request, categoryId)
		return
	}

	// handle products
	if len(urlArray) == 3 {
		categoryId, err := strconv.Atoi(urlArray[1])
		if err != nil { data.Log(err, writer); return }

		subCategoryId, err := strconv.Atoi(urlArray[2])
		if err != nil { data.Log(err, writer); return }
	
		methods.ProductsDispatcher(db, writer, request, categoryId, subCategoryId)
		return
	}

	http.Error(writer, "Страницата не съществува", 404)
}
