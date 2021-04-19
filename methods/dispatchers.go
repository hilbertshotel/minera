package methods

import (
	"net/http"
	"database/sql"
	"log"
	"html/template"
	"errors"
)

func CategoriesDispatcher(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	db *sql.DB,
	tmp *template.Template) error {

	switch r.Method {

		case http.MethodGet:
			categories, err := GetCategories(log, db)
			if err != nil {
				return err
			}

			if err := tmp.ExecuteTemplate(w, "categories.html", categories); err != nil {
				log.Println("ERROR:", err)
				return err
			}

		case http.MethodPost:
			return postCategory(db, r, log)
		case http.MethodPut:
			return putCategory(db, r, log)
		case http.MethodDelete:
			return deleteCategory(db, r, log)
		default:
			log.Printf("ERROR: unsupported method `%v`\n", r.Method)
			return errors.New("err")
	}

	return nil
}


func SubCategoriesDispatcher(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	db *sql.DB,
	tmp *template.Template,
	catId int) error {

	switch r.Method {
		
		case http.MethodGet:
			subCategories, err := GetSubCategories(log, db, catId)
			if err != nil {
				return err
			}

			err = tmp.ExecuteTemplate(w, "sub_categories.html", subCategories)
			if err != nil {
				log.Println("ERROR:", err)
				return err
			}

		case http.MethodPost:
			return postSubCategory(db, r, log)
		case http.MethodPut:
			return putSubCategory(db, r, log)
		case http.MethodDelete:
			return deleteSubCategory(db, r, log)
		default:
			log.Printf("ERROR: unsupported method `%v`\n", r.Method)
			return errors.New("err")
	}

	return nil
}


func ProductsDispatcher(
	w http.ResponseWriter,
	r *http.Request,
	log *log.Logger,
	db *sql.DB,
	tmp *template.Template,
	catId int,
	subId int,
	imgDir string) error {

	switch r.Method {

		case http.MethodGet:
			products, err := GetProducts(log, db, catId, subId)
			if err != nil {
				return err
			}

			err = tmp.ExecuteTemplate(w, "products.html", products)
			if err != nil {
				log.Println("ERROR:", err)
				return err
			}

		case http.MethodPost:
			return postProduct(db, r, log, imgDir)
		case http.MethodPut: 
			return putProduct(db, r, log, imgDir)
		case http.MethodDelete: 
			return deleteProduct(db, r, log)
		default: 	
			log.Printf("ERROR: unsupported method `%v`\n", r.Method)
			errors.New("err")
	}

	return nil
}
