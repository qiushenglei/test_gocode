package myfile

import (
	"io"
	"net/http"
	"os"
)

func HttpServer() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(resp http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(32 << 20)

	file, fileHeader, err := request.FormFile("file")

	if err != nil {
		return
		panic(err)
	}

	newFile, err := os.Create(fileHeader.Filename)
	defer newFile.Close()

	// 方法1
	_, err = io.Copy(newFile, file)
	if err != nil {
		panic(err)
	}

	// 方法2
	bytes := make([]byte, fileHeader.Size)
	_, err = file.Read(bytes)
	_, err = newFile.Write(bytes)
}
