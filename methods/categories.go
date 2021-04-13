package methods

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"minera/data"
)

func GetCategories(db *sql.DB, writer http.ResponseWriter) ([]Category, error) {
	// query database
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY id ASC")
	if err != nil {
		data.LogErr(err, writer)
		return nil, err
	}
	defer rows.Close()

	// package data
	var categories []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			data.LogErr(err, writer)
			return nil, err
		}
		categories = append(categories, category)
	}
	
	return categories, nil
}


func postCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var newCategoryName string
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.LogErr(err, writer); return }
	json.Unmarshal(requestData, &newCategoryName)
	
	// query database
	_, err = db.Exec(`INSERT INTO categories (name, added)
	VALUES ($1, now())`, newCategoryName)
	if err != nil { data.LogErr(err, writer) }
}


func putCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var categoryData Category
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.LogErr(err, writer); return }
	json.Unmarshal(requestData, &categoryData)
	
	// edit category
	_, err = db.Exec(`UPDATE categories SET name = $1
	WHERE id = $2`, categoryData.Name, categoryData.Id) 
	if err != nil { data.LogErr(err, writer) }
}


func deleteCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var id int
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.LogErr(err, writer); return }
	json.Unmarshal(requestData, &id)
	
	// delete query 
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil { data.LogErr(err, writer) }
}