package routes

import (
	"io"
	"os"
	"log"
	"net/http"
)

func FileTransfer(w http.ResponseWriter, r *http.Request, log *log.Logger, imgDir string) {
	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}

	files := r.MultipartForm.File["files"]
	
	folderList, err := listFolder(imgDir)
	if err != nil {
		http.Error(w, "Backend Error", 502)
		log.Println("ERROR:", err)
		return
	}

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
			return
		}
		defer file.Close()

		filename := files[i].Filename
		if contains(folderList, filename) {
			continue
		}

		out, err := os.Create(imgDir + filename)
		if err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Backend Error", 502)
			log.Println("ERROR:", err)
		}
	}
}


func listFolder(folder string) ([]string, error) {
	file, err := os.Open(folder)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	list, err := file.Readdirnames(0)
	if err != nil {
		return []string{}, err
	}

	return list, nil
}


func contains(list []string, filename string) bool {
	for _, file := range list {
		if file == filename {
			return true
		}
	}

	return false
}