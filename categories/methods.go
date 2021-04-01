package categories

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"minera/data"
	"minera/logs"
)

type category struct {
	Id int
	Name string
}


func get(writer http.ResponseWriter) {
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
	var categories []category
	for rows.Next() {
		cat := category{}
		err = rows.Scan(&cat.Id, &cat.Name)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		categories = append(categories, cat)
	}

	// return template
	err = data.EditorTemplates.ExecuteTemplate(writer, "categories.html", categories)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func post(writer http.ResponseWriter, request *http.Request) {
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


func put(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var categoryData category
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


func delete(writer http.ResponseWriter, request *http.Request) {
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
	
	// delete all products in all sub_categories and all sub_categories in this current category
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}
