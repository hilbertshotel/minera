package main

import (
	"fmt"
	"net/http"
	"minera/backend"
)

var ADDRESS = "127.0.0.1:5252"

func main() {
	catalog := http.FileServer(http.Dir("./frontend/catalog"))
	http.Handle("/", catalog)

	editor := http.StripPrefix("/editor/", http.FileServer(http.Dir("./frontend/editor")))
	http.Handle("/editor/", editor)

	test := http.StripPrefix("/test/", http.FileServer(http.Dir("./frontend/test")))
	http.Handle("/test/", test)

	http.HandleFunc("/Catalog/", backend.Catalog)
	http.HandleFunc("/editor/Editor/", backend.Editor)
	http.HandleFunc("/editor/Authentication", backend.Authentication)
	http.HandleFunc("/editor/FileTransfer", backend.FileTransfer)

	fmt.Println("Now serving @", ADDRESS)
	http.ListenAndServe(ADDRESS, nil)
}
