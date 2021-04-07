package main

import (
	"fmt"
	"net/http"
	"minera/data"
	"minera/routes"
)

func main() {
	editorStaticFiles := http.StripPrefix("/static/editor/", http.FileServer(http.Dir("./static/editor/")))
	http.Handle("/static/editor/", editorStaticFiles)

	catalogStaticFiles := http.StripPrefix("/static/catalog/", http.FileServer(http.Dir("./static/catalog/")))
	http.Handle("/static/catalog/", catalogStaticFiles)

	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images/")))
	http.Handle("/images/", images)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", routes.Catalog)
	http.HandleFunc("/editor/", routes.Editor)
	http.HandleFunc("/authentication", routes.Authentication)
	http.HandleFunc("/files", routes.FileTransfer)

	fmt.Println("Now serving @ " + data.Address)
	http.ListenAndServe(data.Address, nil)
}
