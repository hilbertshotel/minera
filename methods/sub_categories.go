package methods

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
)

func GetSubCategories(
	log *log.Logger,
	db *sql.DB,
	catId int) (SubTempData, error) {
	
	var subCategoriesData SubTempData
	
	// query database
	rows, err := db.Query(`SELECT id, name FROM sub_categories
	WHERE category_id = $1 ORDER BY id ASC`, catId)
	if err != nil {
		log.Println("ERROR:", err)
		return subCategoriesData, err
	}
	defer rows.Close()

	// package data
	var subCategories []SubCategory
	for rows.Next() {
		sub := SubCategory{}
		err = rows.Scan(&sub.Id, &sub.Name)
		if err != nil {
			log.Println("ERROR:", err)
			return subCategoriesData, err
		}
		subCategories = append(subCategories, sub)
	}

	// get parent name
	var categoryName string
	err = db.QueryRow(`SELECT name FROM categories WHERE id = $1`, catId).Scan(&categoryName)
	if err != nil {
		log.Println("ERROR:", err)
		return subCategoriesData, err
	}

	subCategoriesData = SubTempData{catId, categoryName, subCategories}
	return subCategoriesData, nil
}


func postSubCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var newSubCategory SubCategory
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &newSubCategory)
	
	// post data
	_, err = db.Exec(`INSERT INTO sub_categories (category_id, name, added)
	VALUES ($1, $2, now())`, newSubCategory.Id, newSubCategory.Name)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func putSubCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var subCategoryData SubCategory
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &subCategoryData)

	// edit data
	_, err = db.Exec(`UPDATE sub_categories SET name = $1
	WHERE id = $2`, subCategoryData.Name, subCategoryData.Id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func deleteSubCategory(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &id)
	
	// delete data
	_, err = db.Exec(`DELETE FROM sub_categories WHERE id = $1`, id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}
