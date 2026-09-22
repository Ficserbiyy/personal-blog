package handlers

import "net/http"

func (s *BlogService) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		}
	}
}
