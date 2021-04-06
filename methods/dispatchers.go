package methods

import (
	"net/http"
	"minera/data"
	"database/sql"
)


func CategoriesDispatcher(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	switch request.Method {

		case http.MethodGet:
			categories, err := GetCategories(db, writer)
			if err != nil { return }
			err = data.EditorTemplates.ExecuteTemplate(writer, "categories.html", categories)
			if err != nil { data.Log(err, writer) }

		case http.MethodPost: postCategory(db, writer, request)
		case http.MethodPut: putCategory(db, writer, request)
		case http.MethodDelete: deleteCategory(db, writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}


func SubCategoriesDispatcher(db *sql.DB, writer http.ResponseWriter, request *http.Request, categoryId int) {
	switch request.Method {
		case http.MethodGet: getSubCategories(db, writer, categoryId)
		case http.MethodPost: postSubCategory(db, writer, request, categoryId)
		case http.MethodPut: putSubCategory(db, writer, request)
		case http.MethodDelete: deleteSubCategory(db, writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}


func ProductsDispatcher(db *sql.DB, writer http.ResponseWriter, request *http.Request, categoryId int, subCategoryId int) {
	switch request.Method {
		case http.MethodGet: getProducts(db, writer, categoryId, subCategoryId)
		case http.MethodPost: postProduct(db, writer, request)
		// case http.MethodPut: put(writer, request)
		case http.MethodDelete: deleteProduct(db, writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}
