package handlers

import (
	"net/http"

	"github.com/Ficserbiyy/personal-blog/internal/config"
	"github.com/Ficserbiyy/personal-blog/internal/models"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, "index", nil)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	config.RenderTemplate(w, "new", nil)
}

func (s *BlogService) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new Post in the database
		// and redirect to /home.
		if r.Method == http.MethodPost {

			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			post := models.Post{
				Title: r.FormValue("title"),
				Body:  r.FormValue("body"),
			}

			if err := s.DB.WithContext(r.Context()).Create(&post).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}

		// Retrieve all posts from the database
		// and render home.html.
		if r.Method == http.MethodGet {

			var rows []models.Post

			if err := s.DB.Find(&rows).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var articles []models.PostPage
			for i := range rows {
				articles = append(articles, rows[i].ResponseModel())
			}

			config.RenderTemplate(w, "home", articles)
		}
	}
}
