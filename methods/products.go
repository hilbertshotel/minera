package methods

import (
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"github.com/lib/pq"
	"minera/data"
)

func GetProducts(db *sql.DB, writer http.ResponseWriter, categoryId, subCategoryId int) (ProductTempData, error) {
	var productsData ProductTempData
	
	// query database
	rows, err := db.Query(`SELECT id, name, description, images FROM products
	WHERE sub_category_id = $1 ORDER BY id ASC`, subCategoryId)
	if err != nil {
		data.Log(err, writer)
		return productsData, err
	}
	defer rows.Close()

	// package data
	var products []Product
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, pq.Array(&product.Images))
		if err != nil {
			data.Log(err, writer)
			return productsData, err
		}
		products = append(products, product)
	}

	// get parent name
	var subCategoryName string
	err = db.QueryRow(`SELECT name FROM sub_categories
	WHERE id = $1`, subCategoryId).Scan(&subCategoryName)
	if err != nil {
		data.Log(err, writer)
		return productsData, err
	}

	productsData = ProductTempData{categoryId, subCategoryId, subCategoryName, products}
	return productsData, nil
}


func postProduct(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var newProduct Product
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &newProduct)

	// add path to images
	var images []string
	for _, imgName := range newProduct.Images {
		images = append(images, data.ImageDir + imgName)
	}

	// query database
	_, err = db.Exec(`INSERT INTO products (sub_category_id, name, description, images, added)
	VALUES ($1, $2, $3, $4, now())`, newProduct.Id, newProduct.Name, newProduct.Description, pq.Array(images))
	if err != nil { data.Log(err, writer) }
}


func putProduct(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var productData Product
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &productData)

	// edit product
	_, err = db.Exec(`UPDATE products SET name = $1, description = $2
	WHERE id = $3`, productData.Name, productData.Description, productData.Id) 
	if err != nil { data.Log(err, writer) }
}


func deleteProduct(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get request data
	var id int
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil { data.Log(err, writer); return }
	json.Unmarshal(requestData, &id)
	
	// delete images from image folder
	var images []string
	err = db.QueryRow(`SELECT images FROM products
	WHERE id = $1`, id).Scan(pq.Array(&images))
	if err != nil { data.Log(err, writer); return }

	folderList, err := data.ListFolder(data.ImageDir)
	if err != nil { data.Log(err, writer); return }

	for _, image := range images {
		filename := image[7:]
		if data.Contains(folderList, filename) {
			err := os.Remove(data.ImageDir + filename)
			if err != nil { data.Log(err, writer); return }
		}
	}

	// delete query
	_, err = db.Exec(`DELETE FROM products WHERE id = $1`, id) 
	if err != nil { data.Log(err, writer) }
}
