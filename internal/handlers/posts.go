package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Ficserbiyy/personal-blog/internal/config"
	"github.com/Ficserbiyy/personal-blog/internal/models"
	"gorm.io/gorm"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, "index", nil)
}

func NewPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	config.RenderTemplate(w, "new", nil)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.PathValue("id"))
	page := models.PostPage{ID: uint(id)}

	config.RenderTemplate(w, "edit", page)
}

func (s *BlogService) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

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

			if err := CreatePost(post, s.DB, ctx); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		// Retrieve all posts from the database
		// and render home.html.
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var rows []models.Post
		if err := s.DB.WithContext(ctx).Find(&rows).Error; err != nil {
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

func (s *BlogService) ArticlePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hiddenMethod string
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, "article ID must be a positive integer", http.StatusBadRequest)
			return
		}

		// This allows only GET, DELETE and PATCH requests.
		if r.Method == http.MethodPost {
			// Parse form data to read hidden fields
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			hiddenMethod = r.FormValue("_method")
			if hiddenMethod != http.MethodDelete && hiddenMethod != http.MethodPatch {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}

		ctx := r.Context()
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

		// Handle PATCH request
		if hiddenMethod == http.MethodPatch {
			newPost := models.Post{
				Title: r.FormValue("title"),
				Body:  r.FormValue("body"),
			}

			if err := updatePost(post, newPost, s.DB, ctx); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/articles/"+idStr, http.StatusSeeOther)
			return
		}

		// Handle DELETE request
		if err := deletePost(post, s.DB, ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
