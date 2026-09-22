package handlers

import "gorm.io/gorm"

type BlogService struct {
	DB *gorm.DB
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{DB: db}
}
