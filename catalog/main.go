package main

import (
	"fmt"
	"net/http"
	"catalog/handlers"
)

var IP = "127.0.0.1:5151"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/LoadCategories", handlers.LoadCategories)
	http.HandleFunc("/LoadItems", handlers.LoadItems)

	fmt.Println("Now serving @", IP)
	http.ListenAndServe(IP, nil)
}
