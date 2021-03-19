package backend

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"minera/backend/utils"
)

func FileTransfer(writer http.ResponseWriter, request *http.Request) {
	// VALIDATOR
	cookie, err := request.Cookie(utils.CookieName)
	if err != nil {	utils.Logger.Println(err); return }

	if cookies[cookie.Name] != cookie.Value  {
		utils.Logger.Println("no such cookie `" + cookie.Name + "` in cookie jar")
		return
	}

	// INSERT FILES IN IMG FOLDER
	request.ParseMultipartForm(32 << 20) 
	files := request.MultipartForm.File["files"]
	
	folderList, err := listFolder(utils.ImageDir)
	if err != nil {	utils.Logger.Println(err); return }

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil { utils.Logger.Println(err); return }
		defer file.Close()

		// VERIFY IF IMAGE IS ACTUALLY AN IMAGE

		filename := files[i].Filename
		if contains(folderList, filename) { continue }

		fmt.Println(utils.ImageDir + files[i].Filename)

		out, err := os.Create(utils.ImageDir + files[i].Filename)
		if err != nil { utils.Logger.Println(err); return }
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil { utils.Logger.Println(err); return }
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
