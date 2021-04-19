package routes

import (
	"log"
	"strconv"
	"strings"
	"net/http"
	"database/sql"
	"html/template"
	"minera/methods"
	"minera/conf"
)

func Editor(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	cfg *conf.Config,
	tmp *template.Template) {

	// Database Connection
	// ==================================================

	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}
	defer db.Close()


	// Validation
	// ==================================================

	cookie, err := r.Cookie(cfg.CookieName)
	if err != nil {
		if err := tmp.ExecuteTemplate(w, "login.html", nil); err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}
		return
	}

	var id string
	if err := db.QueryRow(`SELECT session_id FROM sessions
	WHERE session_id = $1`, cookie.Value).Scan(&id); err != nil {
		if err := tmp.ExecuteTemplate(w, "login.html", nil); err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}
		return
	}


	// Handle Categories
	// ==================================================

	url := r.URL.Path[1:]

	if url == "editor/" {
		err = methods.CategoriesDispatcher(w, r, log, db, tmp)
		if err != nil {
			http.Error(w, "Backend Error", 502)
		}
		return
	}


	// Handle Sub Categories
	// ==================================================

	urlArray := strings.Split(url, "/")

	if len(urlArray) == 2 {
		catId, err := strconv.Atoi(urlArray[1])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		err = methods.SubCategoriesDispatcher(w, r, log, db, tmp, catId)
		if err != nil {
			http.Error(w, "Backend Error", 502)
		}
		return
	}


	// Handle Products
	// ==================================================

	if len(urlArray) == 3 {
		catId, err := strconv.Atoi(urlArray[1])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		subId, err := strconv.Atoi(urlArray[2])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}
	
		err = methods.ProductsDispatcher(w, r, log, db, tmp, catId, subId, cfg.ImgDir)
		if err != nil {
			http.Error(w, "Backend Error", 502)
		}
		return
	}


	http.Error(w, "Not Found", 404)
}
