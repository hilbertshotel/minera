package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return }
	json.Unmarshal(request, &id)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()
	
	// edit category
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil { ErrorLogger.Println(err); return }
}
