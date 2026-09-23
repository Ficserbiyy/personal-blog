package handlers

import (
	"context"

	"github.com/Ficserbiyy/personal-blog/internal/models"
	"gorm.io/gorm"
)

// This function returns Post
// if found in the database,
// otherwise gorm.ErrRecordNotFound.
func getPostByID(
	id uint,
	db *gorm.DB,
	ctx context.Context,
) (models.Post, error) {
	var post models.Post

	err := db.WithContext(ctx).
		Where("id = ?", id).
		First(&post).Error

	return post, err
}

// This function removes Post
// from the database.
func deletePost(post models.Post, db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).
		Delete(&post).Error
}
