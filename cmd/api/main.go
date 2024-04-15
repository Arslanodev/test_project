package main

import (
	"example/blog-app/cmd/db"
	"example/blog-app/cmd/handlers"
	"example/blog-app/cmd/repositories"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Main execution
func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setting up database connection
	db_connection, err := db.ConnectToDatabase()
	if err != nil {
		panic("Database connection error")
	}

	postRepo := repositories.NewPostRepository(db_connection)
	postHandler := handlers.NewPostHandler(postRepo)

	userRepo := repositories.NewUserRepository(db_connection)
	userHandler := handlers.NewUserHandler(userRepo)

	auth_route := r.Group("api/v1/user")
	auth_route.POST("/register", userHandler.RegisterUser)
	auth_route.POST("/login", userHandler.LoginUser)

	post_route := r.Group("api/v1/post", userHandler.RequireAuth)
	post_route.GET("/", postHandler.GetPosts)
	post_route.GET("/:id", postHandler.GetPostByID)
	post_route.POST("/", userHandler.IsAdmin, postHandler.CreatePost)
	post_route.DELETE("/:id", userHandler.IsAdmin, postHandler.DeletePost)

	fmt.Println("Server has started on: http://localhost:8080/")
	r.Run("localhost:8080")
}
