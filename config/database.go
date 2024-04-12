package config

import (
	"example/blog-app/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("primary.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Blog{}, &models.User{})

	return db, nil
}
