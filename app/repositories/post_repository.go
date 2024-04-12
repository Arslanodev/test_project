package repositories

import (
	"example/blog-app/app/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) GetPosts() ([]models.Post, error) {
	var blogs []models.Post

	err := r.db.Find(&blogs).Error

	return blogs, err
}

func (r *PostRepository) GetPostByID(id int) (models.Post, error) {
	var blog models.Post

	err := r.db.Find(&blog, id).Error

	return blog, err
}

func (r *PostRepository) CreatePost(blog models.Post) (models.Post, error) {

	err := r.db.Create(&blog).Error

	return blog, err
}

func (r *PostRepository) DeletePost(id int) error {

	err := r.db.Delete(&models.Post{}, id).Error
	return err
}
