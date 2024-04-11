package controllers

import (
	"example/blog-app/app/models"
	"example/blog-app/app/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController interface {
	GetBlogs() ([]models.Blog, error)
	GetBlogById() (models.Blog, error)
	CreateBlog() (models.Blog, error)
	DeleteBlog() error
}

type blog_controller struct {
	repo repositories.BlogRepository
}

func NewBlogControllers(repo repositories.BlogRepository) *blog_controller {
	return &blog_controller{repo: repo}
}

func (c *blog_controller) GetBlogs(ctx *gin.Context) {
	blogs, err := c.repo.GetBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (c *blog_controller) GetBlogByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	blog, err := c.repo.GetBlogByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog by id"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *blog_controller) CreateBlog(ctx *gin.Context) {
	var blog models.Blog
	ctx.BindJSON(&blog)
	fmt.Println(blog)
	blog, err := c.repo.CreateBlog(blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a blog"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *blog_controller) DeleteBlog(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.repo.DeleteBlog(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog was successfully deleted"})
}
