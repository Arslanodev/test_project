package repositories

import (
	"example/blog-app/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error

	return user, err
}

func (r *UserRepository) RegisterUser(user models.User) error {

	err := r.db.Create(&user).Error

	return err
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username= ?", username).Error

	return user, err
}

func (r *UserRepository) DeleteUser(id int) error {
	err := r.db.Delete(&models.User{}, id).Error

	return err
}
