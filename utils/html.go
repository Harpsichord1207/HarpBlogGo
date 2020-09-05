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

func GeneratePagination(total int) template.HTML {
	articlesPerPage := 5
	html := "<nav aria-label=\"Page navigation\">"
	html += "<ul class=\"pagination\">"
	html += "<li class=\"page-item\"><a class=\"page-link\" href=\"#\">&laquo;</a></li>"
	for i:=0; i*articlesPerPage<total; i++ {
		html += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"\\?page=%d\">%d</a></li>", i+1, i+1)
	}
	html += "<li class=\"page-item\"><a class=\"page-link\" href=\"#\">&raquo;</a></li>"
	html += "</ul></nav>"
	return template.HTML(html)
}