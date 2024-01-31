package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	var path string
	fmt.Print("Введите полный путь: ")
	fmt.Scan(&path)
	fileName := filepath.Base(path)
	fileExt := strings.TrimPrefix(filepath.Ext(path), ".")
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
