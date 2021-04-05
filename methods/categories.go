package methods

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"minera/data"
	"minera/logs"
)

func getCategories(writer http.ResponseWriter) {
	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// query database
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY	id ASC")
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer rows.Close()

	// package data
	var categories []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		categories = append(categories, category)
	}

	// return template
	err = data.EditorTemplates.ExecuteTemplate(writer, "categories.html", categories)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func postCategory(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var newCategoryName string
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &newCategoryName)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()
	
	// query database
	_, err = db.Exec(`INSERT INTO categories (name, added)
	VALUES ($1, now())`, newCategoryName)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func putCategory(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var categoryData Category
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &categoryData)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()
	
	// edit category
	_, err = db.Exec(`UPDATE categories SET name = $1
	WHERE id = $2`, categoryData.Name, categoryData.Id) 
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func deleteCategory(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var id int
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &id)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()
	
	// delete query 
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}