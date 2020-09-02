package main

import (
	"HarpBlog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()

	//r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})

	r.GET("/articles/:articleId", func(c *gin.Context) {
		articleId, err := strconv.ParseUint(c.Param("articleId"), 10, 16)
		if err != nil {
			panic(err)
		}
		c.Data(200, "text/html; charset=utf-8", utils.GetArticle(articleId))
	})

	_ = r.Run()
}
