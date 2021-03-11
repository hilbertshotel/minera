package main

import (
	"fmt"
	"net/http"
	"api/handlers"
)

var IP = "127.0.0.1:5151"

func main() {
	// catalog handlers
	http.HandleFunc("/LoadCategories", handlers.LoadCategories)
	http.HandleFunc("/LoadItems", handlers.LoadItems)

	// editor handlers
	http.HandleFunc("/VerifyPassword", handlers.VerifyPassword)

	fmt.Println("Now serving @", IP)
	http.ListenAndServe(IP, nil)
}
