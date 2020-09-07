package main

import (
	"HarpBlog/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.LoadHTMLFiles("templates/article.html", "templates/home.html")

	r.GET("/", func(c *gin.Context) {
		pageS := c.Query("page")
		page := 1
		if pageS != "" {
			page, _ = strconv.Atoi(pageS)
		}
		c.HTML(200, "home.html", gin.H{
			"articles": utils.GetArticles(page),
			"nav": utils.GetHTMLComponent("nav.html"),
			"calendar": utils.GetHTMLComponent("calendar.html"),
			"pagination": utils.GeneratePagination(page),
		})
	})

	r.GET("/articles/:articleId", func(c *gin.Context) {
		articleId, err := strconv.ParseUint(c.Param("articleId"), 10, 16)
		if err != nil {
			panic(err)
		}
		c.HTML(200, "article.html", gin.H{
			"content": template.HTML(utils.GetArticleContent(articleId)),
			"info": utils.GetArticleMeta(articleId),
			"nav": utils.GetHTMLComponent("nav.html"),
			"calendar": utils.GetHTMLComponent("calendar.html"),
		})
	})

	r.GET("/search", func(c *gin.Context) {

		keyword := c.Query("keyword")

		if keyword != "" {
			c.HTML(200, "home.html", gin.H{
				"articles": utils.SearchArticles(keyword),
				"nav": utils.GetHTMLComponent("nav.html"),
				"calendar": utils.GetHTMLComponent("calendar.html"),
			})
		} else {
			c.Redirect(302,"/")
		}

	})

	_ = r.Run(":6969")
}
