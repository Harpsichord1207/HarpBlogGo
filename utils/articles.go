package utils

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
)

func GetArticle(id uint64) []uint8 {
	filePath := fmt.Sprintf("%s%d%s", "./data/", id, ".md")
	b, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	return blackfriday.Run(b)
}
