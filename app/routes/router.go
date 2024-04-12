package routes

import (
	"example/blog-app/app/controllers"
	"example/blog-app/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	blogRepo := repositories.NewBlogRepository(db)
	blogController := controllers.NewBlogControllers(*blogRepo)

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	blog_route := r.Group("api/v1/blog", userController.RequireAuth)

	blog_route.GET("/", blogController.GetBlogs)
	blog_route.GET("/:id", blogController.GetBlogByID)
	blog_route.POST("/", userController.IsAdmin, blogController.CreateBlog)
	blog_route.DELETE("/:id", userController.IsAdmin, blogController.DeleteBlog)

	auth_route := r.Group("api/v1/user")
	auth_route.POST("/register", userController.RegisterUser)
	auth_route.POST("/login", userController.LoginUser)

	return r
}
