package handlers

import (
	"net/http"

	"github.com/Ficserbiyy/personal-blog/internal/config"
	"github.com/Ficserbiyy/personal-blog/internal/models"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, "index", nil)
}

// Create method creates a new Post
// in the database and redirects to /home.
func (s *BlogService) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// ListAll method retrieves all posts from the database.
func (s *BlogService) ListAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
