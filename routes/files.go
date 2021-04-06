package routes

import (
	"io"
	"os"
	"net/http"
	"minera/data"
)

func FileTransfer(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(128 << 20)
	if err != nil { data.Log(err, writer); return }

	files := request.MultipartForm.File["files"]
	
	folderList, err := data.ListFolder(data.ImageDir)
	if err != nil { data.Log(err, writer); return }

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil { data.Log(err, writer); return }
		defer file.Close()

		filename := files[i].Filename
		if data.Contains(folderList, filename) { continue }

		out, err := os.Create(data.ImageDir + filename)
		if err != nil { data.Log(err, writer); return }
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil { data.Log(err, writer) }
	}
}
