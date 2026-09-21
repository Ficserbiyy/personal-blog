package config

import (
	"net/http"
	"text/template"
)

type Post struct {
	ID        uint
	Title     string
	CreatedAt string
	Body      []byte
}

var (
	templates = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/new.html",
	))
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Post) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
