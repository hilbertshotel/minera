package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const IP = "127.0.0.1:8000"
var templates = template.Must(template.ParseGlob("templates/*"))

type Category struct {
	Name string
	Path string
}

type Item struct {
	Name string
	Description string
	Images []string
}


func init() {
	// 1 RESET ALL PATHS IN DATABASE WITH NEW RANDOM STRING VALUES
	// 2 INIT ALL PATHS DYNAMICALLY


	// SELECT path FROM paths;
	paths := []string{"/345", "/127", "/325"} 
	for _, path := range paths {
		http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

			// METHOD MATCHER

			// IF DELETE MAKE NEW dynamic route WITH SAME ADDRESS BUT RETURNS 404
			// RESTART API SERVICE ONCE A NIGHT

			// IF GET
			// SELECT child, id_list FROM links WHERE path = path;
			child := "SUB"
			// id_list := []int{ 1, 2, 3, 4, 5 }

			switch child {
			case "SUB":
				// SELECT name, path FROM sub WHERE.. FOR LOOP id_list
				sub_categories := []Category{
					Category{ Name: "Sub Category 1", Path: "http://" + IP + "/771" },
					Category{ Name: "Sub Category 2", Path: "http://" + IP + "/379" },
					Category{ Name: "Sub Category 3", Path: "http://" + IP + "/694" },
				}
				
				writer.Header().Set("Content-Type", "text/html")
				err := templates.ExecuteTemplate(writer, "index.html", sub_categories)
				if err != nil { fmt.Println(err) } // LOG

			case "ITEM":
				// SELECT name, description, images FROM items WHERE .. FOR LOOP id_list
				items := []Item{
					Item{ Name: "Item 1", Description: "Description 1", Images: []string{} },
					Item{ Name: "Item 2", Description: "Description 2", Images: []string{} },
					Item{ Name: "Item 3", Description: "Description 3", Images: []string{} },
				}
				
				writer.Header().Set("Content-Type", "text/html")
				err := templates.ExecuteTemplate(writer, "items.html", items)
				if err != nil { fmt.Println(err) } // LOG
			}
		})
	}
}


func index(writer http.ResponseWriter, request *http.Request) {
	// SELECT name, path FROM categories;
	categories := []Category{
		Category{ Name: "Category 1", Path: "http://" + IP + "/345" },
		Category{ Name: "Category 2", Path: "http://" + IP + "/127" },
		Category{ Name: "Category 3", Path: "http://" + IP + "/325" },
	}

	writer.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(writer, "index.html", categories)
	if err != nil { fmt.Println(err) } // LOG
}


func router(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		index(writer, request)
		return
	}
	http.NotFound(writer, request)
}


func main() {
	static_files := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	http.Handle("/static/", static_files)

	http.HandleFunc("/", router)

	fmt.Println("Now serving @ " + IP)
	http.ListenAndServe(IP, nil)
}
