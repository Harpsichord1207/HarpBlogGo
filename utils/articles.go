package utils

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
)

type Article struct {
	
}


func GetArticle(id uint64) string {
	filePath := fmt.Sprintf("%s%d%s", "./data/", id, ".md")
	b, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	return string(blackfriday.Run(b))
}
