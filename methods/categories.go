package methods

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
)

func GetCategories(log *log.Logger, db *sql.DB) ([]Category, error) {
	// query database
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY id ASC")
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer rows.Close()

	// package data
	var categories []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			log.Println("ERROR:", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	
	return categories, nil
}


func postCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var newCategoryName string
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &newCategoryName)
	
	// post data
	_, err = db.Exec(`INSERT INTO categories (name, added)
	VALUES ($1, now())`, newCategoryName)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func putCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var categoryData Category
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &categoryData)
	
	// edit category
	_, err = db.Exec(`UPDATE categories SET name = $1
	WHERE id = $2`, categoryData.Name, categoryData.Id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func deleteCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &id)
	
	// delete data 
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}