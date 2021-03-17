package main

import (
	"fmt"
	"net/http"
	"editor/handlers"
)

var IP = "127.0.0.1:5252"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/LoadCategories", handlers.LoadCategories)
	http.HandleFunc("/LoadItems", handlers.LoadItems)
	http.HandleFunc("/VerifyUser", handlers.VerifyUser)
	http.HandleFunc("/NewCategory", handlers.NewCategory)
	http.HandleFunc("/EditCategoryName", handlers.EditCategoryName)
	http.HandleFunc("/DeleteCategory", handlers.DeleteCategory)
	http.HandleFunc("/NewItem", handlers.NewItem)
	http.HandleFunc("/NewItemImages", handlers.NewItemImages)

	fmt.Println("Now serving @", IP)
	http.ListenAndServe(IP, nil)
}
