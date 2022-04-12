package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
)

type ArticleInfo struct {
	Id       uint64   `json:"id"`
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Time     string   `json:"time"`
	Tags     []string `json:"tags"`
	Abstract string   `json:"abstract"`
}

func GetArticlesNumber() int {
	rd, _ := ioutil.ReadDir("./data/articles")
	cnt := 0
	for _, f := range rd {
		if !f.IsDir() {
			cnt++
		}
	}
	return cnt
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

func GetArticles(page int) []ArticleInfo {

	var articles []ArticleInfo
	cnt := GetArticlesNumber()
	articlesPerPage := 5

	cnt -= articlesPerPage * (page - 1)
	for ; articlesPerPage > 0 && cnt > 0; articlesPerPage-- {
		articles = append(articles, GetArticleMeta(uint64(cnt)))
		cnt--
	}
	return articles
}

func SearchArticles(keyword string) []ArticleInfo {
	var articles []ArticleInfo
	cnt := GetArticlesNumber()

	keywordSL := ""
	keywordSL = strings.Replace(keyword, " ", "", -1)
	keywordSL = strings.ToLower(keywordSL)

	articlesCnt := 0
	for i := 1; i <= cnt; i++ {
		article := GetArticleMeta(uint64(i))
		title := strings.ToLower(article.Title)
		abstract := strings.ToLower(article.Abstract)
		tags := strings.ToLower(strings.Join(article.Tags, ","))
		if strings.Contains(title, keywordSL) || strings.Contains(abstract, keywordSL) || strings.Contains(tags, keywordSL) {
			articles = append(articles, article)
			articlesCnt += 1
		}
	}

	// reverse to order by timestamp desc
	for i := 0; i < articlesCnt/2; i++ {
		j := articlesCnt - i - 1
		articles[i], articles[j] = articles[j], articles[i]
	}

	return articles
}
