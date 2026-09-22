package config

import (
	"log"
	"net/http"
	"text/template"
)

var (
	templates = template.Must(template.ParseFiles(
		"index.html",
		"templates/home.html",
		"templates/new.html",
	))
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
