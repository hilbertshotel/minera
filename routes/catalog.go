package routes

import (
	"log"
	"net/http"
	"database/sql"
	"strings"
	"strconv"
	"html/template"
	"minera/methods"
)

func Catalog(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	conn string,
	tmp *template.Template) {

	// Database Connection
	// ==================================================

	db, err := sql.Open("postgres", conn)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return	
	}
	defer db.Close()
	

	// Handle Categories
	// ==================================================

	url := r.URL.Path[1:]

	if url == "" {
		categories, err := methods.GetCategories(log, db)
		if err != nil {
			http.Error(w, "Backend Error", 502)
			return
		}
		
		if err := tmp.ExecuteTemplate(w, "categories.html", categories); err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}

		return
	}


	// Handle Sub Categories
	// ==================================================

	urlArray := strings.Split(url, "/")

	if len(urlArray) == 1 {
		catId, err := strconv.Atoi(urlArray[0])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		subCategories, err := methods.GetSubCategories(log, db, catId)
		if err != nil {
			http.Error(w, "Backend Error", 502)
			return
		}

		if err = tmp.ExecuteTemplate(w, "sub_categories.html", subCategories); err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}

		return
	}


	// Handle Products
	// ==================================================

	if len(urlArray) == 2 {
		catId, err := strconv.Atoi(urlArray[0])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		subId, err := strconv.Atoi(urlArray[1])
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}
	
		products, err := methods.GetProducts(log, db, catId, subId)
		if err != nil {
			http.Error(w, "Backend Error", 502)
			return
		}

		if err := tmp.ExecuteTemplate(w, "products.html", products); err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}
		
		return
	}

	
	http.Error(w, "Not Found", 404)
}
