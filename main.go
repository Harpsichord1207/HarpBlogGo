package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func main() {
	r := gin.Default()

	r.GET("/", home)

}

func home(c *gin.Context) {
	c.HTML()
}