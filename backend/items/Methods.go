package items

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"minera/backend/utils"
)


func Get(w http.ResponseWriter, r *http.Request, id int) {
	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()

	// query database
	rows, err := db.Query(`SELECT id, name, description, images
	FROM items WHERE category_id = $1 ORDER BY id ASC`, id)
	if err != nil { utils.Logger.Println(err); return }
	defer rows.Close()

	// insert data into Item struct
	var response []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Id, &item.Name, &item.Description, pq.Array(&item.Images))
		if err != nil { utils.Logger.Println(err); return }
		response = append(response, item)
	}

	// return response
	w.Header().Set("content-type", "application/json")
	output, err := json.Marshal(response)
	if err != nil { utils.Logger.Println(err); return }
	w.Write(output)
}


func Post(w http.ResponseWriter, r *http.Request) {
}


func Put(w http.ResponseWriter, r *http.Request) {
}


func Delete(w http.ResponseWriter, r *http.Request) {
}
