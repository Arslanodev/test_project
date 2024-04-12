package routes

import (
	"example/blog-app/app/controllers"
	"example/blog-app/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	postRepo := repositories.NewPostRepository(db)
	postController := controllers.NewPostControllers(*postRepo)

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	post_route := r.Group("api/v1/post", userController.RequireAuth)

	post_route.GET("/", postController.GetPosts)
	post_route.GET("/:id", postController.GetPostByID)
	post_route.POST("/", userController.IsAdmin, postController.CreatePost)
	post_route.DELETE("/:id", userController.IsAdmin, postController.DeletePost)

	auth_route := r.Group("api/v1/user")
	auth_route.POST("/register", userController.RegisterUser)
	auth_route.POST("/login", userController.LoginUser)

	return r
}
