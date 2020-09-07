package utils

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

func GetHTMLComponent(filename string) template.HTML {
	b, e := ioutil.ReadFile("./templates/components/" + filename)
	if e != nil {
		panic(e)
	}
	return template.HTML(string(b))
}

func GeneratePagination(page int) template.HTML {
	total := GetArticlesNumber()
	articlesPerPage := 5
	html := "<nav aria-label=\"Page navigation\">"
	html += "<ul class=\"pagination\">"

	if page == 1 {
		html += "<li class=\"page-item disabled\"><a class=\"page-link\" href=\"#\">&laquo;</a></li>"
	} else {
		html += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"\\?page=%d\">&laquo;</a></li>", page-1)
	}

	i := 0
	for ; i*articlesPerPage<total; i++ {
		if page == i + 1 {
			html += fmt.Sprintf("<li class=\"page-item active\"><a class=\"page-link\" href=\"\\?page=%d\">%d</a></li>", i+1, i+1)
		} else {
			html += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"\\?page=%d\">%d</a></li>", i+1, i+1)
		}
	}

	if page >= i {
		html += "<li class=\"page-item disabled\"><a class=\"page-link\" href=\"#\">&raquo;</a></li>"
	} else {
		html += fmt.Sprintf("<li class=\"page-item disable\"><a class=\"page-link\" href=\"\\?page=%d\">&raquo;</a></li>", page+1)
	}

	html += "</ul></nav>"

	return template.HTML(html)
}