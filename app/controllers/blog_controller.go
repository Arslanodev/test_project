package controllers

import (
	"example/blog-app/app/models"
	"example/blog-app/app/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	repo repositories.PostRepository
}

func NewPostControllers(repo repositories.PostRepository) *PostController {
	return &PostController{repo: repo}
}

func (c *PostController) GetPosts(ctx *gin.Context) {
	blogs, err := c.repo.GetPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	if len(blogs) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "There are no posts yet"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (c *PostController) GetPostByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	blog, err := c.repo.GetPostByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog by id"})
		return
	}

	if blog.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "There is no post at that id"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var blog models.Post
	ctx.BindJSON(&blog)
	blog, err := c.repo.CreatePost(blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a blog"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, _ := c.repo.GetPostByID(id)
	if post.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "There is no post at that id to delete",
		})
		return
	}
	err := c.repo.DeletePost(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post was successfully deleted"})
}
