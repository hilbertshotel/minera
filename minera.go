package main

import (
	"fmt"
	"net/http"
	"minera/backend"
)

func main() {
	editor_static_files := http.StripPrefix("/static/editor/", http.FileServer(http.Dir("./static/editor/")))
	http.Handle("/static/editor/", editor_static_files)

	http.HandleFunc("/editor/", backend.Editor)
	http.HandleFunc("/authentication", backend.Authentication)

	http.Handle("favicon.ico", http.NotFoundHandler())

	fmt.Println("Now serving @ " + backend.Address)
	http.ListenAndServe(backend.Address, nil)
}
