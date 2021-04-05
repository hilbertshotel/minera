package methods

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"minera/data"
	"minera/logs"
)

func getSubCategories(writer http.ResponseWriter, categoryId int) {
	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// query database
	rows, err := db.Query(`SELECT id, name FROM sub_categories
	WHERE category_id = $1 ORDER BY id ASC`, categoryId)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer rows.Close()

	// package data
	var subCategories []SubCategory
	for rows.Next() {
		sub := SubCategory{}
		err = rows.Scan(&sub.Id, &sub.Name)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		subCategories = append(subCategories, sub)
	}

	// get parent name
	var categoryName string
	err = db.QueryRow(`SELECT name FROM categories WHERE id = $1`, categoryId).Scan(&categoryName)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	// return template
	temp := SubTempData{categoryId, categoryName, subCategories} 
	err = data.EditorTemplates.ExecuteTemplate(writer, "subCategories.html", temp)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func postSubCategory(writer http.ResponseWriter, request *http.Request, categoryId int) {
	// get request data
	var newSubCategoryName string
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &newSubCategoryName)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()
	
	// query database
	_, err = db.Exec(`INSERT INTO sub_categories (category_id, name, added)
	VALUES ($1, $2, now())`, categoryId, newSubCategoryName)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func putSubCategory(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var SubCategoryData SubCategory
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &SubCategoryData)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// edit category
	_, err = db.Exec(`UPDATE sub_categories SET name = $1
	WHERE id = $2`, SubCategoryData.Name, SubCategoryData.Id) 
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func deleteSubCategory(writer http.ResponseWriter, request *http.Request) {
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
	_, err = db.Exec(`DELETE FROM sub_categories WHERE id = $1`, id) 
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}
