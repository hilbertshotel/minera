package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

type EditCategoryData struct {
	Id int `json:id`
	NewName string `json:newName`
}

func EditCategoryName(w http.ResponseWriter, r *http.Request) {
	// get request data
	var data EditCategoryData
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return }
	json.Unmarshal(request, &data)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()
	
	// edit category
	_, err = db.Exec(`UPDATE categories SET name = $1
	WHERE id = $2`, data.NewName, data.Id) 
	if err != nil { ErrorLogger.Println(err); return }

	// get edited category data
	var category Category
	err = db.QueryRow(`SELECT id, name FROM categories
	WHERE id = $1`, data.Id).Scan(&category.Id, &category.Name)
	if err != nil { ErrorLogger.Println(err); return }

	// return response
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(category)
	if err != nil { ErrorLogger.Println(err); return }
	w.Write(output)
}
