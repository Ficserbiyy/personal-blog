package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Ficserbiyy/personal-blog/internal/config"
	"github.com/Ficserbiyy/personal-blog/internal/models"
	"gorm.io/gorm"
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

func (s *BlogService) ArticlePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		ctx := r.Context()

		if err != nil || id <= 0 {
			http.Error(w, "article ID must be a positive integer", http.StatusBadRequest)
			return
		}

		post, err := getPostByID(uint(id), s.DB, ctx)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Handle GET request
		if r.Method == http.MethodGet {
			config.RenderTemplate(w, "article", post.ResponseModel())
			return
		}

		// Handle DELETE request
		if r.Method == http.MethodPost {
			// Parse form data to read hidden fields
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if r.FormValue("_method") != http.MethodDelete {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			if err := deletePost(post, s.DB, ctx); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
}
