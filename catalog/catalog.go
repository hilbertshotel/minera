package main

import (
	"fmt"
	"net/http"
	"catalog/backend"
)

func main() {
	static_files := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	http.Handle("/static/", static_files)

	http.HandleFunc("/", backend.Router)

	fmt.Println("Now serving @ " + backend.Address)
	http.ListenAndServe(backend.Address, nil)
}
