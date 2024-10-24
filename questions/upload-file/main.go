package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

var uploadTemplate = template.Must(template.ParseFiles("questions/upload-file/templates/upload.html"))

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /upload", getUploadHandler)
	mux.HandleFunc("POST /upload", postUploadHandler)

	http.ListenAndServe(":8080", mux)
}

func getUploadHandler(writer http.ResponseWriter, request *http.Request) {
	uploadTemplate.Execute(writer, nil)
}

func postUploadHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(32 << 20)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	file, header, err := request.FormFile("img")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	defer file.Close()

	if header.Header.Get("Content-Type") != "image/jpeg" && header.Header.Get("Content-Type") != "image/png" {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("only image/jpeg and image/png is supported"))
		return
	}

	uploadedFile, err := os.Create("questions/upload-file/uploads/" + header.Filename)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("<progress id='progress' value='100' max='100'></progress>"))
}
