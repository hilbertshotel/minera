package methods

import "net/http"

func CategoriesDispatcher(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case http.MethodGet: getCategories(writer)
		case http.MethodPost: postCategory(writer, request)
		case http.MethodPut: putCategory(writer, request)
		case http.MethodDelete: deleteCategory(writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}

func SubCategoriesDispatcher(writer http.ResponseWriter, request *http.Request, id int) {
	switch request.Method {
		case http.MethodGet: getSubCategories(writer, id)
		case http.MethodPost: postSubCategory(writer, request, id)
		case http.MethodPut: putSubCategory(writer, request)
		case http.MethodDelete: deleteSubCategory(writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}

func ProductsDispatcher(writer http.ResponseWriter, request *http.Request, categoryId int, subCategoryId int) {
	switch request.Method {
		case http.MethodGet: getProducts(writer, categoryId, subCategoryId)
		case http.MethodPost: postProduct(writer, request)
		// case http.MethodPut: put(writer, request)
		// case http.MethodDelete: delete(writer, request)
		default: http.Error(writer, "Възникна грешка", 405)
	}
}
