package routes

import (
	"io"
	"os"
	"net/http"
	"minera/data"
	"minera/logs"
)

func FileTransfer(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(128 << 20)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	files := request.MultipartForm.File["files"]
	
	folderList, err := listFolder(data.ImageDir)
	if err != nil {
		logs.Errors.Println(err)
		http.Error(writer, "Възникна грешка", 502)
		return
	}

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		defer file.Close()

		// VERIFY IF IMAGE IS ACTUALLY AN IMAGE

		filename := files[i].Filename
		if contains(folderList, filename) { continue }

		out, err := os.Create(data.ImageDir + files[i].Filename)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			logs.Errors.Println(err)
			http.Error(writer, "Възникна грешка", 502)
			return
		}
	}
}


func listFolder(folder string) ([]string, error) {
	file, err := os.Open(folder)
	if err != nil { return []string{}, err }
	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil { return []string{}, err }
	return list, nil
}


func contains(list []string, filename string) bool {
	for _, file := range list {
		if file == filename { return true }
	}
	return false
}