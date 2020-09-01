package utils

import (
	"os"
	"path/filepath"
	"github.com/russross/blackfriday/v2"
)

func MD2HtmlBytes(s string) []uint8 {
	b := []byte(s)
	o := blackfriday.Run(b)
	return o
}

func MD2HtmlString(s string) string {
	return string(MD2HtmlBytes(s))
}

func listArticlesFiles() []string {
	var files []string
	dataPath := "./data"

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
