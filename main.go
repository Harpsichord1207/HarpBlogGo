package main

import (
	"HarpBlog/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", home)

	r.GET("/md", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", utils.GetArticle(1))
	})

	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	_ = r.Run()
}

func home(c *gin.Context) {
	c.HTML(200, "base.html", gin.H{})
}
