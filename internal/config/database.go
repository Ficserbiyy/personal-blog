package config

import (
	"fmt"
	"os"

	"github.com/Ficserbiyy/personal-blog/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	postgresPort = "5432"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbName     = os.Getenv("DB_NAME")
	dbPassword = os.Getenv("DB_PASSWORD")
)

// The ConnectToDatabase function
// establishes a database connection.
func ConnectToDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		postgresPort,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to connect to database: %w",
			err,
		)
	}

	err = db.AutoMigrate(
		&models.Post{},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to run auto migration: %w",
			err,
		)
	}

	return db, nil
}
