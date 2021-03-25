package backend

import (
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	
	path := request.URL.Path
	if path != "/" {
		http.NotFound(writer, request)
		return
	}

	categories, err := get_categories()
	if err != nil {
		http.Error(writer, "Bad Gateway", 502)
		return
	}
		
	err = templates.ExecuteTemplate(writer, "categories.html", categories)
	if err != nil {
		Logger.Println(err)
		http.Error(writer, "Bad Gateway", 502)
	}
}


func Categories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	
	path := request.URL.Path
	if path == "/cat/" {
		http.NotFound(writer, request)
		return
	}
	
	sub_categories, err, exists := get_sub_categories(path)
	if err != nil {
		http.Error(writer, "Bad Gateway", 502)
		return
	}
	if !exists {
		http.NotFound(writer, request)
		return
	}
	
	err = templates.ExecuteTemplate(writer, "sub_categories.html", sub_categories)	
	if err != nil {
		Logger.Println(err)
		http.Error(writer, "Bad Gateway", 502)
	}
}


func SubCategories(writer http.ResponseWriter, request *http.Request) {

}


func Products(writer http.ResponseWriter, request *http.Request) {

}

