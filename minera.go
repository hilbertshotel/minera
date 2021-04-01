package main

import (
	"fmt"
	"net/http"
	"minera/data"
	"minera/routes"
)

func main() {
	editor_static_files := http.StripPrefix("/static/editor/", http.FileServer(http.Dir("./static/editor/")))
	http.Handle("/static/editor/", editor_static_files)

	http.HandleFunc("/", routes.Catalog)
	http.HandleFunc("/editor/", routes.Editor)
	http.HandleFunc("/authentication", routes.Authentication)

	http.Handle("favicon.ico", http.NotFoundHandler())

	fmt.Println("Now serving @ " + data.Address)
	http.ListenAndServe(data.Address, nil)
}
