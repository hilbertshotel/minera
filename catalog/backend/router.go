package backend

import (
	"log"
	"net/http"
)

func Router(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	
	path := request.URL.Path
	if path == "/" {
		err := templates.ExecuteTemplate(writer, "categories.html", get_categories())
		if err != nil { log.Println(err) } // LOG
		return
	}

	function_name, section, id, ok := search_path(path)
	if !ok {
		http.NotFound(writer, request)
		return
	}

	function := function_map[function_name]
	function()
	// id := 7
	// section := "sub"
	// if section == "sub" {
	// 	// SELECT name, path FROM sub_categories WHERE id = id;
	// 	data := function()
	// 	err := templates.ExecuteTemplate(writer, "sub_categories.html", data) // replace data with function
	// 	if err != nil { log.Println(err) } // LOG
	// 	return
	// } 

	// // SELECT name, description, images FROM items WHERE id = id;
	// data := function()
	// err := templates.ExecuteTemplate(writer, "items.html", data) // replace data with function
	// if err != nil { log.Println(err) } // LOG
}

