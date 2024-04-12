package main

import (
	"example/blog-app/app/routes"
	"example/blog-app/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Main execution
func main() {
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.ConnectToDatabase()
	if err != nil {
		panic("Database connection error")
	}

	r := routes.SetupRouter(db)

	fmt.Println("Server has started on: http://localhost:8080/")
	r.Run("localhost:8080")

}
