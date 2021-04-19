package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"os"
	"minera/routes"
	"minera/conf"
)

func main() {
	if err := service(); err != nil {
		log.Println("ERROR:", err)
		log.Println("service shutting down due to fatal error")
		os.Exit(1)
	}
}

func service() error {

	fmt.Println()
	log.Println(": SERVICE START")

	
	// Logging
	// ==================================================

	log := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("logging initiated")


	// Configuration
	// ==================================================

	cfg := conf.NewConfig()
	if err := cfg.Parse(); err != nil {
		log.Println("ERROR:", err)
		log.Println("config init failed - using default values")
	} else {
		log.Println("config initiated")
	}


	// Templates
	// ==================================================

	tmpC, err := template.ParseGlob(cfg.Tmp.Catalog) 
	if err != nil {
		return err
	}

	tmpE, err := template.ParseGlob(cfg.Tmp.Editor)
	if err != nil {
		return err
	}

	log.Println("templates initiated")


	// File Servers
	// ==================================================

	editor := http.StripPrefix("/static/editor/", http.FileServer(http.Dir("./static/editor/")))
	http.Handle("/static/editor/", editor)

	catalog := http.StripPrefix("/static/catalog/", http.FileServer(http.Dir("./static/catalog/")))
	http.Handle("/static/catalog/", catalog)

	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images/")))
	http.Handle("/images/", images)


	// Route Handlers
	// ==================================================

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		routes.Catalog(w, r, log, cfg.ConnStr, tmpC)
	})

	http.HandleFunc("/editor/", func(w http.ResponseWriter, r *http.Request) {
		routes.Editor(w, r, log, cfg, tmpE)
	})

	http.HandleFunc("/authentication", func(w http.ResponseWriter, r *http.Request) {
		routes.Authentication(w, r, log, cfg)
	})

	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		routes.FileTransfer(w, r, log, cfg.ImgDir)
	})


	// Server
	// ==================================================

	log.Println("Now serving on " + cfg.HostAddr)
	if err := http.ListenAndServe(cfg.HostAddr, nil); err != nil {
		return err
	}

	
	return nil
}