package routes

import (
	"net/http"
	"database/sql"
	// "strings"
	// "strconv"
	"minera/data"
	"minera/methods"
)

func Catalog(writer http.ResponseWriter, request *http.Request) {
	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil { data.Log(err, writer); return }
	defer db.Close()
	
	// handle categories
	url := request.URL.Path[1:]
	if url == "" {
		categories, err := methods.GetCategories(db, writer)
		if err != nil { return }
		
		err = data.CatalogTemplates.ExecuteTemplate(writer, "categories.html", categories)
		if err != nil { data.Log(err, writer) }

		return
	}

	// // handle sub categories
	// urlArray := strings.Split(url, "/")
	// if len(urlArray) == 2 {
	// 	id, err := strconv.Atoi(urlArray[1])
	// 	if err != nil { data.Log(err, writer); return }
	// 	methods.SubCategoriesDispatcher(writer, request, id)
	// 	return
	// }

	// // handle products
	// if len(urlArray) == 3 {
	// 	categoryId, err := strconv.Atoi(urlArray[1])
	// 	if err != nil { data.Log(err, writer); return }

	// 	subCategoryId, err := strconv.Atoi(urlArray[2])
	// 	if err != nil { data.Log(err, writer); return }
	
	// 	methods.ProductsDispatcher(writer, request, categoryId, subCategoryId)
	// 	return
	// }

	http.Error(writer, "Страницата не съществува", 404)
}
