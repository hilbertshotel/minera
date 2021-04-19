package methods

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"github.com/lib/pq"
)

func GetProducts(
	log *log.Logger,
	db *sql.DB,
	catId,
	subId int) (ProductTempData, error) {
	
	var productsData ProductTempData
	
	// query database
	rows, err := db.Query(`SELECT id, name, description, images FROM products
	WHERE sub_category_id = $1 ORDER BY id ASC`, subId)
	if err != nil {
		log.Println("ERROR:", err)
		return productsData, err
	}
	defer rows.Close()

	// package data
	var products []Product
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, pq.Array(&product.Images))
		if err != nil {
			log.Println("ERROR:", err)
			return productsData, err
		}
		products = append(products, product)
	}

	// get parent name
	var subCategoryName string
	err = db.QueryRow(`SELECT name FROM sub_categories
	WHERE id = $1`, subId).Scan(&subCategoryName)
	if err != nil {
		log.Println("ERROR:", err)
		return productsData, err
	}

	productsData = ProductTempData{catId, subId, subCategoryName, products}
	return productsData, nil
}


func postProduct(db *sql.DB, r *http.Request, log *log.Logger, imgDir string) error {
	// get request data
	var newProduct Product
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &newProduct)

	// add folder path to images
	var images []string
	for _, imgName := range newProduct.Images {
		images = append(images, imgDir + imgName)
	}

	// post data
	_, err = db.Exec(`INSERT INTO products (sub_category_id, name, description, images, added)
	VALUES ($1, $2, $3, $4, now())`, newProduct.Id, newProduct.Name, newProduct.Description, pq.Array(images))
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func putProduct(db *sql.DB, r *http.Request, log *log.Logger, imgDir string) error {
	// get request data
	var productData Product
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &productData)

	// edit product without new images
	if len(productData.Images) == 0 {
		_, err = db.Exec(`UPDATE products SET name = $1, description = $2
		WHERE id = $3`, productData.Name, productData.Description, productData.Id) 
		if err != nil {
			log.Println("ERROR:", err)
			return err
		}

		return nil
	}

	// add folder path to images
	var images []string
	for _, imgName := range productData.Images {
		images = append(images, imgDir + imgName)
	}

	// edit product with new images
	_, err = db.Exec(`UPDATE products SET name = $1, description = $2, images = $3
	WHERE id = $4`, productData.Name, productData.Description, pq.Array(images), productData.Id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}


func deleteProduct(db *sql.DB, r *http.Request, log *log.Logger) error {
	// get request data
	var id int
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	json.Unmarshal(request, &id)

	// delete data
	_, err = db.Exec(`DELETE FROM products WHERE id = $1`, id) 
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}
