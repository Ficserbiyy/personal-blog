package models

import "time"

type (
	Post struct {
		ID        uint   `gorm:"primaryKey"`
		Title     string `gorm:"not null"`
		Body      string `gorm:"type:text;not null"`
		CreatedAt time.Time
	}

	// PostPage is a page view model
	// sent to the template renderer.
	PostPage struct {
		ID        uint
		Title     string
		CreatedAt string
		Body      []byte // Convenient for html/template or raw rendering
	}
)
