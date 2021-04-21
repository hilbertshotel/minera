package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"os"
	"minera/routes"
	"minera/conf"
	"database/sql"
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


	// Database connection
	// ==================================================
	
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("database connection established")


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


	// Server
	// ==================================================

	api := http.Server{
		Addr: cfg.HostAddr,
		Handler: routes.Mux(log, cfg, db, tmpC, tmpE),
		ReadTimeout: cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	
	log.Println("api initiated")

	log.Println("Now listening on " + cfg.HostAddr)
	if err := api.ListenAndServe(); err != nil {
		return err
	}

	
	return nil
}