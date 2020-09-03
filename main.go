package main

import (
	"HarpBlog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

func main() {
	r := gin.Default()

	//r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})

	r.GET("/articles/:articleId", func(c *gin.Context) {
		articleId, err := strconv.ParseUint(c.Param("articleId"), 10, 16)
		if err != nil {
			panic(err)
		}
		fmt.Println(utils.GetArticle(articleId))
		c.HTML(200, "article.html", gin.H{
			"content": template.HTML(utils.GetArticle(articleId)),
		})
	})

	_ = r.Run()
}
