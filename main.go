package main

import (
	"example/blog-app/app/routes"
	"example/blog-app/config"

	"gorm.io/gorm"
)

var db *gorm.DB

// Main execution
func main() {
	db, err := config.ConnectToDatabase()
	if err != nil {
		panic("Database connection error")
	}

	r := routes.SetupRouter(db)

	r.Run("localhost:8080")
}
