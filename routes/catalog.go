package routes

import (
	"net/http"
	"database/sql"
	"strings"
	"strconv"
	"minera/data"
	"minera/methods"
)

func Catalog(writer http.ResponseWriter, request *http.Request) {
	// connect to database
	db, err := sql.Open("postgres", data.ConnectionString)
	if err != nil { data.LogErr(err, writer); return }
	defer db.Close()
	
	url := request.URL.Path[1:]

	// handle categories
	if url == "" {
		categories, err := methods.GetCategories(db, writer)
		if err != nil { return }
		
		err = data.CatalogTemplates.ExecuteTemplate(writer, "categories.html", categories)
		if err != nil { data.LogErr(err, writer) }

		return
	}

	urlArray := strings.Split(url, "/")

	// handle sub categories
	if len(urlArray) == 1 {
		categoryId, err := strconv.Atoi(urlArray[0])
		if err != nil { data.LogRequest(writer, request); return }

		subCategories, err := methods.GetSubCategories(db, writer, categoryId)
		if err != nil { return }

		err = data.CatalogTemplates.ExecuteTemplate(writer, "sub_categories.html", subCategories)
		if err != nil { data.LogErr(err, writer) }

		return
	}

	// handle products
	if len(urlArray) == 2 {
		categoryId, err := strconv.Atoi(urlArray[0])
		if err != nil { data.LogRequest(writer, request); return }

		subCategoryId, err := strconv.Atoi(urlArray[1])
		if err != nil { data.LogRequest(writer, request); return }
	
		products, err := methods.GetProducts(db, writer, categoryId, subCategoryId)
		if err != nil { return }

		err = data.CatalogTemplates.ExecuteTemplate(writer, "products.html", products)
		if err != nil { data.LogErr(err, writer) }
		
		return
	}

	http.Error(writer, "Страницата не съществува", 404)
}
