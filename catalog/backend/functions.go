package backend

import (
	"database/sql"
	_ "github.com/lib/pq"
)

////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// GLOBAL ////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

func get_categories() ([]Category, error) {
	empty_slice := []Category{}
	
	// connect to database
	db, err := sql.Open("postgres", connection_string)
	if err != nil { 
		Logger.Println(err)
		return empty_slice, err
	}
	defer db.Close()

	// response data
	query := `SELECT name, path, function
				FROM categories
				WHERE function = 'get_sub_categories'
				ORDER BY added ASC`
	rows, err := db.Query(query)
	if err != nil {
		Logger.Println(err)
		return empty_slice, err
	}
	defer rows.Close()

	// package categories
	var categories []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Name, &category.Path, &category.Function)
		if err != nil {
			Logger.Println(err)
			return empty_slice, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}


////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// CATEGORIES //////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

func get_sub_categories(parent string) ([]SubCategory, error, bool) {
	empty_slice := []SubCategory{}

	// connect to database
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		Logger.Println(err)
		return empty_slice, err, false
	}
	defer db.Close()

	// query database
	query := `SELECT name, path, function
				FROM sub_categories
				WHERE parent = $1
				AND function = 'get_products'
				ORDER BY added ASC`
	rows, err := db.Query(query, parent)
	if err != nil {
		Logger.Println(err)
		return empty_slice, nil, false
	}

	// package data
	var sub_categories []SubCategory
	for rows.Next() {
		sub_category := SubCategory{}
		err = rows.Scan(&sub_category.Name, &sub_category.Path, &sub_category.Function)
		if err != nil {
			Logger.Println(err)
			return empty_slice, err, false
		}
		sub_categories = append(sub_categories, sub_category)
	}

	return sub_categories, nil, true
}


////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// SUB CATEGORIES ////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////



////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////// PRODUCTS ///////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////