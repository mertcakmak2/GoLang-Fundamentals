package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("creating zip archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening second file")
	file, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("writing second file to archive...")
	w2, err := zipWriter.Create(file.Name())
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w2, file); err != nil {
		panic(err)
	}
	fmt.Println("closing zip archive...")
	zipWriter.Close()

}
