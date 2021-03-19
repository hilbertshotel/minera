package categories

import (
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"minera/backend/utils"
)

func Get(w http.ResponseWriter) {
	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()

	// query database
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY	id ASC")
	if err != nil { utils.Logger.Println(err); return }
	defer rows.Close()

	// insert data into Category struct
	var response []Category
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil { utils.Logger.Println(err); return }
		response = append(response, category)
	}

	// return response to frontend
	w.Header().Set("content-type", "application/json")
  	output, err := json.Marshal(response)
  	if err != nil { utils.Logger.Println(err); return }
  	w.Write(output)
}


func Post(w http.ResponseWriter, r *http.Request) {
	// get request data
	var newCategoryName string
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { utils.Logger.Println(err); return }
	json.Unmarshal(request, &newCategoryName)

	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()
	
	// query database
	_, err = db.Exec(`INSERT INTO categories (name)
	VALUES ($1)`, newCategoryName)
	if err != nil { utils.Logger.Println(err); return }
}


func Put(w http.ResponseWriter, r *http.Request) {
	// get request data
	var data PutCategory
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { utils.Logger.Println(err); return }
	json.Unmarshal(request, &data)

	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()
	
	// edit category
	_, err = db.Exec(`UPDATE categories SET name = $1
	WHERE id = $2`, data.NewName, data.Id) 
	if err != nil { utils.Logger.Println(err); return }
}


func Delete(w http.ResponseWriter, r *http.Request) {
	// get request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil { utils.Logger.Println(err); return }
	json.Unmarshal(request, &id)

	// connect to database
	db, err := sql.Open("postgres", utils.ConnStr)
	if err != nil { utils.Logger.Println(err); return }
	defer db.Close()
	
	// edit category
	_, err = db.Exec(`DELETE FROM items WHERE category_id = $1`, id) 
	if err != nil { utils.Logger.Println(err); return }
	_, err = db.Exec(`DELETE FROM categories WHERE id = $1`, id) 
	if err != nil { utils.Logger.Println(err); return }
}
