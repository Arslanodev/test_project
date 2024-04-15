package repositories

import (
	"example/blog-app/cmd/data"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]data.User, error)
	GetUserByID(id int) (data.User, error)
	CreateNewUser(user data.User) error
	GetUserByUsername(username string) (data.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]data.User, error) {
	var users []data.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserByID(id int) (data.User, error) {
	var user data.User
	err := r.db.Find(&user, id).Error

	return user, err
}

func (r *userRepository) CreateNewUser(user data.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *userRepository) GetUserByUsername(username string) (data.User, error) {
	var user data.User
	err := r.db.First(&user, "username= ?", username).Error

	return user, err
}

func (r *userRepository) DeleteUser(id int) error {
	err := r.db.Delete(&data.User{}, id).Error

	return err
}
