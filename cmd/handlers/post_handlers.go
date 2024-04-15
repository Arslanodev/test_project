package handlers

import (
	"example/blog-app/cmd/data"
	"example/blog-app/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	repo repositories.PostRepository
}

func NewPostHandler(repo repositories.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) GetPosts(ctx *gin.Context) {
	blogs, err := h.repo.GetPosts()
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

func (h *PostHandler) GetPostByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	blog, err := h.repo.GetPostByID(id)
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

func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var blog data.Post
	ctx.BindJSON(&blog)
	blog, err := h.repo.CreatePost(blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a blog"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (h *PostHandler) DeletePost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, _ := h.repo.GetPostByID(id)
	if post.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "There is no post at that id to delete",
		})
		return
	}
	err := h.repo.DeletePost(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post was successfully deleted"})
}
