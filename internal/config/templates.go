package config

import (
	"net/http"
	"text/template"

	"github.com/Ficserbiyy/personal-blog/internal/models"
)

var (
	templates = template.Must(template.ParseFiles(
		"index.html",
		"templates/new.html",
	))
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p *models.PostPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
