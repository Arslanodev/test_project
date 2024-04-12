package config

import (
	"example/blog-app/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("primary.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Blog{}, &models.User{})

	return db, nil
}
