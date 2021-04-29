package routes

import (
	"net/http"
	"minera/config"
	"database/sql"
	"html/template"
	"log"
)

func Mux(
	log *log.Logger,
	cfg *config.Config,
	db *sql.DB,
	tmpC *template.Template,
	tmpE *template.Template) *http.ServeMux {

	mux := http.NewServeMux()

	// File Servers
	// ==================================================

	edit := http.StripPrefix("/static/editor/", http.FileServer(http.Dir("./static/editor/")))
	mux.Handle("/static/editor/", edit)

	cat := http.StripPrefix("/static/catalog/", http.FileServer(http.Dir("./static/catalog/")))
	mux.Handle("/static/catalog/", cat)

	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images/")))
	mux.Handle("/images/", images)


	// Route Handlers
	// ==================================================

	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		catalog(w, r, log, cfg.ConnStr, tmpC, db)
	})

	mux.HandleFunc("/editor/", func(w http.ResponseWriter, r *http.Request) {
		editor(w, r, log, cfg, tmpE, db)
	})

	mux.HandleFunc("/authentication", func(w http.ResponseWriter, r *http.Request) {
		authentication(w, r, log, cfg, db)
	})

	mux.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		fileTransfer(w, r, log, cfg.ImgDir)
	})

	mux.Handle("/minera.log", http.FileServer(http.Dir("./logs/")))

	return mux
}