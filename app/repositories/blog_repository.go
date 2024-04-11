package repositories

import (
	"example/blog-app/app/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog

	err := r.db.Find(&blogs).Error

	return blogs, err
}

func (r *BlogRepository) GetBlogByID(id int) (models.Blog, error) {
	var blog models.Blog

	err := r.db.Find(&blog, id).Error

	return blog, err
}

func (r *BlogRepository) CreateBlog(blog models.Blog) (models.Blog, error) {

	err := r.db.Create(&blog).Error

	return blog, err
}

func (r *BlogRepository) DeleteBlog(id int) error {

	err := r.db.Delete(&models.Blog{}, id).Error
	return err
}
