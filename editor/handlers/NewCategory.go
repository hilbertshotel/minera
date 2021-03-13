package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

func NewCategory(w http.ResponseWriter, r *http.Request) {
	// get request data
	var newCategoryName string
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return }
	json.Unmarshal(request, &newCategoryName)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()
	
	// query database
	_, err = db.Exec(`INSERT INTO categories (name)
	VALUES ($1)`, newCategoryName)
	if err != nil { ErrorLogger.Println(err); return }
}
