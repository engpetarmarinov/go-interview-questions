package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./questions/zip/test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	zipFile, err := os.Create("./questions/zip/test.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	zippedFile, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(zippedFile, f); err != nil {
		panic(err)
	}
}
