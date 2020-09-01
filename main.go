package main

import (
	"HarpBlog/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", home)

	r.GET("/md", func(c *gin.Context) {
		m := "### A Title"
		c.Data(200, "text/html; charset=utf-8", utils.MD2HtmlBytes(m))
	})

	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	_ = r.Run()
}

func home(c *gin.Context) {
	c.String(200, "Hello World")
}
