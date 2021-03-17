package handlers

import (
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

func LoadCategories(w http.ResponseWriter, r *http.Request) {
	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()

	// query database
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY	id ASC")
	if err != nil { ErrorLogger.Println(err); return }
	defer rows.Close()

	// insert data into Category struct
	var response []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil { ErrorLogger.Println(err); return }
		response = append(response, category)
	}

	// return response to frontend
	w.Header().Set("content-type", "application/json")
  	output, err := json.Marshal(response)
  	if err != nil { ErrorLogger.Println(err); return }
  	w.Write(output)
}