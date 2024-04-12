package controllers

import (
	"example/blog-app/app/models"
	"example/blog-app/app/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	repo repositories.BlogRepository
}

func NewBlogControllers(repo repositories.BlogRepository) *BlogController {
	return &BlogController{repo: repo}
}

func (c *BlogController) GetBlogs(ctx *gin.Context) {
	blogs, err := c.repo.GetBlogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	blog, err := c.repo.GetBlogByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog by id"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
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

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.repo.DeleteBlog(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog was successfully deleted"})
}
