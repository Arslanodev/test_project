package repositories

import (
	"example/blog-app/cmd/data"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPosts() ([]data.Post, error)
	GetPostByID(id int) (data.Post, error)
	CreatePost(post data.Post) (data.Post, error)
	DeletePost(id int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetPosts() ([]data.Post, error) {
	var blogs []data.Post

	err := r.db.Find(&blogs).Error

	return blogs, err
}

func (r *postRepository) GetPostByID(id int) (data.Post, error) {
	var post data.Post

	err := r.db.Find(&post, id).Error

	return post, err
}

func (r *postRepository) CreatePost(post data.Post) (data.Post, error) {

	err := r.db.Create(&post).Error

	return post, err
}

func (r *postRepository) DeletePost(id int) error {

	err := r.db.Delete(&data.Post{}, id).Error
	return err
}
