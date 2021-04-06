package methods

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"minera/data"
)

func getSubCategories(db *sql.DB, writer http.ResponseWriter, categoryId int) {
	// query database
	rows, err := db.Query(`SELECT id, name FROM sub_categories
	WHERE category_id = $1 ORDER BY id ASC`, categoryId)
	if err != nil { data.Log(err, writer); return }
	defer rows.Close()

	// package data
	var subCategories []SubCategory
	for rows.Next() {
		sub := SubCategory{}
		err = rows.Scan(&sub.Id, &sub.Name)
		if err != nil { data.Log(err, writer); return }
		subCategories = append(subCategories, sub)
	}

	// get parent name
	var categoryName string
	err = db.QueryRow(`SELECT name FROM categories WHERE id = $1`, categoryId).Scan(&categoryName)
	if err != nil { data.Log(err, writer); return }

	// return template
	temp := SubTempData{categoryId, categoryName, subCategories} 
	err = data.EditorTemplates.ExecuteTemplate(writer, "subCategories.html", temp)
	if err != nil { data.Log(err, writer) }
}


func postSubCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request, categoryId int) {
	// get request data
	var newSubCategoryName string
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &newSubCategoryName)
	
	// query database
	_, err = db.Exec(`INSERT INTO sub_categories (category_id, name, added)
	VALUES ($1, $2, now())`, categoryId, newSubCategoryName)
	if err != nil { data.Log(err, writer) }
}


func putSubCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var SubCategoryData SubCategory
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &SubCategoryData)

	// edit category
	_, err = db.Exec(`UPDATE sub_categories SET name = $1
	WHERE id = $2`, SubCategoryData.Name, SubCategoryData.Id) 
	if err != nil { data.Log(err, writer) }
}


func deleteSubCategory(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var id int
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &id)
	
	// delete query
	_, err = db.Exec(`DELETE FROM sub_categories WHERE id = $1`, id) 
	if err != nil { data.Log(err, writer) }
}
