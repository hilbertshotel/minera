package handlers

import (
	"io"
	"os"
	"net/http"
	"io/ioutil"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
)

func NewItemImages(w http.ResponseWriter, r *http.Request) {
	// parse files
	r.ParseMultipartForm(32 << 20) 
	files := r.MultipartForm.File["files"]
	
	var response string
	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil { ErrorLogger.Println(err); return }
		defer file.Close()

		// VERIFY IF IMAGE IS ACTUALLY IN IMAGE

		out, err := os.Create(IMAGEDIR + files[i].Filename) // PERSONALIZE FILENAMES
		if err != nil { ErrorLogger.Println(err); return }

		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil { ErrorLogger.Println(err); return }
	}
}


func NewItem(w http.ResponseWriter, r *http.Request) {
		// get request data
		var newItem Item
		request, err := ioutil.ReadAll(r.Body)
		if err != nil { ErrorLogger.Println(err); return }
		json.Unmarshal(request, &newItem)
	
		// connect to database
		db, err := sql.Open("postgres", connStr)
		if err != nil { ErrorLogger.Println(err); return }
		defer db.Close()
		
		// add path to images
		var images []string
		for _, img := range newItem.Images {
			images = append(images, "images/" + img)
		}

		// query database
		_, err = db.Exec(`INSERT INTO items (category_id, name, description, images)
		VALUES ($1, $2, $3, $4)`, newItem.Id, newItem.Name, newItem.Description, pq.Array(images))
		if err != nil { ErrorLogger.Println(err); return }
}