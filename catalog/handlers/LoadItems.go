package handlers

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
)

func LoadItems(w http.ResponseWriter, r *http.Request) {
	// pickup request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { ErrorLogger.Println(err); return } 
	json.Unmarshal(request, &id)

	// connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil { ErrorLogger.Println(err); return }
	defer db.Close()

	// query database
	rows, err := db.Query(`SELECT name, description, images
	FROM items WHERE category_id = $1 ORDER BY id ASC`, id)
	if err != nil { ErrorLogger.Println(err); return }
	defer rows.Close()

	// insert data into Item struct
	var response []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Name,	&item.Description, pq.Array(&item.Images))
		if err != nil { ErrorLogger.Println(err); return }
		response = append(response, item)
	}

	// return response
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { ErrorLogger.Println(err); return }
	w.Write(output)
}

