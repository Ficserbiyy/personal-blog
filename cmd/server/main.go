package main

import (
	"log"
	"net/http"

	"github.com/Ficserbiyy/personal-blog/internal/config"
	"github.com/Ficserbiyy/personal-blog/internal/handlers"
)

func main() {
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	repo := handlers.NewBlogService(db)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetIndex)
	mux.HandleFunc("GET /home", repo.ListAll())
	mux.HandleFunc("POST /home", repo.Create())

	log.Println("Server listening on http://127.0.0.1:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
