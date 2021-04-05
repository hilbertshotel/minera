package methods

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"github.com/lib/pq"
	"minera/data"
	"minera/logs"
)

func getProducts(writer http.ResponseWriter, categoryId, subCategoryId int) {
	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()

	// query database
	rows, err := db.Query(`SELECT id, name, description, images FROM products
	WHERE sub_category_id = $1 ORDER BY id ASC`, subCategoryId)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer rows.Close()

	// package data
	var products []Product
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, pq.Array(&product.Images))
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		products = append(products, product)
	}

	// get parent name
	var subCategoryName string
	err = db.QueryRow(`SELECT name FROM sub_categories
	WHERE id = $1`, subCategoryId).Scan(&subCategoryName)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	// return template
	temp := ProductTempData{categoryId, subCategoryId, subCategoryName, products} 
	err = data.EditorTemplates.ExecuteTemplate(writer, "products.html", temp)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}


func postProduct(writer http.ResponseWriter, request *http.Request) {
	// get request data
	var newProduct Product
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	json.Unmarshal(requestData, &newProduct)

	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}
	defer db.Close()
	
	// add path to images
	var images []string
	for _, img := range newProduct.Images {
		images = append(images, "images/" + img)
	}

	// query database
	_, err = db.Exec(`INSERT INTO products (sub_category_id, name, description, images, added)
	VALUES ($1, $2, $3, $4, now())`, newProduct.Id, newProduct.Name, newProduct.Description, pq.Array(images))
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
	}
}