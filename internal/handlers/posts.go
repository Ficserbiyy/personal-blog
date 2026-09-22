package handlers

import (
	"net/http"

	"github.com/Ficserbiyy/personal-blog/internal/models"
)

// Create method creates a new Post
// in the database and redirects to /tasks.
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

			http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		}
	}
}
