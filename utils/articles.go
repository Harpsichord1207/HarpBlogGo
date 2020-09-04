package utils

import (
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"os"
)

type ArticleInfo struct {
	Id       uint64     `json:"id"`
	Title    string		`json:"title"`
	Category string		`json:"category"`
	Time     string		`json:"time"`
	Tags     []string	`json:"tags"`
}

func GetArticleContent(id uint64) string {
	filePath := fmt.Sprintf("%s%d%s", "./data/articles/", id, ".md")
	b, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	return string(blackfriday.Run(b))
}

func GetArticleMeta(id uint64) ArticleInfo {
	filePath := fmt.Sprintf("%s%d%s", "./data/infos/", id, ".json")
	filePtr, e := os.Open(filePath)
	if e != nil {
		panic(e)
	}
	var article ArticleInfo
	decoder := json.NewDecoder(filePtr)
	_ = decoder.Decode(&article)
	return article
}
