package main

import (
	"log"
	"net/http"
	"html/template"
	"os"
	"database/sql"
	"syscall"
	"os/signal"

	"minera/routes"
	"minera/config"
)

func main() {

	if err := service(); err != nil {
		log.Println("ERROR:", err)
		log.Println("SERVICE STOP\n")
		os.Exit(1)
	} else {
		log.Println("SERVICE STOP\n")
		os.Exit(0)
	}

}

func service() error {

	// Logging
	// ==================================================

	log := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	log.Println("SERVICE START")
	log.Println("logging initiated")


	// Configuration
	// ==================================================

	cfg := config.New()

	if err := cfg.Parse(); err != nil {
		log.Println("ERROR:", err)
		log.Println("config parsing failed: loading default values")
	} else {
		log.Println("config initiated")
	}


	// Database connection
	// ==================================================
	
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		return err
	}

	defer log.Println("closing database connection")
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


	// Channels
	// ==================================================

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	
	serverError := make(chan error, 1)

	log.Println("channels initiated")


	// Api
	// ==================================================

	api := http.Server{
		Addr: cfg.HostAddr,
		Handler: routes.Mux(log, cfg, db, tmpC, tmpE),
		ReadTimeout: cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	
	log.Println("api initiated")


	// Server
	// ==================================================

	go func() {
		log.Println("service listening on", cfg.HostAddr)
		serverError <- api.ListenAndServe()
	}()


	// Shutdown
	// ==================================================
	
	select {

	case <-shutdown:
		log.Println("SHUTDOWN initiated")
		return nil

	case err := <-serverError:
		return err

	}

}