package tools

import (
	"html/template"
	"log"
	"net/http"
)

var templatesDir = "./static/templates"
var layoutsTpl *template.Template

func init() {
	layoutsTpl = template.Must(template.ParseGlob(templatesDir + "/layouts/*.html")).Funcs(template.FuncMap{
		"calcPrice": calcPrice,
	})
}

func Render(w http.ResponseWriter, page string, data interface{}) {
	layoutClone := template.Must(layoutsTpl.Clone())
	pageTpl := template.Must(layoutClone.ParseFiles(templatesDir + "/pages/" + page))
	if err := pageTpl.ExecuteTemplate(w, page, data); err != nil {
		log.Fatal(err)
	}
}

func calcPrice(quantity int, price float64) float64 {
	return float64(quantity) * price
}
